serve
=====

A golang command line utility file server.

## Install

You must have the go tool installed to compile serve.

```sh
go build serve.go
mv ./serve /usr/local/bin
```

## Usage

Usage is very simple, open a terminal, cd into the directory you would like to serve, and run `serve`.
```sh
# serve current directory at :8080
$ serve
```
You can also optionally add a port number. __Note:__ you must be root to use port 80.
```sh
# optionally with port number
$ sudo serve 80
```
There are also other options, -p for port number, -d for directory, and -s for an address sub-dir.
```sh
# serve ~/go at :8080/h/
$ serve -p 8080 -d ~/go -s /h/
```

## License

If you've never seen an MIT license, check out the LICENSE.

## Authors

+ Connor Taffe (@cptaffe)
