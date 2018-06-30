PREFIX=/usr/local
DESTDIR=
GOFLAGS=
BINDIR=${PREFIX}/bin

BLDDIR = build/${GOOS}
EXT=
ifeq (${GOOS},windows)
	EXT=.exe
endif

APPS=psdui psdui-gen-exml

all: $(APPS)

$(BLDDIR)/psdui:        $(wildcard cmd/psdui/*.go psdui/*.go)
$(BLDDIR)/psdui-gen-exml:  $(wildcard cmd/psdui-gen-exml/*.go psdui/*.go)

$(BLDDIR)/%:
	@mkdir -p $(dir $@)
	go build ${GOFLAGS} -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	rm -fr $(BLDDIR)

.PHONY: install clean all
.PHONY: $(APPS)

install: $(APPS)
	@echo ${BLDDIR}
	install -m 755 -d ${DESTDIR}${BINDIR}
	for APP in $^ ; do install -m 755 ${BLDDIR}/$$APP ${DESTDIR}${BINDIR}/$$APP${EXT} ; done
