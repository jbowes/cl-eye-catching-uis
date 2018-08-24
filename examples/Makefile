PLATFORMS=darwin linux windows
EXAMPLES=$(filter-out Makefile $(PLATFORMS), $(wildcard * ))

all: $(PLATFORMS:=-all)

$(PLATFORMS:=-all): %-all: $(EXAMPLES:%=\%/%)

$(EXAMPLES:%=darwin/%): darwin/%: $(wildcard %/*.go)
	GOOS=darwin go build -o $@ ./$*

$(EXAMPLES:%=linux/%): linux/%: $(wildcard %/*.go)
	GOOS=linux go build -o $@ ./$*

$(EXAMPLES:%=windows/%): windows/%: $(wildcard %/*.go)
	GOOS=windows go build -o $@.exe ./$*
