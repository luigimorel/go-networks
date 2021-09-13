package serverconnect

import (
	"fmt"
	"net"
)


func resolveAddresses() {
	/*
	*The resolution of the domain name from the IP address can be done with
	*the LookupAddr function from the net package. To find out the IP address from the
	*domain name, the LookupIP function is applied.
	*/
	//Resolve by IP
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _ , addr := range addrs {
		fmt.Println(addr)
	}

	//Resolve by Address 
	ips, err := net.LookupAddr("localhost")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}