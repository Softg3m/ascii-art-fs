# ASCII Art FS

A command-line application written in **Go** that converts input text into ASCII art using different banner templates. This project demonstrates file system handling by dynamically reading banner files.



##  Features

* Convert text into ASCII art
* Supports multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Handles:

  * Single argument (defaults to `standard`)
  * Two arguments (`STRING` + `BANNER`)
* Proper error handling for invalid inputs
* Supports multi-line input (`\n`)


## Usage

```bash
go run . [STRING] [BANNER]
```

### Examples

```bash
go run . "hello"
go run . "hello" standard
go run . "Hello There!" shadow
go run . "Hello\nWorld" thinkertoy
```


##  Invalid Usage

Any incorrect input format will display:

```bash
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```


##  Project Structure

```
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── go.mod
├── main_test.go
├── main.go
├── README.md
```

##  How It Works

1. The program reads command-line arguments.
2. Selects a banner (default is `standard`).
3. Loads the corresponding banner file from the `banners/` directory.
4. Maps each character in the input string to its ASCII representation.
5. Prints the final ASCII art output.

##  Technologies Used

* Go (Golang)
* Standard library only (`os`, `fmt`, `strings`)
