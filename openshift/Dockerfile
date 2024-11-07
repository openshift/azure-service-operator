FROM registry.ci.openshift.org/ocp/builder:rhel-9-golang-1.22-openshift-4.18 AS builder

WORKDIR /build
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=${GOOS} GOPROXY=${GOPROXY} go build \
  -ldflags="-w -s -X 'main.version=${VERSION}'" \
  -o=azure-service-operator \
  ./cmd/controller

FROM registry.ci.openshift.org/ocp/4.18:base-rhel9

LABEL description="Azure Service Operator"

COPY --from=builder /build/azure-service-operator /bin/azure-service-operator

ENTRYPOINT [ "/bin/azure-service-operator" ]
