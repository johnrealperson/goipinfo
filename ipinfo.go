package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("IP: ")
	scanner.Scan()
	ip := scanner.Text()
	url := "https://www.virustotal.com/api/v3/ip_addresses/" + ip

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-apikey", "c241bec690d0f217240381d069b6d41ee61bbf2f88f0f4829a15d4a2e3b4bfdd")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
