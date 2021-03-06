# romaine-head
The background service for Romaine, a crouton packaged app for Chrome OS chrooting

This project was started on a Celeron Acer C720
and is being continued on a Pixel 2015.

## Requirements
* Your device must be in dev mode.
* Download [Crouton](https://goo.gl/fd3zc) if you didn't already. Just make sure it's in ~/Downloads.
* x64 (so, not ARM) Chromebook or Chromebox. I didn't compile for ARM, but the chroot experience is harder then anyway.

## Setup
1. Download a binary from the [releases page](https://github.com/danopia/romaine-head/releases).
2. Open crosh (Ctrl-Alt-T) and run `shell`
3. If you haven't use crouton yet, enter `sudo mkdir /usr/local/bin`
4. Copy the bin into place with `sudo cp ~/Downloads/romaine-x64 /usr/local/bin/romaine`
5. Make it executable with `sudo chmod +x /usr/local/bin/romaine`

## Running
1. Run `sudo romaine` in a crosh `shell`.
2. Open the [Romaine app](https://chrome.google.com/webstore/detail/romaine/akmgdkngbndhjenanchcijaappbglfgn).

## Development
Run something like this in a chronos shell to build and run from git:

```shell
sudo enter-chroot sh go/src/github.com/danopia/romaine-head/build.sh && sh ~/Downloads/romaine-head.run
```

You can store the command in a cronos shell script for ease of use.
