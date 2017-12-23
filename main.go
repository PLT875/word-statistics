package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"

	"github.com/PLT875/word-statistics/aggregator"
	"github.com/PLT875/word-statistics/api"
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
	}
}

func main() {
	// read the port from the command line flag, default is 5555
	var ingest = flag.String("ingest", "5555", "default port to ingest words")
	var stats = flag.String("stats", "8080", "default port to ingest words")
	flag.Parse()
	ingestPort := *ingest
	statsPort := *stats

	// create new Aggregator
	agg := aggregator.NewAggregator()

	// run the API
	router := api.Router(agg)
	go func() {
		router.Run(":" + statsPort)
	}()

	// register the port to accept incoming connections
	in, err := net.Listen("tcp", "localhost:"+ingestPort)
	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
		os.Exit(1)
	}

	defer in.Close()

	log.Printf("Listening / ready to ingest at port: %s", ingestPort)
	for {
		// listen for an incoming connection.
		c, err := in.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s", err.Error())
			os.Exit(1)
		}

		go handleIngestion(c, agg)
	}
}
