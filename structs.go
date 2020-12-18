package main

// Result of Imgur API
type Result struct {
	Data []struct {
		Link string `json:"link"`
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
	Data []struct {
		Link string `json:"link"`
	} `json:"data"`
}

// RedditAPIResult of Reddit API
type RedditAPIResult struct {
	Data struct {
		Children []struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data,omitempty"`
		} `json:"children"`
	} `json:"data"`
}
