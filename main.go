package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"./src/occupation"
)

const (
	DBNAME = "./database/onet.db" // onet database
)


// Context which holds the database instance
type HandlerContext struct {
	db *sql.DB // the database
}

// NewHandlerContext creates a new HandlerContext
func NewHandlerContext(database *sql.DB) *HandlerContext {
	if database == nil {
		panic("Database is null")
	}
	return &HandlerContext{database}
}

// API handler for endpoint: /
func (ctx *HandlerContext) getMainEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to VocationNation API v.0.0.1")
}

// API handler for endpoint: /occupations
func (ctx *HandlerContext) getAllOccupations(w http.ResponseWriter, r *http.Request) {
	var AllOccupations []occupation.Occupation
	AllOccupations = occupation.GetOccupations(ctx.db)
	err := json.NewEncoder(w).Encode(AllOccupations)
	if err != nil {
		fmt.Println("JSON encoding failed, err: ", err)
	}
}

// API handler for endpoint: /occupation/code/{onetcode}
func (ctx *HandlerContext) getOccupationById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["onetcode"]

	Occupation := occupation.GetOccupationById(ctx.db, key)
	err := json.NewEncoder(w).Encode(Occupation)
	if err != nil {
		fmt.Println("JSON encoding failed, err:", err)
	}
}

// API handler for endpoint: /occupation/title/{title}
func (ctx *HandlerContext) getOccupationsByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["title"]

	Occupations := occupation.GetOccupationsByTitle(ctx.db, key)
	err := json.NewEncoder(w).Encode(Occupations)
	if err != nil {
		fmt.Println("JSON encoding failed, err:", err)
	}
}

// Channels the http requests to appropriate handler functions
func handleRequests(hctx *HandlerContext) {

	// creates a new mux router
	router := mux.NewRouter().StrictSlash(true)

	// Add API endpoints
	router.HandleFunc("/", hctx.getMainEndpoint)
	router.HandleFunc("/occupations", hctx.getAllOccupations)
	router.HandleFunc("/occupation/code/{onetcode}", hctx.getOccupationById)
	router.HandleFunc("/occupation/title/{title}", hctx.getOccupationsByTitle)

	// serve the API at :5000 port
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {

	db, err := sql.Open("sqlite3", DBNAME)
	if err != nil {
		fmt.Printf("Database cannot connect, err:", err)
	}

	hctx := NewHandlerContext(db)
	handleRequests(hctx)
}
