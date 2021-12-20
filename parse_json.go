package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Id      uint
	Name    string
	Email   string
	Address Address
	Phone   string
	Website string
	Company Company
	Posts   []Post
}

type Company struct {
	Name        string
	CatchPhrase string
	Bs          string
}

type Address struct {
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     Geo
}

type Geo struct {
	Lat float32 `json:",string"`
	Lng float32 `json:",string"`
}

type Post struct {
	UserId uint
	Id     uint
	Title  string
	Body   string
}

func main() {
	// Urls slice
	urls := []string{
		"https://jsonplaceholder.typicode.com/users", "https://jsonplaceholder.typicode.com/posts",
	}

	var wg sync.WaitGroup

	content := load(urls)

	var users []User
	var posts []Post

	wg.Add(2)
	go decode(content[0], &users, &wg)
	go decode(content[1], &posts, &wg)
	wg.Wait()

	// to hash map
	usersHashMap := make(map[uint]User)
	for _, user := range users {
		usersHashMap[user.Id] = user
	}

	for _, post := range posts {
		tmpUser := usersHashMap[post.UserId]
		tmpUser.Posts = append(usersHashMap[post.UserId].Posts, post)
		usersHashMap[post.UserId] = tmpUser
	}

	fmt.Printf("%+v\n", usersHashMap)
}

func load(urls []string) []*http.Response {
	res := make([]*http.Response, len(urls))
	var wg sync.WaitGroup
	for k, url := range urls {
		wg.Add(1)
		go func(i int, urlToLoad string, res []*http.Response) {
			res[i], _ = http.Get(urlToLoad)
			wg.Done()
		}(k, url, res)
	}
	wg.Wait()
	return res
}

func decode(response *http.Response, toJsonStruct interface{}, group *sync.WaitGroup) {
	err := json.NewDecoder(response.Body).Decode(toJsonStruct)
	if err != nil {
		fmt.Println(err.Error())
	}
	group.Done()
	defer response.Body.Close()
}
