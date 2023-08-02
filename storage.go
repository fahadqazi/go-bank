package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostGresStore() (*PostgresStore, error) {
	conStr := "user=postgres dbname=postgres password=go-bank sslmode=disable"
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account(
  id serial primary key,
  first_name varchar(50),
  last_name varchar(50),
  number serial,
  balance serial,
  created_at timestamp)`

	_, err := s.db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account
  (first_name, last_name, number, balance, created_at)
  values($1, $2, $3, $4, $5)`
	_, err := s.db.Query(query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	query := `select * from account`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {

		account := new(Account)
		err := rows.Scan(&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
