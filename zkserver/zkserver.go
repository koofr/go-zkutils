package main

import (
	"flag"
	"github.com/koofr/go-zkutils"
	"log"
)

func main() {
	var port = flag.Int("port", 2181, "ZooKeeper port for listening")

	flag.Parse()

	server, err := zkutils.NewTestServer(*port)

	if err != nil {
		log.Fatal(err)
	}

	defer server.Stop()

	log.Printf("ZooKeeper server running on 127.0.0.1:%d", *port)

	<-make(chan int)
}
