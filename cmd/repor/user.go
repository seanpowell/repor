package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
	"fmt"
	rdb "github.com/dancannon/gorethink"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type User struct {
	Username  string
	Password  []byte
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// HashPassword takes a plaintext password and hashes/salts if with bcrypt and assigns it to the user's password field
func (u *User) HashPassword(password string) []byte {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashPass
}

// LogInUser takes the http response, request and other parameters and matched against the database. If a user is found a session is created.
func LogInUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) string {
	rdbSession := GetConn()
	fmt.Println("User: ", p.ByName("username"))
	un := p.ByName("username")
	u := User{}
	res, err := rdb.Table("users").GetAllByIndex("Username", un).Run(rdbSession)
	if err != nil {
		log.Println(err)
	}
	res.One(&u)
	fmt.Println("RDB Res: ", u)
	bytePass := []byte(p.ByName("password"))
	if bcrypt.CompareHashAndPassword(u.Password, bytePass) != nil {
		errors.New("Username or password is incorrect.")
		fmt.Println("Username or password is incorrect.")
		return ""
	}
	fmt.Println("All good.")
	return u.Username
}

// CreatUser creates a new user in the database
// func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	db, err := gorm.Open("postgres", "postgres://rzkxuetyfxhmao:UgKNrGtfe1SnM8sUHDZ5hVkbBL@ec2-54-235-99-22.compute-1.amazonaws.com:5432/df0oj5tcrc8erb")
// 	if err != nil {
// 		errors.New("Could not connect to the DB")
// 	}
// 	// Future: build out invite system
//
// 	u := User{
// 		Username:  p.ByName("username"),
// 		Password:  HashPassword(p.ByName("password")),
// 		Role:      "admin",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	db.Create(&u)
// }
