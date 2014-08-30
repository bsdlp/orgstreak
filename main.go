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

	if v, ok := i[0].(string); ok {
		c.Date, err = time.Parse("2006-01-02", v)
	}
	if err != nil {
		return err
	}

	if v, ok := i[1].(float64); ok {
		c.Num = int(v)
	}

	return nil
}

func getContributions(user github.User, userData chan<- []Contribution) {
	login := github.Stringify(user.Login)
	url := "https://github.com/users/" + login[1:len(login)-1] + "/contributions"
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
	if data == nil {
		log.Fatalln("No data returned.")
	}

	err = json.Unmarshal(data, &contribData)
	if err != nil {
		log.Fatal(err)
	}

	userData <- contribData
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
	users := getOrgMembers("linode")
	userData := make(chan []Contribution, len(users))
	for _, user := range users {
		go getContributions(user, userData)
	}

	for contribution := range userData {
		fmt.Println(contribution)
	}
}
