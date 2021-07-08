package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head><title>Gupshup</title>
				<body>
					<h1>Let's do some quick gupshup!</h1>
				</body>
			</html>
		`))
	})

	log.Println("starting the gupshup server")

	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
