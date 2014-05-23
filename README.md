go-zkutils
==================

Go ZooKeeper utils.

[![GoDoc](https://godoc.org/github.com/koofr/go-zkutils?status.png)](https://godoc.org/github.com/koofr/go-zkutils)

## Install

    go get github.com/koofr/go-zkutils

## Dependencies

Dependencies for examples and testing:

    wget http://www.apache.org/dist/zookeeper/zookeeper-3.4.6/zookeeper-3.4.6.tar.gz
    tar xf zookeeper-3.4.6.tar.gz -C /opt

## Example

    go get github.com/koofr/go-zkutils/zkserver
    ZKROOT=/opt/zookeeper-3.4.6 zkserver

## Testing

    go get -t
    ZKROOT=/opt/zookeeper-3.4.6 go test
