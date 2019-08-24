NAME := n
GOFLAGS := -mod=vendor

default: ${NAME}
${NAME}: vendor
	@echo "Building..."
	@echo "------------------------------------------------------------------------"
	go build ${GOFLAGS} -o bin/${NAME} .
	

mech: vendor
	go build ${GOFLAGS} -o bin/mech pkg/net/mech/example/mech.go

vendor:
	go mod vendor

test: ${NAME}
	@echo -e "\nRunning all go tests:"
	@echo -e "------------------------------------------------------------------------"
	go test ${GOFLAGS} ./pkg/arch/tar
	go test ${GOFLAGS} ./pkg/arch/zip
	go test ${GOFLAGS} ./pkg/bin
	go test ${GOFLAGS} ./pkg/cli
	go test ${GOFLAGS} ./pkg/errs
	go test ${GOFLAGS} ./pkg/net
	go test ${GOFLAGS} ./pkg/opt
	go test ${GOFLAGS} ./pkg/sys
	go test ${GOFLAGS} ./pkg/term
	go test ${GOFLAGS} ./pkg/time
	go test ${GOFLAGS} ./pkg/tmpl
	go test ${GOFLAGS} ./pkg/unit

bench: ${NAME}
	@echo -e "\nRunning all go benchmarks:"
	@echo -e "------------------------------------------------------------------------"
	go test ${GOFLAGS} -bench=.

clean:
	rm -rf ./vendor
	rm -f ./bin/${NAME}
