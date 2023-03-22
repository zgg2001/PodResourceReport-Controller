FROM golang:1.19 as builder

# Copy in the go src
COPY kube_config /root/.kube/config
WORKDIR /go/src/github.com/zgg2001/PodResourceReport-Controller
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=off go build -a -o /user/local/bin/reporter github.com/zgg2001/PodResourceReport-Controller/cmd/

ENTRYPOINT ["/user/local/bin/reporter", "-kubeconfig=/root/.kube/config"]