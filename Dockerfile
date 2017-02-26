# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/rest-go/drinks

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go install github.com/rest-go/drinks

# Run the drinks command by default when the container starts.
ENTRYPOINT ["/go/bin/drinks"]

# Document that the service listens on port 10000.
EXPOSE 10000
