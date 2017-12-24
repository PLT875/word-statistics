# Word Statistics

Just a demo application.

## Pre-requisites
* Go 1.8 or above
* [Optional] Glide 0.11.0 or above (the vendor package is commited)

## Clean, test, building and running the application
* Firstly clone the repository into $GOPATH/src/github.com/PLT875
* Take a look at the Makefile for the corresponding steps.

## Interacting with the application

Sending over text on port 5555, e.g.
```
telnet localhost 5555
Trying ::1...
telnet: connect to address ::1: Connection refused
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
hello world
how are you today
hello world
where are you from
hello hello
```

Retrieving statistics on the GET /stats endpoint, e.g.
```
curl -XGET localhost:8080/stats
{"count":14,"top_5_words":["hello","are","you","world","today"],"top_5_letters":["o","l","e","r","h"]}
```