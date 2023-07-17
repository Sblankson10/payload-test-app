package controllers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"payload-app/api/utils"
)

type Server struct {
	Db     *sql.DB
	Router *mux.Router
}

//func (s *Server) Init() {
//	s.Db = utils.DbConnect()
//	utils.RabbitMqConnect()
//	s.initRoutes()
//}

func (s *Server) Init() {
	db, err := utils.DbConnect()
	if err != nil {
		panic(err)
	}
	s.Db = db

	utils.RabbitMqConnect()
	s.initRoutes()
}

func (s *Server) Run() {
	port := ":" + os.Getenv("PORT")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "content-type", "content-length", "accept-encoding", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST"})

	log.Println("Listening on port ", port)

	if err := http.ListenAndServe(port, handlers.CORS(origins, headers, methods)(s.Router)); err != nil {
		log.Println("Unable to start app because ", err)
	}
}
