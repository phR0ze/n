NAME := nub
DEPDIR := vendor

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
	go test

clean:
	rm -rf ./vendor
	rm -f ./bin/$(NAME)
