package main

import (
	"encoding/csv"
	"io"
	"os"
	"send_email_go/mailers"
)

//User is for save information from csv
type User struct {
	Name  string
	Email string
}

func main() {
	users, _ := parseUser("./user.csv")
	for _, user := range users {
		subject := "Prueba Sunapp es gratis"
		receiver := []string{user.Email}
		r := mailers.NewRequest(receiver, subject)
		r.Send("templates/template.html", map[string]string{"username": user.Name})
	}
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
