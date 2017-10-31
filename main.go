package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type serviceDescriptor struct {
	Name string
}

type beveragePrice struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type specialOffer struct {
	Promotion       string `json:"promotion"`
	PercentDiscount int    `json:"percent_discount"`
}

var templates = template.Must(template.ParseFiles("views/index.html"))

func main() {

	http.HandleFunc("/", statusPage)
	http.HandleFunc("/prices", beveragePrices)
	http.HandleFunc("/specialoffers", specialOffers)

	http.ListenAndServe(":8081", nil)
}

func statusPage(w http.ResponseWriter, r *http.Request) {
	tmpl := "index"
	sd := serviceDescriptor{Name: "Coffee Shop Pricing"}

	// Render template...
	err := templates.ExecuteTemplate(w, tmpl+".html", sd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Sends a list of drinks and prices to the client as JSON
func beveragePrices(w http.ResponseWriter, r *http.Request) {

	priceList := []beveragePrice{
		{Name: "Cappucino", Price: 299},
		{Name: "Latte", Price: 299},
		{Name: "Espresso", Price: 250},
		{Name: "Americano", Price: 280},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(priceList)
}

// Send information about any special offers that apply for the request
func specialOffers(w http.ResponseWriter, r *http.Request) {

	// We should have some logic in here that determines if
	// we return a specialOffer.
	//
	// For now we will just always return a fixed offer...

	specialOffers := []specialOffer{
		{Promotion: "Every Day is a Good Day!", PercentDiscount: 10},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(specialOffers)

}
