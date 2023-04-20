package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Разделить по логике на файлы
// Применить ООП с использованием интрфейсов, ресиверов (методов), наследование

type User struct {
	Username string
	Password string
}
type Str struct {
	Users []User
}

type Item struct {
	Name   string
	Price  float64
	Rating float64
}
type It struct {
	Items []Item
}

func ReadJStr() Str {
	rawDataIn, err := os.ReadFile("Usr.json")
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}
	var settings Str
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
	return settings
}

func WriteJStr(settings Str) {
	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}
	err = os.WriteFile("Usr.json", rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}

func ReadJIt() It {
	rawDataIn, err := os.ReadFile("Itm.json")
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}
	var settings It
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
	return settings
}

func WriteJIt(settings It) {
	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = os.WriteFile("Itm.json", rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}

func ToHash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}

func PrintItems(item Item) {
	fmt.Println("Item name:", item.Name, "\nItem price:", item.Price, "\nItem rating:", item.Rating, "\n---------------")
}

func main() {
	fmt.Println("\nWelcome to my Store\n")

	for true {
		fmt.Print("1. Login\n2. Register\n3. Break\nEnter your choice: ")

		var index string
		fmt.Scan(&index)
		switch index {
		case "1":
			Login()
		case "2":
			Register()
		case "3":
			fmt.Println("See you soon!")
			return
		case "admin":
			fmt.Println("You not admin!!!")
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func Register() {
	fmt.Print("Enter your username: ")
	var name, pass string
	fmt.Scan(&name)
	fmt.Print("Enter your password: ")
	fmt.Scan(&pass)
	hashingPass := ToHash(pass)

	newUser := User{name, hashingPass}

	settings := ReadJStr()

	settings.Users = append(settings.Users, newUser)

	WriteJStr(settings)

	fmt.Println("Successfully registered.")
}

func Login() {
	fmt.Print("Enter your username: ")
	var name, pass string
	fmt.Scan(&name)
	fmt.Print("Enter your password: ")
	fmt.Scan(&pass)
	hashingPass := ToHash(pass)
	settings := ReadJStr()

	for _, user := range settings.Users {
		if strings.ToLower(user.Username) == strings.ToLower(name) && user.Password == hashingPass {
			fmt.Println("Successfully logged in.")
			for true {
				fmt.Print("1. Add item\n2. Searching items\n3. Filtering items\n4. Giving rating\n5. List of items\n0. Exit\nEnter your choice: ")

				var index string
				fmt.Scan(&index)
				switch index {
				case "1":
					AddItem()
				case "2":
					Searching()
				case "3":
					Filter()
				case "4":
					GiveRating()
				case "5":
					ListItem()
				case "0":
					return
				}
			}
		}
	}
	fmt.Println("Incorrect username or password.")
}

func AddItem() {
	var name string
	var price, rating float64
	fmt.Print("Enter item name: ")
	fmt.Scan(&name)
	fmt.Print("Enter item price: ")
	fmt.Scan(&price)
	fmt.Print("Enter item rating: ")
	fmt.Scan(&rating)
	newItem := Item{name, price, rating}

	settings := ReadJIt()

	settings.Items = append(settings.Items, newItem)

	WriteJIt(settings)

	fmt.Println("You have successfully added the item.")
}

func Searching() {
	var name string
	fmt.Print("Enter item name: ")
	fmt.Scan(&name)
	settings := ReadJIt()

	for _, item := range settings.Items {
		if item.Name == name {
			PrintItems(item)
			return
		}
	}
	fmt.Println("This item is not exist.")
}

func Filter() {
	var price, price2, rating, rating2 float64
	fmt.Println("1. Sort by price?")
	fmt.Println("2. Sort by rating?")
	fmt.Println("3. Sort by price and rating?")
	settings := ReadJIt()
	var a string
	fmt.Scan(&a)
	switch a {
	case "1":
		fmt.Print("Enter item start price: ")
		fmt.Scan(&price)
		fmt.Print("Enter item end price: ")
		fmt.Scan(&price2)
		for _, item := range settings.Items {
			if item.Price >= price && item.Price <= price2 {
				PrintItems(item)
			}
		}
	case "2":
		fmt.Print("Enter item start rating: ")
		fmt.Scan(&rating)
		fmt.Print("Enter item end rating: ")
		fmt.Scan(&rating2)
		for _, item := range settings.Items {
			if item.Rating >= rating && item.Rating <= rating2 {
				PrintItems(item)
			}
		}
	case "3":
		fmt.Print("Enter the maximum price of the item: ")
		fmt.Scan(&price)
		fmt.Print("Enter the minimum rating of the item: ")
		fmt.Scan(&rating)
		for _, item := range settings.Items {
			if item.Price <= price && item.Rating >= rating {
				PrintItems(item)
			}
		}
	}

}
func GiveRating() {
	var name string
	var rating float64
	fmt.Print("Enter item name: ")
	fmt.Scan(&name)
	fmt.Print("Enter item rating: ")
	fmt.Scan(&rating)
	settings := ReadJIt()
	for _, item := range settings.Items {
		if item.Name == name {
			item.Rating = rating
			fmt.Println("Rating edited successfully.")
			PrintItems(item)
			return
		}
	}
}
func ListItem() {
	settings := ReadJIt()
	for _, item := range settings.Items {
		PrintItems(item)
	}
}
