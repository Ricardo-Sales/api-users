package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Ricardo-Sales/api-users/models"
	"github.com/gorilla/mux"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario
	var err error

	if usuarios, err = models.GetAll(); err != nil {
		w.Write([]byte("erro ao buscar o usuario"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("erro ao "))
		return
	}
	fmt.Println(usuarios)
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	var err error
	// pegar o id no path
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("erro ao fazer o parse do id"))
	}
	usuario.ID = uint32(ID)

	if err = usuario.GetOne(); err != nil {
		w.Write([]byte("erro ao buscar o usuario"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("erro ao realizar o encode do usuario"))
		return
	}
	fmt.Println(usuario)
}

func PostUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("erro ao ler o corpo da requisição"))
		return
	}
	if err = json.Unmarshal(body, &usuario); err != nil {
		w.Write([]byte("erro ao realizar unmarshall do body"))
		return
	}
	if err = usuario.Save(); err != nil {
		w.Write([]byte("erro ao salvar o usuario no banco"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("erro ao realizar o encode do usuario"))
	}
	fmt.Println(usuario)
}
func PutUsuario(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario
	param := mux.Vars(r)
	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		w.Write([]byte("erro ao fazer o parse do ID"))
		return
	}
	usuario.ID = uint32(ID)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("erro ao ler o corpo da requisicao"))
		return
	}
	if err = json.Unmarshal(body, &usuario); err != nil {
		w.Write([]byte("erro ao realizar unmarshal para usuario"))
		return
	}

	if err = usuario.Update(); err != nil {
		w.Write([]byte("erro ao atualizar o usuario"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("erro ao realizar o encode do usuario atualizado"))
	}
	fmt.Println(usuario)
}
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	// ler o id
	param := mux.Vars(r)
	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		w.Write([]byte("erro ao obter o id da requisição"))
		return
	}
	usuario.ID = uint32(ID)
	err = usuario.Delete()
	if err != nil {
		w.Write([]byte("erro ao deletar o usuario"))
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
