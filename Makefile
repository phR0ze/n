NAME := n
PKGROOT := github.com/phR0ze/$(NAME)
GOFLAGS := -mod=vendor

default: $(NAME)
$(NAME): vendor
	@echo "Building..."
	@echo "------------------------------------------------------------------------"
	go build ${GOFLAGS} -o bin/$(NAME) .
	

mech: vendor
	go build ${GOFLAGS} -o bin/mech pkg/net/mech/example/mech.go

vendor:
	go mod vendor

test: $(NAME)
	@echo -e "\nRunning all go tests:"
	@echo -e "------------------------------------------------------------------------"
	go test ${GOFLAGS} $(PKGROOT)
	go test ${GOFLAGS} $(PKGROOT)/pkg/arch/tar
	go test ${GOFLAGS} $(PKGROOT)/pkg/arch/zip
	go test ${GOFLAGS} $(PKGROOT)/pkg/bin
	go test ${GOFLAGS} $(PKGROOT)/pkg/cli
	go test ${GOFLAGS} $(PKGROOT)/pkg/errs
	go test ${GOFLAGS} $(PKGROOT)/pkg/net
	go test ${GOFLAGS} $(PKGROOT)/pkg/opt
	go test ${GOFLAGS} $(PKGROOT)/pkg/sys
	go test ${GOFLAGS} $(PKGROOT)/pkg/term
	go test ${GOFLAGS} $(PKGROOT)/pkg/time
	go test ${GOFLAGS} $(PKGROOT)/pkg/tmpl
	go test ${GOFLAGS} $(PKGROOT)/pkg/unit

bench: $(NAME)
	@echo -e "\nRunning all go benchmarks:"
	@echo -e "------------------------------------------------------------------------"
	go test ${GOFLAGS} -bench=.

clean:
	rm -rf ./vendor
	rm -f ./bin/$(NAME)
