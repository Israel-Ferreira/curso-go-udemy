package repositories

import (
	"database/sql"
	"go-crud-api/models"
)

type UsuarioRepo struct {
	Db *sql.DB
}


func NewUsuarioRepo(db *sql.DB) UsuarioRepo {
	return UsuarioRepo{Db: db}
}

func (u UsuarioRepo) CreateUser(user models.Usuario) error {
	stmt, err := u.Db.Prepare(
		`insert into usuarios(nome, email) values (?, ?)`,
	)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		user.Nome,
		user.Email,
	)

	if err != nil {
		return err
	}

	if err = stmt.Close(); err != nil {
		return err
	}

	return nil

}
