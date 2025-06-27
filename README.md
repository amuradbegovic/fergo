# fergo - a [_šatrovački_](https://en.wikipedia.org/wiki/%C5%A0atrova%C4%8Dki) Gopher server 

*fergo* is a Gopher server written in the Go programming language. It was my high school senior year graduation project. 

*fergo* lets you share files from a directory over Gopher by automatically generating menus based on file type and location. It also allows the administrator to create custom menus using a user-friendly syntax (like [geomyidae](http://r-36.net/scm/geomyidae/file/README.html)). It provides its user with the typical configuration options, like choosing which TCP port to listen to, directory to serve content from and so on...

This repository used to be (besides variable, function and type names) completely in Bosnian language which was required since it was a final project in a Bosnian high school. I've decided to translate it to English because I want to continue improving the program, and share the improvements I make with international friends. Old commit messages will stay as they were, since I don't think anyone cares for them anyway.

## Building

To build *fergo*, clone the repository and run `go build`. Alternatively, you can install the program by running: 

`go install github.com/amuradbegovic/fergo@latest`

Of course, it is assumed that you have Go 1.20 or newer installed on your machine. You can download the Go for most major operating systems from [go.dev](https://go.dev).

I've tested building and running *fergo* on Linux, Microsoft Windows and OpenBSD. Some day, I'll probably test it on Plan 9 too...   

## Running

If you want to run *fergo* with default settings; i.e. set the current working directory as server's root and listen to port 70, simply run the executable. You may encounter a `bind: permission denied` error. That means you don't have the permissions to open port 70. Either run the program as a user with proper permissions or specify a different port number (ports numbered between 1024 and 49151 should be accessible to unprivileged users).

To set a different port number, use the `-p` option like so: `fergo -p 7070` (sets the port number to 7070).

To serve content from a directory other than the current one, use the `-d` option like so: `fergo -d /path/to/dir`.

To see the complete (but short) list of command line options, run *fergo* with `-h`.

I'll probably create a systemd service for this.

## Security

As stated before, this is my amateur project from high school and there is much to improve. I don't guarantee anything security-wise, but the server sends an error in case client sends a selector containing ".." in an attempt to access parent directory of the one served from. Ideally, you should ru n this program as a special user that has access only to the directory that you want to share and the port that you want to listen to. I don't plan to support TLS as that's beyond the scope of this. I'll probably write a Dockerfile to support running this in a container.

## License

BSD 2-Clause