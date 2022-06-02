package dbhandler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cip8/autoname"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/j4real2208/golang-db/directory"
	"github.com/j4real2208/golang-db/memcached"

	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB = getDbClient()
var mc, newError  = memcached.IntializeMEM()





func ShowDb(w http.ResponseWriter, r *http.Request){
	var newdir  []directory.Directory
	Db.Select(&newdir,"SELECT * from person")
	for _, v := range newdir {
		fmt.Fprintf(w , "The cust id is %d name is :  %s and  the Aadhar Number is : %d \n",v.Customer_id,v.Name,v.Aadhar)
	}
}
func AddDb(w http.ResponseWriter, r *http.Request  ){
	tx , _  := Db.Begin()
	for i := 0; i < 5; i++ {
		name:= autoname.Generate()
		name = strings.ReplaceAll(name,"_","")
		tx.Exec(`INSERT INTO person (name, aadhar_id) VALUES (?, ?)`, name, int64(uuid.New().ID()))		
	}		
	tx.Commit()	
	// var newdir directory.Directory;
    // err := Db.Get(&newdir,"SELECT * FROM person WHERE name = ? ", "Jason")
	// _ = err
	// fmt.Fprintf(w , "The name is :  %s and  the Aadhar Number is : %d \n",newdir.Name,newdir.Aadhar)
	//logger.Info(err.Error())
}

func UserGetByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id:= vars["aadhar_id"]
	var newUser directory.Directory
	Db.Get(&newUser,"SELECT * FROM person where aadhar_id = ?" , id)
	fmt.Fprintf(w , "The name is :  %s",newUser.Name)
}

func UserGetByName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id:= vars["name"]
	var newUser directory.Directory	
	
	Db.Get(&newUser,"SELECT * FROM person where name = ?" , id)	

	fmt.Fprintf(w , "The name is :  %d",newUser.Aadhar)
}

func AddNewUserWeb(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name:= vars["name"]
	id:= vars["aadhar_id"]

	tx , _  := Db.Begin()
	tx.Exec(`INSERT INTO person (name, aadhar_id) VALUES (?, ?)`, name, id)		
	tx.Commit()	

	fmt.Fprintf(w , "The name Added is :  %s and %s ",name ,id )
}


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



/*

23:03:40.036678 +0530 IST m=+1.714738637 
23:03:40.042969 +0530 IST m=+1.721029868

23:03:47.30516 +0530 IST m=+8.983037492 
23:03:47.305175 +0530 IST m=+8.983052910 


 */