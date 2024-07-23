package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Wali Yar khan"
}

func fetchLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11
	wg.Done()
}

func fetchMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- "Maryam"
	wg.Done()
}

func main() {
	start := time.Now()

	userName := fetchUser()

	respch := make(chan any, 3)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchLikes(userName, respch, wg)
	go fetchMatch(userName, respch, wg)

	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println("response: ", resp)
	}

	fmt.Println("took: ", time.Since(start))
}
