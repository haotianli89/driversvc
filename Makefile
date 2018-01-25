.PHONY: pb data lint run

current_dir = $(shell pwd)

pb:
	for f in pb/*.proto; do \
		protoc -I${GOPATH}/src --go_out=plugins=micro:${GOPATH}/src ${current_dir}/pb/driversvc.proto; \
	done
