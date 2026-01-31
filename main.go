// package main

// import (
// //	"fmt"
// 	"log"

// 	"my-employee-api/storage"
// )

//	func main() {
//		dbstore, err := storage.NewDBStore("mongodb://localhost:27017","TeamGoal","books")
//		if err!=nil{
//			log.Fatalf("Not able to connect with mongodb %v",err)
//		}
//		server:=NewServer(dbstore)
//		server.Start()
//	}
package main

import (
	//"log"
	"log"
	"my-employee-api/storage"
)

//"my-employee-api/storage"

func main() {
	storage.Loadsecret()
	dbStore, err := storage.NewDBStore("mongodb://localhost:27017","TeamGoal","books")
	if err != nil {
		log.Fatalf("Not able to connect with mongodb: %v", err)
	}


	// filestore,err:=storage.NewFileStorage("data")
	// 	if err != nil {
	// 	log.Fatalf("could not create file storage: %v", err)
	// }

	// memorystore:=storage.NewMemoryStorage()



	server := NewServer(dbStore)
	server.Start()
}