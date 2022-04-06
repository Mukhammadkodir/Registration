package repo

import (
	pb "github/RegisterTask/register_service/genproto/register_service"
)

//RergisterStorageI ...
type RegisterStorageI interface {
	Create(*pb.User) (*pb.User, error)
	Get(*pb.ById) (*pb.User, error)
	Login(*pb.Loginn) (*pb.User, error)
}
