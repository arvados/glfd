#!/bin/bash
#
# Requires asm_ukk.a from the 'asmukk' repository
#

#libdir="/data-sde/scripts/lib/htslib-1.3.1"
#incdir="/data-sde/scripts/lib/htslib-1.3.1"
#gcc -L$libdir -I$incdir -lhts cbgz.c -o cbgz -lhts

go build glfd.go glfd_web.go glfd_js.go glfd_init.go bgzfw.go asmukkgo.go

