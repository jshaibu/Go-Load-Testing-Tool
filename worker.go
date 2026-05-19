package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"io"
	"fmt"
)

func Worker(
	id int,
	client *http.Client,
	token string,
	products []Product,
	metrics *Metrics,
	jobs <-chan int,
	ordersURL string,
) {

	for range jobs {

		start := time.Now()

		size := rand.Intn(5) + 2
		items := make([]OrderItem, 0, size)
		used := make(map[string]bool)

		total := 0.0

		for i := 0; i < size; i++ {

			p := products[rand.Intn(len(products))]

			if used[p.Id] {
				i--
				continue
			}
			used[p.Id] = true

			qty := rand.Intn(4) + 1

			items = append(items, OrderItem{
				ProductId: p.Id,
				Quantity:  qty,
			})

			total += p.Price * float64(qty)
		}

		reqBody := OrderRequest{
			PaymentMethod:   2,
			DiscountPercent: 0,

			ProviderId:      "8236c425-19b5-4fff-b383-8f0f5be83f5f",
			CurrencySymbol:  "MK",

			AmountTendered:  total + float64(rand.Intn(100)), // safer margin

			Items:           items,
		}

		b, _ := json.Marshal(reqBody)

		req, err := http.NewRequest("POST", ordersURL, bytes.NewBuffer(b))
		if err != nil {
			metrics.AddFail("REQUEST_BUILD_ERROR")
			continue
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)

		lat := int(time.Since(start).Milliseconds())
		metrics.AddLatency(lat)

		if err != nil {
			metrics.AddFail("NETWORK")
			continue
		}

		resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			metrics.AddSuccess()
		} else {
			metrics.AddFail("HTTP_" + http.StatusText(resp.StatusCode))

			// IMPORTANT: debug output
			fmt.Println("FAILED REQUEST RESPONSE:")
			fmt.Println(string(body))
		}
	}
}