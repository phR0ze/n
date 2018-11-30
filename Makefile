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
	go test $(PKGROOT)/pkg/tarn
	go test $(PKGROOT)/pkg/timen
	go test $(PKGROOT)/pkg/tmpl
	go test $(PKGROOT)/pkg/tracen

bench: $(NAME)
	@echo -e "\nRunning all go benchmarks:"
	@echo -e "------------------------------------------------------------------------"
	go test -bench=.

clean:
	rm -rf ./vendor
	rm -f ./bin/$(NAME)
