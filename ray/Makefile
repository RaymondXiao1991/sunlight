all: build

deps:
	go get -d -v ray.com/...

vet:
	@ for f in $$(find . -type f -iregex '.*\.go' | grep -v "vendor"); do \
		go fmt $$f; \
	done

	@ for f in $$(find . -type f -iregex '.*\.go'  | grep -v "vendor"); do \
		go tool vet $$f; \
	done

updatedeps:
	go get -d -v -u -f ray.com/...

testdeps:
	go get -d -v -t ray.com/...

updatetestdeps:
	go get -d -v -t -u -f ray.com/...

build:
	bash/build.sh

gen:
	@ if ! which thrift > /dev/null; then \
		echo "error: thrift not installed" >&2; \
		exit 1; \
	fi

	@ for f in $$(find proto -type f -iregex '.*\.thrift'); do \
		thrift -out proto -r --gen go $$f; \
	done

	@ for d in $$(find proto/* -type 'd' | grep "_sevice-remote"); do \
		rm -rf $$d; \
	done

	@ for f in $$(find . -type f -iregex '.*\.go'  | grep 'consts.go\|GoUnusedProtection__.go'); do \
		rm -f $$f; \
	done

test: testdeps
	go test -v -cpu 1,4 ray.com/...

testrace: testdeps
	go test -v -race -cpu 1,4 ray.com/...

clean:
	@ #go clean -i ray.com/...
	@ rm -rf _output

	@ for d in $$(find proto/* -type 'd' | grep -v "thrift"); do \
		rm -rf $$d; \
	done

coverage: testdeps
	./coverage.sh --coveralls

.PHONY: \
	all \
	deps \
	vet	\
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	gen \
	test \
	testrace \
	clean \
	coverage
