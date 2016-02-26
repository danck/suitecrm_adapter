FROM golang

ADD . /go/src/HAWAI/repos/hawai-crm

RUN go install HAWAI/repos/hawai-crm

ENTRYPOINT /go/bin/hawai-crm \
	-suitecrm-addr http://192.168.29.131/service/v4_1/rest.php

EXPOSE 8080

