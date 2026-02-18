package storage

type Employee struct {
    ID       int    `bson:"_id" json:"id"`
    Name     string `bson:"name" json:"name"`
    Age      int    `bson:"age" json:"age"`
    Salary   int    `bson:"salary" json:"salary"`
    Username string `bson:"username" json:"username"`
    Password string `bson:"password" json:"password"`  
}

type EmployeeStorage interface { 
    CreateEmployee(emp Employee) (int,error)     
    GetEmployee(id int) (Employee, error)   
    UpdateEmployee(id int, salary int) error     
    DeleteEmployee(id int) error  
    GetEmployeeByUsername(username string) (Employee, error)
}