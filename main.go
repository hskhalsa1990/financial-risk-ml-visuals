package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func score(client *http.Client, base string, age, income, cards, loans int) (string, error) {
	q := url.Values{}
	q.Set("age", fmt.Sprint(age))
	q.Set("income", fmt.Sprint(income))
	q.Set("numberOfCard", fmt.Sprint(cards))
	q.Set("numberOfCarLoan", fmt.Sprint(loans))
	resp, err := client.Get(base + "/score?" + q.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	customers := [][4]int{{23, 30000, 2, 1}, {45, 60000, 1, 0}, {31, 18000, 3, 2}}
	var wg sync.WaitGroup
	for _, c := range customers {
		wg.Add(1)
		go func(c [4]int) {
			defer wg.Done()
			out, err := score(client, "http://127.0.0.1:5000", c[0], c[1], c[2], c[3])
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			fmt.Printf("%v -> %s", c, out)
		}(c)
	}
	wg.Wait()
}
