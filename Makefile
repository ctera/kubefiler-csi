# *** WARNING: Targets are meant to run in a build container - Use skipper make ***

all: pylint flake8 coverage

flake8:
	flake8 ctera_gateway_openapi tests/ut

pylint:
	mkdir -p reports/
	PYLINTHOME=reports/ pylint -r n ctera_gateway_openapi tests/ut

test:
	# Run the unittests and create a junit-xml report
	mkdir -p reports/
	nose2 --config=tests/ut/nose2.cfg --verbose --project-directory . $(TEST)

coverage: test
	# Create a coverage report and validate the given threshold
	coverage html --fail-under=90 -d reports/coverage

build: force
	skipper build gateway-openapi
	docker tag gateway-openapi:$(shell git rev-parse HEAD) gateway-openapi:last_build

client:
	GO_POST_PROCESS_FILE="/usr/local/go/bin/gofmt -w -s" \
	java -jar /jars/openapi-generator-cli.jar generate \
	-i https://raw.githubusercontent.com/ctera/ctera-gateway-openapi/master/ctera_gateway_openapi/api.yml \
	-g go \
	-o pkg/ctera \
	--additional-properties packageName=ctera,packageVersion=1.0.0,isGoSubmodule=true \
    --type-mappings=object=interface{} \
	--enable-post-process-file

clean:
	# Clean any generated files
	rm -rf build dist .coverage .cache reports

force: