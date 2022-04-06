package storage

import (
	"github/RegisterTask/register_service/storage/postgres"
	"github/RegisterTask/register_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	User() repo.RegisterStorageI
}

type storagePg struct {
	db       *sqlx.DB
	UserRepo repo.RegisterStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		UserRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.RegisterStorageI {
	return s.UserRepo
}
