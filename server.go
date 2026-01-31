package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	// "strings"
	"my-employee-api/storage"
)
func Authmiddleware(next http.HandlerFunc) http.HandlerFunc{
    return func(w http.ResponseWriter,r *http.Request){
        auth:=r.Header.Get("Authorization")
        if auth==""{
           http.Error(w,"Invalid Authorization",http.StatusUnauthorized)
           return
        }
        token:=strings.TrimPrefix(auth,"Bearer ")
        username,err:=storage.ValidateToken(token)
        if err!=nil{
            http.Error(w,"No Authorization",http.StatusUnauthorized)
            return
        }
        r.Header.Set("X-Username", username)
        next(w,r)
    }
}

type Server struct {
    store storage.EmployeeStorage
}

func NewServer(s storage.EmployeeStorage) *Server {
    return &Server{store: s}
}
func(s *Server) Login(w http.ResponseWriter,r *http.Request){
    if r.Method!=http.MethodPost{
        http.Error(w,"Method Not allowed",http.StatusMethodNotAllowed)
        return 
    }
   type loginreq struct{
    Username string `json:"username"`
    Password string `json:"password"`
}
var input loginreq
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
    http.Error(w,"JSON decode error: "+err.Error(),http.StatusBadRequest)
    return 
}
    if (input.Username=="" ||input.Password==""){
        http.Error(w,"Input invalid",http.StatusBadRequest)
        return
    }
     employee, err := s.store.GetEmployeeByUsername(input.Username)
    if err != nil {
        http.Error(w,"Invalid credentials",http.StatusUnauthorized)
        return 
    }
    if input.Password != employee.Password {  
        http.Error(w,"Invalid credentials",http.StatusUnauthorized)
        return 
    } 
    storage.Loadsecret()
    token,eror:=storage.GenerateToken(input.Username)
    if eror!=nil{
        http.Error(w,"Not able to generate token ",http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (s *Server) Start() {
    http.HandleFunc("/employee/login",s.Login)
	http.HandleFunc("/employee/create", s.CreateEmployeehandler)
	http.HandleFunc("/employee/get", Authmiddleware(s.GetEmployeehandler))
	http.HandleFunc("/employee/update", Authmiddleware(s.UpdateEmployeehandler))
	http.HandleFunc("/employee/delete", Authmiddleware(s.DeleteEmployeehandler))

	fmt.Println("Server is running on port 8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
func (s *Server) CreateEmployeehandler(w http.ResponseWriter, r *http.Request) {
    var emp storage.Employee
    err := json.NewDecoder(r.Body).Decode(&emp)
    if err != nil {
        http.Error(w, "Not able to read request body", http.StatusBadRequest) 
        return
    }
    id, err := s.store.CreateEmployee(emp)
    if err != nil {
        http.Error(w, "Not able to create employee", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(fmt.Sprintf("Employee created with ID %d", id)))
}

func (s *Server) GetEmployeehandler(w http.ResponseWriter, r *http.Request ) {
     idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID parameter is missing", http.StatusBadRequest)
        return
    }
    id,err:=strconv.Atoi(idStr)
    if err!=nil {
   http.Error(w, "Not able to converet", http.StatusNotFound)
        return
    }
    employee, err := s.store.GetEmployee(id)
    if err != nil {
        http.Error(w, "Employee not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(employee)
}

func (s *Server) UpdateEmployeehandler(w http.ResponseWriter, r *http.Request) {
      idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID parameter is missing", http.StatusBadRequest)
        return
    }
    id,err:=strconv.Atoi(idStr)
    if err!=nil {
   http.Error(w, "Not able to converet", http.StatusNotFound)
        return
    }
    var temp struct {
        Salary int `json:"salary"`
    }
    err = json.NewDecoder(r.Body).Decode(&temp)
    if err != nil {
        http.Error(w, "Not able to decode request", http.StatusBadRequest)
        return
    }
    err = s.store.UpdateEmployee(id,temp.Salary)
    if err != nil {
        http.Error(w, "Not able to update employee", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Employee with ID %d has been updated", id)))
}

func (s *Server) DeleteEmployeehandler(w http.ResponseWriter, r *http.Request) {
      idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID parameter is missing", http.StatusBadRequest)
        return
    }
    id,err:=strconv.Atoi(idStr)
    if err!=nil {
   http.Error(w, "Not able to converet", http.StatusNotFound)
        return
    }
    err = s.store.DeleteEmployee(id)
    if err != nil {
        http.Error(w, "Not able to delete employee", http.StatusNotFound)
        return
    }
    w.Write([]byte(fmt.Sprintf("Employee with ID %d has been deleted", id)))
}