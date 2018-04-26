# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Install dependencies
#RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/nicolaszon/2b_todayslesson/todayslessons

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,

WORKDIR /go/src/github.com/nicolaszon/2b_todayslesson/todayslessons

# either manually or with a tool like "godep".)
RUN go install github.com/nicolaszon/2b_todayslesson/todayslessons

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/todayslesson

# Document that the service listens on port 4002.
EXPOSE 4002
