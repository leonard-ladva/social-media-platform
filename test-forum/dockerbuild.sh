#!/bin/bash
echo "Building docker image"
sudo docker build -t forum .
echo ""
echo "Pruning docker images"
sudo docker system prune
echo ""
echo "Showing docker images"
sudo docker images
echo "Running docker container"
sudo docker run -p 8000:8000 --name forumcontainer -it --rm forum
