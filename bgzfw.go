package main

import "fmt"

/*
#cgo CFLAGS: -I/data-sdd/scripts/lib/htslib-1.3.1
#cgo LDFLAGS: -L/data-sdd/scripts/lib/htslib-1.3.1 -lhts
#include <stdio.h>
#include <stdlib.h>
#include "htslib/bgzf.h"
*/
import "C"

import "unsafe"

type BGZFh struct {
  H *C.BGZF
}

func BGZFOpen(fn string, mode string) (h BGZFh,e error) {
  var fp *C.BGZF

  c_ifn := C.CString(fn)
  defer C.free(unsafe.Pointer(c_ifn))
  c_mod := C.CString(mode)
  defer C.free(unsafe.Pointer(c_mod))

  fp = C.bgzf_open(c_ifn, c_mod)

  if fp==nil {
    e = fmt.Errorf("could not open file")
    return
  }

  h.H = fp

  return
}

func (bgz *BGZFh) IndexLoad(fn, sfx string) (e error) {
  c_ifn := C.CString(fn)
  defer C.free(unsafe.Pointer(c_ifn))
  c_sfx := C.CString(sfx)
  defer C.free(unsafe.Pointer(c_sfx))

  r := C.bgzf_index_load(bgz.H, c_ifn, c_sfx)
  if r < 0 { e = fmt.Errorf("could not load index"); return }

  return
}

func (bgz *BGZFh) USeek(off int) (e error) {

  c_off := C.long(off)

  C.bgzf_useek(bgz.H, c_off, C.SEEK_SET)

  return
}

func (bgz *BGZFh) Read(b []byte) (e error) {
  sz := C.size_t(len(b))

  //z := C.malloc(sz)
  //defer C.free(unsafe.Pointer(z))

  zbuf := unsafe.Pointer(&b[0])

  r := C.bgzf_read(bgz.H, zbuf, sz)
  if C.size_t(r) != C.size_t(sz) { e = fmt.Errorf("bad read"); return }

  return
}


func (bgz *BGZFh) Close() (e error) {
  r := C.bgzf_close(bgz.H)
  if r < 0 { e = fmt.Errorf("close error") ; return }

  return
}


/*
func main() {
  //fn := "/mnt/tmpfs/0005.tar.gz"
  fn := "0005.tar.gz"
  fmt.Printf("ok\n")

  //bgz,e := bgzf_open("0003.tar.gz", "rb")
  //bgz,e := BGZFOpen("0005.tar.gz", "rb")
  bgz,e := BGZFOpen(fn, "rb")
  if e!=nil { panic(e) }
  defer bgz.Close()

  //e = bgz.IndexLoad("0003.tar.gz", ".gzi")
  //e = bgz.IndexLoad("0005.tar.gz", ".gzi")
  e = bgz.IndexLoad(fn, ".gzi")
  if e!=nil { panic(e) }

  //b := make([]byte, 340)
  //bgz.Seek(11776)
  b := make([]byte, 121472)
  e = bgz.USeek(230312960)
  if e!=nil { panic(e) }

  e = bgz.Read(b)
  if e!=nil { panic(e) }

  for i:=0; i<10; i++ {
    fmt.Printf(" %02x", b[i])
  }
  fmt.Printf("\n")


  //bgz.Close()

}
*/
