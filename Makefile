################################################################################
# Build Variables
################################################################################

BUILDFLAGS = -tags $(GTK_VERSION)
LDFLAGS    = -ldflags="-X main.appVersion=$(APP_VERSION)"
POTFLAGS   = --from-code=utf-8 \
             -kgt -kgtn \
             --package-name="$(APP)"

################################################################################
# Install Variables
################################################################################

prefix  = /usr
bindir  = $(prefix)/bin
datadir = $(prefix)/share

################################################################################
# Application Variables
################################################################################

APP = com.github.rkoesters.xkcd-gtk

EXE_NAME     = $(APP)
ICON_NAME    = $(APP).svg
DESKTOP_NAME = $(APP).desktop
APPDATA_NAME = $(APP).appdata.xml
POT_NAME     = $(APP).pot

EXE_PATH     = $(EXE_NAME)
ICON_PATH    = data/$(ICON_NAME)
DESKTOP_PATH = data/$(DESKTOP_NAME)
APPDATA_PATH = data/$(APPDATA_NAME)
POT_PATH     = po/$(APP).pot

################################################################################
# Automatic Variables
################################################################################

GO_SOURCES  = $(shell find . -name '*.go' -type f)
IN_SOURCES  = $(shell find . -name '*.ui' -type f)
GEN_SOURCES = $(patsubst %,%.go,$(IN_SOURCES))
SOURCES     = $(GO_SOURCES) $(GEN_SOURCES)
DEPS        = $(shell tools/list-imports.sh ./...)
PO          = $(shell find po -name '*.po' -type f)
MO          = $(patsubst %.po,%.mo,$(PO))

APP_VERSION = $(shell tools/app-version.sh)
GTK_VERSION = $(shell tools/gtk-version.sh)

# If GOPATH isn't set, then just use the current directory.
ifeq "$(shell go env GOPATH)" ""
export GOPATH = $(shell pwd)
endif
ifeq "$(shell go env GOPATH)" "/nonexistent/go"
export GOPATH = $(shell pwd)
endif

################################################################################
# Targets
################################################################################

all: $(EXE_PATH) $(DESKTOP_PATH) $(APPDATA_PATH) $(POT_PATH) $(MO)

deps:
	go get -u $(BUILDFLAGS) $(DEPS)

$(EXE_PATH): Makefile $(SOURCES)
	go build -o $@ $(BUILDFLAGS) $(LDFLAGS) ./cmd/xkcd-gtk

$(POT_PATH): $(shell cat po/POTFILES)
	xgettext -o $@ $(POTFLAGS) $^

%.ui.go: %.ui
	tools/go-wrap.sh $< >$@

%.desktop: %.desktop.in
	msgfmt -o $@ --desktop -d po --template $<

%.xml: %.xml.in
	msgfmt -o $@ --xml -d po --template $<

%.mo: %.po
	msgfmt -o $@ $<

check:
	-go fmt ./...
	-go vet ./...
	-golint ./...

clean:
	-rm -f $(EXE_PATH) $(GEN_SOURCES) $(DESKTOP_PATH) $(APPDATA_PATH) $(MO)

install: $(EXE_PATH)
	mkdir -p $(DESTDIR)$(bindir)
	install $(EXE_PATH) $(DESTDIR)$(bindir)
	mkdir -p $(DESTDIR)$(datadir)/icons/hicolor/scalable/apps
	cp $(ICON_PATH) $(DESTDIR)$(datadir)/icons/hicolor/scalable/apps
	mkdir -p $(DESTDIR)$(datadir)/applications
	cp $(DESKTOP_PATH) $(DESTDIR)$(datadir)/applications
	mkdir -p $(DESTDIR)$(datadir)/metainfo
	cp $(APPDATA_PATH) $(DESTDIR)$(datadir)/metainfo

uninstall:
	rm $(DESTDIR)$(bindir)/$(EXE_NAME) \
	   $(DESTDIR)$(datadir)/icons/hicolor/scalable/apps/$(ICON_NAME) \
	   $(DESTDIR)$(datadir)/applications/$(DESKTOP_NAME) \
	   $(DESTDIR)$(datadir)/metainfo/$(APPDATA_NAME)

.PHONY: all check clean deps install uninstall
