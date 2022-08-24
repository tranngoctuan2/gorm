package controllers

import (
	"encoding/json"
	"fmt"
	"gorm/database"
	"gorm/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func Init() {
	var server = Server{}
	server.DB = database.InitDB()
	server.DB.AutoMigrate(&models.User{})
	server.Router = mux.NewRouter()
	// server initialization
	server.initializeRoutes()
	server.Run(":8080")
	return
}

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/users", server.CreateUser).Methods("POST")
	server.Router.HandleFunc("/users", server.GetUsers).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.GetUserByID).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")
}
func (server *Server) Run(addr string) {
	fmt.Println("Listen to port" + addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

// create user
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", 500)
		return
	}
	// convert json sang object go
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Bad convert json to object users", 500)
		return
	}
	err = models.CreateUser(server.DB, &user)
	if err != nil {
		http.Error(w, "Error create user", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// get user
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	err := models.GetUsers(server.DB, &user)
	if err != nil {
		http.Error(w, "Cannot get all users or null", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// get user id
func (server *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Bad request", 500)
		return
	}
	var user models.User
	err = models.GetUserByID(server.DB, &user, int32(uid))
	if err != nil {
		http.Error(w, "can not get user by ID", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// update user by id
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Bad request", 500)
		return
	}
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", 500)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Bad convert json to object users", 500)
		return
	}
	err = models.UpdateUser(server.DB, &user, int32(uid))
	if err != nil {
		http.Error(w, "Cannot update user", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Bad request", 500)
		return
	}
	var user models.User
	err = models.DeleteUser(server.DB, &user, int32(uid))
	if err != nil {
		http.Error(w, "Cannot delete users", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
