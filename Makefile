deps:
	go get github.com/jinzhu/gorm
	go get github.com/go-martini/martini
	go get github.com/martini-contrib/sessions
	go get github.com/martini-contrib/render
	go get github.com/martini-contrib/binding
	go get github.com/lib/pq
	bower install

test:
	@find . -name "*_test.go" | xargs dirname | xargs go test $(O)
