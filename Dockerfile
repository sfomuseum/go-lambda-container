FROM golang:1.16-alpine

RUN mkdir /build

COPY . /build/go-lambda-container

RUN apk update && apk upgrade \
    && cd /build/go-lambda-container \
    && go build -mod vendor -o /main cmd/hello-world/main.go \    
    && cd && rm -rf /build

# (Optional) Add Lambda Runtime Interface Emulator and use a script in the ENTRYPOINT for simpler local runs
# # https://docs.aws.amazon.com/lambda/latest/dg/go-image.html#go-image-clients

ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod 755 /usr/bin/aws-lambda-rie
COPY Dockerfile.entry.sh /entry.sh
RUN chmod 755 /entry.sh

ENTRYPOINT [ "/entry.sh" ] 