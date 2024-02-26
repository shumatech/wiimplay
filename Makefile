VERSION?=$(shell git describe --tags --always)
BUILD:=$(shell date +%FT%T)
FLAGS=-tags=gtk_3_10 -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"
Q?=@
SRC=main.go \
    config/config.go \
    wiimdev/types.go \
    wiimdev/wiim.go \
    wiimdev/fake.go \
    upnp/schema.go \
    upnp/wiim.go \
    dcache/dcache.go \
    mpris/server.go \
    ui/controller.go \
    ui/controls.go \
    ui/gotk3.go \
    ui/settings.go \
    ui/view.go \
    ui/res/res.go \
    ui/util.go

wiimplay: $(SRC)
	$(Q)go build $(FLAGS) .
	$(Q)strip $@
	$(Q)upx $@

clean:
	$(Q)rm -f wiimplay
