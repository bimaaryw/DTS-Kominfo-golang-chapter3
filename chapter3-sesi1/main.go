package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

)

func main() {

	ticktack := time.NewTicker(15 * time.Second)

	for s := range ticktack.C {
		fmt.Println(s)
		placeHolder()
		randomNum()
	}

}

func placeHolder() {
	randomUser := rand.Intn(10)
	data := map[string]interface{}{
		"userId": randomUser,
	}

	reqJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

func randomNum()  {
	var (
		randomWater = rand.Intn(100)
		randomWind = rand.Intn(100)
		statusWater = "aman"
		statusWind = "aman"
	)

	type field struct {
		Water int `json:"water"`
		Wind int `json:"wind"`
	}

	if randomWater >= 6 && randomWater <= 8 {
		statusWater = "siaga"
	} else if randomWater > 8 {
		statusWater = "bahaya"
	}

	if randomWind >= 7 && randomWind <= 15 {
		statusWind = "siaga"
	} else if randomWater > 8 {
		statusWind = "bahaya"
	}

	jsonString := field{
		Water: randomWater,
		Wind: randomWind,
	}

	res, err := json.MarshalIndent(jsonString, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
	fmt.Println()
	fmt.Println("status water: ", statusWater)
	fmt.Println("status wind: ", statusWind)
}