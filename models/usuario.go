package models

import (
	"github.com/Ricardo-Sales/api-users/banco"
)

type Usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func (usuario *Usuario) GetOne() error {
	db, err := banco.Conectar()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("Select nome, email from usuarios where id =?", usuario.ID)
	if err != nil {
		return err
	}
	if rows.Next() {
		if err = rows.Scan(&usuario.Nome, &usuario.Email); err != nil {
			return err
		}
	}
	return nil
}

func GetAll() ([]Usuario, error) {
	var usuarios []Usuario

	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("Select id, nome, email from usuarios")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var usuario Usuario
		if err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (usuario *Usuario) Save() error {
	db, err := banco.Conectar()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		return err
	}

	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		return err
	}
	ID, err := insert.LastInsertId()
	if err != nil {
		return err
	}
	usuario.ID = uint32(ID)

	return nil
}

func (usuario *Usuario) Update() error {

	db, err := banco.Conectar()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) Delete() error {
	db, err := banco.Conectar()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id  = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(usuario.ID)
	if err != nil {
		return err
	}

	return nil
}
