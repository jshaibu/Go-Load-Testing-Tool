package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

//
// ===================== CONFIG =====================
//

const (
	BaseURL      = "http://localhost:5115"
	LoginURL     = BaseURL + "/api/Auth/login"
	ProductsURL  = BaseURL + "/api/Products?page=1&pageSize=200&includeInactive=false"
	OrdersURL    = BaseURL + "/api/Orders"

	Email    = "admin@pos.local"
	Password = "Admin@1234!"
)

//
// ===================== API LAYER =====================
//

func login() string {

	payload := map[string]string{
		"email":    Email,
		"password": Password,
	}

	b, _ := json.Marshal(payload)

	resp, err := http.Post(LoginURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var res map[string]any
	json.Unmarshal(body, &res)

	token, _ := res["accessToken"].(string)
	return token
}

func fetchProducts(token string) []Product {

	req, _ := http.NewRequest("GET", ProductsURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data struct {
		Items []Product `json:"items"`
	}

	json.Unmarshal(body, &data)

	return data.Items
}

//
// ===================== ENTRY POINT =====================
//

func main() {

	token := login()
	products := fetchProducts(token)

	client := &http.Client{}
	metrics := NewMetrics()

	jobs := make(chan int, 1000)

	go Dashboard(metrics)

	var wg sync.WaitGroup
	workerCount := 10

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Worker(id, client, token, products, metrics, jobs, OrdersURL)
		}(i)
	}

	for i := 0; i < 1000; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}