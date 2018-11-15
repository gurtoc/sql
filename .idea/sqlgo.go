package main


import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	fmt.Println("Łączenie do mysql")

	db, err := sql.Open("mysql", "root:tomasson1@tcp(127.0.0.1:3308)/testdb")
	if err != nil {
		panic(err.Error() )

	}
	defer db.Close()

	fmt.Println("podlaczono")

	insert, err := db.Query("INSERT INTO users VALUES('JAN') ")

	if err != nil {
		panic(err.Error() )

	}

	defer insert.Close()

}
