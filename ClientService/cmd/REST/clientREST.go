package main
import (
	"log"
	"net/http"
	"io/ioutil"
)
func main() {
	resp, err := http.Get("http://localhost:8080/products")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}