# walgur
Randomly set your desktop wallpaper from Imgur Galleries!

<img src="https://goreportcard.com/badge/github.com/M4cs/walgur">


<p align="center">
  <img src="https://raw.githubusercontent.com/M4cs/walgur/master/preview.gif">
</p>

## Installation

```bash
go get github.com/M4cs/walgur
go install github.com/M4cs/walgur
```

or build from source

```bash
git clone github.com/M4cs/walgur
cd walgur
go build
# Move the walgur binary to your $PATH
```

## Usage

```
walgur -u <imgur URL>
```

#### Supports:

- Albums
- Tags
- Subreddits
- Galleries

This will randomly pick an image from said URL and set it as your wallpaper.
