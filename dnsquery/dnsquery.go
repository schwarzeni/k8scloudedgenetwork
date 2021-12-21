package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "10.211.55.2:53530")
		},
	}
	ip, err := r.LookupHost(context.Background(), "test.edge.zone")
	if err != nil {
		log.Fatal(err)
	}

	print(ip[0])
}
