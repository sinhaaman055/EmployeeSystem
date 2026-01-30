package storage

type Employee struct {
    ID   int  `bson:"_id"`
    Name string  `bson:"name"`
    Age  int     `bson:"age"`
	Salary int   `bson:"salary"`
}

type EmployeeStorage interface { 
    CreateEmployee(emp Employee) (int,error)     
    GetEmployee(id int) (Employee, error)   
    UpdateEmployee(id int, salary int) error     
    DeleteEmployee(id int) error         
}