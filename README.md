# netpipe

netpipe listens for incoming connections and runs a command on a connection
using `net.Conn` as `stdin`.

## Getting Started

```bash
go get github.com/koorgoo/netpipe
```

## Usage

Copy a directory.

```bash
# Destination host
netpipe -l 8000 tar -xvf - -C dst

# Source host
tar -cvf - .   | nc destination 8000
find . -type f | parallel -j5 "tar -cvf - {} | nc destination 8000"
```
