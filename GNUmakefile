PROJECT=go-dbt1
VERSION=1.0.0
PREFIX=/usr/local
all:
clean:
install:

## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/csv-split$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/csv-split$(EXE)
install-go:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp  $(BUILDDIR)/csv-split$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/csv-split$(EXE)
##
$(BUILDDIR)/csv-split$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/csv-split
## -- BLOCK:go --
