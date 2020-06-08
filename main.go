package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/reujab/wallpaper"
)

var (
	fileName string
	url      string
)

type Result struct {
	Data struct {
		Images []struct {
			Tag string `json:"id"`
		} `json:"images"`
	} `json:"data"`
}

func main() {
	var args []string = os.Args
	if len(args) <= 1 {
		fmt.Println("Please Add An Imgur Gallery Hash: wallgur.exe {hash}")
		os.Exit(1)
	}
	var hash string = args[1]
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.imgur.com/3/gallery/album/"+hash, nil)
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
	var r Result
	err = json.Unmarshal([]byte(string(body)), &r)
	if err != nil {
		fmt.Println("Unable to decode JSON from Imgur. Please try again!")
	}
	rand.Seed(time.Now().Unix())
	name := fmt.Sprint(r.Data.Images[rand.Intn(len(r.Data.Images))].Tag, ".jpg")
	url := fmt.Sprint("https://imgur.com/", name)
	err = wallpaper.SetFromURL(url)
	if err != nil {
		fmt.Println("Unable to set wallpaper. Please try again or it may not work with your OS.")
	}
}
