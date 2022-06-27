//Written in Golang by Harvey Coombs - http://harveycoombs.com/
package main 

import (
	"fmt"
	"net/http"
	"math/rand"
	"io"
	"log"
)

const kobore string = "http://cli.kobore.co/"

func main() {
	fmt.Println("KOBORE CLI | 2022 | HTTPS://KOBORE.CO/")
	fmt.Println("\nEnter an Alias:")

	var alias string
	fmt.Scanln(&alias)

	connect(alias)
}

func connect(alias string) {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	token := ""

	for j := 1; j <= 16; j++ {
		i := rand.Intn(len([]rune(chars)))
		token += string([]rune(chars)[i])
	}

	subject := kobore + "/connect.php?auth=" + token + "&alias=" + string(alias)
	res, err := http.Get(subject)

	if err != nil {
		log.Fatal(err)
	} else {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		final := string(body)

		if err == nil {
			fmt.Println("No Issues. " + final)
		}
	}
}

func message() {
	//to-do
}

func fetch() {
	//to-do
}