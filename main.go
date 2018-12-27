package main

import ( 
	"log"
	"fmt"
	// "log"
	"strings"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"encoding/json"
)

type abc struct {
	StoreId json.Number `json:"StoreId,Number"`
}

func main() {
	file, err := ioutil.ReadFile("resources/va_stores.json")
	if err != nil {
		log.Fatal(err)
	}

	var stores []abc
  err = json.Unmarshal(file, &stores)
  if err != nil {
    log.Fatal(err)
  }

	for _, id := range stores {
		id := fmt.Sprintf("%v", id)		//convert from type abc to string
		id = strings.Replace(id, "{", "", -1)
		id = strings.Replace(id, "}", "", -1)
		queryStore(string(id))
		// println("https://www.abc.virginia.gov/api/stores/inventory/mystore/" + id + "/027100")
	}

	// fmt.Printf("%v\n", stores)
	// return
}

func queryStore(storeId string) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("https://www.abc.virginia.gov/api/stores/inventory/mystore/" + storeId + "/027100")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
