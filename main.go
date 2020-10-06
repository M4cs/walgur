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
	reddit := parser.String("r", "reddit", &argparse.Options{Required: false, Help: "Reddit Link", Default: "None"})
	imgur := parser.String("i", "imgur", &argparse.Options{Required: false, Help: "Imgur Link", Default: "None"})
	var show *bool = parser.Flag("s", "show-bg", &argparse.Options{Required: false, Help: "Show where background is stored", Default: false})
	var ver *bool = parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display walgur Version", Default: false})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *ver {
		fmt.Println("walgur v1.3.0")
	}

	if *reddit != "None" && *imgur == "None" {
		//Reddit
		body := makeQueryRequest(*reddit, "reddit")
		changeWallpaper("reddit", body, show)
	} else if *imgur != "None" && *reddit == "None" {

		query, typeOfQuery := getQuery(*imgur)
		body := makeQueryRequest(query, "imgur")
		changeWallpaper(typeOfQuery, body, show)
	}
}
