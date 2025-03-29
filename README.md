# Soundlib - music management system

![GitHub CI status](https://img.shields.io/github/actions/workflow/status/thuongtruong109/soundlib/ci.yml)
![Go report](https://goreportcard.com/badge/github.com/thuongtruong109/soundlib)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thuongtruong109/soundlib)
![GitHub License](https://img.shields.io/github/license/thuongtruong109/soundlib?color=orange)

## Preview

![Main Menu](/public/1.png)
![Read data](/public/2.png)

## Description

This is a music management system that allows you to interect and manage data through cli. Program is written in Go and uses JSON as store data format. Practicing and handle usecases such as: programing logic, database relation, file I/O, data structure, error handling, etc.

## ✨ Features

✅ Save/Load/Read/Write data from/to file

✅ Create/Get(all,one)/Update/Delete to Album/Artist/Genre/Playlist/Tracks

✅ Add/Get/Remove songs to playlists/albums

## Architecture

![ERD](/docs/erd.png)

## Usage

This program is designed to run on the command line interface (CLI)

For user: Download [Soundlib](https://github.com/thuongtruong109/soundlib/blob/main/soundlib) app
At the same dir, open Command Prompt (cmd) or Terminal and run:

```bash
soundlib
```

Then, you can use the following commands to interact with the program

For developer: Install Go on your computer at [Golang](https://golang.org/dl/). Then, clone this repository and run the following commands:

```bash
git clone https://github.com/thuongtruong109/soundlib.git
cd soundlib
go mod tidy
go run main.go
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE) License © 2023-PRESENT Tran Nguyen Thuong Truong

<!-- ## References

[Read x write JSON](https://www.developer.com/languages/json-files-golang/)

[Append to Json](https://dev.to/evilcel3ri/append-data-to-json-in-go-5gbj)

[symbol](https://www.cutesymbols.net/p/dot.html) -->
