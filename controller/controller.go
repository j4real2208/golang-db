package controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/j4real2208/golang-db/directory"
	"github.com/j4real2208/golang-db/error"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)


var logger *zap.Logger = error.Getlogger()
var dir  []directory.Directory = *directory.Initdirectory()
var Db *sqlx.DB = getDbClient()

func sanityCheck()  {
	if ( os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWD") == ""|| os.Getenv("DB_NAME") == "" ) {
		logger.Info("Enviornment variables not set or defined .....")		
	}
}



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

func ShowDb(w http.ResponseWriter, r *http.Request){
	var newdir  []directory.Directory
	Db.Select(&newdir,"SELECT * from person")
	for _, v := range newdir {
		fmt.Fprintf(w , "The name is :  %s and  the Aadhar Number is : %d \n",v.Name,v.Aadhar)
	}

}


func InitHandlers()  {
	
	sanityCheck()			
	r := mux.NewRouter()
	r.HandleFunc("/",homepage)
	r.HandleFunc("/showName",printName)
	r.HandleFunc("/sd",sidepage)
	r.HandleFunc("/addNew",addNew)
	r.HandleFunc("/del",DeleteNew)
	r.HandleFunc("/showDb",ShowDb)

	addr := os.Getenv("SERVER_ADDRESS")
	prt := os.Getenv("SERVER_PORT")

	
	logger.Sugar().Error(http.ListenAndServe(fmt.Sprintf("%s:%s",addr,prt), r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	
	//logger.DPanic(http.ListenAndServe("localhost:3030",r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	//logger.Error(http.ListenAndServe("localhost:3030",r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	//logger.Fatal(http.ListenAndServe("localhost:3030",r) , zap.Error(errors.New("unable to serve at localhost:3030")))
	
}

func getDbClient() *sqlx.DB{
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbName := os.Getenv("DB_NAME")
	addr := os.Getenv("SERVER_ADDRESS")
	dbPrt := os.Getenv("DB_PORT")
	
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser,dbPasswd,addr,dbPrt,dbName)	
	
	client, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 5)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}