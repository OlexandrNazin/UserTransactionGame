package main

import (
	"UserTransactionGame/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/create", handler.UserCreate).Methods("POST")
	r.HandleFunc("/user/get", handler.UserDeposit).Methods("POST")
	r.HandleFunc("/user/deposit", handler.UserGet).Methods("POST")
	r.HandleFunc("/transaction", handler.Transaction).Methods("POST")
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
