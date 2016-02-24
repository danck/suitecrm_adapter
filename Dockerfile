FROM golang

ADD . /go/src/gitlab.com/danck/hawai-suitecrm

RUN go install gitlab.com/danck/hawai-suitecrm

ENTRYPOINT /go/bin/hawai-suitecrm \
	-suitecrm-addr http://192.168.29.131/service/v4_1/rest.php

EXPOSE 8080

