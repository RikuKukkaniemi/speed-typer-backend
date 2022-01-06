# Speed Typer 

Speed Typer is a game that runs in a web browser. The idea is to type the words you as precisely as you can before timer runs out. The points are determined on how precisely you type the words.

Game can be played at https://speed-typer-demo.herokuapp.com/

## Backend

This is a simple backend API for Speed Typer application. This backend stores and forwards information about the highscores and provides list of words requested language.

Application is dockerized and runs in GCP.

Dev stack:
- Go
- MongoDB
- Gin
- Docker

Speed Typer game contains also a frontend: https://github.com/RikuKukkaniemi/speed-typer-frontend
