package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start to call on D")
	fmt.Fprintf(w, "Calling Service D")
	fmt.Println("打印Header参数列表：")
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	req, err := http.NewRequest("GET", "http://service-d:80/", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-B3-Spanid", r.Header.Get("X-B3-Spanid"))
	req.Header.Add("X-Request-Id", r.Header.Get("X-Request-Id"))
	req.Header.Add("X-B3-Traceid", r.Header.Get("X-B3-Traceid"))
	req.Header.Add("X-B3-Sampled", r.Header.Get("X-B3-Sampled"))
	req.Header.Add("X-B3-Parentspanid", r.Header.Get("Upstream-Spanid"))
	req.Header.Add("Upstream-Spanid", r.Header.Get("X-B3-Spanid"))
	fmt.Println("my request id: ", r.Header.Get("X-Request-Id"), ";", "my span id:", r.Header.Get("X-B3-Spanid"))
	fmt.Println("发出的Header参数列表：")
	if len(req.Header) > 0 {
		for k, v := range req.Header {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, string(body))
	fmt.Fprintf(w, "Hello from service C")

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
