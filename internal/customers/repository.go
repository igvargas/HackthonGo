package internal

import (
	"log"

	"github.com/igvargas/HackthonGo/internal/models"
	"github.com/igvargas/HackthonGo/pkg/db"
)

const (
	queryStore = "INSERT INTO mydb(id, nombre, apellido,condicion) VALUES( ?, ?, ?, ?)"
)

type RepositorySQL interface {
	Store(customer models.Customer) (models.Customer, error)
}

type repositorySQL struct {
}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(customer models.Customer) (models.Customer, error) {
	db := db.StorageDB
	stmt, err := db.Prepare(queryStore)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(customer.ID, customer.Nombre, customer.Apellido, customer.Condicion)
	if err != nil {
		return models.Customer{}, err
	}
	idCreado, err := result.LastInsertId()
	customer.ID = int(idCreado)
	return customer, nil
}
