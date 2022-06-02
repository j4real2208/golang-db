package controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	dbHandler "github.com/j4real2208/golang-db/controller/dbHandler"
	oldhandler "github.com/j4real2208/golang-db/controller/oldHandler"
	"github.com/j4real2208/golang-db/error"

	"go.uber.org/zap"
)


var logger *zap.Logger = error.Getlogger()


func sanityCheck()  {
	if ( os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWD") == ""|| os.Getenv("DB_NAME") == "" ) {
		logger.Info("Enviornment variables not set or defined .....")		
	}
}

func InitHandlers()  {
	
	sanityCheck()			
	r := mux.NewRouter()


	// Old Handler Static 
	r.HandleFunc("/",oldhandler.Homepage)
	r.HandleFunc("/showName",oldhandler.PrintName)
	r.HandleFunc("/sd",oldhandler.Sidepage)
	r.HandleFunc("/addNew",oldhandler.AddNew)
	r.HandleFunc("/del",oldhandler.DeleteNew)


	// Db based controller
	r.HandleFunc("/showDb",dbHandler.ShowDb)
	r.HandleFunc("/AddDb",dbHandler.AddDb)
	r.HandleFunc("/usr/{aadhar_id:[0-9]+}", dbHandler.UserGetByID).Methods(http.MethodGet)
	r.HandleFunc("/usr/{name:[a-zA-Z]+}", dbHandler.UserGetByName).Methods(http.MethodGet)
	r.HandleFunc("/usr/{name:[a-zA-Z]+}/{aadhar_id:[0-9]+}", dbHandler.AddNewUserWeb).Methods(http.MethodGet)
	r.HandleFunc("/cust/{customer_id}",dbHandler.NewIdQueryMem).Methods(http.MethodGet)	
	
	addr := os.Getenv("SERVER_ADDRESS")
	prt := os.Getenv("SERVER_PORT")

	
	logger.Sugar().Error(http.ListenAndServe(fmt.Sprintf("%s:%s",addr,prt), r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	
	
}

