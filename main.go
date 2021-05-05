// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
// like a class, used for OOP in GO
// type Book struct {
// 	//json field that we fetch
// 	ID     string  `json:"id"`
// 	Isbn   string  `json:"isbn"`
// 	Title  string  `json:"title"`
// 	Author *Author `json:"author"`
// }

// // Author Struct
// type Author struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

//Init books var as a slice Book struct
//slice is a variable like arrays
// var books []Book

// Get Shoplink
// func getShoplink(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// params := mux.Vars(r)
// 	var ID int
// 	ID = rand.Intn(10000000)

// 	var str1 string
// 	var shoplink string
// 	str1 = "https://shop-links.co/"
// 	shoplink = str1 + strconv.Itoa(ID)
// 	json.NewEncoder(w).Encode(shoplink)
// }

// swagger:route POST /api/narrativ/getShoplink shopLink newShopLink
// returns a unique shopLink

//Link Struct
type Link struct {
	//json field that we fetch
	ProductURL string `json:"productUrl"`
}

// Smartlink Struct
type Smartlink struct {
	// json field that we fetch
	// in: body
	ProductURL string `json:"productUrl"`
	Shoplink   string `json:"shoplink"`
}

// SmartlinkDB Struct
type SmartlinkDB struct {
	ProductURL string
	Shoplink   string
}

//Change to hash
//how does the client know what endpoints we are exposing?
// var smartlinks []Smartlink

// smartlinks2 := make(map[string]string)
var smartlinks2 map[string]string = map[string]string{}

var db *sql.DB

// func insertData(db *sql.DB, productURL string, shoplink string) error {
// 	q := "INSERT INTO smartlinks VALUES(?, ?)"
// 	insert, err := db.Prepare(q)
// 	defer insert.Close()

// 	if err != nil {
// 		return err
// 	}

// 	_, err = insert.Exec(productURL, shoplink)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func queryLink(db *sql.DB, productURL string) (shoplink []SmartlinkDB, err error) {
// 	q := "SELECT shoplink FROM smartlinks WHERE productURL = ?"
// 	resp, err := db.Query(q, productURL)
// 	defer resp.Close()

// 	if err != nil {
// 		return shoplink, err
// 	}

// 	for resp.Next() {
// 		var link SmartlinkDB
// 		err = resp.Scan(&link.ProductURL, &link.Shoplink)
// 		if err != nil {
// 			return shoplink, err
// 		}
// 		shoplink = append(shoplink, link)
// 	}

// 	return shoplink, nil
// }

//Function dbConn opens connection with MySQL driver
//sends the parameter `db *sq.DB` to be used by another function
func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/GO")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getShoplink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getShoplink")
	w.Header().Set("Content-Type", "application/json")

	//does use param but needs access to productUrl
	var smartlink Smartlink
	//decode body and store it into smartlink variable; & -> address of; * -> value at
	_ = json.NewDecoder(r.Body).Decode(&smartlink)

	// var existing bool
	// existing = false

	//**
	// if smartlinks2[smartlink.ProductURL] {
	// var existingLink string
	db := dbConn()
	q := `SELECT shoplink FROM smartlinks WHERE productURL=?;`

	var slink string
	row := db.QueryRow(q, smartlink.ProductURL)

	switch err := row.Scan(&slink); err {
	case sql.ErrNoRows:
		fmt.Println("else statement")
		var ID int
		ID = rand.Intn(10000000)

		var str1 string
		var shoplink string
		str1 = "https://shop-links.co/"
		shoplink = str1 + strconv.Itoa(ID)
		q := "INSERT INTO smartlinks VALUES(?, ?)"
		insert, err := db.Prepare(q)
		defer insert.Close()

		if err != nil {
			panic(err.Error())
		}

		_, err = insert.Exec(smartlink.ProductURL, shoplink)
		if err != nil {
			panic(err.Error)
		}
		// smartlinks2[smartlink.ProductURL] = shoplink
		json.NewEncoder(w).Encode(shoplink)
	case nil:
		json.NewEncoder(w).Encode(slink)
	default:
		panic(err)
	}
}

//**

// func getShoplink(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	//does use param but needs access to productUrl
// 	var smartlink Smartlink
// 	_ = json.NewDecoder(r.Body).Decode(&smartlink)
// for _, item := range smartlinks {
// 	if item.ProductURL == smartlink.ProductURL {
// 		json.NewEncoder(w).Encode(item)
// 		existing = true
// 		return
// 	}
// }

// if existing == true {
// 	json.NewEncoder(w).Encode(&Smartlink{})
// } else {
// 	var ID int
// 	ID = rand.Intn(10000000)

// 	var str1 string
// 	var shoplink string
// 	str1 = "https://shop-links.co/"
// 	shoplink = str1 + strconv.Itoa(ID)
// 	smartlink.Shoplink = shoplink
// 	smartlinks = append(smartlinks, smartlink)
// 	json.NewEncoder(w).Encode(shoplink)
// 	// json.NewEncoder(w).Encode(smartlinks)
// }
// }

// func getShoplink(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	var link Link
// 	_ = json.NewDecoder(r.Body).Decode(&link)

// 	for _, item := range table {
// 		if item.ProductURL == params["productUrl"] {
// 			json.NewEncoder(w).Encode(item)
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		} else {

// 		}
// 	}

// 	var ID int
// 	ID = rand.Intn(10000000)

// 	var str1 string
// 	var shoplink string
// 	str1 = "https://shop-links.co/"
// 	shoplink = str1 + strconv.Itoa(ID)
// 	json.NewEncoder(w).Encode(shoplink)
// }

//Get All Books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	//set header value to serve as json not as text
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

//Get Single Book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // Get params
// 	//Loop through books and find with id
// 	//like enumerate
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// // Create a New Book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID - not safe
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }

//combination of delete and create
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	//after deleting book will return all remaining books
// 	json.NewEncoder(w).Encode(books)
// }

func main() {
	//Init Router
	//:= does type inference
	r := mux.NewRouter()

	//Mock Data - @todo -implement DB
	// books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "476868", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/GO")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	// insert, err := db.Query("INSERT INTO smartlinks VALUES('https://www.bakedgoods.com/fudge', 'https://shop-links.co/555321')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	// fmt.Println("Successfully inserted into smartlinks table")

	//Route Handlers /Endpoints
	// r.HandleFunc("/api/books", getBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/api/books", createBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/narrativ/getShoplink", getShoplink).Methods("POST")

	// use redoc handler for serving swagger
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	//sh is handler
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	log.Fatal(http.ListenAndServe(":8000", r))

}
