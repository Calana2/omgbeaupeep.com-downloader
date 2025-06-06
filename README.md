# popsensei-scraper

![2025-06-06-171225_422x123_scrot](https://github.com/user-attachments/assets/56b93862-070f-4d25-a382-d26ae53b114a)

A web scraper to download comics from [popsensei](https://www.popsensei.com/request-a-comic-book/) ( previously [omgbeaupeep](https://www.omgbeaupeep.com/))

## Build and Install

### Linux
 ``` 
  git clone https://github.com/Calana2/popsensei-scraper
  cd popsensei-scraper
  go build -o popscrap main.go && sudo cp omg-dl /usr/bin/popscrap
```

### Others
There are a couple of binaries for Linux and Windows amd64 in the releases.
<br/><br/>

## Usage
```
└─$ popscrap
A tool to download comics from https://www.popsensei.com/request-a-comic-book/

Usage: popscrap [OPTIONS] [comicURL|issueURL]

Options:
-h, --help  Print help
-lc, --list-comics  List all available comics
-li, --list-issues [issueURL]  List all issues of a comic
--pdf Convert issues to pdf
```



