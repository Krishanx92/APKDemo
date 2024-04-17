package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Booking struct {
	Id             string `json:"id"`
	Passenger_name string `json:"passenger_name"`
	Destination    string `json:"destination"`
}

var booking = []Booking{
	{Id: "1234123", Passenger_name: "Mrs. Heily Feyers", Destination: "Birmingham"},
	{Id: "23451234", Passenger_name: "Mr. Brendon MacSmith", Destination: "Southampton"},
	{Id: "34561234", Passenger_name: "Mr. Peter Queenslander", Destination: "London"},
	{Id: "45671243", Passenger_name: "Miss. Liza MacAdams", Destination: "Bristol"},
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if jsonData, err := json.MarshalIndent(booking, "", "    "); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

func addBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking Booking
	if err := json.NewDecoder(r.Body).Decode(&newBooking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	booking = append(booking, newBooking)
	json.NewEncoder(w).Encode(newBooking)
}

func editBooking(w http.ResponseWriter, r *http.Request) {
	employeeId := strings.TrimPrefix(r.URL.Path, "/booking/")
	for i, emp := range booking {
		if emp.Id == employeeId {
			var updatedEmployee Booking
			if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			booking[i] = updatedEmployee
			json.NewEncoder(w).Encode(updatedEmployee)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingId := strings.TrimPrefix(r.URL.Path, "/booking/")
	for i, book := range booking {
		if book.Id == bookingId {
			booking = append(booking[:i], booking[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/booking", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getBooking(w, r)
		case "POST":
			addBooking(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/booking/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			editBooking(w, r)
		case "DELETE":
			deleteBooking(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":81", nil)
}
