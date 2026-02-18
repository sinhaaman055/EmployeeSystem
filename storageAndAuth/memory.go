package storage

import "fmt"

type MemoryStorage struct {
	employee map[int]Employee
	counter  int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		employee: make(map[int]Employee),
		counter:  0,
	}
}
func (s *MemoryStorage) GetEmployee(id int) (Employee, error) {
	emp, ok := s.employee[id]
	if !ok  {
		return Employee{}, fmt.Errorf("Not able to find the employee")
	}
	return emp,nil
}
func(s* MemoryStorage) CreateEmployee(emp Employee)(int , error){
	s.counter++
	emp.ID=s.counter
	s.employee[s.counter]=emp
    return emp.ID,nil
}
func(s* MemoryStorage) DeleteEmployee(id int) error{
	_,ok:=s.employee[id] 
	if !ok{
		return  fmt.Errorf("ID %d not found",id)
	}
	delete(s.employee,id)
	return nil
}
func(s* MemoryStorage)UpdateEmployee(id int, salary int )error{
	existemp,ok:=s.employee[id]
	if !ok{
		return fmt.Errorf("Not able to find employee")
	}

	existemp.Salary = salary
	
	s.employee[id]=existemp
	return nil
}