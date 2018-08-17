FROM golang:1.10-alpine3.8
LABEL maintainer="Ryota Egusa <gedorinku@yahoo.co.jp>"

RUN apk --no-cache add git
RUN go get -u -v github.com/ProgrammingLab/toshokan/... && \
    go install github.com/ProgrammingLab/toshokan/cmd/toshokan-server && \
    rm -rf /go/src/github.com/

CMD [ "/go/bin/toshokan-server", "--port=53830", "--host=0.0.0.0" ]
