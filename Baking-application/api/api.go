package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Tabed23/Fun-Application-Projects/tree/main/Baking-application/repository"
	"github.com/Tabed23/Fun-Application-Projects/tree/main/Baking-application/types"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)


//API Server Struct
type APIServer struct {
	listenAddr string
	store      repository.Store
}

//NewAPIServer creates a new APIServer
func NewAPIServer(listenAddr string, storage repository.Store) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      storage,
	}
}
//Run the APIServer
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("health check ", r.URL.String())
		WriteJson(w, http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})
	router.HandleFunc("/account", makeHttpHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandler(s.handleGetAccountById))

	logger.Info("API Server is listening on %s ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}


//Handle Account Routes
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "PATCH" {
		return s.handleUpdateAccount(w, r)
	}

	return fmt.Errorf("method not supported %s", r.Method)
}

//Handler Get Accounts
func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	logger.Info("Get Account ", "Handler")
	accounts, err := s.store.GetAccounts()
	if err != nil {
		logger.Error("API-Server", "Get Account", err.Error())
		return err
	}

	return WriteJson(w, http.StatusOK, accounts)
}


//HandlePostAccount
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	logger.Info("Post Account ", "Handler")
	ac := types.CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		logger.Error("error creating account", err)
		return err

	}

	account, err := types.NewAccount(ac.FirstName, ac.LastName, ac.Password)
	if err != nil {
		logger.Error("error creating account service layer", err)
		return err
	}

	if err := s.store.CreateAccount(account); err != nil {
		logger.Error("error creating account store layer", err)
		return nil
	}

	logger.Info("Post Account ", "created account", account)

	return WriteJson(w, http.StatusOK, account)
}


//handleGetAccountById
func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getID(r)
		if err != nil {
			logger.Error("error getting account Id", err)
			return err
		}

		account, err := s.store.GetAccountByID(id)
		if err != nil {
			logger.Error("error getting account", err)
            return err
		}

		return WriteJson(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}


	return fmt.Errorf("method not allowed %s", r.Method)

}

//HandleDeleteAccount
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	logger.Info("Delete Account ", "Handler")
	id, err := getID(r)
	if err != nil {
		logger.Error("DeleteAccount", "error getting  id: %v", err)
		return err
	}

	if err := s.store.DeleteAccount(id); err != nil {
		logger.Error("Delete Account ", "error deleting account: %v", err)
		return err
	}

	return WriteJson(w, http.StatusOK, map[string]int{"deleted": id})
}


//Handler Update Account
func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	logger.Info("Update Account", "Handler")

	ac := types.UpdateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		logger.Error("error Updating account", err)
		return err

	}

	logger.Info(ac)

	upac := types.UpdateAccount(int(ac.ID),ac.FirstName, ac.LastName, ac.Password, ac.Balance,ac.Number)

	logger.Info(*upac)

	if err := s.store.UpdateAccount(upac); err != nil {
		logger.Error("error Updating account", err)
        return err
	}
	return WriteJson(w, http.StatusOK,  map[string]any{"is_update":true, "data": upac })
}

func (s *APIServer) handleTransferPayment(w http.ResponseWriter, r *http.Request) error {
	return nil
}

//Write Json Response
func WriteJson(w http.ResponseWriter, status int, v any) error {

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}


//Making a Type of Handler
type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)

		if err != nil {

			WriteJson(w, http.StatusInternalServerError, err.Error())
		}
	}
}

//Get ID Func
func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}
