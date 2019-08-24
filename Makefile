NAME := n

default: ${NAME}
${NAME}: vendor
	@echo "Building..."
	@echo "------------------------------------------------------------------------"
	go build -o bin/${NAME} .
	

mech: vendor
	go build -o bin/mech pkg/net/mech/example/mech.go

vendor:
	dep ensure -v

test: ${NAME}
	@echo -e "\nRunning all go tests:"
	@echo -e "------------------------------------------------------------------------"
	go test ./pkg/arch/tar
	go test ./pkg/arch/zip
	go test ./pkg/bin
	go test ./pkg/cli
	go test ./pkg/errs
	go test ./pkg/net
	go test ./pkg/opt
	go test ./pkg/sys
	go test ./pkg/term
	go test ./pkg/time
	go test ./pkg/tmpl
	go test ./pkg/unit

bench: ${NAME}
	@echo -e "\nRunning all go benchmarks:"
	@echo -e "------------------------------------------------------------------------"
	go test -bench=.

clean:
	rm -rf ./vendor
	rm -f ./bin/${NAME}
