package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a handler for the chatbot web page
	http.HandleFunc("/chatbot", func(w http.ResponseWriter, r *http.Request) {
		// Write the HTML for the chatbot page
		fmt.Fprintf(w, `
			<html>
				<head>
					<title>Chatbot</title>
				</head>
				<body>
					<h1>Chatbot</h1>
					<form action="/chatbot" method="post">
						<input type="text" name="message" placeholder="Type your message here...">
						<input type="submit" value="Send">
					</form>
				</body>
			</html>
		`)
	})

	// Create a handler for the chatbot's response
	http.HandleFunc("/chatbot", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the user's message
		message := r.FormValue("message")
		// Generate a response
		response := GenerateResponse(message)
		// Write the response
		fmt.Fprintf(w, response)
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
