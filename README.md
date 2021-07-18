# Comic Sticks

Comic Sticks (`xkcd-gtk`) is a simple xkcd comic viewer written in Go using
GTK+3.

<a href="https://appcenter.elementary.io/com.github.rkoesters.xkcd-gtk"><img height="51" alt="Get it on AppCenter" src="https://appcenter.elementary.io/badge.svg"/></a>
<a href="https://flathub.org/apps/details/com.github.rkoesters.xkcd-gtk"><img height="51" alt="Download on Flathub" src="https://flathub.org/assets/badges/flathub-badge-en.svg"/></a>
<a href="https://aur.archlinux.org/packages/xkcd-gtk/"><img height="51" alit="Availabe in AUR" 
src="https://archlinux.org/static/logos/archlinux-logo-light-scalable.1ae4cc2e2469.svg"/></a>

![screenshot](screenshots/screenshot-1@2x.png)

## Building from source

### Requirements

To build this program, you will need Go (version >= 1.11, something like
`golang` or `go`) and GTK+ development files (version >= 3.20, something like
`libgtk-3-dev` or `gtk3-devel`).

### Building

Just run `make` from the root of the repo:

```shell
$ make
```

#### Testing

Run static analysis with:

```shell
$ make check
```

Run tests with:

```shell
$ make test
```

### Installing

After you have compiled the application, you can install it.

To install for all users:

```shell
$ sudo make install prefix=/usr/local
```

To install for the current user only (you may need to add `$HOME/.local/bin` to
your `$PATH`):

```shell
$ make install prefix="$HOME/.local"
```

### Uninstalling

To uninstall for all users:

```shell
$ sudo make uninstall prefix=/usr/local
```

To uninstall for the current user:

```shell
$ make uninstall prefix="$HOME/.local"
```

## License

This program comes with absolutely no warranty. See the [GNU General Public
License, version 3 or later](LICENSE) for details.
