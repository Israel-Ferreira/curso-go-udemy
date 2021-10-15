package repositories

import (
	"database/sql"
	"fmt"
	"go-crud-api/models"
)

type UsuarioRepo struct {
	Db *sql.DB
}

func NewUsuarioRepo(db *sql.DB) UsuarioRepo {
	return UsuarioRepo{Db: db}
}

func (u UsuarioRepo) CloseDbConnection() {
	u.Db.Close()
}

// FindById Retorna um usuario pelo id informado
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

// FindAll retorna todos os usuarios do banco de dados
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

// CreateUser cria o usuario no banco de dados
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

func (u UsuarioRepo) FindByIdAndUpdate(id int, user models.Usuario) error {

	stmt, err := u.Db.Prepare(
		"update usuarios set nome =  ?, email = ? where id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println("Usuario: ", user)
	fmt.Println("ID: ", id)

	_, err = stmt.Exec(
		user.Nome,
		user.Email,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u UsuarioRepo) DeleteById(id int) error {
	stmt, err := u.Db.Prepare(
		"delete from usuarios where id = ?",
	)


	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		id,
	)

	if err != nil {
		return err
	}


	return	nil
}
