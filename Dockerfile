FROM golang:alpine
#set workdir
WORKDIR /app

#download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd

CMD [ "/bin/app" ]
