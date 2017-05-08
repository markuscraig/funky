FUNKY_GO_EXECUTABLE ?= go
DIST_DIRS := find * -type d -exec
VERSION ?= $(shell git describe --tags)
VERSION_INCODE = $(shell perl -ne '/^var version.*"([^"]+)".*$$/ && print "v$$1\n"' funky.go)
VERSION_INCHANGELOG = $(shell perl -ne '/^\# Release (\d+(\.\d+)+) / && print "$$1\n"' CHANGELOG.md | head -n1)

build:
	${FUNKY_GO_EXECUTABLE} build -o funky -ldflags "-X main.Version=${VERSION}" cmd/funky/main.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./funky ${DESTDIR}/usr/local/bin/funky

test:
	${FUNKY_GO_EXECUTABLE} test . ./project

clean:
	rm -f ./funky.test
	rm -f ./funky
	rm -rf ./dist

bootstrap-dist:
	${FUNKY_GO_EXECUTABLE} get -u github.com/franciscocpg/gox
	cd ${GOPATH}/src/github.com/franciscocpg/gox && git checkout dc50315fc7992f4fa34a4ee4bb3d60052eeb038e
	cd ${GOPATH}/src/github.com/franciscocpg/gox && ${FUNKY_GO_EXECUTABLE} install

build-all:
	gox -verbose \
	-ldflags "-X main.Version=${VERSION}" \
	-os="linux darwin windows freebsd openbsd netbsd" \
	-arch="amd64 386 armv5 armv6 armv7 arm64" \
	-osarch="!darwin/arm64" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf funky-${VERSION}-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r funky-${VERSION}-{}.zip {} \; && \
	cd ..

verify-version:
	@if [ "$(VERSION_INCODE)" = "v$(VERSION_INCHANGELOG)" ]; then \
		echo "funky: $(VERSION_INCHANGELOG)"; \
	elif [ "$(VERSION_INCODE)" = "v$(VERSION_INCHANGELOG)-dev" ]; then \
		echo "funky (development): $(VERSION_INCHANGELOG)"; \
	else \
		echo "Version number in funky.go does not match CHANGELOG.md"; \
		echo "funky.go: $(VERSION_INCODE)"; \
		echo "CHANGELOG : $(VERSION_INCHANGELOG)"; \
		exit 1; \
	fi

.PHONY: build test install clean bootstrap-dist build-all dist integration-test verify-version
