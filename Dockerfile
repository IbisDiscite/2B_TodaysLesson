# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Install dependencies
#RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

RUN mkdir -p /go/src/github.com/nicolaszon/2b_todayslesson/todayslesson

WORKDIR /go/src/github.com/nicolaszon/2b_todayslesson/todayslesson


# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/nicolaszon/2b_todayslesson/todayslesson

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/nicolaszon/2b_todayslesson/todayslesson

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/todayslesson

# Document that the service listens on port 4002.
#EXPOSE 4002
