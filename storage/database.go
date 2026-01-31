package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbStore struct {
	client *mongo.Client
	collection *mongo.Collection
	nextID int
}
func NewDBStore(connectionString, dbName, collectionName string) (EmployeeStorage, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}


	collection := client.Database(dbName).Collection(collectionName)

	return &dbStore{
		client:   client,
		collection: collection,
		nextID: 1,
		}, nil
}


func (s *dbStore) CreateEmployee(emp Employee) (int, error) {
    emp.ID = s.nextID
    _, err := s.collection.InsertOne(context.TODO(), emp)
    if err != nil {
        return 0, fmt.Errorf("not able to insert: %w", err)
    }
    s.nextID++
    return emp.ID, nil 
}
func (s *dbStore) DeleteEmployee(id int)(error){
	filter:=bson.M{"_id": id}
	resp,err:=s.collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return fmt.Errorf("not able to delete %w",err)
	}
	if resp.DeletedCount == 0 {
 return fmt.Errorf("Not deleted")
}else {
	return nil
}
}
func (s*dbStore) GetEmployee(id int )( Employee,error){
	var emp Employee
	filter := bson.M{"_id": id}
	err:= s.collection.FindOne(context.TODO(),filter).Decode(&emp)
	if err!=nil{
		return Employee{},fmt.Errorf("Employee not found") 
	}
	return emp, nil
}
func (s *dbStore) UpdateEmployee(id int, salary int) error {
    filter := bson.M{"_id": id}

  update := bson.D{
	{Key: "$set", Value: bson.D{
		{Key: "salary", Value: salary},
	}},
}

    res, err := s.collection.UpdateOne(context.Background(), filter, update)

    if err != nil {
        return fmt.Errorf("failed to update employee: %w", err)
    }
    if res.MatchedCount == 0 {
        return fmt.Errorf("Not able to Update")
    }
    return nil
}
func(s *dbStore) GetEmployeeByUsername(username string) (Employee,error){
	filter:=bson.M{"username":username}
	var emp Employee
	err := s.collection.FindOne(context.TODO(), filter).Decode(&emp) 
	if err!=nil{
		return Employee{}, fmt.Errorf("user not found")
	}
	return emp,nil
}
