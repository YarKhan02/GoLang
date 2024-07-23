package main

import (
	"fmt"
	"sync"
	"time"
)

type App interface {
	fetchUser()
	fetchLikes()
	fetchMatch()
}

type data struct {
	user_name string
	likes     int
	match     string
}

func CreateData(name string, likes int, match string) *data {
	return &data{
		user_name: name,
		likes:     likes,
		match:     match,
	}
}

func (d *data) fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return d.user_name
}

func (d *data) fetchLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respch <- d.likes
	wg.Done()
}

func (d *data) fetchMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- d.match
	wg.Done()
}

func main() {
	start := time.Now()

	person := CreateData("wali", 11, "maryam")

	userName := person.fetchUser()

	respch := make(chan any, 3)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go person.fetchLikes(userName, respch, wg)
	go person.fetchMatch(userName, respch, wg)

	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println("response: ", resp)
	}

	fmt.Println("took: ", time.Since(start))
}
