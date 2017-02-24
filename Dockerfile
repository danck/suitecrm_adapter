FROM golang

ADD . /go/src/gitlab.com/danck/hawai-crm

RUN go install gitlab.com/danck/hawai-crm

RUN go get gitlab.com/danck/hawai-runner

RUN go install gitlab.com/danck/hawai-runner

ENV EXTERNAL_HOST_ADDRESS=192.168.29.142
ENV EXTERNAL_HOST_PORT=32001
ENV LOG_FILE=hawai-crm.log
ENV SERVICE_COMMAND="/go/bin/hawai-crm -suitecrm-addr http://192.168.29.131/service/v4_1/rest.php -listen-addr :32001"
ENV SERVICE_LABEL=crm
ENV REGISTRY_ADDRESS='http://192.168.29.142:32000/service'


ENTRYPOINT /go/bin/hawai-runner
EXPOSE 32001 

