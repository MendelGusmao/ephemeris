run:
	@fresh

deps:
	go get -u github.com/go-martini/martini
	go get -u github.com/lib/pq
	go get -u github.com/martini-contrib/binding
	go get -u github.com/martini-contrib/render
	go get -u github.com/martini-contrib/sessions
	go get -u github.com/MendelGusmao/go-testdb
	go get -u github.com/MendelGusmao/gorm
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/ginkgo/ginkgo
	go get -u github.com/onsi/gomega
	go get -u github.com/ae0000/fresh

frontend: deps	
	bower install

test:
	@ginkgo $(O) -race `make find_tests`

t: test

cov:
	@find $(PWD) -name "*.coverprofile" -delete
	@ginkgo -cover -race $(O) `make find_tests`
	@echo "mode: set" > ephemeris.coverprofile
	@find $(PWD) -name "*.coverprofile" | \
		xargs -I@ cat "@" | \
		grep -v "mode:" | \
		sort -r | \
		awk '{if($$1 != last) {print $$0;last=$$1}}' >> ephemeris.coverprofile
	@go tool cover -html=ephemeris.coverprofile

find_tests:
	@find $(PWD) -name "*_test.go" | xargs dirname | uniq
