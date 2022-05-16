FROM golang:1.18

WORKDIR /var/www/html

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

EXPOSE 8080

CMD ["air"]
