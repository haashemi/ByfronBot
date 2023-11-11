# ByfronBot

A simple personal Telegram bot based on [tgo](https://github.com/haashemi/tgo) framework.

### Commands:

- `me` shows a basic user information, nothing special
- `arz` returns the current IRT exchange rates based on [bonbast](https://bonbast.com)
- `stp` downloads a sticker and sends it back as a document
- `ptss` converts a photo to a 512p png
- `time` shows current (bruh) and current Jalali date.
- `server` shows server's uptime and resource usage

### Usage:

1- Modify the config file however you prefer

```bash
$ cp config.example.yaml config.yaml
$ # modify the config.yaml however you want here
```

2- Compile and run it!

```
$ go build .
$ ./ByfronBot
```
