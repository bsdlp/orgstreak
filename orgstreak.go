package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/go-github/github"
)

type contribCalendar struct {
	XMLName xml.Name       `xml:"svg"`
	Group   containerGroup `xml:"g"`
}

type containerGroup struct {
	Groups []contribGroup `xml:"g"`
}

type contribGroup struct {
	Rects []Contribution `xml:"rect"`
}

type Contribution struct {
	Date time.Time `xml:"data-date,attr"`
	Num  int       `xml:"data-count,attr"`
}

const dateFormat string = "2006-01-02"

func (c *Contribution) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i interface{}
	err := d.DecodeElement(&i, &start)
	if err != nil {
		return err
	}

	c.Date, err = time.Parse(dateFormat, i.date)
	if err != nil {
		return err
	}

	c.Num, err = strconv.Atoi(i.count)
	if err != nil {
		return err
	}

	if v, ok := i.date.(string); ok {
		c.Date, err = time.Parse(dateFormat, v)
	}
	if err != nil {
		return err
	}

	if v, ok := i.count.(string); ok {
		c.Num, err = strconv.Atoi(v)
	}
	if err != nil {
		return err
	}

	return nil
}

func getContributions(user github.User) []Contribution {
	login := github.Stringify(user.Login)
	url := "https://github.com/users/" + login[1:len(login)-1] + "/contributions"
	var contribData contribCalendar

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

	err = xml.Unmarshal(data, &contribData)
	if err != nil {
		log.Fatal(err)
	}
	var contributions []Contribution

	for _, group := range contribData.Group.Groups {
		contributions = append(contributions, group.Rects...)
	}

	return contributions
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
		fmt.Println("Usage: orgstreak <orgName>")
		os.Exit(1)
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
	fmt.Printf("%s\n", j)
}
