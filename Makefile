deps:
	git submodule init
	git submodule update
	bower install

test:
	@find . -name "*_test.go" | xargs dirname | xargs go test $(O)