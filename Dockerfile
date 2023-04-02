FROM golang:alpine
WORKDIR /workspace
RUN apk add --no-cache --quiet git && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/ramya-rao-a/go-outline@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v github.com/rogpeppe/godef@latest
