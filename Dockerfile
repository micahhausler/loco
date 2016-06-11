# To build:
# $ GOOS=linux go build -o loco main.go
# $ docker build -t micahhausler/loco .
#
# To run:
# $ docker run -v $(pwd):/loco micahhausler/loco -u user -p password -r <registry>
# or
# $ docker run -v micahhausler/loco -u user -p password -r <registry> -o - | tee docker.tgz

FROM busybox

MAINTAINER Micah Hausler, <hausler.m@gmail.com>

ADD ./loco /bin/loco

WORKDDIR /loco
VOLUME /loco

ENTRYPOINT ["/bin/loco"]
