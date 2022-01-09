# How to contribute

## Installation

`git clone git@github.com:RikuKukkaniemi/speed-typer-backend.git`

After cloning the project

1. Create `.env` file
2. Copy `.env.example` content to `.env` file
3. Run `go run main.go`

## Create new Docker image

`docker build -t speed-typer-backend:{VERSION.NUMBER} .`

## Deploy to GCP

`docker tag speed-typer-backend:{VERSION.NUMBER} gcr.io/speed-typer-app/speed-typer-backend:{VERSION.NUMBER}`

`docker push gcr.io/speed-typer-app/speed-typer-backend:{VERSION.NUMBER}`
