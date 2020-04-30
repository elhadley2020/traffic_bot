FROM alpine:latest
RUN apk add --no-cache git go chromium
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN go get -u github.com/chromedp/chromedp github.com/chromedp/cdproto github.com/chromedp/chromedp-proxy
COPY main.go $GOPATH/src/github.com/elhadley2020/traffic_bot/
RUN cd /go/src/github.com/elhadley2020/traffic_bot/ && go build
CMD /go/src/github.com/elhadley2020/traffic_bot/traffic_bot
