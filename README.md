# omgbeaupeep.com-downloader
Download comics from https://www.omgbeaupeep.com

## How to use
1. Find the Comic Book and the Issue in the link.

![test1](https://github.com/user-attachments/assets/bffb32ea-083f-4383-9fbb-ed147c3097fc)

2. Replace the ComicRoute var with this data in `main.go`.

``` go
// main.go
package main

import (
 "os"
 "scraper/scraper"
)

func main() { 
  ComicRoute := "/Avatar_The_Last_Airbender/001/"
  if len(os.Args) == 2 {
    ComicRoute = os.Args[1]
  }

  scraper.DownloadComic(ComicRoute)
}
```

3. Run `go run main.go`. The result will be saved in ./output/\<comic_book\>/\<issue\>/ folder.

``` 
> go run main.go 
Creating directory: ./output/Avatar_The_Last_Airbender/001/
Starting task: Download /Avatar_The_Last_Airbender/001/
Image downloaded: /comics/mangas/Avatar The Last Airbender/001 - Avatar The Last Airbender - The Promise Part 1 (2012)/read-avatar-the-last-airbender-comics-online-free-001.jpg
Image downloaded: /comics/mangas/Avatar The Last Airbender/001 - Avatar The Last Airbender - The Promise Part 1 (2012)/read-avatar-the-last-airbender-comics-online-free-002.jpg
Image downloaded: /comics/mangas/Avatar The Last Airbender/001 - Avatar The Last Airbender - The Promise Part 1 (2012)/read-avatar-the-last-airbender-comics-online-free-003.jpg
Image downloaded: /comics/mangas/Avatar The Last Airbender/001 - Avatar The Last Airbender - The Promise Part 1 (2012)/read-avatar-the-last-airbender-comics-online-free-004.jpg
Image downloaded: /comics/mangas/Avatar The Last Airbender/001 - Avatar The Last Airbender - The Promise Part 1 (2012)/read-avatar-the-last-airbender-comics-online-free-005.jpg
```

You can also just pass the data as a command-line argumment, for example: `go run main.go /Avatar_The_Last_Airbender/001/`.
