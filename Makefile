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
	bower install

test:
	@find . -name "*_test.go" | xargs dirname | xargs go test $(O)

integration-test:
	@echo not implemented yet

coverage:
	@$(PWD)/coverage.py
