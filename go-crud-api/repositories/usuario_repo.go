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

func (u UsuarioRepo) FindById(id uint64) (models.Usuario, error) {
	linha, err := u.Db.Query("select * from usuarios where id = ?", id)

	if err != nil {
		return models.Usuario{}, err
	}

	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil

}

func (u UsuarioRepo) FindAll() ([]models.Usuario, error) {
	lns, err := u.Db.Query("select * from usuarios")

	if err != nil {
		return nil, err
	}

	defer lns.Close()

	var usuarios []models.Usuario

	for lns.Next() {
		var usuario models.Usuario

		if erro := lns.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
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
