package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
)

type Contribution struct {
	Date time.Time
	Num  int
}

func (c *Contribution) UnmarshalJSON(data []byte) error {
	var i []interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	c.Date, err = time.Parse("2006-01-02", i[0].(string))
	if err != nil {
		return err
	}

	if v, ok := i[1].(float64); ok {
		c.Name = int(v)
	}

	return nil
}

func getContributions(user string, userData chan []Contribution) {
	url := "https://github.com/users/" + user + "/contributions"
	var contribData []Contribution

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if data != nil {
		fmt.Println(string(data[:len(data)]))
		err = json.Unmarshal(data, &contribData)
	}

	if err != nil {
		log.Fatal(err)
	}

	userData <- contribData
	return
}

func getOrgMembers(orgName string) []github.User {
	client := github.NewClient(nil)

	users, _, err := client.Organizations.ListMembers(orgName, nil)
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func main() {
	userData := make(chan []Contribution)
	go getContributions("fly", userData)
	data := <-userData

	for x := range data {
		fmt.Println(data[x])
	}
}
