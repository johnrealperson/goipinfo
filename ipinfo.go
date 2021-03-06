/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//Found a cool video streamed a while back on parsing JSON raw data with golang
//but I think the video was running an older version of golang possibly. I had
//trouble with most of the syntax from the video, and have commented out the parts
//where I couldn't quite get it functioning.

//I noticed my API response from VT was actually more updated naturally without most
//ff the required structures from the video and it was entirely plaintext json format
//still need to learn parsing json if it is returned as string(body) or if there is
//a better way to return it for easier parsing later.

//Some of the commented code can be found here: hxxps://www[.]youtube[.]com/watch?v=Jhexd8YuiSU&
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	//"log"
	//"github.com/joho/godotenv"
)

/*type VTIPResponse struct {
	LastAnalysisStats [string]reputation
}

type reputation struct {
	Harmless int 'json:"harmless"'
	Malicious int 'json:"malicious"'
	Suspicious int 'json:"suspicious"'
	Undetected int 'json:"undetected"'
}
*/
func main() {
	//enverr := godotenv.Load()
	//if enverr != nil {
	//	log.Fatal("Error loading .env file")
	//}

	//apiKey := os.Getenv("API_KEY")
	apiKey := "YOUR_API_KEY_HERE "
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("IP: ")
	scanner.Scan()
	input := scanner.Text()
	ip := input
	checkIPAddress(ip)
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

	//fmt.Println(res)
	fmt.Println(string(body))
	/*var repResponse = new(VTIPResponse)
	err = json.Unmarshal(body, &reputation)
	if err != nil {
		panic(err)
	}
	fmt.Println(repResponse)
	*/
}

func checkIPAddress(ip string) {
	if net.ParseIP(ip) == nil {
		fmt.Printf("IP Address: %s - Invalid\n", ip)
	} else {
		fmt.Printf("IP Address: %s - Valid\n", ip)
	}
}
