FROM golang:1.23

WORKDIR /jwt-auth-test

COPY . . 

RUN go build -o jwt-auth-test .

EXPOSE 3000

CMD ["./jwt-auth-test"]