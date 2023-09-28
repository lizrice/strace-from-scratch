.POSIX:
.SUFFIXES:
NAME = strace

default: help
	exit 1
help:
	@ printf 'Usage:\n\n\
	    make <command>\n\n\
	Commands:\n\n\
	    build      build program\n\
	    clean      remove built files\n\
	    docker     build Docker image\n\
	    goversion  print Go version (from go.mod)\n\
	    help       print usage\n\
	    release    publish\n\
	    test       run tests\n'
all: build
install: build
build:
	go build -o $(NAME) .
clean: cleantag
	rm -f $(NAME)
docker: goversion
	docker build -t $(NAME) --build-arg GO_VERSION=$$(make goversion) .
release: test cleantag tag
	git tag "$$(cat tag)" # https://go.dev/doc/modules/publishing
	git push origin "$$(cat tag)"
	GOPROXY=proxy.golang.org go list -m \
	    "github.com/tingstad/strace@$$(cat tag)"
tag:
	git tag --list
	@ echo "Please enter new tag:"\
	  ; read tag\
	  ; echo "$$tag" > tag
cleantag:
	rm -f tag
test:
	go test ./...
goversion:
	@ awk '$$1 == "go" && $$2 ~ /^[1-9]/ && !n++ { print $$2 } END{ if(!n)\
	    print "ERROR finding Go version in go.mod" | "cat >&2"; exit !n }' go.mod
