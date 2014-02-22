# uhttpd

A laughably-small HTTP server.

The only thing *uhttpd* is good for, is serving static content. It has no fancy
bells, or whistles, like virtual host, or CGI support. All it does is host
static files out of a directory for you.

## Where would this be useful?

If you do not want to set up Apache, nginx, or any other web server, just to
serve up some local content you are working on, or if you do not want to have
to pay for a VPS to (S)FTP your files to, as you are working on them, then
*uhttpd* would be perfect for you!

## Running uhttpd

In the name of simplicity, *uhttpd* does not require any configuration files!
Any configurable values are passed to *uhttpd* as command-line flags. For
example, the following will serve up the `www` folder in your home directory,
and can be browsed by going to `http://localhost:8080`:

```
$ uhttpd -dir=/home/user/www -addr="127.0.0.1:8080"
```

### Command-line flags

If you run *uhttpd* with the `-help` flag, you will get the following output:

	Usage of uhttpd:
	  -addr=":80": The ADDRESS:PORT to listen on
	  -dir="": The directory to serve up
	  -log=false: Enable/disable logging
	  -log-path="": Log to file (leave blank for STDOUT)
	  -log-prefix="uhttpd": Set the logging prefix

`-addr="..."` lets you set the IPv4 address and port *uhttp* listens on.

`-dir="..."` tells *uhttpd* which directory to serve files out of.

`-log` tells *uhttpd* to log requests for files.

`-log-path="..."` is an optional flag, and lets you specify the name of a file
to write your request logs to. If you do not specify a log file, the request
logs are written to standard output (STDOUT).

`-log-prefix="..."` is another optional flag, that lets you specify the logging
prefix; by default, the logging prefix is "uhttpd".

## Installing uhttpd

If you have the [Go](http://golang.org) toolchain installed, you can simply
run:

	$ go get github.com/nesv/uhttpd


## License

MIT, baby.
