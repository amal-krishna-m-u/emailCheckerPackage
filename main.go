package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"regexp"
)

// Template for the web form
const formTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Email Domain Verifier</title>
</head>
<body>
	<h1> Amalmullangath</h1>
	<p>This is a email verifier ,
	This will verify the existance of a email</p>
    <h2>Email Domain Verifier</h2>
    <form method="post" action="/verify">
        Email: <input type="email" name="email">
        <input type="submit" value="Verify">
    </form>
    {{if .}}
        <p>Result: {{.}}</p>
    {{end}}
</body>
</html>
`

// validateEmailDomain checks if the email's domain has MX records
func validateEmailDomain(email string) bool {
	parts := regexp.MustCompile(`@`).Split(email, 2)
	if len(parts) != 2 {
		return false
	}

	_, err := net.LookupMX(parts[1])
	return err == nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("form").Parse(formTemplate))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	email := r.FormValue("email")
	result := "Invalid or unreachable email domain."
	if validateEmailDomain(email) {
		result = "Email domain is valid and reachable."
	}

	tmpl.Execute(w, result)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
