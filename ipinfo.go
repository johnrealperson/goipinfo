package main

import (
	"bufio"
	"fmt"
	"io/ioutil"

	//"log"
	"net/http"
	"os"
	//"github.com/joho/godotenv"
)

func main() {
	//enverr := godotenv.Load()
	//if enverr != nil {
	//	log.Fatal("Error loading .env file")
	//}

	//apiKey := os.Getenv("API_KEY")
	apiKey := "c241bec690d0f217240381d069b6d41ee61bbf2f88f0f4829a15d4a2e3b4bfdd"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("IP: ")
	scanner.Scan()
	ip := scanner.Text()
	url := "https://www.virustotal.com/api/v3/ip_addresses/" + ip

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-apikey", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println(string(body))

}
