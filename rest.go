package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home page")
	fmt.Println("Endpoint hit")
}

func handleRequest() {
	log.Println("Server Started:")
	log.Println("press ctrl and c to quit the server")
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homepage)
	myrouter.HandleFunc("/new-booking", createNewBooking).Methods("GET")
	myrouter.HandleFunc("/all-bookings", returnAllBookings).Methods("GET")
	myrouter.HandleFunc("/total-bookings", returnTotalBookings).Methods("GET")
	myrouter.HandleFunc("/booking/{id}", returnSingleBooking).Methods("GET")
	myrouter.HandleFunc("/update-booking/{id}/{user}", updateBooking).Methods("GET")
	myrouter.HandleFunc("/deleteBooking/{id}", deleteBooking).Methods("GET")
	log.Fatal(http.ListenAndServe(":8085", myrouter))
}

func main() {

	db, err = gorm.Open("mysql", "manish:manish@tcp(localhost:3306)/student?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connetion Failed")
	} else {
		log.Println("Connetion established")
	}

	db.AutoMigrate(&Booking{})
	handleRequest()

}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err == nil {
		var booking Booking
		er := json.Unmarshal(reqBody, &booking)
		if er == nil {
			db.Create(&booking)
			fmt.Println("Endpoint Hit: Creating New Booking")
			json.NewEncoder(w).Encode(booking)
		} else {
			fmt.Println(er)
		}
	}
}

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	bookings := []Booking{}
	db.Find(&bookings)
	fmt.Println("Endpoint Hit: returnAllBookings")
	json.NewEncoder(w).Encode(bookings)
}

func returnTotalBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	bookings := []Booking{}
	db.Raw("Select * from bookings").Scan(&bookings)
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []Booking{}
	s, err := strconv.Atoi(key)
	// db.Find(&bookings)
	// for _, booking := range bookings {
	// 	// string to int
	// 	s, err := strconv.Atoi(key)
	// 	if err == nil {
	// 		if booking.Id == s {
	// 			fmt.Println(booking)
	// 			fmt.Println("Endpoint Hit: Booking No:", key)
	// 			json.NewEncoder(w).Encode(booking)
	// 		}
	// 	}
	// }
	if err == nil {
		db.Raw("Select * from bookings where id= ?", s).Scan(&bookings)
		json.NewEncoder(w).Encode(bookings)
	}

}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key1 := vars["id"]
	key2 := vars["user"]
	bookings := []Booking{}
	s, err := strconv.Atoi(key1)
	if err == nil {
		db.Model(&bookings).Where("id = ?", s).Update("user", key2)
		fmt.Println("Updated")
	} else {
		fmt.Print(err)
	}

}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []Booking{}
	s, err := strconv.Atoi(key)
	if err == nil {
		db.Where("id = ?", s).Delete(&bookings)
		fmt.Println("Deleted")
	} else {
		fmt.Print(err)
	}
}
