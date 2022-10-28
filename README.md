
Lemonade
========

forked from lemonade-command/lemonade

## **Why another lemonade fork?**

lemonade-command/lemonade over TCP.

yaocccc/lemonade over HTTP.

**Sometimes, HTTP is easier to use than TCP.**

------------

remote...lemote...lemode......Lemonade!!! :lemon: :lemon:

Lemonade is a remote utility tool.
(copy, paste and open browser) over TCP.

[![Build Status](https://travis-ci.org/lemonade-command/lemonade.svg?branch=master)](https://travis-ci.org/lemonade-command/lemonade)

Maintainers wanted
==========

See https://github.com/lemonade-command/lemonade/issues/25


Installation
------------

```sh
go get -d github.com/lemonade-command/lemonade
cd $GOPATH/src/github.com/lemonade-command/lemonade/
make install
```

Or download from [latest release](https://github.com/lemonade-command/lemonade/releases/latest)


Example of use
----------------

![Example](http://f.st-hatena.com/images/fotolife/P/Pocke/20150823/20150823173041.gif)

For example, you use a Linux as a virtual machine on Windows host.
You connect to Linux by SSH client(e.g. PuTTY).
When you want to copy text of a file on Linux to Windows, what do you do?
One solution is doing `cat file.txt` and drag displayed text.
But this answer is NOT elegant! Because your hand leaves from the keyboard to use the mouse.

Another solution is using the Lemonade.
You input `cat file.txt | lemonade copy`. Then, lemonade copies text of the file to clipboard of the Windows!

In addition to the above, lemonade supports pasting and opening URL.


Usage
--------

```sh
Usage: lemonade [options]... SUB_COMMAND [arg]
Sub Commands:
  copy [text]                 Copy text.
  paste                       Paste text.
  server                      Start lemonade server.

Options:
  --port=2489                 TCP port number
  --line-ending               Convert Line Ending(CR/CRLF)
  --host="localhost"          Destination hostname          [Client only]
  --help                      Show this message
```


### On server (in the above, Windows)

```sh
$ lemonade server
```


### Client (in the above, Linux)


```sh
# You want to copy a text
$ cat file.txt | lemonade copy

# You want to paste a text from the clipboard of Windows
$ lemonade paste

# You want to open an URL to a browser on Windows.
$ lemonade open 'http://google.com'
```


Configuration
---------------

You can override command line options by configuration file.
There is configuration file at `~/.config/lemonade.toml`.

### Server

```toml
port = 1234
line-ending = 'crlf'
```

- `port` is a listening port of TCP.


### Client

```toml
port = 1234
host = '192.168.x.x'
line-ending = 'crlf'
```

- `port` is a port of server.
- `host` is a hostname of server.

Detail of `trans-loopback` and `trans-localfile` are described Advanced Usage.


Advanced Usage
-----------------

### line-ending

Default: "" (NONE)

This options works with `copy` and `paste` command only.

If this option is `lf` or `crlf`, lemonade converts the line ending of text to the specified.


### Alias

You can use lemonade as a `xdg-open`, `pbcopy` and `pbpaste`.


For example.

```sh
$ ln -s /path/to/lemonade /usr/bin/xdg-open
$ xdg-open  'http://example.com' # Same as lemonade open 'http://example.com'
```


### Secure TCP Connection

Lemonade doesn't provide encryption and authorization.
However we can use `SSH Port forwarding` for the purpose.

lemonade server

```sh
$ lemonade server -allow 127.0.0.1 &
$ ssh -R 2489:127.0.0.1:2489 user@host
```

See:

- [SSH/OpenSSH/PortForwarding - Community Help Wiki](https://help.ubuntu.com/community/SSH/OpenSSH/PortForwarding)
- [WOW! and security? · Issue #14 · lemonade-command/lemonade](https://github.com/lemonade-command/lemonade/issues/14#)



Links
-------

- https://speakerdeck.com/pocke/remote-utility-tool-lemonade
- [リモートのPCのブラウザやクリップボードを操作するツール Lemonade を作った - pockestrap](http://pocke.hatenablog.com/entry/2015/07/04/235118)
- [リモートユーティリティーツール、Lemonade v0.2.0 をリリースした - pockestrap](http://pocke.hatenablog.com/entry/2015/08/23/221543)
- [lemonade v1.0.0をリリースした - pockestrap](http://pocke.hatenablog.com/entry/2016/04/19/233423)
