FROM golang:1.18-alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -buildvcs=false -o /impactify-api

EXPOSE 8080
CMD [ "/impactify-api" ]