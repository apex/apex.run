
SRC = $(wildcard views/*.html)
SRC += $(wildcard cmd/build/*.go)
SRC += $(wildcard ../apex/docs/*.md)

index.html: $(SRC)
	@go run cmd/build/build.go > $@

commit: index.html
	git commit -a -m "Build docs for $(shell go run ../apex/cmd/apex/main.go version)"
	git push
.PHONY: commit

clean:
	rm -f index.html
.PHONY: clean
