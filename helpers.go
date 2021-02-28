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
func getQuery(url string) (query string, typeOfQuery string) {
	urlSplit := strings.Split(url, "/")
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
	case "a":
		query = "album/" + urlSplit[4] + "/images"
		break
	default:
		fmt.Println("Your URL seems to be invalid. Please enter a subreddit, gallery, or tag URL.")
		os.Exit(1)
	}
	return query, typeOfQuery
}

// Make Request to Imgur
func makeQueryRequest(query string, website string) (body string) {
	client := &http.Client{}
	if website == "imgur" {
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
	} else if website == "reddit" {
		if query[len(query)-1:] == "/" {
			query = query + ".json"
		} else {
			query = query + "/.json"
		}
		req, err := http.NewRequest("GET", query, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if res.StatusCode == 404 {
			fmt.Println("Error: Reddit Not Found")
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		body = string(b)
	}
	return body
}

// changeWallpaper on Desktop
func changeWallpaper(typeOfQuery string, body string, show *bool) {
	var (
		name string
		err  error
	)
	switch typeOfQuery {
	case "r":
		var r SubredditResult
		err = json.Unmarshal([]byte(string(body)), &r)
		rand.Seed(time.Now().Unix())
		name = fmt.Sprintf("https://imgur.com/%s", fmt.Sprint(r.Data[rand.Intn(len(r.Data))].Tag, ".jpg"))
		break
	case "t":
		var r TagResult
		err = json.Unmarshal([]byte(string(body)), &r)
		rand.Seed(time.Now().Unix())
		name = r.Data[rand.Intn(len(r.Data))].Link
		err = wallpaper.SetFromURL(name)
		break
	case "gallery":
		var r Result
		err = json.Unmarshal([]byte(string(body)), &r)
		if err != nil {
			fmt.Println("Unable to decode JSON from Imgur. Please try again!")
		}
		rand.Seed(time.Now().Unix())
		name = r.Data[rand.Intn(len(r.Data))].Link
		break
	case "a":
		var r Result
		err = json.Unmarshal([]byte(string(body)), &r)
		if err != nil {
			fmt.Println("Unable to decode JSON from Imgur. Please try again!")
		}
		rand.Seed(time.Now().Unix())
		name = r.Data[rand.Intn(len(r.Data))].Link
		break
	case "album":
		var r Result
		err = json.Unmarshal([]byte(string(body)), &r)
		if err != nil {
			fmt.Println("Unable to decode JSON from Imgur. Please try again!")
		}
		rand.Seed(time.Now().Unix())
		name = r.Data[rand.Intn(len(r.Data))].Link
		break
	case "reddit":
		var r RedditAPIResult
		err = json.Unmarshal([]byte(string(body)), &r)
		if err != nil {
			fmt.Println("Unable to decode JSON from Reddit. Please try again!")
		}
		maxLen := len(r.Data.Children)
		rand.Seed(time.Now().Unix())
		rand.Intn(maxLen)
		fileTypes := [2]string{"jpg", "png"}
		image := false
		for !image {
			name = r.Data.Children[rand.Intn(maxLen)].Data.URL
			ending := name[len(name)-3:]
			for _, value := range fileTypes {
				if ending == value {
					image = true

				}
			}
		}
		break
	default:
		fmt.Println("I don't know how you got to this point in all honesty.")
		os.Exit(1)
	}
	err = wallpaper.SetFromURL(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable To Set Image from URL:%s", err)
	}
	if *show {
		background, err := wallpaper.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't grab background")
			os.Exit(1)
		}
		fmt.Println(background)
	}
}
