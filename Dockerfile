FROM golang:1.8
COPY . "$GOPATH/src/github.com/FlyingFeather/Agenda-server"
RUN cd "$GOPATH/src/github.com/FlyingFeather/Agenda-server/cli" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/FlyingFeather/Agenda-server/service" && go get -v && go install -v
WORKDIR /
EXPOSE 8080
VOLUME ["/data"]
