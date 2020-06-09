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

	"github.com/reujab/wallpaper"
)

// getQuery function returns API query type from URL
func getQuery(url *string) (query string, typeOfQuery string) {
	urlSplit := strings.Split(*url, "/")
	typeOfQuery = urlSplit[3]
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
	return query, typeOfQuery
}

// Make Request to Imgur
func makeQueryRequest(query string) (body string) {
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
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body = string(b)
	return body
}

// changeWallpaper on Desktop
func changeWallpaper(typeOfQuery string, body string, show *bool) {
	switch typeOfQuery {
	case "r":
		var r SubredditResult
		err := json.Unmarshal([]byte(string(body)), &r)
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
		err := json.Unmarshal([]byte(string(body)), &r)
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
		err := json.Unmarshal([]byte(string(body)), &r)
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
		err := json.Unmarshal([]byte(string(body)), &r)
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
