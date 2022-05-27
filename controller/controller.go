package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j4real2208/golang-db/directory"
	"github.com/j4real2208/golang-db/error"
	"go.uber.org/zap"
)


var logger *zap.Logger = error.Getlogger()
var dir  []directory.Directory = *directory.Initdirectory()


func homepage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w , "Welcome to our Dev Page ")	
	logger.Info("You have hit the Dev page endpoint ")
}

func sidepage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w , "Welcome to our Side Page ")	
	logger.Info("You have hit the Side page endpoint ")
}


func printName(w http.ResponseWriter, r *http.Request )  {

	fmt.Fprintf(w , "Welcome to our Print Name page and len is %d  " , len(dir))	
	var displayString string = ""
	for _, v := range dir {

		displayString += "The name is : " + v.Name + "and  the Aadhar Number is " + strconv.Itoa(int(v.Aadhar)) +"\n"

	}
	fmt.Fprintf(w , "The list of Names are \n  %s  " , displayString)	
	logger.Info("You have hit the Print page endpoint ")
}
func addNew(w http.ResponseWriter, r *http.Request)  {
	dir   = *directory.AddNewEntry(dir)
	logger.Info("You have hit the Add New name page endpoint ")
	x := len(dir)
	fmt.Fprintf(w , "The New added Names is : %s and Aaadhar is : %d \n " , dir[x-1].Name , dir[x-1].Aadhar)
}

func DeleteNew(w http.ResponseWriter, r *http.Request)  {
	dir   = *directory.DeleteEntry(dir)
	logger.Info("You have hit the Delete name page endpoint ")	
	fmt.Fprintf(w , "The Last Name was deleted and number of entries are is : %d \n " , len(dir))
}


func InitHandlers()  {
	r := mux.NewRouter()
	r.HandleFunc("/",homepage)
	r.HandleFunc("/showName",printName)
	r.HandleFunc("/sd",sidepage)
	r.HandleFunc("/addNew",addNew)
	r.HandleFunc("/del",DeleteNew)
	logger.Sugar().Error(http.ListenAndServe("localhost:3030",r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	
	//logger.DPanic(http.ListenAndServe("localhost:3030",r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	//logger.Error(http.ListenAndServe("localhost:3030",r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	//logger.Fatal(http.ListenAndServe("localhost:3030",r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	
}