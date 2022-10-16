

<br />



<h3 align="center">
  <img src="https://user-images.githubusercontent.com/58083563/193971505-32402d48-42cc-4ac2-be84-9f0a7a470784.svg" alt="Go" width="100" />
</h3><br />
<h3 align="center">Golang CLI</h3>
<br>
<p align="center">
  <a href="#wrench-Configuration">Flags</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#gift-Others-commands">Others</a>&nbsp;&nbsp;&nbsp;
</p>


## :wrench: Flags


```
Usage: CLI Template [OPTIONS]

Options:
	-u, --user get all repositories, to export as csv use -u [USER]).
	-h, --help print all cli options.
	-e, --export export as csv
	-w, --number of workers to export concurrently
```

#### 1. Flags availble on CLI

```sh
# Print all github repositories of user
$ go run main.go -u "dannielss"

# Export all repositories as csv (sequentially)
$ go run main.go -u "dannielss" -e

# Export all repositories as multiples csv (concurrently)
$ go run main.go -u "dannielss" -e -w=4

# Print help on terminal
$ go run main.go -h
```


## :gift: Others commands

```sh
# Build application
$ go build
