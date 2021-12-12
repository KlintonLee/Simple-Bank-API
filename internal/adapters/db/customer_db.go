package db

import (
	"database/sql"
	"time"

	"github.com/klintonlee/simple-bank-api/internal/modules/customers"
	_ "github.com/mattn/go-sqlite3"
)

type CustomerDb struct {
	db *sql.DB
}

func NewCustomerDb(db *sql.DB) *CustomerDb {
	return &CustomerDb{db: db}
}

func (c *CustomerDb) FindByCpf(cpf string) (customers.CustomerInterface, error) {
	var customer customers.Customer
	stmt, err := c.db.Prepare(`
		SELECT id, name, cpf, birth
		FROM customers
		WHERE cpf = ?;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(cpf).Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Birth)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerDb) Save(customer customers.CustomerInterface) (customers.CustomerInterface, error) {
	var rows int
	c.db.QueryRow("SELECT COUNT(id) FROM customers WHERE id = ?;", customer.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := c.create(customer)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := c.update(customer)
		if err != nil {
			return nil, err
		}
	}
	return customer, nil
}

func (c *CustomerDb) create(customer customers.CustomerInterface) (customers.CustomerInterface, error) {
	stmt, err := c.db.Prepare(`
		INSERT INTO customers(id, name, cpf, birth, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?);
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		customer.GetID(),
		customer.GetName(),
		customer.GetCpf(),
		customer.GetBirth(),
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerDb) update(customer customers.CustomerInterface) (customers.CustomerInterface, error) {
	_, err := c.db.Exec(`
		UPDATE customers
		SET name = ?, cpf = ?, birth = ?
		WHERE id = ?
	`, customer.GetName(), customer.GetCpf(), customer.GetBirth(), customer.GetID())
	if err != nil {
		return nil, err
	}
	return customer, nil
}
