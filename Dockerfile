FROM centos

MAINTAINER light4d
ADD GOPATH/src/github.com/light4d/lightpan/lightpan .
EXPOSE 9002
VOLUME /root/home/config.json
CMD  ["./lightpan" ]
