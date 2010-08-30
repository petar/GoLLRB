# Copyright 2010 Petar Maymounkov. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all:	install

install:
	cd llrb && make install

clean:
	cd llrb && make clean

nuke:
	cd llrb && make nuke
