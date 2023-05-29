package config

import (
	"flag" // Paquete encargadop de gestionar los argumentos por terminal
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Args es la estructura que contiene los argumentos pasados por terminal o por variables de entorno
type Args struct {
	//Address es la dirección en la que el servidor escuchara en formato host:port
	Address string

	DBName string
	DBUser string
	DBPass string
}

func (args *Args) CheckDefaults() *Args {
	default_value := "gabthebox"
	if strings.TrimSpace(args.Address) == "" {
		args.Address = "0.0.0.0:8080"
	}
	if strings.TrimSpace(args.DBName) == "" {
		args.DBName = default_value
	}
	if strings.TrimSpace(args.DBUser) == "" {
		args.DBUser = default_value
	}
	if strings.TrimSpace(args.DBPass) == "" {
		args.DBPass = default_value
	}
	return args
}

// parseArgs parsea los argumentos que llegan por terminal
func parseArgs() *Args {
	// cargamos el archivo .env
	_ = godotenv.Load()

	// obtenemos el argumento "-address" o ":8080" si no se ha espeecificado por terminal
	address := flag.String("address", os.Getenv("GTB_ADDR"), "Address to listen on")
	dbname := flag.String("dbname", os.Getenv("DATABASE_NAME"), "Database name")
	dbuser := flag.String("dbuser", os.Getenv("DATABASE_USER"), "Database user")
	dbpass := flag.String("dbpass", os.Getenv("DATABASE_PASS"), "Database password")
	flag.Parse()

	// creamos la estructura de argumentos
	args := &Args{
		Address: *address,
		DBName:  *dbname,
		DBUser:  *dbuser,
		DBPass:  *dbpass,
	}
	// comprobamos que los argumentos no esten vacios y devolvemos el resultado
	return args.CheckDefaults()
}
