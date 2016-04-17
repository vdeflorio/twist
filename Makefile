SUFFIX=.exe
INSTALLDIR=/home/deflorio/cygdrive/c/bin
CCOPTS=-O3
INSTALL=/usr/bin/install -b -v

all:	twist

twist:	twist.c
	gcc ${CCOPTS} -o twist twist.c

install:
	${INSTALL} twist${SUFFIX} ${INSTALLDIR}

dist:
	rar a twist-1.0.rar twist.c
