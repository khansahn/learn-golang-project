package main

import (
	"fmt"
	"project/stringutil"

	"database/sql"
	_ "github.com/lib/pq"

//	"log"
//	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "inksnippet"
	dbname   = "inksnippet"
)

type CodeSnippet struct {
	ID          int
	Title       string
	Codelines   string
	Tag         string
	Author      string
	ContentType string
	CreatedAt   string
	UpdatedAt   string
}

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
	/*
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
	//fmt.Println(id)

	// GET
	/*
	    sqlStatementGet := `
	    SELECT id, title
	    FROM codesnippet
	    WHERE id=$1;
	    `
	    var id int
	    var title string
	    row := db.QueryRow(sqlStatementGet, 3)
	    switch err = row.Scan(&id, &title); err {
	    case sql.ErrNoRows:
		    fmt.Println("no rows returned??>?>?")
	    case nil:
		    fmt.Println(id, title)
	    default:
		    panic(err)
	    }
	*/

	/*
	    // multiple GET
	    sqlStatementGet2 := `
	    SELECT * FROM codesnippet WHERE id=$1
	    `
	    var codesnippet CodeSnippet
	    row2 := db.QueryRow(sqlStatementGet2, 1)
	    err = row2.Scan(&codesnippet.ID, &codesnippet.Title, &codesnippet.Codelines, &codesnippet.Tag, &codesnippet.Author, &codesnippet.ContentType, &codesnippet.CreatedAt, &codesnippet.UpdatedAt)
	    switch err {
	    case sql.ErrNoRows:
		    fmt.Println("gggada")
		    return
	    case nil:
		    fmt.Println(codesnippet)
	    default:
		    panic(err)
	    }
	*/
	// MULTIPLE GET
	rows, err := db.Query("SELECT id, title, tag, created_at FROM codesnippet")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		res := CodeSnippet{}
		var id int
		var title string
		var tag string
		var createdAt string
		err = rows.Scan(&id, &title, &tag, &createdAt)
		if err != nil {
			panic(err)
		}
		res.ID = id
		fmt.Println(id, title, tag, createdAt)
		fmt.Println("HAAIII", res.ID)
	}
	// get error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	/*
	   // DELETE
	   sqlStatementDelete := `
	   DELETE FROM codesnippet
	   WHERE id=$1;
	   `
	   _, err = db.Exec(sqlStatementDelete, 4)
	   if err != nil {
	   	panic(err)
	   }
	*/
}
