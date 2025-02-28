### `The site has changed to https://www.popsensei.com/request-a-comic-book-or-manga-omgbp/ and apparently offers the comics via email now, making this useless.` 

# omgbeaupeep.com-downloader
Download comics from https://www.omgbeaupeep.com

![beau-peep-logo](https://github.com/user-attachments/assets/77f159a4-4cfb-486f-be6b-0aeac57803d4)

## How to install

### Linux
 ``` bash
  git clone https://github.com/Calana2/omgbeaupeep.com-downloader
  cd omgbeaupeep.com-downloader
  go build -o omg-dl main.go && sudo cp omg-dl /usr/bin/omg-dl
```

### Windows
 ``` cmd
  :: As administrator
  git clone https://github.com/Calana2/omgbeaupeep.com-downloader
  cd omgbeaupeep.com-downloader
  go build -o omg-dl.exe main.go
  mkdir "%ProgramFiles"\omg-dl
  copy omg-dl.exe "%ProgramFiles%"\omg-dl
  setx PATH "%PATH%;%ProgramFiles%\omg-dl"
```

## How to use

- #### List all comics:

&nbsp;  `omg-dl --list-comics`

<h2></h2>

- #### List all issues of a comic:

&nbsp;  `omg-dl --list-issues <comicURL>`

<h2></h2>

- #### Download an issue

&nbsp; `omg-dl <issueURL>`

&nbsp; Example: `omg-dl https://www.omgbeaupeep.com/comics/Adventurers_(1986)/01.001.01/`

<h2></h2>

- #### Download all the issues of a comic

&nbsp; `omg-dl <comicURL>`

&nbsp; Example: `omg-dl https://www.omgbeaupeep.com/comics/Adventurers_(1986)/`

<h2></h2>

- #### Convert images to PDF:
 
&nbsp;  `omg-dl --pdf <issueURL|comicURL>`

&nbsp; Example: `omg-dl --pdf https://www.omgbeaupeep.com/comics/Adventurers_(1986)/`


## Notes
- The files will be saved in the 'output' directory automatically.
- If you don't have a Go compiler you can use Nix or one of the released binaries.
