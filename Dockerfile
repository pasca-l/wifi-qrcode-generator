FROM golang:1.24-bookworm

ENV HOME=/home/local/
WORKDIR /home/local/src/

RUN apt-get update && apt-get upgrade -y

COPY */go.mod */go.sum ./
