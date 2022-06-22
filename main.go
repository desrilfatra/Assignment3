package main

import (
	entity "Assignment3/Entity"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var Cuaca entity.Status

func main() {
	go GoRandomCuaca()
	mux := http.NewServeMux()
	endpoint := http.HandlerFunc(statuscuaca)
	mux.Handle("/Cuaca", MiddlewareCuaca(endpoint))
	fmt.Println("Connect port")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func MiddlewareCuaca(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func statuscuaca(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	context := entity.Status{
		Water:       Cuaca.Water,
		Wind:        Cuaca.Wind,
		StatusWater: Cuaca.StatusWater,
		StatusWind:  Cuaca.StatusWind,
	}
	tpl.Execute(w, context)

	file, _ := ioutil.ReadFile("status.json")
	json.Unmarshal(file, &Cuaca)

}

func GoRandomCuaca() {
	for {
		Cuaca.Water = rand.Intn(20)
		Cuaca.Wind = rand.Intn(20)
		Cuaca.StatusWater = "Aman"
		Cuaca.StatusWind = "Aman"

		if Cuaca.Water <= 5 {
			Cuaca.StatusWater = "Aman"
		} else if Cuaca.Water >= 6 && Cuaca.Water <= 8 {
			Cuaca.StatusWater = "Siaga"
		} else {
			Cuaca.StatusWater = "Bahaya"
		}

		if Cuaca.Wind <= 6 {
			Cuaca.StatusWind = "Aman"
		} else if Cuaca.Wind >= 7 && Cuaca.Wind <= 15 {
			Cuaca.StatusWind = "Siaga"
		} else {
			Cuaca.StatusWind = "Bahaya"
		}

		jsonString, _ := json.Marshal(&Cuaca)
		ioutil.WriteFile("status.json", jsonString, os.ModePerm)
		time.Sleep(5 * time.Second)
	}

}
