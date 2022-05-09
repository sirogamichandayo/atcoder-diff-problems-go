FROM golang:1.18.0-bullseye

RUN apt-get update && apt-get install -y unzip

# プロジェクト用ディレクトリ
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./ /go/src/app

RUN go mod tidy
RUN go get github.com/rubenv/sql-migrate/sql-migrate
RUN go install -v github.com/rubenv/sql-migrate/sql-migrate
