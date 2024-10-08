all: build
.PHONY: all

# Include the library makefile
include $(addprefix ./vendor/github.com/openshift/build-machinery-go/make/, \
	golang.mk \
	targets/openshift/deps-gomod.mk \
	targets/openshift/operator/profile-manifests.mk \
)

# This will include additional actions on the update and verify targets to ensure that profile patches are applied
# to manifest files
# $0 - macro name
# $1 - target name
# $2 - profile patches directory
# $3 - manifests directory
$(call add-profile-manifests,manifests,./profile-patches,./manifests)

# Run core verification and all self contained tests.
#
# Example:
#   make check
check: | verify test-unit golangci-lint
.PHONY: check

golangci-lint:
	golangci-lint run --verbose --print-resources-usage --modules-download-mode=vendor --timeout=5m0s
.PHONY: golangci-lint

install.tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b ${GOPATH}/bin
.PHONY: install.tools

.PHONY: update-codegen
update-codegen:
	hack/update-codegen.sh

.PHONY: verify-update-codegen
verify-update-codegen: update-codegen
	git diff --exit-code

.PHONY: verify
verify: verify-update-codegen

clean:
	$(RM) cluster-network-operator cluster-network-check-endpoints cluster-network-check-target
.PHONY: clean

GO_TEST_PACKAGES :=./pkg/... ./cmd/...
