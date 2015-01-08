package main

import (
	r "github.com/dancannon/gorethink"
	"log"
	"os"
	"time"
)

// InitDb setups the Database for the applications first use.
func InitDb() *r.Session {
	dbName := os.Getenv("DB_NAME")

	session, err := r.Connect(r.ConnectOpts{
		Address:     os.Getenv("RETHINKDB_URL"),
		Database:    dbName,
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	})
	if err != nil {
		log.Println(err)
	}

	err = r.DbCreate(dbName).Exec(session)
	if err != nil {
		log.Println(err)
	}

	_, err = r.Db(dbName).TableCreate("users").RunWrite(session)
	if err != nil {
		log.Println(err)
	}

	_, err = r.Db(dbName).TableCreate("sessions").RunWrite(session)
	if err != nil {
		log.Println(err)
	}

	var baseUser User
	_, err = r.Db(dbName).Table("users").Insert(User{Username: "sean@engageyourcause.com", Password: baseUser.HashPassword("powell5112")}).RunWrite(session)
	if err != nil {
		log.Println(err)
	}

	return session
}

func GetConn() *r.Session {
	dbName := os.Getenv("DB_NAME")

	session, err := r.Connect(r.ConnectOpts{
		Address:     os.Getenv("RETHINKDB_URL"),
		Database:    dbName,
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	})
	if err != nil {
		log.Println(err)
	}

	return session
}
