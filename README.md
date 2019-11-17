# netpipe

netpipe listens for incoming connections and runs a command on a connection
using `net.Conn` as stdin.

## Getting Started

```bash
go get github.com/koorgoo/netpipe
```

## Usage

Copy a directory over network.

```bash
# Destination host
netpipe -l 8000 sudo tar -xvf - -C dst

# Source host
find src -type f | tar -cvf - -T - | nc destination 8000
find src -type f | parallel "tar -cvf - {} | nc destination 8000"
```
