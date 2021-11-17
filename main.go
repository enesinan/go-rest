package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Food struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IsSpicy string `json:"isSpicy"`
}

var data []Food
var foo []Food

var timer = time.NewTimer(5 * time.Second)

// Main
func main() {
	r := mux.NewRouter()

	data = append(data, Food{ID: "1", Name: "Kebap", IsSpicy: "Yes"})
	data = append(data, Food{ID: "2", Name: "Pide", IsSpicy: "No"})
	// data = append(data, foo)
	// tempfoo append to data
	r.HandleFunc("/get", getAllData).Methods("GET")
	r.HandleFunc("/get/{id}", getData).Methods("GET")
	r.HandleFunc("/create", createData).Methods("POST")
	r.HandleFunc("/update/{id}", updateData).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteData).Methods("DELETE")
	r.HandleFunc("/flush", Flush)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: ServerLog(r),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	//	var foo [3]Food
	//	copy(foo[:3], data)
	//  return foo

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}

// GetAll
func getAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Get
func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range data {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Food{})
}

// Add
func createData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d Food
	_ = json.NewDecoder(r.Body).Decode(&d)
	d.ID = d.ID
	fmt.Println("süre başladı")
	<-timer.C
	data = append(data, d)
	json.NewEncoder(w).Encode(d)
	fmt.Println("bitti başarılı")
}

// Update
func updateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range data {
		if item.ID == params["id"] {
			data = append(data[:index], data[index+1:]...)
			var d Food
			_ = json.NewDecoder(r.Body).Decode(&d)
			d.ID = params["id"]
			<-timer.C
			data = append(data, d)
			json.NewEncoder(w).Encode(d)
			return
		}
	}
}

//Flush
func Flush(w http.ResponseWriter, r *http.Request) {
	var f Food
	f.ID = ""
	f.Name = ""
	f.IsSpicy = ""
	fmt.Fprintf(w, f.Name)
	json.NewEncoder(w).Encode(f)
}

// Delete
func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range data {
		if item.ID == params["id"] {
			data = append(data[:index], data[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data)
}

func ServerLog(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}
