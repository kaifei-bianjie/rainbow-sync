FROM alpine:3.8

# Set up dependencies
ENV PACKAGES go make git libc-dev bash

# Set up path
ENV BINARY_NAME rainbow-sync
ENV GOPATH       /root/go
ENV REPO_PATH    $GOPATH/src/github.com/irisnet/rainbow-sync/service/cosmos
ENV PATH         $GOPATH/bin:$PATH

RUN mkdir -p $GOPATH $REPO_PATH

COPY . $REPO_PATH
WORKDIR $REPO_PATH

VOLUME $REPO_PATH/logs

# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && \
    cd $REPO_PATH && make all && \
    mv $REPO_PATH/$BINARY_NAME $GOPATH/bin && \
    rm -rf $REPO_PATH/vendor && \
    rm -rf $GOPATH/src/github.com/golang $GOPATH/bin/dep $GOPATH/pkg/* && \
    apk del $PACKAGES

CMD $BINARY_NAME