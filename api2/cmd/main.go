package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/family-mmyr/testWebAPIServer/api2/domain/model"
)

func main() {
	var u url.URL
	u.Scheme = "http"
	u.Host = "localhost:8080"
	u.Path = "students"

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		fmt.Println("1. ", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("2.", err)
		return
	}

	defer resp.Body.Close()

	var respData model.ResponseStudents

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		fmt.Println("3. ", err)
		return
	}

	fmt.Printf("%#v\n", respData)
}
