FROM golang:1.11 as oven
WORKDIR /go/src/github.com/xocasdashdash/nappy/
COPY nappy.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o nappy . && chmod +x nappy

FROM       scratch  
COPY --from=oven /go/src/github.com/xocasdashdash/nappy/nappy /nappy
ENTRYPOINT [ "/nappy" ]
CMD [ "1s" ]