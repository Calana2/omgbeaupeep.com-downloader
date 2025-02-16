# omgbeaupeep.com-downloader
Download comics from https://www.omgbeaupeep.com

![beau-peep-logo](https://github.com/user-attachments/assets/77f159a4-4cfb-486f-be6b-0aeac57803d4)

## How to install

### Linux
 ``` bash
  git clone https://github.com/Calana2/omgbeaupeep.com-downloader
  cd omgbeaupeep.com-downloader
  go build -o omgb main.go && sudo cp omgb /usr/bin/omgb
```

### Windows
 ``` cmd
  :: As administrator
  git clone https://github.com/Calana2/omgbeaupeep.com-downloader
  cd omgbeaupeep.com-downloader
  go build -o omgb.exe main.go
  mkdir "%ProgramFiles"\omgb
  copy omgb.exe "%ProgramFiles%"\omgb
  setx PATH "%PATH%;%ProgramFiles%\omgb"
```

## How to use

- #### List all comics:

&nbsp;  `omgb --list-comics`

<h2></h2>

- #### List all issues of a comic:

&nbsp;  `omgb --list-issues <comicURL>`

<h2></h2>

- #### Download an issue

&nbsp; `omgb <issueURL>`

&nbsp; Example: `omgb https://www.omgbeaupeep.com/comics/Adventurers_(1986)/01.001.01/`

<h2></h2>

- #### Download all the issues of a comic

&nbsp; `omgb <comicURL>`

&nbsp; Example: `omgb https://www.omgbeaupeep.com/comics/Adventurers_(1986)/`

<h2></h2>

- #### Convert images to PDF:
 
&nbsp;  `omgb --pdf <issueURL|comicURL>`

&nbsp; Example: `omgb --pdf https://www.omgbeaupeep.com/comics/Adventurers_(1986)/`


## Notes
- The files will be saved in the 'output' directory automatically.
- If you don't have a Go compiler you can use Nix or one of the released binaries.
