package main

import (
	"encoding/json"
	"fmt"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//var jwtKey = []byte("example_jwt")
//
//type Jwt struct {
//	Username string `json:"username"`
//	jwt.StandardClaims
//}


type User struct {
	Username string `json:"username"`
	Password string `json:password`
}

type Mark struct {
	MaxSpeed int `json:"max_speed"`
	Distance int `json:"distance"`
	Handler  string `json:"handler"`
	Stock    string `json:"stock"`
}


var userStore = make(map[string]User)

var markStore = make(map[string]Mark)

func Register(w http.ResponseWriter,r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	_,ok := userStore[user.Username]

	//code 201
	if !ok {
		userStore[user.Username]=user
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Message : User created. Try to auth"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: User already exists"))
	}



	//code 400

	fmt.Println(userStore)
}

func CreateNewMark(w http.ResponseWriter,r *http.Request){
	var mark Mark
	vars := mux.Vars(r)
	m := vars["mark"]
	err := json.NewDecoder(r.Body).Decode(&mark)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_,ok := markStore[m]

	if !ok {
		markStore[m] = mark
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Message Auto created"))
	}else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Auto with that mark exists"))
	}
	fmt.Println(markStore)
}


func GetMark(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	m := vars["mark"]

	_,ok := markStore[m]

	if ok {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type","application/json")
		j,_ := json.Marshal(markStore[m])
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Auto with that mark not found"))
	}

}


func UpdateMark(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	m := vars["mark"]

	var updMark Mark
	err := json.NewDecoder(r.Body).Decode(&updMark)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_,ok := markStore[m]
	if ok {
		delete(markStore,m)
		markStore[m]=updMark
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Auto with that mark not found"))
	}
	//{"Error" : "Auto with that mark not found"}
}

func DeleteMark(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	m := vars["mark"]

	var updMark Mark
	err := json.NewDecoder(r.Body).Decode(&updMark)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_,ok := markStore[m]
	if ok {
		delete(markStore,m)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Auto with that mark not found"))
	}
}

func Stock(w http.ResponseWriter, r *http.Request){


	if len(markStore) > 0 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type","application/json")
		j,_ := json.MarshalIndent(markStore,"","  ")
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: No one autos found in DataBase"))
	}
}



func main(){
	r := mux.NewRouter()

	r.HandleFunc("/register",Register).Methods("POST")
	r.HandleFunc("/stock",Stock).Methods("GET")
	r.HandleFunc("/auto/{mark}",CreateNewMark).Methods("POST")
	r.HandleFunc("/auto/{mark}",GetMark).Methods("GET")
	r.HandleFunc("/auto/{mark}",UpdateMark).Methods("PUT")
	r.HandleFunc("/auto/{mark}",DeleteMark).Methods("DELETE")

	log.Println("Listening...")
	server := &http.Server{
		Addr:
		":8081",
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}
