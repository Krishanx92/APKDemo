package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Hotel struct {
	Id        string `json:"id"`
	Name      string `json:"hotel_name"`
	Passenger string `json:"passenger"`
	Date      string `json:"date"`
}

var reservation = []Hotel{
	{Id: "1234123", Name: "The Landmark", Passenger: "Mrs. Heily Feyers", Date: "2024/05/07"},
	{Id: "23451234", Name: "The Ritz", Passenger: "Mr. Brendon MacSmith", Date: "2024/07/08"},
	{Id: "34561234", Name: "St. Athans", Passenger: "Mr. Peter Queenslander", Date: "2024/04/09"},
	{Id: "45671243", Name: "The Langham", Passenger: "Miss. Liza MacAdams", Date: "2024/05/10"},
}

func getReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if jsonData, err := json.MarshalIndent(reservation, "", "    "); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

func addReservation(w http.ResponseWriter, r *http.Request) {
	var newReservation Hotel
	if err := json.NewDecoder(r.Body).Decode(&newReservation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reservation = append(reservation, newReservation)
	json.NewEncoder(w).Encode(newReservation)
}

func editReservation(w http.ResponseWriter, r *http.Request) {
	reservationId := strings.TrimPrefix(r.URL.Path, "/reservation/")
	for i, reserv := range reservation {
		if reserv.Id == reservationId {
			var updatedReserv Hotel
			if err := json.NewDecoder(r.Body).Decode(&updatedReserv); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			reservation[i] = updatedReserv
			json.NewEncoder(w).Encode(updatedReserv)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteReservation(w http.ResponseWriter, r *http.Request) {
	reservationId := strings.TrimPrefix(r.URL.Path, "/reservation/")
	for i, reserv := range reservation {
		if reserv.Id == reservationId {
			reservation = append(reservation[:i], reservation[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/reservation", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getReservation(w, r)
		case "POST":
			addReservation(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/reservation/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			editReservation(w, r)
		case "DELETE":
			deleteReservation(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":82", nil)
}
