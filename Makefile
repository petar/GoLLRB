# Copyright 2010 Petar Maymounkov. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all:	install

install:
	cd src && make install

clean:
	cd src && make clean

nuke:
	cd src && make nuke
