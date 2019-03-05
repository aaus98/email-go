package main

import (
	"fmt"
	"send_email_go/mailers"
)

func main() {
	fmt.Println("asdasdas")
	subject := "Get latest Tech News directly to your inbox"
	receiver := []string{"adri_ubi_9@hotmail.com", "luis.pch01@gmail.com", "paulov01324@gmail.com"}
	r := mailers.NewRequest(receiver, subject)
	r.Send("templates/template.html", map[string]string{"username": "Paulo"})
}
