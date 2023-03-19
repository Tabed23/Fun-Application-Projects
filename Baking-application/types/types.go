package types

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//Account Structure
type Account struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Number int64 `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance int64 `json:"balance" min:"0"`
	CreatedAt         time.Time `json:"createdAt"`
}


//New Account
func NewAccount(firstName string, lastName string, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
        FirstName: firstName,
        LastName:  lastName,
		Number: int64(rand.Intn(1000000)),
		Balance: 0,
		EncryptedPassword: string(encpw),
		CreatedAt: time.Now().UTC(),
	}, nil
}
//New Account
func UpdateAccount(id int, firstName , lastName , password string, balance , number int64) *Account {
	log.Println("Update Account", "Types", "Returning Account")
	return &Account{
		ID: id,
        FirstName: firstName,
        LastName:  lastName,
		Balance: balance,
		Number: number,
		EncryptedPassword: password,
	}
}




type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type UpdateAccountRequest struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
}

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}
