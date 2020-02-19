OUT := mgr-clbd
DISTRO := $(shell cat /etc/os-release | grep '^NAME=' | sed -e 's/.*=//')
GCCGO :=
GO :=
ifeq ($(DISTRO),"Ubuntu")
	GCCGO += gccgo
	GO += go
else ifeq (($DISTRO),"openSUSE Leap")
	GCCGO += gccgo-8
	GO += go-8
endif

SOURCES = \
	*.go

all: update-deps swagger mainbuild strip

mainbuild: $(SOURCES)
	$(GO) build -x -o $(OUT) $(SOURCES)

gcc:
ifneq ($GCCGO,)
	$(GO) build -compiler $(GCCGO) -gccgoflags '-static-libgo' -o $(OUT)
else
	@echo "Cannot find gccgo"
endif

strip:
	strip $(OUT)

update-deps:
	$(GO) get -u

deps:
	$(GO) get -v ./... 2>/dev/null ||:
strip:
	strip $(OUT)

clean:
	$(GO) clean -x -i

flush-cache:
	$(GO) clean -cache -modcache -i -r

swagger:
ifeq (, $(shell which swag))
	$(GO) get -u github.com/swaggo/swag/cmd/swag
endif
	cd ..;swag init -g cmd/mgr-clbd.go -o ./docs

.PHONY: all install clean
