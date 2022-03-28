package main

import (
	"fmt"
	"log"
	"net/http"
)

//notewe use Fprintf here alot cos it prints on the stream where the frontend browser of client can recieve the message
func formHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "POST" {
		http.Error(res, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(res, "Post request successful \n")

	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(res, "Name is %[1]s and address is %[2]s \n", name, address)

}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "method not supported", http.StatusNotFound)
		return

	}

	fmt.Fprintf(res, "hello!")
}

func main() {

	//check out the static directory
	fileServer := http.FileServer(http.Dir("./static"))

	//handle "/" route.. here we want the user to be redirected to the home page index.html..mind u golang automatically knows the "/" route should lead to index.html so it implicitly automatically looks out for the index.html
	http.Handle("/", fileServer)

	//we are using HandleFunc here because this is a route that directs to a controller function
	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println(" server starting at port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
