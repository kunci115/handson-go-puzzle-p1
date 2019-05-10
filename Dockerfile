FROM golang:latest
LABEL maintainer "Rino Alfian <rino.alpin@gmail.com>"

# for install go packages RUN go get /path
RUN go get github.com/mattn/go-sqlite3

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/kunci115/myproject

# build executable
RUN go install github.com/kunci115/myproject

# execute 
ENTRYPOINT /go/bin/myproject

# Document that the service listens on port 8080.
EXPOSE 8080
