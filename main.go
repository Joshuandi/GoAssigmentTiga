package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	auto_reload "GoAssigmentTiga/handler"

	"GoAssigmentTiga/model"

	"github.com/gorilla/mux"
)

var PORT = ":8088"
var status model.Web

const JsonPath = "web/weather.json"

func main() {
	go jsonParsing()
	r := mux.NewRouter()
	// userHandler := user_handler.NewUserHandler(database.Db)
	// r.HandleFunc("/users", userHandler.UserLoginHandler)
	// r.HandleFunc("/users/{id}", userHandler.UserLoginHandler)

	// middleware := middleware.NewAuthMiddleware(&cfg, googleTokenValidator)
	// r.Use(middleware.GetToken)
	// //r.Use(middleware.LoginMiddleware)
	autoreload := auto_reload.NewAutoReload(&model.Web{})
	r.HandleFunc("/", autoreload.AutoReload)
	fmt.Println("Now Loading on Port 0.0.0.0" + PORT)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8088",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func jsonParsing() {
	for {
		status.StatusWater = "aman"
		status.StatusWind = "aman"
		status.Water = rand.Intn(20)
		status.Wind = rand.Intn(20)
		if status.Water >= 6 && status.Water <= 8 {
			status.StatusWater = "Siaga"
		}
		if status.Wind >= 7 && status.Wind <= 15 {
			status.StatusWind = "Siaga"
		}
		if status.Water > 8 {
			status.StatusWater = "Bahaya"
		}
		if status.Wind > 15 {
			status.StatusWind = "Bahaya"
		}

		//write json file
		jsonS, _ := json.Marshal(&status)
		ioutil.WriteFile(JsonPath, jsonS, os.ModePerm)
		time.Sleep(10 * time.Second)
	}
}
