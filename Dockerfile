# image with actual code to build binary or dev/test
FROM golang:1.16 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    LD_FLAGS="-w -s"

WORKDIR /usr/src/app

COPY go.mod /usr/src/app/
COPY go.sum /usr/src/app/

RUN go mod download &&\
    go mod verify 

COPY cmd /usr/src/app/cmd
COPY internal /usr/src/app/internal

ENV COFFEGB_LISTEN=0.0.0.0:8123 \
    COFFEGB_LOGLEVEL=DEBUG \
    COFFEGB_LOGREQUESTS=true \
    DATABASE_URL="user=pguser password=pgpass host=db port=5432 dbname=GOLAB sslmode=disable" \
    DATABASE_QUERYTIMEOUT=20s

RUN go build -ldflags="$LD_FLAGS" -o /tmp/build/apiserver ./cmd/apiserver/*.go

# final image to run app
FROM scratch as app
ARG VERSION
LABEL Author="Vitaly vvsh.msk@gmail.com"\
      Version=${VERSION}

ENV COFFEGB_LISTEN=0.0.0.0:8123 \
    COFFEGB_LOGLEVEL=DEBUG \
    COFFEGB_LOGREQUESTS=true \
    DATABASE_URL="user=pguser password=pgpass host=db port=5432 dbname=GOLAB sslmode=disable" \
    DATABASE_QUERYTIMEOUT=20s

COPY --from=builder /tmp/build/apiserver /usr/bin/

EXPOSE 8123

ENTRYPOINT [ "apiserver" ]
