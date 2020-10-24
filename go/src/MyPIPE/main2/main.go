package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func ticker() {
	t := time.NewTicker(5 * time.Second) //1秒周期の ticker
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
			url := "http://localhost:8080/health"

			resp, _ := http.Get(url)

			byteArray, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(byteArray)) // htmlをstringで取得
			resp.Body.Close()
		}
	}
}

func main() {
	ticker()
}
