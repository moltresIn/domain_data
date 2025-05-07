package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
)

type PageData struct {
	Domain  string
	IP      string
	Message string
}
func resolveIP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))

	if r.Method == http.MethodPost {
		r.ParseForm()
		domain := r.FormValue("domain")

		// Resolve the IP address
		ips, err := net.LookupIP(domain)
		data := PageData{Domain: domain}

		if err != nil {
			data.Message = fmt.Sprintf("Failed to resolve IP for domain: %s", domain)
		} else if len(ips) > 0 {
			data.IP = ips[0].String()
			data.Message = "IP Address resolved successfully!"
		} else {
			data.Message = "No IP addresses found for the domain."
		}

		tmpl.Execute(w, data)
		return
	}

	tmpl.Execute(w, PageData{})
}

func main() {
	http.HandleFunc("/", resolveIP)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
