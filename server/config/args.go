package config

import "flag" // Paquete encargadop de gestionar los argumentos por terminal

// Args es la estructura que contiene los argumentos pasados por terminal o por variables de entorno
type Args struct {
	//Address es la dirección en la que el servidor escuchara en formato host:port
	Address string
	//Port    string
    //User    string
    //Password string
    //Database string
}

// parseArgs parsea los argumentos que llegan por terminal 
func parseArgs() *Args {
	// obtenemos el argumento "-address" o ":8080" si no se ha espeecificado por terminal
	address := flag.String("address", ":8080", "Address to listen on")
	flag.Parse()
	
	// devolvemos los argumentos
	return &Args{
        Address: *address,
    }
}