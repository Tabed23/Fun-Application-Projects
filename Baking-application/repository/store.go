package repository

import (
	"database/sql"
	"fmt"

	"github.com/Tabed23/Fun-Application-Projects/tree/main/Baking-application/types"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)


//Database Operations
type Store interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccountByID(int) (*types.Account, error)
	GetAccounts() ([]*types.Account, error)
	GetAccountByNumber(int) (*types.Account, error)
}


var (
	logger = logrus.New()
)

type PostgresStore struct {
	db  *sql.DB
}

//NewPostgresStore - Connection to Postgres
func NewPostgresStore() (*PostgresStore, error) {

	connStr := "user=postgres dbname=postgres password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	logger.Println("database connection established")

	return &PostgresStore{
		db: db,
	}, nil
}

//Initialize the database
func (s *PostgresStore) Init() error {
	logger.Info("init", "create table postgres")
	return s.createAccountTable()
}

//Create a Table to store
func (s *PostgresStore) createAccountTable() error {
	logger.Info("create table postgres", "account table")

	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		number serial,
		encrypted_password varchar(100),
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	if err!= nil {
		logger.Error("error for db exec", err)

        return err
    }

	return nil
}

//Create a Account to Store
func (p *PostgresStore) CreateAccount(acc *types.Account) error {
	logger.Info("Store", "Create Account Function")
	query := `insert into account
	(first_name, last_name, number, encrypted_password, balance, created_at)
	values ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		logger.Error("error creating account", "Create Account Function", "Store", err)
		return err
	}

	return nil
}

//Delete a Account from Store
func (p *PostgresStore) DeleteAccount(id int) error {
	logger.Info("Store", "Delete Account Function")
	_, err := p.db.Query("delete from account where id = $1", id)
	if err != nil {
		logger.Error("error deleting account", "Delete Account Function", "Store", err)
		return err
	}
	return nil
}

//Update a Account in Store
func (p *PostgresStore) UpdateAccount(acc *types.Account) error {
	logger.Info("Store", "Update Account Function")

	_, err := p.GetAccountByID(acc.ID)
	if err != nil {
		logger.Error("error updating account", "cannot get the account", err)
		return  err
	}
	stmt, err := p.db.Prepare("UPDATE account SET first_name=$1, last_name=$2, number=$3, encrypted_password=$4, balance=$5  WHERE id= $6")
	if err != nil {
		logger.Error("Updated account", "cannot update the account, query invalid ", err)
		return err
	}

	_, err = stmt.Exec(acc.FirstName, acc.LastName,acc.Number, acc.EncryptedPassword, acc.Balance, acc.ID)
	if err != nil{

		logger.Error("Updated account", "cannot update the account", err)
		return err

	}
	logger.Info("Updated account", "update the account", acc)

	return nil
}

//Get Account By Number Store
func (p *PostgresStore) GetAccountByNumber(num int) (*types.Account, error) {
	logger.Info("Store", "Get By Number Account Function")
	rows, err := p.db.Query("select * from account where number = $1", num)
	if err != nil {
		logger.Info("error getting account", "Get Account By Number Function", "Store", err)
		return nil, err
	}

	ac := new(types.Account)
	for rows.Next() {
		ac, err = MapToStruct(rows)
		if err != nil {
			logger.Error("error getting account", "Get Account  By Number Function", "Store", err)
			return nil, err
		}

	}

	if ac != nil {
		return ac, nil
	} else {
		return nil, fmt.Errorf("account %d not found", num)

	}
}

//Get Account By ID to Store
func (p *PostgresStore) GetAccountByID(id int) (*types.Account, error) {
	logger.Info("Store", "Get By ID Account Function")
	rows, err := p.db.Query("select * from account where id = $1", id)
	if err != nil {
		logger.Error("error getting account", "Get Account By ID Function", "Store", err)
		return nil, err
	}

	ac := new(types.Account)
	for rows.Next() {
		ac, err = MapToStruct(rows)
		if err != nil {
			logger.Error("error getting account", "Get Account By ID Function", "Store", err)
			return nil, err
		}

	}

	if ac != nil {
		return ac, nil
	} else {
		return nil, fmt.Errorf("account %d not found", id)

	}
}

//Get Accounts from Store
func (p *PostgresStore) GetAccounts() ([]*types.Account, error) {
	logger.Info("Store", "Get All Account Function")
	rows, err := p.db.Query("select * from account")
	if err != nil {
		logger.Error("Error getting accounts", "Get Accounts Function", "Store", err)
		return nil, err
	}

	accounts := make([]*types.Account, 0)

	for rows.Next() {
		account, err := MapToStruct(rows)
		if err != nil {
			logger.Error("cannot get account", "Get Accounts Function", "Store")
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}


//Convert Database Value to Account Value
func MapToStruct(rows *sql.Rows) (*types.Account, error) {
	acc := &types.Account{}
	err := rows.Scan(
		&acc.ID,
		&acc.FirstName,
		&acc.LastName,
		&acc.Number,
		&acc.EncryptedPassword,
		&acc.Balance,
		&acc.CreatedAt,
	)

	if err != nil {
		logger.Fatal("Error getting account", "MapToStruct", err)
		return nil, err
	}

	return acc, nil
}
