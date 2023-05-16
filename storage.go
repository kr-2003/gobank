package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByNumber(int64) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USER := os.Getenv("DB_USER")
	connStr := fmt.Sprintf("host=database port=5432 user=%s dbname=%s password=%s sslmode=disable", DB_USER, DB_NAME, DB_PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query :=
		`
	create table if not exists account (
		id serial primary key,
		first_name varchar(50), 
		last_name varchar(50),
		number serial,
		balance int,
		created_at timestamp
	)

	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query :=
		`
		insert into account (first_name, last_name, number, balance, created_at)
		values
		($1, $2, $3, $4, $5)
		`
	resp, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	query :=
		`
	delete from account where id = $1
	`
	resp, err := s.db.Query(query, id)
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("account not found")
	}
	return nil
}

func (s *PostgresStore) GetAccountByNumber(number int64) (*Account, error) {
	query :=
		`
	select * from account where id = $1
	`

	rows, err := s.db.Query(query, number)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", rows)
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	if len(accounts) != 1 {
		return nil, fmt.Errorf("Account %d not found", number)
	}
	return accounts[0], nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query :=
		`
	select * from account where id = $1
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", rows)
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	if len(accounts) != 1 {
		return nil, fmt.Errorf("Account %d not found", id)
	}
	return accounts[0], nil
}
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	query := "select * from account"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}
