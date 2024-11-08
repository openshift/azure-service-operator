FROM registry.ci.openshift.org/ocp/builder:rhel-9-golang-1.22-openshift-4.18 AS builder

WORKDIR /build
# NOTE: We're copying v2 in like this rather than using v2 as the context
# to avoid issues with ci-operator's working directory behavior
COPY v2/ ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=${GOOS} GOPROXY=${GOPROXY} go build \
  -ldflags="-w -s -X 'main.version=${VERSION}'" \
  -o=azure-service-operator \
  ./cmd/controller

FROM registry.ci.openshift.org/ocp/4.18:base-rhel9

LABEL description="Azure Service Operator Controller Manager"

COPY --from=builder /build/azure-service-operator /bin/azure-service-operator

ENTRYPOINT [ "/bin/azure-service-operator" ]
