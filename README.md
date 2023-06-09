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

# TODO
- [X] Mejorar cliente de docker (Funciones personalizadas)
- [X] Mejorar cliente de SQL (Funciones personalizadas)
- [ ] Gestion de usuarios
    - [X] Registro
    - [X] Login
    - [X] Logout
    - [ ] Admin {opcional}
- [X] Desplegar en puerto aleatorio
- [X] Link lab a user
- [X] Comprobar el numero de contenedores desplegados por usuario (max 1)
- [X] Cambiar estado a Deployed o !Deployed
- [X] Si ya ha deployeado una maquina no permitir deployear otra
- [ ] Revisar CSS

- [ ] Generar contenedores personalizados
- [ ] go (paralelismo)
- [ ] Añadir test
- [ ] Añadir documentación
- [ ] Añadir benchmarks
- [ ] Pasar linter
- [ ] Pruebas de race condition
- [ ] Middlewares

## IMPORTANTE
- [ ] Añadir archivos web a ejecutable