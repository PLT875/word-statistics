package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/PLT875/word-statistics/aggregator"
)

// handleIngestion handler.
func handleIngestion(conn net.Conn, agg *aggregator.Aggregator) {
	for {
		req, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			log.Fatalf("Error reading: %s", err.Error())
			return
		}

		agg.IngestWords(string(req))
		fmt.Println(agg.GetWordCounts())
	}
}

func main() {
	// read the port from the command line flag, default is 5555
	var p = flag.String("p", "5555", "default port to ingest words")
	flag.Parse()
	port := *p

	// create new Aggregator
	agg := aggregator.NewAggregator()

	// register the port to accept incoming connections
	ingest, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
		os.Exit(1)
	}

	defer ingest.Close()

	log.Printf("Listening / ready to ingest at port: %s", port)
	for {
		// listen for an incoming connection.
		c, err := ingest.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s", err.Error())
			os.Exit(1)
		}

		go handleIngestion(c, agg)
	}
}
