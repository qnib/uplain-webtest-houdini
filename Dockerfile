FROM golang AS build

WORKDIR /go/src/github.com/qnib/policy/
RUN go get github.com/julienschmidt/httprouter
COPY main.go .
RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM qnib/uplain-cuda8:8.0.61-1.3
ENV HTTP_HOST=""
ENV HTTP_PORT="9999"
COPY --from=build /go/src/github.com/qnib/policy/main /main
CMD ["/main"]
