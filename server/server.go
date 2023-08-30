package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"taxis/database/auth/xauth"
	"taxis/graph"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func conexion() *sql.DB {
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	loc := "America%2FLa_Paz"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=%s", dbuser, dbpass, dbhost, dbname, loc)
	db, err := sql.Open("mysql", dsn)
	fmt.Println(dsn)

	if err != nil {
		panic(err.Error())
	}

	er := db.Ping()
	if er != nil {
		panic(er.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

const defaultPort = "8016"

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := conexion()
	router := chi.NewRouter()
	router.Use(xauth.AuthMiddleware(db))
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	show_playground := os.Getenv("PLAYGROUND")
	if show_playground == "1" {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

/* func main() {
	godotenv.Load()
	b64 := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAgCAIAAACdAM/hAAAAA3NCSVQICAjb4U/gAAAFg0lEQVQ4jYVUe1BUVRg/59zXvl+810V5JKJioIKpYIoUpZmJKDqNomIqJpnYaI6DTiUNDoivdFBnzNGyGZupZpg0p+jhOxEwUMlWgYAFWRYWVtjde/fee05/4FR7aabvjzNnzvl+8/u+3/eArhsRbq/M6pnmpn6fQC1eZKwoH6YpsH4191QYMLGGN9b7T+1nIqIDhXtjv6x+CP5lCGmkkn2+6Gg6MUbratfbG4dXrRJW5wmHT/kvXmKxij1zRI4ZRxUWAwoBpf16NqK3MbLzJ83JXSpfx/y1i1S7N0f4m4ydV2hXg2H5q7oNS7SeWqbjJixYMZEEG9Xdw+csUA+7B7UGzt78eEEmlzyZ7Cn1GQ1M4gRDtMmTl8vt+BhIXuJ4alqc804Qs9gZV/NF6OFtdMARf+e89qMtGn9rZFcNCNi1y2ajA1sZvpVuu6xy1ZnmpScqmFFvuzeUkxfnqA5V9Pbx1O5d5vKyPlZvbLwvl+4CRe9aKstI00NBp+VNJqJIGZmNnstXcXiE5eW5/i6HdOfaUNFmYI4ApYdIjxt22Pk1+cxrOZacVQGBV+oF3bUGt9f3oIH5s5fetEn3ydEB61h2ZpKfpWSzAVRUoRArU7CScj0Rth+M+/pSUKko3k0yM4w2q9TmEGPHcFazmDYV5e+QkifrXLx+zgv8rJkgeznweiRZFfJ6TlEQtes6vHFOM9FKd1/lLpRpv6q0uO9FO29CsSOqKFdzdp9OctjarwPHL9o5MyYoBKMZXaQxbCgrHfHYaBvblzwd7Szpys7QRzkH95VoGCyufLNrbS6TMV8Ya1UKRqkoueGuVHEQnz0ZmJSgYxBZtJQk2MDpz4GrR3ouJWzh7KGkWZFVx4Yb20Ly1wWFTZ0ot7Q+4tMSQwwhvtg4sWArrmtCM6dbxkUNzc5gtm3vwxTb3+dLmwKbWs1LlgWDs1O5dYWawmI3BlR8rDozIzA+Hp06OSwCLtwmLMzGU1PoskphmNd29umW5gWB4a3zlpQp/JOh8J42p4ilW7dJ4Qb18OAwkOlpefjEXibGxlvjoqC3d3lxwpXa5qAmScu0lOyXbvzYHhlvphlx+/ua778TsayRVdpvj6PU6WxNLbxW0+fsZ6PHYGWX2KuR765OfGwpylPPe552/BxW+w2SuuOyptHH9mjc9dyVz1hfI1V9lM59ZbJyqmgEurtoxKKEaLm4yFR1emBiAvugwbsmD6QkBZruUzVXcVQk4+mX7T2mpSuCckY7txoKCkmnPdD4CHp8zJa1Gl2o8N4B4dMLSIZGk178sBSrGb61F3AqZZ3pupuidbxpxrTBrDC0udh5oCKs+bZ0uYrXRZLfbvH9HtQ/AFta0NtvqTaVKFNGC5Zw9vtOiy28vl4u30M0yFd9kb/4AxlwAatNyHoJmNTyrFTcVM/L8ihw5XFv7Hg9gJ6qc1JYjFV46j/yAZ2/HtXXCTpWLQxREWFchA1PStIFhFHg5ERo1vkb7gyeOUI7Ha7T58G9P2BXC0pPZVRqn2dQCrci/wDT6/Sq1MoNSE+KE1itzW5/0tEiUYy8bQsNGBp7xf5+Sm+moRsjlpawv+13hmBlnRHvUxPeMzedGRNnFHwcQZpOOyTIDCA3MMCo9JzbKYWG616cp6VpRhk3+V+Tn50bN25UzjPGmBACCQQIEJlACgIMCCQQQgAAkQnFUAADMHrjP3sb+fvbY+QOACEEIIAxxgATWdkhAAA04kFkQiAB+FkWFEMBABBCCCFCCITwv5kDgQCRiUxkSZIAAlJAAgDIskwIwRLGGAMAMMZYwoSMak+v18uyrCRJHMcF+ACiEZCAKIpqTv0PGwaQgiMqKMFEJgABnucxxiiA1Fq1IAgsy4o+UaVRjWgxorkC/BejsQ173ANBJQAAAABJRU5ErkJggg=="
	_, er := firestore.SubirImagen(b64, "demo2")
	if er != nil {
		fmt.Println(er)
	}
}
*/
