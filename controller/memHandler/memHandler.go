package memhandler

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	dbHandler "github.com/j4real2208/golang-db/controller/dbHandler"
	"github.com/j4real2208/golang-db/directory"
	"github.com/j4real2208/golang-db/error"
)



var Db = dbHandler.Db
var mc  = dbHandler.Mc


func NewIdQueryMem(w http.ResponseWriter, r *http.Request){


	/*
	
	23:03:40.036678 +0530 IST m=+1.714738637 
	23:03:40.042969 +0530 IST m=+1.721029868
	
	Net Time Taken for db recovery of data --> .042969 - 036678 = .006291 
	
	23:03:47.30516 +0530 IST m=+8.983037492 
	23:03:47.305175 +0530 IST m=+8.983052910 --> .305175 - .30516 = 0.000015 
	
	Net Time Taken for db recovery of data --> .305175 - .30516 = 0.000015 
	
	On an apprrox 420x better in our current  evaluation 
	
	 */
	
		error.Logger.Info("We have the new memCached setup ready")
		id := mux.Vars(r)["customer_id"]		
		
		//fmt.Printf()
		val , err := mc.GetName(id)	
	
		fmt.Fprintf(w," Start time stamp %v \n",time.Now())
	
		if err == nil {
			 
			fmt.Fprintf(w , "The cust_id %d name from Memchached is :  %s and %d  And time taken is  %v \n ",val.Customer_id, val.Name ,val.Aadhar , time.Now() )
		}
	
		var newUser directory.Directory	
		Db.Get(&newUser,"SELECT * FROM person where customer_id = ?" , id)
		
		fmt.Fprintf(w , "The cust_id %d name from DB is :  %s and %d And time taken is  %v \n ",newUser.Customer_id, newUser.Name ,newUser.Aadhar , time.Now() )
	
		_ = mc.SetName(newUser)
	
	}
	