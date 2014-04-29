serve
=====

A golang command line utility file server.

## Install

To run this tool first go to golang.org and download go.
Once you've set up the go compiler, you can type
```sh
go build serve.go
mv ./serve /usr/local/bin
```
in a terminal to compile `serve.go` to `serve` and move it to your `/usr/local/bin`.

To remove this tool, type
```sh
rm /usr/local/bin/serve
```
and be rid of it.

## Usage

Usage is very simple, open a termianl, cd into the directory you would like to serve, and run `serve`.
```sh
$ cd ~/dev/
$ serve
```
You can also optionally add a port number. __Note:__ you must be root to use port 80.
```sh
# optionally with port number
$ sudo serve 80
