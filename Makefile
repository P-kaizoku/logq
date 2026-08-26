.PHONY: logs run

run:
	@go run .

logs:
	@go run ./gen/gen.go > testdata/gen.log
