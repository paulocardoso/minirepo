FROM        golang:1.13
MAINTAINER  Paulo Cardoso <paulinhocru@gmail.com>

ENV SERVER_PORT  8080
ENV LIB_PATH "/tmp/minirepo"

# Setting up working directory
WORKDIR     /go/src/minirepo-container
ADD         . /go/src/minirepo-container


RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

EXPOSE 8080

RUN ls

CMD ["minirepo"]
