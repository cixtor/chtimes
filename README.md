# Chtimes [![GoReport](https://goreportcard.com/badge/github.com/cixtor/chtimes)](https://goreportcard.com/report/github.com/cixtor/chtimes) [![GoDoc](https://godoc.org/github.com/cixtor/chtimes?status.svg)](https://godoc.org/github.com/cixtor/chtimes)

> MAC times are pieces of file system metadata which record when certain events pertaining to a computer file occurred most recently. The events are usually described as "modification" (the data in the file was modified), "access" (some part of the file was read), and "metadata change" (the file's permissions or ownership were modified), although the acronym is derived from the "mtime", "atime", and "ctime" structures maintained by Unix file systems.
>
> A file's modification time describes when the content of the file most recently changed. Because most file systems do not compare data written to a file with what is already there, if a program overwrites part of a file with the same data as previously existed in that location, the modification time will be updated even though the contents did not technically change.
> 
> -- https://en.wikipedia.org/wiki/MAC_times

## Installation

```
go get -u github.com/cixtor/chtimes
```

## Usage

For a folder with files apparently created on July 2nd, 2021:

```shell
$ ls -l
-rw-r--r--  Jul  2  2021  IMG_20210101_031550.jpg
-rw-r--r--  Jul  2  2021  IMG_20210204_124927.png
-rw-r--r--  Jul  2  2021  IMG_20210309_224507.gif
-rw-r--r--  Jul  2  2021  VID_20210411_123940.mp4
-rw-r--r--  Jul  2  2021  VID_20210515_131809.mov
-rw-r--r--  Jul  2  2021  AUD_20210618_131342.opus
-rw-r--r--  Jul  2  2021  TXT_20210720_201457.json
-rw-r--r--  Jul  2  2021  DOC_20210826_095022.pdf
```

Use `chtimes *.*` to fix the modification time of all files.

```shell
$ chtimes *.*
IMG_20210101_031550.jpg  >>> 2021-01-01 03:15:50 -0800
IMG_20210204_124927.png  >>> 2021-02-04 12:49:27 -0800
IMG_20210309_224507.gif  >>> 2021-03-09 22:45:07 -0800
VID_20210411_123940.mp4  >>> 2021-04-11 12:39:40 -0800
VID_20210515_131809.mov  >>> 2021-05-15 13:18:09 -0800
AUD_20210618_131342.opus >>> 2021-06-18 13:13:42 -0800
TXT_20210720_201457.json >>> 2021-07-20 20:14:57 -0800
DOC_20210826_095022.pdf  >>> 2021-08-26 09:50:22 -0800
```
