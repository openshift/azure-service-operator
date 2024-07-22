BIN_DIR := bin
TOOLS_DIR := tools

$(RELEASE_DIR):
	mkdir -p $(RELEASE_DIR)/

KUSTOMIZE := tools/bin/kustomize
MANIFESTS_GEN := go run ./vendor/github.com/openshift/cluster-capi-operator/manifests-gen/

$(KUSTOMIZE):
	./tools/ensure-kustomize.sh
