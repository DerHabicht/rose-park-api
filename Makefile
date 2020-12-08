default: bin/rose-park

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
SRC = $(basename $(wildcard */*.go))
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

bin/rose-park: main.go routes.go $(foreach f, $(SRC), $(f).go)
	swag init
	go build ${LDFLAGS} -o bin/rose-park

.PHONY: install
install: bin/rose-park
	-@rm ${GOPATH}/bin/rose-park
	cp bin/rose-park ${GOPATH}/bin/

.PHONY: run
run: bin/rose-park
	air -d -c .air.conf

.PHONY: clean
clean:
	-@rm bin/rose-park
