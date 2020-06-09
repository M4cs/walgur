package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var (
	fileName string
	url      string
)

func main() {
	parser := argparse.NewParser("walgur", "Set your wallpaper randomly from Imgur Galleries, Albums, and Subreddits.")
	url := parser.String("u", "url", &argparse.Options{Required: false, Help: "Imgur URL to grab from.", Default: "Empty"})
	var show *bool = parser.Flag("s", "show-bg", &argparse.Options{Required: false, Help: "Show where background is stored", Default: false})
	var ver *bool = parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display walgur Version", Default: false})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	if *ver {
		fmt.Println("walgur v1.3.0")
	}
	if *url != "Empty" {
		query, typeOfQuery := getQuery(url)
		body := makeQueryRequest(query)
		changeWallpaper(typeOfQuery, body, show)
	}
}
