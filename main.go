package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
	"github.com/reujab/wallpaper"
)

var (
	fileName string
	url      string
)

// Result of Imgur API
type Result struct {
	Data struct {
		Images []struct {
			Tag string `json:"id"`
		} `json:"images"`
	} `json:"data"`
}

// SubredditResult of Imgur API
type SubredditResult struct {
	Data []struct {
		Tag string `json:"id"`
	} `json:"data"`
}

// TagResult of Imgur API
type TagResult struct {
	Data struct {
		Images []struct {
			Tag string `json:"id"`
		} `json:"items"`
	} `json:"data"`
}

func main() {
	parser := argparse.NewParser("walgur", "Set your wallpaper randomly from Imgur Galleries, Albums, and Subreddits.")
	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "Imgur URL to grab from."})
	var show *bool = parser.Flag("s", "show-bg", &argparse.Options{Required: false, Help: "Show where background is stored", Default: false})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	urlSplit := strings.Split(*url, "/")
	typeOfQuery := urlSplit[3]
	var query string
	switch typeOfQuery {
	case "t":
		query = "gallery/t/" + urlSplit[4] + "/top/all"
		break
	case "gallery":
		query = "gallery/album/" + urlSplit[4]
		break
	case "r":
		query = "gallery/r/" + urlSplit[4] + "/top/all"
		break
	case "album":
		query = "album/" + urlSplit[4] + "/images"
		break
	default:
		fmt.Println("Your URL seems to be invalid. Please enter a subreddit, gallery, or tag URL.")
		os.Exit(1)
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.imgur.com/3/"+query, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Client-ID ec55a6eb90d209b")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if res.StatusCode == 404 {
		fmt.Println("Error: Gallery Not Found")
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch typeOfQuery {
	case "r":
		var r SubredditResult
		err = json.Unmarshal([]byte(string(body)), &r)
		rand.Seed(time.Now().Unix())
		name := fmt.Sprint(r.Data[rand.Intn(len(r.Data))].Tag, ".jpg")
		imageURL := fmt.Sprint("https://imgur.com/", name)
		err = wallpaper.SetFromURL(imageURL)
		if err != nil {
			fmt.Println("Unable To Set Image from URL:", err)
		}
		if *show {
			background, err := wallpaper.Get()
			if err != nil {
				fmt.Println("Couldn't grab background!")
				os.Exit(1)
			}
			fmt.Println("Background Stored At:", background)
		}
		break
	case "t":
		var r TagResult
		err = json.Unmarshal([]byte(string(body)), &r)
		rand.Seed(time.Now().Unix())
		name := fmt.Sprint(r.Data.Images[rand.Intn(len(r.Data.Images))].Tag, ".jpg")
		imageURL := fmt.Sprint("https://imgur.com/", name)
		err = wallpaper.SetFromURL(imageURL)
		if err != nil {
			fmt.Println("Unable To Set Image from URL:", err)
		}
		if *show {
			background, err := wallpaper.Get()
			if err != nil {
				fmt.Println("Couldn't grab background!")
				os.Exit(1)
			}
			fmt.Println("Background Stored At:", background)
		}
		break
	case "gallery":
		var r Result
		err = json.Unmarshal([]byte(string(body)), &r)
		if err != nil {
			fmt.Println("Unable to decode JSON from Imgur. Please try again!")
		}
		rand.Seed(time.Now().Unix())
		name := fmt.Sprint(r.Data.Images[rand.Intn(len(r.Data.Images))].Tag, ".jpg")
		imageURL := fmt.Sprint("https://imgur.com/", name)
		err = wallpaper.SetFromURL(imageURL)
		if err != nil {
			fmt.Println("Unable to set wallpaper. Please try again or it may not work with your OS.")
		}
		if *show {
			background, err := wallpaper.Get()
			if err != nil {
				fmt.Println("Couldn't grab background!")
				os.Exit(1)
			}
			fmt.Println("Background Stored At:", background)
		}
		break
	case "album":
		var r Result
		err = json.Unmarshal([]byte(string(body)), &r)
		if err != nil {
			fmt.Println("Unable to decode JSON from Imgur. Please try again!")
		}
		rand.Seed(time.Now().Unix())
		name := fmt.Sprint(r.Data.Images[rand.Intn(len(r.Data.Images))].Tag, ".jpg")
		imageURL := fmt.Sprint("https://imgur.com/", name)
		err = wallpaper.SetFromURL(imageURL)
		if err != nil {
			fmt.Println("Unable to set wallpaper. Please try again or it may not work with your OS.")
		}
		if *show {
			background, err := wallpaper.Get()
			if err != nil {
				fmt.Println("Couldn't grab background!")
				os.Exit(1)
			}
			fmt.Println("Background Stored At:", background)
		}
		break
	default:
		fmt.Println("I don't know how you got to this point in all honesty.")
		os.Exit(1)
	}

}
