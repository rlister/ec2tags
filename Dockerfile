FROM scratch
MAINTAINER Ric Lister <rlister@gmail.com>

## tls needs root CA
ADD https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt /etc/ssl/ca-bundle.pem

ADD ec2tags ec2tags

ENTRYPOINT [ "/ec2tags" ]
