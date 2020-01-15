targets:=\
	aix/ppc64 \
	android/386 \
	android/amd64 \
	android/arm \
	android/arm64 \
	darwin/386 \
	darwin/amd64 \
	darwin/arm \
	darwin/arm64 \
	dragonfly/amd64 \
	freebsd/386 \
	freebsd/amd64 \
	freebsd/arm \
	illumos/amd64 \
	js/wasm \
	linux/386 \
	linux/amd64 \
	linux/arm \
	linux/arm64 \
	linux/ppc64 \
	linux/ppc64le \
	linux/mips \
	linux/mipsle \
	linux/mips64 \
	linux/mips64le \
	linux/s390x \
	netbsd/386 \
	netbsd/amd64 \
	netbsd/arm \
	openbsd/386 \
	openbsd/amd64 \
	openbsd/arm \
	openbsd/arm64 \
	plan9/386 \
	plan9/amd64 \
	plan9/arm \
	solaris/amd64 \
	windows/386 \
	windows/amd64

build-all:
	$(MAKE) build-all-target -j 4

build-all-target: $(patsubst %,build/%,$(targets))

build/%:
	-GOOS=$(*D) GOARCH=$(*F) \
			 go build -o $(CURDIR)/binaries/$(*)/watcher

download-links:
	@find binaries | grep watcher | xargs -I{} sh -c 'echo - [{}]\(https://github.com/u110/watcher/raw/master/{}\)' | sort
