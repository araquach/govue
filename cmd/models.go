package main

type User struct {
	ID        	uint       	`json:"id" gorm:"primaryKey"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email" gorm:"uniqueIndex"`
	Password 	string 		`json:"password"`
	IsAdmin		bool 		`json:"is_admin"`
	Token		string 		`json:"token" gorm:"-"`
}

type Test struct {
	Id			uint 	`json:"id" gorm:"primaryKey"`
	FirstName 	string 	`json:"first_name"`
	LastName	string 	`json:"last_name"`
}

