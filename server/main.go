package main

import (
	"github.com/mediocregopher/radix.v2/redis"
	"flag"
)

var (
	redis_client *redis.Client
	redis_prot *string = flag.String("redis_prot", "tcp", "Redis protocol (tcp/udp)")
	redis_addr *string = flag.String("redis_addr", "localhost:6379", "Redis address (default localhost:6379)")

	address *string = flag.String("address", "127.0.0.1", "Address to bind listener to")
	port *int = flag.Int("port", 7171, "Port to bind listener to")
)

func main() {

	flag.Parse()

	var err error

	redis_client, err = redis.Dial(*redis_prot, *redis_addr)

	if err != nil {
		fatal("Unable to connect to redis", err)
	}

	go cleaner(redis_client)
	startListening(*address, *port)
}

