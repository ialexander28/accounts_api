package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Account struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Balance   string `json:"Balance"`
	Direction string `json:"Direction"`
}

var Accounts []Account

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Account page!")
}

func getAllAccounts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Accounts)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, account := range Accounts {
		if account.Id == key {
			json.NewEncoder(w).Encode(account)
		}
	}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	account.Id = vars["id"]
	Accounts = append(Accounts, account)
	json.NewEncoder(w).Encode(account)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, account := range Accounts {
		if account.Id == vars[id] {
			Accounts = append(Accounts[:index], Accounts[index+1:]...)
		}
	}
}

//curl POST script as follows:  curl -d '{"Id":"five","Name":"BITCH","Direction":"DEBIT"}' -X POST http://localhost:3000/account/
func apiRequests() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", homePage)
	route.HandleFunc("/account", getAllAccounts)
	route.HandleFunc("/account/{id}", createAccount).Methods("POST")
	route.HandleFunc("/account/{id}", deleteAccount).Methods("DELETE")
	route.HandleFunc("/account/{id}", getAccount)
	log.Fatal(http.ListenAndServe(":3000", route))
}

func main() {
	Accounts = []Account{
		//hash value of account name: 5256529315534043551. Use for corresponding Hex value.

		{Id: "0x48f2f18febd5799f",
			Name:      "bridge",
			Direction: "debit",
			Balance:   "0"},
		{Id: "2",
			Name:      "wedge",
			Direction: "credit",
			Balance:   "0"},
		{Id: "3",
			Name:      "swift",
			Direction: "debit",
			Balance:   "100"},
		{Id: "4",
			Name:      "defi_wallet",
			Direction: "credit",
			Balance:   "100"},
	}
	apiRequests()
}
