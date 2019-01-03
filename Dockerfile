FROM centos

MAINTAINER light4d
ADD GOPATH/src/github.com/light4d/lightpan/lightpan .
ADD vueproject/dest dest
EXPOSE 9002
EXPOSE 9003
EXPOSE 9004
VOLUME /root/home/config.json
CMD  ["./lightpan" ]
