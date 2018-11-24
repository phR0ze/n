NAME := nub
DEPDIR := vendor
PKGROOT := github.com/phR0ze/nub

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
	go test $(PKGROOT)/pkg/time
	go test $(PKGROOT)/pkg/trace

clean:
	rm -rf ./vendor
	rm -f ./bin/$(NAME)
