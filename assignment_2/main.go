package main

import (
	"net/http"
)

type User struct {
	Username string
	Password string
}
type Str struct {
	Users []User
}

type Item struct {
	Name   string
	Price  float32
	Rating float32
}
type It struct {
	Items []Item
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/register/", register)
	http.HandleFunc("/home/", home)
	http.HandleFunc("/checking/", checking)
	http.HandleFunc("/add_product/", addProduct)
	http.HandleFunc("/save_product/", saveProduct)
	http.HandleFunc("/search_product/", searchProduct)
	http.HandleFunc("/filter/", Filter)
	http.HandleFunc("/give_rating/", GiveRating)
	http.HandleFunc("/list_product/", ListProduct)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/register/static/", http.StripPrefix("/register/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}
