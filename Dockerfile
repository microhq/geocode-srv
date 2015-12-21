FROM alpine:3.2
ADD geocode-srv /geocode-srv
ENTRYPOINT [ "/geocode-srv" ]
