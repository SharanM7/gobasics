package main

import (
	"log"
	"sync"
	"time"
)

type User struct {
	Id          int
	Name        string
	Connections []int
	Followers   int
}

type Response struct {
	data any
	err  error
}

func main() {
	start := time.Now()

	user, err := getUser(10)
	if err != nil {
		log.Fatal("err >>", err)
	}
	log.Println("got user >>", user)
	log.Println("time took ", time.Since(start))
}

func getUser(id int) (User, error) {
	respch := make(chan Response, 4)
	var wg sync.WaitGroup

	wg.Add(3)
	go getName(id, respch, &wg)
	go getConnections(id, respch, &wg)
	go getFollowers(id, respch, &wg)
	wg.Wait()

	close(respch)
	var user User
	for res := range respch {
		if res.err != nil {
			return User{}, nil
		}
		switch d := res.data.(type) {
		case int:
			user.Followers = d
		case string:
			user.Name = d
		case []int:
			user.Connections = d
		}
	}

	return user, nil
}

func getName(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: "abc",
		err:  nil,
	}
}

func getConnections(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: []int{11, 22, 45},
		err:  nil,
	}
}

func getFollowers(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: 20,
		err:  nil,
	}
}
