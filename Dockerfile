FROM alpine:latest
RUN apk add --no-cache git go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN go get -u github.com/chromedp/chromedp
COPY main.go $GOPATH/src/github.com/elhadley2020/traffic_bot/
