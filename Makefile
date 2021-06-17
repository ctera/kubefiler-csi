# *** WARNING: Targets are meant to run in a build container - Use skipper make ***
PKG=github.com/ctera/ctera-gateway-csi
VERSION=v1.0.0
GIT_COMMIT?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS?="-X ${PKG}/pkg/driver.driverVersion=${VERSION} -X ${PKG}/pkg/cloud.driverVersion=${VERSION} -X ${PKG}/pkg/driver.gitCommit=${GIT_COMMIT} -X ${PKG}/pkg/driver.buildDate=${BUILD_DATE} -s -w"
OPENAPI_FILE?="https://raw.githubusercontent.com/ctera/ctera-gateway-openapi/master/ctera_gateway_openapi/api.yml"

OUTPUT_DIR=out
SOURCES := $(shell find . -name "*.go")

REPORTS_DIR=reports
COVERAGE_REPORT_FILE=${REPORTS_DIR}/coverage.cov
COVERAGE_HTML_FILE=${REPORTS_DIR}/coverage.html
MINIMAL_COVERAGE_RATE=5

all: out/ctera-csi-driver coverage verify

.PHONY: verify
verify:
	hack/verify-all.sh

.PHONY: unit-test
unit-test:
	mkdir -p ${REPORTS_DIR}
	go test -coverprofile=${COVERAGE_REPORT_FILE} ./pkg/... -v

coverage: unit-test
	go tool cover -func=${COVERAGE_REPORT_FILE}
	go tool cover -html=${COVERAGE_REPORT_FILE} -o ${COVERAGE_HTML_FILE}
	hack/verify-coverage.sh ${COVERAGE_REPORT_FILE} ${MINIMAL_COVERAGE_RATE}

.PHONY: gofmt
gofmt:
	hack/update-gofmt.sh

build: ${OUTPUT_DIR}/ctera-csi-driver
	skipper build gateway-csi
	docker tag gateway-csi:${GIT_COMMIT} gateway-csi:last_build

${OUTPUT_DIR}/ctera-csi-driver: ${SOURCES}
	CGO_ENABLED=0 GOOS=linux go build -ldflags ${LDFLAGS} -o $@ ./cmd/

tidy:
	go mod tidy

client:
	GO_POST_PROCESS_FILE="/usr/local/go/bin/gofmt -w -s" \
	java -jar /jars/openapi-generator-cli.jar generate \
	-i ${OPENAPI_FILE} \
	-g go \
	-o pkg/ctera-openapi \
	--additional-properties packageName=cteraopenapi,packageVersion=1.0.0,isGoSubmodule=true \
    --type-mappings=object=interface{} \
	--enable-post-process-file

clean:
	# Clean any generated files
	rm -rf ${OUTPUT_DIR} ${REPORTS_DIR}
