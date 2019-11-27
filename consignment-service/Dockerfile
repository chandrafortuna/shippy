FROM alpine:latest as builder

# RUN curl https://glide.sh/get | sh

# ENV PKG_NAME=github.com/chandrafortuna/shippy/consignment-service
# ENV PKG_PATH=$GOPATH/src/$PKG_NAME
# WORKDIR $PKG_PATH

# COPY glide.yaml glide.lock $PKG_PATH/
# RUN glide install

# COPY . $PKG_PATH

# RUN go build && go install

# WORKDIR $PKG_PATH
# EXPOSE 9487
# CMD ["./consignment-service"]

RUN apk --no-cache add git

RUN apk add --update tzdata \ bash wget curl git

WORKDIR /app/shippy-service-consignment

COPY . .

RUN glide install

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-service-consignment/shippy-service-consignment .

CMD ["./consignment-service"]