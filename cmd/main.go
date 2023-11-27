package main

import (
	"encoding/json"
	"fmt"
	"github/SGDIEGO/UrlShortener/internal/core/db"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type config struct {
	host string
	port string
}

var Cf config

func loadEnv(c *config) {
	c.port = ":"
	// Load data from .env

	if c.host == "" {
		c.host = "localhost"
	}
	if c.port == ":" {
		c.port += "3000"
	}
}

// Error Data
type ErrorResponse struct {
	Message string
	Code    int
}

// Url response

func main() {
	// Load configuration variables
	loadEnv(&Cf)

	// Load database
	db_, err := db.LoadDb()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	// Create server
	server := chi.NewRouter()

	// Configurate server
	server.Use(cors.AllowAll().Handler)
	//	server.Use(cors.Handler(cors.Options{
	//		AllowedOrigins:   []string{"https://*", "http://*", "*"},
	//		AllowedMethods:   []string{"GET", "POST"},
	//		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Origin"},
	//		AllowCredentials: false,
	//	}))

	server.Post("/", func(w http.ResponseWriter, r *http.Request) {
		type mssge struct {
			Original string `json:"original"`
		}

		var m mssge
		var body, err_ = io.ReadAll(r.Body)
		if err_ != nil {
			fmt.Printf("Error: %s", err.Error())
			return
		}
		err = json.Unmarshal(body, &m)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return
		}
		//		err = json.NewEncoder(w).Encode(m)
		//		if err != nil {
		//			fmt.Printf("Error: %s", err.Error())
		//			return
		//		}

		// New id
		id := rand.Intn(1000)
		created := fmt.Sprintf("localhost:3000/%d", id)

		_, err := db_.Exec("insert into url values(?,?,?)", id, m.Original, created)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		log.Printf("New data (Id: %d, Url: %s)", id, m.Original)
		w.Write([]byte(m.Original))
	})

	server.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strings.CutPrefix(r.URL.String(), "/") // Get id
		query := fmt.Sprintf("select original from url where id=%s", id)

		res, err := db_.Query(query)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		var url string
		for res.Next() {
			if err := res.Scan(&url); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}
		}

		// If url does not exist
		if url == "" {
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Url does not exist",
				Code:    http.StatusNotFound,
			})
			return
		}

		// Redirect
		http.Redirect(w, r, url, http.StatusSeeOther)
		//w.Write([]byte(url))
	})

	// Run server
	fmt.Printf("Runnig server on %s%s\n", Cf.host, Cf.port)
	if err := http.ListenAndServe(Cf.port, server); err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
}
