package main

import (
	"fmt"
	"time"
)

func Dashboard(metrics *Metrics) {

	for {

		p50, p95, p99 := metrics.Snapshot()

		fmt.Print("\033[H\033[2J") // clear screen

		fmt.Println("========= POS LOAD ENGINE V4 =========")
		fmt.Printf("Success: %d\n", metrics.success)
		fmt.Printf("Failed : %d\n", metrics.fail)
		fmt.Println("-------------------------------------")
		fmt.Printf("p50: %d ms\n", p50)
		fmt.Printf("p95: %d ms\n", p95)
		fmt.Printf("p99: %d ms\n", p99)
		fmt.Println("-------------------------------------")

		for k, v := range metrics.errors {
			fmt.Printf("%s: %d\n", k, v)
		}

		time.Sleep(1 * time.Second)
	}
}
