FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GO111MODULE=on

WORKDIR /opt/app

RUN apk --no-cache update && \
      apk --no-cache add git ca-certificates && \
      rm -rf /var/cache/apk/*

COPY . ./

RUN go build -a -o app .

FROM alpine

RUN apk --no-cache update && \
      apk --no-cache add ca-certificates && \
      rm -rf /var/cache/apk/* &&\
      mkdir /usr/local/bin/log && \
      touch /usr/local/bin/log/communication.log
      
COPY --from=builder /opt/app/app /usr/local/bin/app
 
# SMTP_USER=it@persianblack.com
ENV LOG_FILE_LOCATION=/usr/local/bin/log/communication.log ATTACHMENT_PATH=/usr/local/bin/attachments/ \
 CLIENT_ID=persianblack SMTP_HOST=us2.smtp.mailhostbox.com \
  SMTP_PORT=25 SMTP_USER=it@persianblack.com SMTP_PASSWORD=Princess4Daprinz


CMD ["/usr/local/bin/app", "--help"]