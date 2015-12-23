package main

import (
	"flag"
	"fmt"
	"github.com/Drpsycho/goquery"
	"html/template"
	"log"
	"net/http"
	"os"
)

var url = flag.String("url", "", "url for parse")
var templ = flag.String("t", "./templ.html", "template file (.html)")
var outputname = flag.String("o", "./tfstat.html", "output file name")

type Player struct {
	Rank        string
	Name        string
	Points      string
	Time_online string
	Kills       string
	Death       string
	Kd          string
	Headshot    string
	Accuracy    string
}

func check(err error) {
	if err != nil {
		log.Println("Error!!!")
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	doc, err := goquery.NewDocument(*url)
	check(err)

	var bar []Player

	doc.Find(".data-table").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(j int, tr *goquery.Selection) {
			var foo Player
			tr.Find("td").Each(func(k int, td *goquery.Selection) {
				switch k + 1 {
				case 1:
					foo.Rank = td.Text()
				case 2:
					foo.Name = td.Text()
				case 3:
					foo.Points = td.Text()
				case 5:
					foo.Time_online = td.Text()
				case 6:
					foo.Kills = td.Text()
				case 7:
					foo.Death = td.Text()
				case 8:
					foo.Kd = td.Text()
				case 9:
					foo.Headshot = td.Text()
				case 10:
					foo.Accuracy = td.Text()
				}
			})
			bar = append(bar, foo)
		})
	})

	t, err := template.ParseFiles(*templ)
	check(err)

	err = t.Execute(w, bar)
	check(err)
}

func main() {
	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	go func() {

		var inputs string
		for {
			fmt.Scanln(&inputs)
			fmt.Println("For quit enter 'q'")
			if inputs == "q" {
				fmt.Println("quit")
				os.Exit(0)
			}
		}
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9009", nil)
}
