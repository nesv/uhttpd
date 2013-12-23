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
$ uhttpd -dir=/home/user/www -addr=":8080"
```

## Installing uhttpd

If you have the [Go](http://golang.org) toolchain installed, you can simply
run:

	$ go get github.com/nesv/uhttpd


## License

MIT, baby.
