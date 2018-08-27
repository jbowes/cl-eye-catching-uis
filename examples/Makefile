PLATFORMS=darwin linux windows
EXAMPLES=$(filter-out Makefile $(PLATFORMS), $(wildcard * ))

all: $(PLATFORMS:=-all)

darwin-all linux-all: %-all: $(EXAMPLES:%=\%/%)
windows-all: %-all: $(EXAMPLES:%=windows/%.exe)

.SECONDEXPANSION:
$(EXAMPLES:%=darwin/%): darwin/%: $$(wildcard %/*.go)
	GOOS=darwin go build -o $@ ./$*

.SECONDEXPANSION:
$(EXAMPLES:%=linux/%): linux/%: $$(wildcard %/*.go)
	GOOS=linux go build -o $@ ./$*

.SECONDEXPANSION:
$(EXAMPLES:%=windows/%.exe): windows/%.exe: $$(wildcard %/*.go)
	GOOS=windows go build -o $@ ./$*
