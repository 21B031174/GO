package main

import (
	"crypto/sha1"
	"encoding/hex"
	"html/template"
	"net/http"
	"strconv"
)

func addProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/add_product.html")
	tmpl.Execute(w, nil)
}
func saveProduct(w http.ResponseWriter, r *http.Request) {
	settings := ReadJIt()
	value, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	float := float32(value)
	value2, _ := strconv.ParseFloat(r.FormValue("rating"), 32)
	float2 := float32(value2)
	newItem := Item{Name: r.FormValue("name"), Price: float, Rating: float2}
	settings.Items = append(settings.Items, newItem)
	settings.WriteJIt()
	http.Redirect(w, r, "/home/", http.StatusSeeOther)
}

func searchProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/search_product.html")
	settings := ReadJIt()
	for _, item := range settings.Items {
		if item.Name == r.FormValue("name") {
			tmpl.Execute(w, item)
			return
		}
	}
	tmpl.Execute(w, nil)
}

func Filter(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/filter_product.html")
	settings := ReadJIt()

	var products []Item
	value, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	float := float32(value)
	value2, _ := strconv.ParseFloat(r.FormValue("rating"), 32)
	float2 := float32(value2)
	for _, item := range settings.Items {
		if item.Price <= float && item.Rating >= float2 {
			products = append(products, item)
		}
	}
	tmpl.Execute(w, products)

}

func GiveRating(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/give_rating.html")
	settings := ReadJIt()
	i := 0
	for _, item := range settings.Items {
		if item.Name == r.FormValue("name") {
			value2, _ := strconv.ParseFloat(r.FormValue("rating"), 32)
			float2 := float32(value2)
			item.Rating = float2
			settings.Items[i].Rating = float2
			settings.WriteJIt()
			tmpl.Execute(w, item)
			return
		}
		i++
	}

	tmpl.Execute(w, nil)
}
func ListProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/list_product.html")
	settings := ReadJIt()
	var products []Item
	for _, item := range settings.Items {
		products = append(products, item)
	}
	tmpl.Execute(w, products)
}

func ToHash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}
