FROM alpine
MAINTAINER Kelsey Hightower <kelsey.hightower@gmail.com>
ADD gopath/bin/secrets /secrets
ENTRYPOINT ["/secrets"]
