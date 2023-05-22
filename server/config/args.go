package config

import "flag"

type Args struct {
	//Address string
	Address string
	//Port    string
    //User    string
    //Password string
    //Database string
}

func parseArgs() *Args {
	address := flag.String("address", ":8080", "Address to listen on")
	flag.Parse()
	
	return &Args{
        Address: *address,
    }
}