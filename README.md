# omgbeaupeep.com-downloader
Download comics from https://www.omgbeaupeep.com

![beau-peep-logo](https://github.com/user-attachments/assets/77f159a4-4cfb-486f-be6b-0aeac57803d4)

## How to use
1. #### Get a valid issue URL and add it as an argument

&nbsp; `go run main.go <issueURL>`

&nbsp; Example: `go run main.go https://www.omgbeaupeep.com/comics/Adventurers_(1986)/01.001.01/`

<h2></h2>

2. #### Download all the issues of a comic

&nbsp; `go run main.go <comicURL>`

&nbsp; Example: `go run main.go https://www.omgbeaupeep.com/comics/Adventurers_(1986)/`

<h2></h2>

3. #### Convert images to PDF:
 
&nbsp;  `go run main.go --pdf <issueURL|comicURL>`

&nbsp; Example: `go run main.go --pdf https://www.omgbeaupeep.com/comics/Adventurers_(1986)/`

4. #### List all comics:

&nbsp;  `go run main.go --list-comics`


## Notes
- The files will be saved in the 'output' directory automatically.
- If the URL contains characters at the end such as parentheses, which can be misinterpreted by the terminal emulator it is recommended to end it with '/' or enclose it in single or double quotes.
- If you don't have a Go compiler you can use Nix or one of the released binaries.
