package directory

import (
	"github.com/cip8/autoname"
	"github.com/google/uuid"
)


type Directory struct {
	Customer_id int64 	`db:"customer_id"`
	Name string 	`db:"name"`
	Aadhar int64 	`db:"aadhar_id"`


}


func  AddNewEntry(dirs []Directory) (*[]Directory )  {
	name := autoname.Generate()
	id := int64(uuid.New().ID())
	dirs = append(dirs,Directory{id+2,name ,id  } )	
	return &dirs
}

func  DeleteEntry(dirs []Directory) (*[]Directory )  {
	
	dirs = dirs[:len(dirs)-1]	
	return &dirs
}




func Initdirectory() *[]Directory {
	entry := Directory{int64(uuid.New().ID())+3 , autoname.Generate(),int64(uuid.New().ID()) }
	dir := []Directory{entry,}
	
	for x:=0 ;x<10 ; x++ {
		dir = append(dir, Directory{int64(uuid.New().ID())+11 ,autoname.Generate(),int64(uuid.New().ID())}) 
	}
	
	return &dir
}