# walgur
:computer: Set your desktop background from randomly picked Imgur/Reddit links! :framed_picture:

<img src="https://goreportcard.com/badge/github.com/M4cs/walgur"> <a href="https://github.com/M4cs/walgur/stargazers"> <img alt="GitHub stars" src="https://img.shields.io/github/stars/M4cs/walgur"></a> <a href="https://github.com/M4cs/walgur/issues"> <img alt="GitHub issues" src="https://img.shields.io/github/issues/M4cs/walgur"></a>

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
usage: walgur [-h|--help] [-r|--reddit "<value>"] [-i|--imgur "<value>"]
              [-s|--show-bg] [-v|--version]

              Set your wallpaper randomly from Imgur Galleries, Albums, and
              Subreddits.

Arguments:

  -h  --help     Print help information
  -r  --reddit   Reddit Link. Default: None
  -i  --imgur    Imgur Link. Default: None
  -s  --show-bg  Show where background is stored. Default: false
  -v  --version  Display walgur Version. Default: false
```

#### Supports:

- Albums
- Tags
- Subreddits
- Galleries
- Reddit Subs
This will randomly pick an image from said URL and set it as your wallpaper.
