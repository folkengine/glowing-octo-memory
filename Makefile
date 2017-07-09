binary = glowing

.DEFAULT_GOAL := install

run: clean test
	$(clean)
	@go fmt
	@go build
	@./$(binary)

test:
	@echo 'Running Tests'
	@go fmt
	@go test
	@go vet
	@golint -set_exit_status
	@echo 'Tests Completed'

clean:
	@rm -f $(binary)
