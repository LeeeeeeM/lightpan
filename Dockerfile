FROM centos

MAINTAINER light4d
ADD GOPATH/src/github.com/light4d/lightpan/lightpan .
EXPOSE 9002
EXPOSE 9003
VOLUME /root/home/config.json
CMD  ["./lightpan" ]
