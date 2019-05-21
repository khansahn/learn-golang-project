package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"net/http"
	"project/structs"
	//"encoding/json"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "inksnippet"
	dbname   = "inksnippet"
)

var db *sql.DB

// dbinit
func dbinit() {
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
	//return c.String(http.StatusOK, "connect ke database yaa")
}

// handler hello
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "HIII")
}



// QUERY GET ALL
func queryGetAll(result *structs.CodeSnippetArray) (res structs.CodeSnippetArray, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	sqlStatement := `
	    SELECT id, title, codelines, tag, author, content_type, created_at
	    FROM codesnippet;
	`
	codesnippetarray := structs.CodeSnippetArray{}
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		codesnippet := structs.CodeSnippet{}
		err = rows.Scan(&codesnippet.ID, &codesnippet.Title, &codesnippet.Codelines, &codesnippet.Tag, &codesnippet.Author, &codesnippet.ContentType, &codesnippet.CreatedAt)
		if err != nil {
			panic(err)
		}
		codesnippetarray.CodeSnippetArray = append(codesnippetarray.CodeSnippetArray, codesnippet)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return codesnippetarray, nil
}

// HANDLER GET ALL
func getAll(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	result := structs.CodeSnippetArray{}

	res, err := queryGetAll(&result)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, res)
}



// QUERY GET PER ID
func queryGetPerId(result *structs.CodeSnippetArray, id int) (res structs.CodeSnippet, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	sqlStatement := `
	    SELECT id, title, codelines, tag, author, content_type, created_at
	    FROM codesnippet
	    WHERE id=$1;
	`
	codesnippet := structs.CodeSnippet{}
	row, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&codesnippet.ID, &codesnippet.Title, &codesnippet.Codelines, &codesnippet.Tag, &codesnippet.Author, &codesnippet.ContentType, &codesnippet.CreatedAt)
		if err != nil {
			panic(err)
		}
	}

	return codesnippet, nil
}

// HANDLER GET PER ID
func getPerId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := structs.CodeSnippetArray{}

	res, err := queryGetPerId(&result, id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, res)
}

// QUERY CREATE
func queryCreate(newcodesnippet *structs.CodeSnippet) (codesnippet structs.CodeSnippet) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	sqlStatement := `
	    INSERT INTO codesnippet (title, codelines, tag, author, content_type)
	    VALUES ($1, $2, $3, $4, $5)
	    RETURNING id, title, codelines, tag, author, content_type
	`
	codesnippet = structs.CodeSnippet{}
	err = db.QueryRow(sqlStatement, newcodesnippet.Title, newcodesnippet.Codelines, newcodesnippet.Tag, newcodesnippet.Author, newcodesnippet.ContentType).Scan(&codesnippet.ID, &codesnippet.Title, &codesnippet.Codelines, &codesnippet.Tag, &codesnippet.Author, &codesnippet.ContentType)
	if err != nil {
		panic(err)
	}	
	return codesnippet
}


// HANDLER CREATE
func createCodeSnippet(c echo.Context) error {
	codesnippet := new(structs.CodeSnippet)
	if err := c.Bind(codesnippet); err != nil {
		return err
	}
	response := queryCreate(codesnippet)
	return c.JSON(http.StatusOK, response)
}


// QUERY UPDATE
func queryUpdate(newcodesnippet *structs.CodeSnippet, id int) (codesnippet structs.CodeSnippet) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	sqlStatement := `
	    UPDATE codesnippet SET
		title = $1,
		codelines = $2,
		tag = $3,
		author = $4,
		content_type = $5
	    WHERE id = $6
	    RETURNING id, title, codelines, tag, author, content_type
	`
	codesnippet = structs.CodeSnippet{}
	err = db.QueryRow(sqlStatement, newcodesnippet.Title, newcodesnippet.Codelines, newcodesnippet.Tag, newcodesnippet.Author, newcodesnippet.ContentType, id).Scan(&codesnippet.ID, &codesnippet.Title, &codesnippet.Codelines, &codesnippet.Tag, &codesnippet.Author, &codesnippet.ContentType)
	if err != nil {
		panic(err)
	}	
	return codesnippet
}


// HANDLER UPDATE
func updateCodeSnippet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	codesnippet := new(structs.CodeSnippet)
	if err := c.Bind(codesnippet); err != nil {
		return err
	}
	response := queryUpdate(codesnippet, id)
	return c.JSON(http.StatusOK, response)
}



// QUERY DELETE
func queryDelete(id int) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	sqlStatement := `
	    DELETE FROM codesnippet
	    WHERE id = $1
	`
	//codesnippet = structs.CodeSnippet{}
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}	
	return nil
}


// HANDLER DELETE
func deleteCodeSnippet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	//codesnippet := new(structs.CodeSnippet)
	//if err := c.Bind(codesnippet); err != nil {
	//	return err
	//}
	_ = queryDelete(id)
	return c.JSON(http.StatusOK, "deleted")
}




func main() {
	//dbinit()

	e := echo.New()

	// routing
	e.GET("/hello", hello)
	e.GET("/getAll", getAll)
	e.GET("/getPerId/:id", getPerId)
	e.POST("/createCodeSnippet", createCodeSnippet)
	e.PUT("/updateCodeSnippet/:id", updateCodeSnippet)
	e.DELETE("/deleteCodeSnippet/:id", deleteCodeSnippet)

	e.Logger.Fatal(e.Start(":1333"))
}
