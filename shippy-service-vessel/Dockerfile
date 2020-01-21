FROM golang:latest as builder

RUN mkdir /app
WORKDIR /app


COPY . .
ENV GO111MODULE=on
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-vessel


FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-vessel .

CMD ["./shippy-service-vessel"]