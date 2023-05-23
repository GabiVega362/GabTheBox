# Flow

- Ejecuta [main()](./main.go)
    - Crea el contexto de la aplicación usando [NewContext()](./server/config/context.go)
        - Obtiene el cliente de Docker usando [NewDocker](./server/docker/client.go)
        - Obtiene los argumentos usando [parseArgs()](./server/config/args.go)
    - Inicia y pone a la escucha el servidor usando [ListenAndServe()](./server/server.go)
        - Crea un nuevo enrutador web 
        - Registra las rutas usando [SetRoutes()](./server/routes/index.go)
            - Las rutas GET estan definidas en [get.go](./server/routes/get.go)
            - Las rutas POST estan definidas en [post.go](./server/routes/post.go)
            - Las plantillas HTML est'an definidas en [templates/](./server/templates/)
        - Pone a la escucha el servidor