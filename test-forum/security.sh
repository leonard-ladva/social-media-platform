#!/bin/bash

# Adding the environment variables
export SERVER_PORT=":8001"
export SERVER_HTTPS_PORT=":8000"
export SQLITE_USERNAME="admin"
export SQLITE_PASSWORD="admin321"
export SQLITE_DB_PATH="./assets/database.db"
export CGO_ENABLED=1

echo Environment variables are added

#Getting the name of directory, where the project is built (same as filename of the binary)
filename=${PWD##*/}

echo building the program with tags

go build --tags "sqlite_userauth linux"
chmod +x $filename
./$filename