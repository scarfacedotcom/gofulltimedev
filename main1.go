package main

import (
	"fmt"
	"fulltimegodev/typess"
	"fulltimegodev/utilss"
)

// type NumberStorer interface {
// 	GetAll() ([]int, error)
// 	Put(int) error
// }

// type PostgressNumberStore struct {
// }

// type MongoDBNumberStore struct {
// }

// func (s PostgressNumberStore) GetAll() ([]int, error) {
// 	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, nil
// }

// func (s PostgressNumberStore) Put(number int) error {
// 	fmt.Println("store the number into the postgress Data base database")
// 	return nil
// }

// func (m MongoDBNumberStore) GetAll() ([]int, error) {
// 	return []int{1, 2, 3}, nil
// }

// func (m MongoDBNumberStore) Put(number int) error {
// 	fmt.Println("store the number into the mongoDB Data base database")
// 	return nil
// }

// type ApiServer struct {
// 	numberStore NumberStorer
// }

func main1() {
	// apiServer := ApiServer{
	// 	numberStore: MongoDBNumberStore{},
	// }

	// apiServer := ApiServer{
	// 	numberStore: PostgressNumberStore{},
	// }

	// if err := apiServer.numberStore.Put(1); err != nil {
	// 	panic(err)
	// }

	// numbers, err := apiServer.numberStore.GetAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(numbers)

	user := typess.User{
		Username: utilss.GetName(),
		Age:      utilss.GetAge(),
	}

	fmt.Println("The users name is:", user.Username, "with an age of", user.Age)
	fmt.Printf("the user is %+v", user)
}
