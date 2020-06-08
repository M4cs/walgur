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

// Authorization: Client-ID ec55a6eb90d209b

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
		panic("Please Add An Imgur Gallery Hash: wallgur.exe {hash}")
	}
	var hash string = args[1]
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.imgur.com/3/gallery/album/"+hash, nil)
	req.Header.Set("Authorization", "Client-ID ec55a6eb90d209b")
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var r Result
	json.Unmarshal([]byte(string(body)), &r)
	rand.Seed(time.Now().Unix())
	name := fmt.Sprint(r.Data.Images[rand.Intn(len(r.Data.Images))].Tag, ".jpg")
	url := fmt.Sprint("https://imgur.com/", name)
	wallpaper.SetFromURL(url)
}
