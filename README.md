# TicTacToe - built using [tictacgo](https://github.com/rajenderK7/tictacgo)
A websocket based online multiplayer game.

## Tech
- Go Echo framework - For the game server
- tictacgo - A tictactoe game engine
- React JS + Vite - The game client

## Running the application
### Dev
Start the client and server separately.

Navigate to the server entry point (From the root):
```
cd server/cmd/tictactoews
go run main.go
```

Navigate to the client (From the root):
The project uses pnpm, however npm can also be used. Just swap `pnpm` with `npm run`
```
cd client
pnpm install
```

Start the client
```
pnpm dev
```


### Using Docker
Install [Docker](https://docker.com)

### 1. With Docker compose
Follow [this](https://github.com/rajenderK7/tictactoews/new/main#2-without-docker-compose) instructions if you want to use docker compose instead.

Navigate to the root of the project and do:
```
docker compose up
```
This command will build the images using the appropriate Dockerfile(s) as specified in `docker-compose.yml` spins up both the _client_ and _server_ containers. Checkout the compose file for more details on the ports used.

Now you can access the client on `localhost:3000`

### 2. Without Docker compose
To build and run containers separately navigate to the client and server directories and follow the below steps:

Create Docker image
```
docker build -t image-name:<optional-tag> .
```

Start the containers
Server
```
docker run -d -p 4000:4000 server-image-name:tag 
```

Client
```
docker run -d -p 3000:80 client-image-name:tag
```

Now you can access the client on `localhost:3000`
