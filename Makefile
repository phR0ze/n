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
	go test -gcflags=-l ./pkg/arch/tar
	go test -gcflags=-l ./pkg/arch/zip
	go test ./pkg/enc/bin
	go test ./pkg/enc/json
	go test ./pkg/enc/unit
	go test ./pkg/enc/yaml
	go test ./pkg/errs
	go test ./pkg/futil
	go test ./pkg/net
	go test ./pkg/opt
	go test -gcflags=-l ./pkg/sys
	go test ./pkg/term
	go test -gcflags=-l ./pkg/time
	go test ./pkg/tmpl

cover: ${NAME}
	@echo -e "\nRunning go coverage tests:"
	@echo -e "------------------------------------------------------------------------"
	go test -gcflags=-l -coverprofile=coverage.out ./pkg/${pkg}
	go tool cover -html=coverage.out

bench: vendor
	@echo -e "\nRunning all go benchmarks:"
	@echo -e "------------------------------------------------------------------------"
	go test -bench=.

clean:
	rm -rf ./vendor
	rm -f ./bin/${NAME}
