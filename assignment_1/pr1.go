package main

import (
	"fmt"
)

//type User struct {
//	Username string
//	Password string
//	Email    string
//}

//type Item struct {
//	Name   string
//	Price  float64
//	Rating float64
//}

func (u *User) Register(username, password, email string) {
	u.Username = username
	u.Password = password

	fmt.Println("User Registered:", u.Username)
}

func (u *User) Login(username, password string) bool {
	if u.Username == username && u.Password == password {
		fmt.Println("User Logged in:", u.Username)
		return true
	}
	fmt.Println("Login Failed")
	return false
}

func (i *Item) GiveRating(rating float64) {
	i.Rating = rating
	fmt.Println("Item Rated:", i.Name)
}

type ItemStore struct {
	Items []Item
}

func (is *ItemStore) Search(name string) []Item {
	var result []Item
	for _, item := range is.Items {
		if item.Name == name {
			result = append(result, item)
		}
	}
	return result
}

func (is *ItemStore) Filter(price, rating float64) []Item {
	var result []Item
	for _, item := range is.Items {
		if item.Price <= price && item.Rating >= rating {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	itemStore := ItemStore{
		Items: []Item{
			Item{Name: "item1", Price: 10.0, Rating: 4.5},
			Item{Name: "item2", Price: 20.0, Rating: 3.5},
			Item{Name: "item3", Price: 30.0, Rating: 5.0},
			Item{Name: "item4", Price: 30.0, Rating: 3.5},
			Item{Name: "item5", Price: 10.0, Rating: 4.0},
			Item{Name: "item6", Price: 20.0, Rating: 5.0},
		},
	}

	user := &User{}
	user.Register("user1", "password", "user1@assignment_3.com")
	user.Login("user1", "password")

	items := itemStore.Search("item4")
	for _, item := range items {
		fmt.Println("Search Result:", item)
	}

	filteredItems := itemStore.Filter(15.0, 4.0)
	for _, item := range filteredItems {
		fmt.Println("Filtered Result:", item)
	}

	item := &Item{Name: "item4", Price: 40.0}
	item.GiveRating(4.5)
}
