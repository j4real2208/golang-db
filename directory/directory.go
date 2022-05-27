package directory

import (
	"github.com/cip8/autoname"
	"github.com/google/uuid"
)


type Directory struct {
	Name string
	Aadhar int64 

}


func  AddNewEntry(dirs []Directory) (*[]Directory )  {
	name := autoname.Generate()
	id := int64(uuid.New().ID())
	dirs = append(dirs,Directory{name ,id } )	
	return &dirs
}

func  DeleteEntry(dirs []Directory) (*[]Directory )  {
	
	dirs = dirs[:len(dirs)-1]	
	return &dirs
}




func Initdirectory() *[]Directory {
	entry := Directory{autoname.Generate(),int64(uuid.New().ID()) }
	dir := []Directory{entry,}
	
	for x:=0 ;x<10 ; x++ {
		dir = append(dir, Directory{autoname.Generate(),int64(uuid.New().ID())}) 
	}
	
	return &dir
}