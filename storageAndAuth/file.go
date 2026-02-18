// package storage

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type FileStorage struct {
// 	filePath string
//     nextIDD int
// }

// func NewFileStorage(filePath string) (*FileStorage, error) {

//	    return &FileStorage{filePath: filePath,
//	                        nextIDD: 1,
//	                       }, nil
//	}
package storage

import (
	"encoding/json"
	"fmt"
	//"my-employee-api/storage"

	//"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	//"vendor/golang.org/x/net/idna"
)

type FileStorage struct {
	directoryPath string
	counter       int
}

func NewFileStorage(directoryPath string) (*FileStorage, error) {
	
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {

		os.MkdirAll(directoryPath, 0755)
	}


	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	maxID := 0
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			idStr := file.Name()[:len(file.Name())-5] 
			id, err := strconv.Atoi(idStr)
			if err == nil && id > maxID {
				maxID = id
			}
		}
	}

	return &FileStorage{
		directoryPath: directoryPath,
		counter:       maxID,
	}, nil
}
func(s *FileStorage) CreateEmployee(emp Employee) (int ,error){
    s.counter++
    emp.ID=s.counter
   jsondata, err:= json.Marshal(emp)
   if(err!=nil){
     return 0, fmt.Errorf("Not able to marshal the data %w",err)
   }
   filepath:=fmt.Sprintf("data/%d.json",emp.ID)
   errr:=os.WriteFile(filepath,jsondata,0644)
   if errr!= nil{
	return 0 , fmt.Errorf("not able to write on file %w",errr)
   }
   return emp.ID,nil
}
func (s *FileStorage) GetEmployee(id int)(Employee,error){
	 filepath:=fmt.Sprintf("data/%d.json",id)
	 data,err:=os.ReadFile(filepath)
	 if err!=nil{
		return Employee{},fmt.Errorf("Not able to read file %w",err)
	 }
	 var emo Employee
	 errr:=json.Unmarshal(data,&emo)
	 if errr!=nil{
		return Employee{},fmt.Errorf("Not able to unmarshal %w",errr)
	 }
	 return emo,nil
}
func(s* FileStorage) UpdateEmployee(id int,salary int )(error){
	filepath:=fmt.Sprintf("data/%d.json",id)
	var existingEmployee Employee
	data , err:=os.ReadFile(filepath)
	if err!=nil{
		return fmt.Errorf("not able to read the file %w",err)
	}
	 err=json.Unmarshal(data,&existingEmployee)
	 if err!=nil{
		return fmt.Errorf("not able to unmarshal %w",err)
	 }
		existingEmployee.Salary = salary
	
	updatedData,err:=json.Marshal(existingEmployee)
	if err!=nil{
		return fmt.Errorf("NOt able to marshal the data %w",err)
	}
	err= os.WriteFile(filepath,updatedData,0644)
	if err!=nil{
		return fmt.Errorf("Not able to write on the file %w",err)
	}
	return nil
}
func(s *FileStorage) DeleteEmployee(id int)error{
	filepath:=fmt.Sprintf("data/%d.json",id)
	err:=os.Remove(filepath)
	if err!=nil{
		return fmt.Errorf("not able to delete %w",err)
	}
	return nil
}

