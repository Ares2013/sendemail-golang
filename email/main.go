package main

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "pengpanpan0924@163.com")
	m.SetHeader("To", "1534157801@qq.com", "184472960@qq.com")
	m.SetAddressHeader("Cc", "pengpanpan0924@163.com", "Ares Peng")
	m.SetHeader("Subject", "Hello! Golang Send email")
	m.SetBody("text/html", "Hello <b>Ares</b> and <i>honey</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.163.com", 25, "pengpanpan0924@163.com", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
