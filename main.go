package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"time"

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
		//code 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: User already exists"))
	}

}

func CreateNewMark(w http.ResponseWriter,r *http.Request){


	//получаем токен из куки
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			//если куки не настроены
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// если другие то 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString,
		claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	var mark Mark
	vars := mux.Vars(r)
	m := vars["mark"]
	err = json.NewDecoder(r.Body).Decode(&mark)
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

	//получаем токен из куки
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			//если куки не настроены
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// если другие то 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString,
		claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}



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

	//получаем токен из куки
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			//если куки не настроены
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// если другие то 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString,
		claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	vars := mux.Vars(r)
	m := vars["mark"]

	var updMark Mark
	err = json.NewDecoder(r.Body).Decode(&updMark)
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

	//получаем токен из куки
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			//если куки не настроены
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// если другие то 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString,
		claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	m := vars["mark"]

	var updMark Mark
	err = json.NewDecoder(r.Body).Decode(&updMark)
	fmt.Println("++",m,err)
	if err != nil && err != io.EOF {
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

	//получаем токен из куки
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			//если куки не настроены
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// если другие то 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString,
		claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

var jwtKey = []byte("default_key")

//второе поле нужно для expiry
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


//выдаем токен
func Auth(w http.ResponseWriter, r *http.Request)  {
	//декодируем полученные имя и пароль
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//наличие пользователя
	_, ok := userStore[user.Username]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//правильность пароля
	outsideUserInfo, ok := userStore[user.Username]
	if !ok || outsideUserInfo.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//определяем время действия токена
	expirationTime := time.Now().Add(3 * time.Minute)

	//для заявок определим имя пользователя и expire
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//header + payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// jwt token string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/register",Register).Methods("POST")
	r.HandleFunc("/stock",Stock).Methods("GET")
	r.HandleFunc("/auth",Auth).Methods("POST")
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
