package service

import (
	"context"

	pb "github/RegisterTask/register_service/genproto/register_service"
	l "github/RegisterTask/register_service/pkg/logger"
	"github/RegisterTask/register_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RegisterService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewRegisterService(db *sqlx.DB, log l.Logger) *RegisterService {
	return &RegisterService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *RegisterService) Login(ctx context.Context, req *pb.Loginn) (*pb.User, error) {

	User, err := s.storage.User().Login(req)
	if err != nil {
		s.logger.Error("Error getting User Profile \n \n", l.Error(err))
		return nil, err
	}

	return User, nil
}

func (s *RegisterService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {

	req.Id = uuid.New().String()
	User, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Error create User \n \n", l.Error(err))
		return nil, err
	}

	return User, nil
}

func (s *RegisterService) Get(ctx context.Context, req *pb.ById) (*pb.User, error) {
	user, err := s.storage.User().Get(req)
	if err != nil {
		s.logger.Error("Failed Get user", l.Error(err))
		return nil, nil
	}

	return user, nil
}