package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type Response struct {
	Number    string   `json:"data"`
	APIStatus string   `json:"apiStatus"`
	Assets    []Assets `json:"assets"`
	Date      string   `json:"dateModified"`
	Name      string   `json:"name"`
	ID        string   `json:"slug"`
	Status    string   `json:"status"`
}

type Assets struct {
	Icon  string `json:"iconUrl"`
	Tile  string `json:"tileUrl"`
	Cover string `json:"coverURL"`
}

func main() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.curseforge.com/v1/games/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", "$2a$10$uH6Iu6LmjTERw/KtEa3kuOcQYvkV.YhKSxuoD3mRJ3KnDu8PqALHW")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("%v\n", responseObject)

}
