package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"send_email_go/mailers"
)

//User is for save information from csv
type User struct {
	Name  string
	Email string
}

type Response struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func main() {
	http.HandleFunc("/user", greet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	email := query.Get("email")
	fmt.Println(name)
	fmt.Println(email)
	// users, _ := parseUser("./user.csv")
	// for _, user := range users {
	// 	subject := "Prueba Sunapp es gratis"
	// 	receiver := []string{user.Email}
	// 	r := mailers.NewRequest(receiver, subject)
	// 	r.Send("templates/template.html", map[string]string{"username": user.Name})
	// }
	subject := "Prueba Sunapp es gratis"
	receiver := []string{email}
	ra := mailers.NewRequest(receiver, subject)
	ra.Send("templates/template.html", map[string]string{"username": name})
	response := Response{}
	response.Status = true
	responseConvert, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseConvert)
	// fmt.Fprintf(w, "Se envió correo a "+email+"")
}

func parseUser(file string) (map[string]*User, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	locations := map[string]*User{}
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return locations, err
		}
		p := &User{}
		p.Name = row[0]
		p.Email = row[1]
		locations[row[0]] = p
	}
}
