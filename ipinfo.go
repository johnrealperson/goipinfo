package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"log"
	//"github.com/joho/godotenv"
)

type VTIPResponse struct {
	LastAnalysisStats [string]reputation
}

type reputation struct {
	Harmless int 'json:"harmless"'
	Malicious int 'json:"malicious"'
	Suspicious int 'json:"suspicious"'
	Undetected int 'json:"undetected"'
}

func main() {
	//enverr := godotenv.Load()
	//if enverr != nil {
	//	log.Fatal("Error loading .env file")
	//}

	//apiKey := os.Getenv("API_KEY")
	apiKey := "API_KEY_HERE"
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

	var repResponse = new(VTIPResponse)
	err = json.Unmarshal(body, &reputation)
	if err != nil {
		panic(err)
	}
	fmt.Println(repResponse)

}
