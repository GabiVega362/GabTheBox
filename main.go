package main

import (
	"fmt" // paquete encargado de la entrada y salida de datos por consola
	"os" // paquete encargado de gestionar el sistema operativo

	"github.com/gabivega362/gabthebox/server"
	"github.com/gabivega362/gabthebox/server/config"
)

// Función de entrada al programa
func main() {
	// obtenemos el contexto del programa, es decir, la configuración con las variables globales(argumentos, variables de entorno, clientes, etc)
	ctx, err := config.NewContext()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	// ponemos el servidor a la escucha pasandole el contexto de la aplicación. Esta función es bloqueante, por lo que el servidor continuara hatsa que el servidor se detenga
	server.ListenAndServe(ctx)
}
