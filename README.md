# real-time-forum
[Task Description](https://github.com/01-edu/public/tree/master/subjects/real-time-forum) | [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/real-time-forum/audit)
---|---
## Running Dockerized App
run
```bash
./dockerbuild.sh
```
to create the docker images, this will take a minute or two

after that has finished run
```bash
docker compose up
```
to create containers and run them, after the frontend has compiled you are ready to visit localhost:3000

## Stopping Containers and removing docker images
open a second terminal, or stop the program with **CTRL+C**, then run
```bash
docker compose down
./removeimages.sh
```