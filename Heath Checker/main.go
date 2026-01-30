// take input form cli and check if

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func ping(url string) string {

	res, err := http.Get(url)

	if err != nil {
		return fmt.Sprintf("Error checking %s: %v\n", url, err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return fmt.Sprintf("%s UP, %d\n", url, res.StatusCode)
	} else {
		return fmt.Sprintf("%s Down %d\n", url, res.StatusCode)
	}
}

func main() {
	arr := []string{}

	if len(os.Args) > 1 {

		for _, arg := range os.Args[1:] {
			if !strings.HasPrefix(arg, "https://") {
				arg = "https://" + arg
			}

			arr = append(arr, arg)
		}
	}

	start := time.Now()
	for _, arg := range arr {
		fmt.Print(ping(arg))
	}
	duration := time.Since(start)

	fmt.Printf("time: %v \n", duration)

	// parallal processing usecase

	ch1 := make(chan string)

	start = time.Now()
	for _, arg := range arr {
		go func(url string) {
			result := ping(url)
			ch1 <- result
		}(arg)
	}

	for range arr {
		fmt.Print(<-ch1)
	}

	duration = time.Since(start)
	fmt.Printf("time: %v\n", duration)

}
