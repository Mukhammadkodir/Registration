package postgres

import (
	pb "github/RegisterTask/register_service/genproto/register_service"

	"github.com/jmoiron/sqlx"
)

type RegisterRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *RegisterRepo {
	return &RegisterRepo{db: db}
}

func (r *RegisterRepo) Create(User *pb.User) (*pb.User, error) {
	query := `
        INSERT INTO 
			register (id, name, password, 
				   access_token, refresh_token)
        VALUES($1,$2,$3,$4,$5) 
		RETURNING id
    `
	err := r.db.DB.QueryRow(query,
		User.Id,
		User.Name,
		User.Password,
		User.AccessToken,
		User.RefreshToken,
	).Scan(&User.Id)

	user, err := r.Get(&pb.ById{Id: User.Id})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *RegisterRepo) Login(L *pb.Loginn) (*pb.User, error) {
	query := `
        SELECT
			id,
			name,
			password, 
			access_token, refresh_token
        FROM register
        WHERE name = $1
    `
	var User pb.User
	err := r.db.DB.QueryRow(query,
		L.Name,
	).Scan(
		&User.Id,
		&User.Name,
		&User.Password,
		&User.AccessToken,
		&User.RefreshToken,
	)
	if err != nil {
		return nil, err
	}

	return &User, nil
}

func (r *RegisterRepo) Get(id *pb.ById) (*pb.User, error) {
	query := `
        SELECT
			id,
			name, 
			password, 
			access_token, refresh_token
        FROM register
        WHERE id = $1
    `
	var User pb.User
	err := r.db.DB.QueryRow(query, id.Id).Scan(
		&User.Id,
		&User.Name,
		&User.Password,
		&User.AccessToken,
		&User.RefreshToken,
	)
	if err != nil {
		return nil, err
	}
	return &User, nil
}
