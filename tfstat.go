package main

import (
	"flag"
	"fmt"
	"github.com/Drpsycho/goquery"
	"html/template"
	"log"
	"os"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>TF</title>
	</head>
	<body>
		{{.}}!
	</body>
</html>`

var url = flag.String("url", "", "url for parse")

func main() {
	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	//"http://91.215.138.164:9000/"
	doc, err := goquery.NewDocument(*url)
	check(err)

	ret, err := doc.Find(".columns").Html()
	check(err)

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	f, err := os.Create("./hlstat.html")
	check(err)

	defer f.Close()
	err = t.Execute(f, template.HTML(ret))
	check(err)

	fmt.Println("all ok")
}
