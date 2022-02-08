.PHONY: tidy
tidy:
	GOPRIVATE=github.com/RafaySystems/* go mod tidy
.PHONY: vendor
vendor:
	go mod vendor

.PHONY: build-proto
build-proto:
	cd components; buf build

.PHONY: gen-proto
gen-proto:
	cd components/common; buf generate
	cd components/adminsrv; buf generate
	cd components/usermgmt; buf generate
	cd components/cluster-scheduler; buf generate

.PHONY: test
test:
	cd components/common; go test ./...
	cd components/adminsrv; go test ./...
	cd components/usermgmt; go test ./...
	cd components/cluster-scheduler; go test ./...

.PHONY: check
check:
	cd components/common; go fmt ./...
	cd components/adminsrv; go fmt ./...
	cd components/usermgmt; go fmt ./...
	cd components/cluster-scheduler; go fmt ./...

	cd components/common; go vet ./...
	cd components/adminsrv; go vet ./...
	cd components/usermgmt; go vet ./...
	cd components/cluster-scheduler; go vet ./...

.PHONY: clean
clean:
	rm -rf components/**/gen
	find . -name "*.pb*" -type f -delete