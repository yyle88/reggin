COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/gormrepo/blob/10d9b8056278298592769425c88c2075857f4848/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
