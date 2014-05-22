deps:
	git submodule init
	bower install

test:
	@find . -name "*_test.go" | xargs dirname | xargs go test $(O)