FROM golang:1.22

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/golang/mock/mockgen@v1.5.0
RUN go install github.com/spf13/cobra-cli@latest

RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache

# Allow www-data to write to /go/src
RUN chown -R www-data:www-data /go/src
RUN chmod -Rf 777 /go

USER www-data

CMD ["tail", "-f", "/dev/null"]