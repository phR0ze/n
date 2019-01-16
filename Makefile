NAME := n
DEPDIR := vendor
PKGROOT := github.com/phR0ze/n

default: $(NAME)
$(NAME): $(DEPDIR)
	go build -o bin/$(NAME) .

install:
	dep ensure -v

update:
	dep ensure -v -update

test: $(NAME)
	@echo -e "\nRunning all go tests:"
	@echo -e "------------------------------------------------------------------------"
	go test $(PKGROOT)
	go test $(PKGROOT)/pkg/ncli
	go test $(PKGROOT)/pkg/nerr
	go test $(PKGROOT)/pkg/nnet
	go test $(PKGROOT)/pkg/ntar
	go test $(PKGROOT)/pkg/ntime
	go test $(PKGROOT)/pkg/ntmpl
	go test $(PKGROOT)/pkg/sys
	go test $(PKGROOT)/pkg/term
	go test $(PKGROOT)/pkg/trace

bench: $(NAME)
	@echo -e "\nRunning all go benchmarks:"
	@echo -e "------------------------------------------------------------------------"
	go test -bench=.

clean:
	rm -rf ./vendor
	rm -f ./bin/$(NAME)
