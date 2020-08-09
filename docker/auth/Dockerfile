FROM golang:1.14

WORKDIR /go/src/auth

RUN go get -u github.com/labstack/echo/... && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/go-sql-driver/mysql && \
    go get -u github.com/oxequa/realize && \
    go get -u github.com/rubenv/sql-migrate/... && \
    go get -u github.com/dgrijalva/jwt-go