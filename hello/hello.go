package main

import (
	"fmt"
	"project/stringutil"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "inksnippet"
	dbname = "inksnippet"
)


func main() {
    fmt.Printf("hello, world\n")
    fmt.Println(stringutil.Reverse("Hello, Go!"))

    // CONNECTING TO DB
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
    	panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
    	panic(err)
    }

    fmt.Println("connectttted")


    // CREATE
    /*
    sqlStatement := `
    INSERT INTO codesnippet (title, codelines, tag, author, content_type)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
    `
    id := 0
    err = db.QueryRow(sqlStatement, "tes hc go", "package main ok", "test", "ink", "go").Scan(&id)
    if err != nil {
    	panic(err)
    }

    fmt.Println("New record id is: ", id)
*/

    // UPDATE
    sqlStatementUpdate := `
    UPDATE codesnippet
    SET title=$2, codelines=$3, tag=$4, author=$5, content_type=$6, updated_at=$7
    WHERE id=$1
    RETURNING id;
    `
    var id int
    err = db.QueryRow(sqlStatementUpdate, 1, "tes hc go update", "haha", "tes", "ink", "go", "2019-05-20 15:55:40.222222").Scan(&id)
    if err != nil {
    	panic(err)
    }
    /*
    count, err := res.RowsAffected()
    if err != nil {
    	panic(err)
    }
    fmt.Println(count)
    */
    fmt.Println(id)
}
