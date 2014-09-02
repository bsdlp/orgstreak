package main

import (
	"encoding/json"
	"flag"
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

const dateFormat string = "2006-01-02"

func (c *Contribution) UnmarshalJSON(data []byte) error {
	var i []interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	if v, ok := i[0].(string); ok {
		c.Date, err = time.Parse(dateFormat, v)
	}
	if err != nil {
		return err
	}

	if v, ok := i[1].(float64); ok {
		c.Num = int(v)
	}

	return nil
}

func getContributions(user github.User) []Contribution {
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

	return contribData
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
	flag.Parse()
	var i interface{} = flag.Arg(0)
	var users []github.User

	var v, ok = i.(string)
	if ok == false || v == "" {
		log.Fatalln("Usage: orgstreak <orgName>")
	} else {
		users = getOrgMembers(v)
	}

	orgContribs := make(map[time.Time]int)
	userContributions := make(chan []Contribution, len(users))

	for _, user := range users {
		go func(u github.User) {
			userContributions <- getContributions(u)
		}(user)
	}

	for y := 0; y < len(users); y++ {
		select {
		case userContribution := <-userContributions:
			for _, contrib := range userContribution {
				orgContribs[contrib.Date] += contrib.Num
			}
		}
	}

	contribData := []interface{}{}
	for k, v := range orgContribs {
		contribData = append(contribData, []interface{}{k.Format(dateFormat), v})
	}
	j, err := json.Marshal(contribData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)
}
