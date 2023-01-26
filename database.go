package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func rowSelect(query string, user string, pass string) string {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())

	}

	rows := db.QueryRow(query, user, pass)

	var id string

	queryError := rows.Scan(&id)

	if queryError != nil {
		//log.Panic(queryError)
		fmt.Println("Nothing was returned Mysql SELECT")
		return ""
	} else {
		fmt.Println("Mysql SELECT was returned")
		return id
	}

}

func accountValidator(query string, user string, pass string) bool {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())

	}

	rows := db.QueryRow(query, user, pass)

	var username string
	var password string

	queryError := rows.Scan(&username, &password)

	if queryError != nil {
		//log.Panic(queryError)
		fmt.Println("This user and/or password found in our system")
		return false
	} else {
		fmt.Println("User: " + username + " logged in successfully")
		return true
	}

}

func querydb(query string) {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())
	}

	q, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer q.Close()
	fmt.Println("Successfully queried the database")
}

func insertUserPass(user string, pass string) {
	query := "INSERT INTO Accounts.AccountLogin (username,password) VALUES ('" + user + "','" + pass + "');"
	fmt.Println(user, pass)
	querydb(query)
}

func cleanupBracket(info []string) string {
	inf := fmt.Sprint(info)
	data := inf[1 : len(inf)-1]
	return data
}

func ValidateUserPass(user string, pass string) bool {
	query := "SELECT username, password FROM Accounts.AccountLogin WHERE username = ? AND password = ?"
	flag := accountValidator(query, user, pass)
	return flag
}

func getID(user string, pass string) string {
	query := "SELECT ID FROM Accounts.AccountLogin WHERE username = ? AND password = ?"
	selectreturn := rowSelect(query, user, pass)
	return selectreturn
}
