run:
	@fresh

deps:
	go get github.com/jinzhu/gorm
	go get github.com/go-martini/martini
	go get github.com/martini-contrib/sessions
	go get github.com/martini-contrib/render
	go get github.com/martini-contrib/binding
	go get github.com/lib/pq
	go get github.com/pilu/fresh
	go get github.com/ae0000/fresh
	go get github.com/erikstmartin/go-testdb
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega
	bower install

test:
	@go test ./... $(O)

integration-test:
	@echo not implemented yet

cov:
	@$(PWD)/coverage.py
