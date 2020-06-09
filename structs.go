package main

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
