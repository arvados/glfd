package main

import "os"
import "fmt"
import "bytes"
import "bufio"
import "strings"
import "strconv"
import "io/ioutil"

import "crypto/md5"

import "time"

import "github.com/aebruno/twobit"

import "github.com/abeconnelly/pasta"
import "github.com/abeconnelly/pasta/gvcf"

type GLFD struct {

  // e.g. Assembly["hg19"][0x2fb][0x30]
  // Maps (assembly) tile path and tile step
  // to ending reference position
  // ~170Mb
  //
  Assembly map[string]map[int]map[int]int

  TilepathToChrom map[int]string

  SeqCache map[int]map[int]map[int]string

  RefV map[string]map[int][]int

  Tagset map[int]string

  // holds only non 1 span tile information
  //
  //TileLibSpan map[int]map[int]map[int]int
  TileLibSpan map[uint64]int

  // Location of library files
  //
  GLFDir string
}

func Md5sum2str(md5sum [16]byte) string {
  var str_md5sum [32]byte
  for i:=0; i<16; i++ {
    x := fmt.Sprintf("%02x", md5sum[i])
    str_md5sum[2*i]   = x[0]
    str_md5sum[2*i+1] = x[1]
  }

  return string(str_md5sum[:])
}

func (glfd *GLFD) TagEnd(tilepath int, tilelibver int, tilestep int) (string, error) {
  _ = tilelibver

  if _,ok := glfd.Tagset[tilepath] ; ok {
    pos := tilestep * 24
    if (pos<0) || ((pos+24)>len(glfd.Tagset[tilepath])) {
      return "", fmt.Errorf("tag out of range")
    }

    return glfd.Tagset[tilepath][pos:pos+24], nil
  }

  return "", fmt.Errorf("tilepath not found")
}

func (glfd *GLFD) TileSpan(tilepath, tilelibver, tilestep, tilevar int) (int, error) {

  /*
  _ = tilelibver
  if _,ok := glfd.TileLibSpan[tilepath] ; ok {
    if _,oks := glfd.TileLibSpan[tilepath][tilestep] ; oks  {
      if _,okv := glfd.TileLibSpan[tilepath][tilestep][tilevar] ; okv {
        return glfd.TileLibSpan[tilepath][tilestep][tilevar],nil
      }
    }
  }
  */

  var key uint64

  key = key | (uint64(tilepath) << (8*6))
  key = key | (uint64(tilelibver) << (8*4))
  key = key | (uint64(tilestep) << (8*2))
  key = key | uint64(tilevar)

  if span,ok := glfd.TileLibSpan[key] ; ok { return span, nil }
  return 1,nil
}


func (glfd *GLFD) TileLibSequences(tilepath int, tilelibver int, tilestep int) (string, error) {
  base_dir := glfd.GLFDir

  step_ifn := fmt.Sprintf("%04x.%02x.%04x.2bit", tilepath, tilelibver, tilestep) ; _ = step_ifn
  gz_ifn := fmt.Sprintf("%s/%04x.tar.gz", base_dir, tilepath) ; _ = gz_ifn
  taridx_ifn := fmt.Sprintf("%s/%04x.tar.tai", base_dir, tilepath) ; _ = taridx_ifn

  idx_fp,err := os.Open(taridx_ifn)
  if err!=nil { return "", err }
  defer idx_fp.Close()

  vbyte_beg := -1
  vbyte_len := -1

  scanner := bufio.NewScanner(idx_fp)
  for scanner.Scan() {
    line := scanner.Text()
    line_parts := strings.Split(line, " ")
    if line_parts[0] != step_ifn { continue }

    vbyte_beg,_ = strconv.Atoi(line_parts[1])
    vbyte_len,_ = strconv.Atoi(line_parts[2])
  }

  bgz,e := BGZFOpen(gz_ifn, "r")
  if e!=nil { return "", e }
  _ = bgz.IndexLoad(gz_ifn, ".gzi")
  defer bgz.Close()

  b := make([]byte, vbyte_len)
  bgz.USeek(vbyte_beg)
  bgz.Read(b)

  b_rdr := bytes.NewReader(b)
  tb_rdr,e := twobit.NewReader(b_rdr)
  if e!=nil { return "", e }

  ret_a := []string{}

  names := tb_rdr.Names()
  for idx:=0; idx<len(names); idx++ {

    seq,e := tb_rdr.Read(names[idx])
    if e!=nil {return "", e}

    m5str := Md5sum2str( md5.Sum(seq) )

    s := fmt.Sprintf(`{"tile-id":"%s","md5sum":"%s","seq":"%s"}`,
      names[idx], m5str, seq)
    ret_a = append(ret_a, s)
  }

  return "[" + strings.Join(ret_a, ",") + "]", nil
}

func (glfd *GLFD) TileSequence(tilepath int, tilelibver int, tilestep int, tilevar int) (string, error) {

  if _,okp := glfd.SeqCache[tilepath] ; okp {
    if _,oks := glfd.SeqCache[tilepath][tilestep] ; oks {
      if seq,ok := glfd.SeqCache[tilepath][tilestep][tilevar] ; ok {

        //DEBUG
        fmt.Printf("hit! (ts) %x.%x\n", tilepath, tilestep)

        return seq, nil
      }
    }
  }

  //DEBUG
  fmt.Printf("miss (ts) %x.%x\n", tilepath, tilestep)

  //base_dir := "/mnt/tmpfs"
  base_dir := glfd.GLFDir

  step_ifn := fmt.Sprintf("%04x.%02x.%04x.2bit", tilepath, tilelibver, tilestep) ; _ = step_ifn
  gz_ifn := fmt.Sprintf("%s/%04x.tar.gz", base_dir, tilepath) ; _ = gz_ifn
  taridx_ifn := fmt.Sprintf("%s/%04x.tar.tai", base_dir, tilepath) ; _ = taridx_ifn

  //ref_tileid := fmt.Sprintf("%04x.%02x.%04x.000", tilepath, tilelibver, tilestep)
  tileid := fmt.Sprintf("%04x.%02x.%04x.%03x", tilepath, tilelibver, tilestep, tilevar)

  idx_fp,err := os.Open(taridx_ifn)
  if err!=nil { return "", err }
  defer idx_fp.Close()

  vbyte_beg := -1
  vbyte_len := -1

  scanner := bufio.NewScanner(idx_fp)
  for scanner.Scan() {
    line := scanner.Text()
    line_parts := strings.Split(line, " ")
    if line_parts[0] != step_ifn { continue }

    vbyte_beg,_ = strconv.Atoi(line_parts[1])
    vbyte_len,_ = strconv.Atoi(line_parts[2])
  }

  bgz,e := BGZFOpen(gz_ifn, "r")
  if e!=nil { return "", e }
  _ = bgz.IndexLoad(gz_ifn, ".gzi")
  defer bgz.Close()

  b := make([]byte, vbyte_len)
  bgz.USeek(vbyte_beg)
  bgz.Read(b)

  b_rdr := bytes.NewReader(b)
  tb_rdr,e := twobit.NewReader(b_rdr)
  if e!=nil { return "", e }

  seq,e := tb_rdr.Read(tileid)
  if e!=nil { return "", e}

  return string(seq), nil
}

// loq_info holds a packed integer array where even entries are start position and odd entries are length of the nocall sequence.
//
func (glfd *GLFD)  TileSequenceLoq(tilepath int, tilelibver int, tilestep int, tilevar int, loq_info []int) (string, error) {

  seq := []byte{}

  if _,okp := glfd.SeqCache[tilepath] ; okp {
    if _,oks := glfd.SeqCache[tilepath][tilestep] ; oks {
      if seqstr,ok := glfd.SeqCache[tilepath][tilestep][tilevar] ; ok {

        //DEBUG
        fmt.Printf("hit! (tslq) %x.%x\n", tilepath, tilestep)

        seq = []byte(seqstr)

      }
    }
  }

  if len(seq)==0 {

    //DEBUG
    fmt.Printf("miss (tslq) %x.%x\n", tilepath, tilestep)


    //base_dir := "/mnt/tmpfs"
    base_dir := glfd.GLFDir

    step_ifn := fmt.Sprintf("%04x.%02x.%04x.2bit", tilepath, tilelibver, tilestep) ; _ = step_ifn
    gz_ifn := fmt.Sprintf("%s/%04x.tar.gz", base_dir, tilepath) ; _ = gz_ifn
    taridx_ifn := fmt.Sprintf("%s/%04x.tar.tai", base_dir, tilepath) ; _ = taridx_ifn

    //ref_tileid := fmt.Sprintf("%04x.%02x.%04x.000", tilepath, tilelibver, tilestep)
    tileid := fmt.Sprintf("%04x.%02x.%04x.%03x", tilepath, tilelibver, tilestep, tilevar)

    idx_fp,err := os.Open(taridx_ifn)
    if err!=nil { return "", err }
    defer idx_fp.Close()

    vbyte_beg := -1
    vbyte_len := -1

    scanner := bufio.NewScanner(idx_fp)
    for scanner.Scan() {
      line := scanner.Text()
      line_parts := strings.Split(line, " ")
      if line_parts[0] != step_ifn { continue }

      vbyte_beg,_ = strconv.Atoi(line_parts[1])
      vbyte_len,_ = strconv.Atoi(line_parts[2])
    }

    bgz,e := BGZFOpen(gz_ifn, "r")
    if e!=nil { return "", e }
    _ = bgz.IndexLoad(gz_ifn, ".gzi")
    defer bgz.Close()

    b := make([]byte, vbyte_len)
    bgz.USeek(vbyte_beg)
    bgz.Read(b)

    b_rdr := bytes.NewReader(b)
    tb_rdr,e := twobit.NewReader(b_rdr)
    if e!=nil { return "", e }

    //seq,e := tb_rdr.Read(tileid)
    seq,err = tb_rdr.Read(tileid)
    if err!=nil { return "", err}
  }

  for i:=0; i<len(loq_info); i+=2 {
    loq_s := loq_info[i]
    loq_n := loq_info[i+1]

    if loq_s < 0 { return "", fmt.Errorf("nocall start position less than 0"); }
    if loq_s+loq_n > len(seq) {
      return "", fmt.Errorf(fmt.Sprintf("nocall sequence oob (%04x.%02x.%04x.%03x, len %d, loq_s %d, loq_n %d)", tilepath, tilelibver, tilestep, tilevar, len(seq), loq_s, loq_n));
    }

    for x:=0; x<loq_n; x++ {
      seq[loq_s + x] = 'n'
    }
  }

  return string(seq), nil
}

func AlignToPasta(ref, x string) (string, error) {
  //var e error

  b := []byte{}

  if len(ref)!=len(x) { return "", fmt.Errorf("length mismatch") }

  for i:=0; i<len(ref); i++ {
    b = append(b, pasta.SubMap[ref[i]][x[i]])
  }

  return string(b), nil
}

func EmitGVCFHeader(outs *bufio.Writer) {
  //outs := bufio.NewWriter(os.Stdout)
  g := gvcf.GVCFRefVar{}
  g.Header(outs)
}

func EmitGVCF(outs *bufio.Writer, chrom string, ref_pos int, ref_seq, x_seq, y_seq string) {
  var e error

  x_ref,x_align,_ := dna_align(ref_seq, x_seq)
  y_ref,y_align,_ := dna_align(ref_seq, y_seq)

  x_pasta,_ := AlignToPasta(x_ref, x_align)
  y_pasta,_ := AlignToPasta(y_ref, y_align)

  x_pasta_rdr := bytes.NewReader([]byte(x_pasta))
  y_pasta_rdr := bytes.NewReader([]byte(y_pasta))
  xy_pasta_wtr := new(bytes.Buffer)

  e = pasta.InterleaveStreams(x_pasta_rdr, y_pasta_rdr, xy_pasta_wtr) ; _ = e
  //if e!=nil { return e }

  //outs := bufio.NewWriter(os.Stdout)

  g := gvcf.GVCFRefVar{}
  xy_pasta_rdr := bufio.NewReader(bytes.NewReader(xy_pasta_wtr.Bytes()))

  g.Chrom(chrom)
  g.Pos(ref_pos)

  e = pasta.InterleaveToDiffInterface(xy_pasta_rdr, &g, outs)
  if e!=nil { panic(e) }

}

/*
func (glfd *GLFD) TileToGVCF_old(tilepath, tilelibver, anchor_tilestep int, varid_knot [][]int, loq_info [][][]int, ref_varid []int) (string, error) {
  seq_a := [2][]string{}
  ref_a := []string{}

  trail_tag := [2]string{}
  trail_ref_tag := ""

  del_step := 0
  for i:=0; i<len(varid_knot); i++ {
    del_step = 0
    for pos:=0; pos<len(varid_knot[i]); pos+=2 {
      s,e := glfd.TileSequence(tilepath, tilelibver, anchor_tilestep + del_step, varid_knot[i][pos])
      if e!=nil { return "", e }

      trail_tag[i] = s[len(s)-24:]

      seq_a[i] = append(seq_a[i], s[:len(s)-24])
      del_step += varid_knot[i][pos+1]
    }
  }

  del_step = 0
  for pos:=0; pos<len(ref_varid); pos+=2 {
    s,e := glfd.TileSequence(tilepath, tilelibver, anchor_tilestep + del_step, ref_varid[pos])
    if e!=nil { return "", e }

    trail_ref_tag = s[len(s)-24:]

    ref_a = append(ref_a, s[:len(s)-24])
    del_step += ref_varid[pos+1]
  }

  seq_a[0] = append(seq_a[0], trail_tag[0])
  seq_a[1] = append(seq_a[1], trail_tag[1])
  ref_a = append(ref_a, trail_ref_tag)

  seq0 := strings.Join(seq_a[0], "")
  seq1 := strings.Join(seq_a[1], "")
  refseq := strings.Join(ref_a, "")

  fmt.Printf("seq0: %s\nseq1: %s\n ref: %s\n", seq0, seq1, refseq)

  chrom := glfd.TilepathToChrom[tilepath]
  ref_pos := 0
  if anchor_tilestep > 0 {
    ref_pos = glfd.Assembly["hg19"][tilepath][anchor_tilestep-1]
  }

  EmitGVCF(chrom, ref_pos, refseq, seq0, seq1)

  return "", nil
}
*/

/*
func (glfd *GLFD) TileToGVCFx(tilepath, tilelibver, anchor_tilestep int, varid_knot [][]int, loq_info [][][]int, ref_varid []int) (string, error) {
  seq_a := [2][]string{}
  ref_a := []string{}

  trail_tag := [2]string{}
  trail_ref_tag := ""

  for i:=0; i<len(varid_knot); i++ {
    for pos:=0; pos<len(varid_knot[i]); pos++ {

      //fmt.Printf(">> allele %i, pos %i, knot %i\n", i, pos, varid_knot[i][pos])

      if (varid_knot[i][pos]<0) { continue; }
      s,e := glfd.TileSequenceLoq(tilepath, tilelibver, anchor_tilestep + pos, varid_knot[i][pos], loq_info[i][pos])
      //s,e := glfd.TileSequence(tilepath, tilelibver, anchor_tilestep + pos, varid_knot[i][pos])
      if e!=nil { return "", e }

      trail_tag[i] = s[len(s)-24:]

      seq_a[i] = append(seq_a[i], s[:len(s)-24])
    }
  }

  for pos:=0; pos<len(ref_varid); pos++ {
    if (ref_varid[pos]<0) { continue; }
    s,e := glfd.TileSequence(tilepath, tilelibver, anchor_tilestep + pos, ref_varid[pos])
    if e!=nil { return "", e }

    trail_ref_tag = s[len(s)-24:]

    ref_a = append(ref_a, s[:len(s)-24])
  }

  seq_a[0] = append(seq_a[0], trail_tag[0])
  seq_a[1] = append(seq_a[1], trail_tag[1])
  ref_a = append(ref_a, trail_ref_tag)

  seq0 := strings.Join(seq_a[0], "")
  seq1 := strings.Join(seq_a[1], "")
  refseq := strings.Join(ref_a, "")

  fmt.Printf("seq0: %s\nseq1: %s\n ref: %s\n", seq0, seq1, refseq)

  chrom := glfd.TilepathToChrom[tilepath]
  ref_pos := 0
  if anchor_tilestep > 0 {
    ref_pos = glfd.Assembly["hg19"][tilepath][anchor_tilestep-1]
  }

  EmitGVCF(chrom, ref_pos, refseq, seq0, seq1)

  return "", nil
}
*/

func (glfd *GLFD) TileToGVCF(outs *bufio.Writer, tilepath, tilelibver, anchor_tilestep int, varid [][]int, loq_info [][][]int, ref_varid []int) (string, error) {
  seq_a := [2][]string{}
  ref_a := []string{}

  trail_tag := [2]string{}
  trail_ref_tag := ""

  if len(varid)!=2 { return "", fmt.Errorf("varid not valid (must have 2 alleles)") }
  if len(varid[0]) != len(varid[1]) { return "", fmt.Errorf("varid allele lengths do not match") }

  n := len(varid[0])
  if len(loq_info) != 2 { return "", fmt.Errorf("loq_info not valid (must have 2 alleles)") }
  if len(loq_info[0])!=len(loq_info[1]) { return "", fmt.Errorf("loq_info allele lengths do not match") }
  if n!=len(loq_info[0]) { return "", fmt.Errorf("loq_info lengths must match varid lengths") }
  if n!=len(ref_varid) { return "", fmt.Errorf("ref_varid length mismatch") }

  for step_idx:=0; step_idx < n; {

    //DEBUG
    fmt.Printf("## %d / %d\n", step_idx, n)

    z := step_idx+1
    for (z<n) && ((varid[0][z]<0) || (varid[1][z]<0)) { z++ }


    for i:=0; i<len(varid); i++ {
      seq_a[i] = seq_a[i][0:0]
      for pos:=step_idx; pos<z; pos++ {

        if (varid[i][pos]<0) { continue; }
        //s,e := glfd.TileSequenceLoq(tilepath, tilelibver, anchor_tilestep + pos, varid_knot[i][pos], loq_info[i][pos])
        s,e := glfd.TileSequenceLoq(tilepath, tilelibver, anchor_tilestep + pos, varid[i][pos], loq_info[i][pos])
        //s,e := glfd.TileSequence(tilepath, tilelibver, anchor_tilestep + pos, varid_knot[i][pos])
        if e!=nil { return "", e }

        trail_tag[i] = s[len(s)-24:]

        seq_a[i] = append(seq_a[i], s[:len(s)-24])
      }
    }

    ref_a = ref_a[0:0]
    for pos:=step_idx; pos<z; pos++ {
      if (ref_varid[pos]<0) { continue; }
      s,e := glfd.TileSequence(tilepath, tilelibver, anchor_tilestep + pos, ref_varid[pos])
      if e!=nil { return "", e }

      trail_ref_tag = s[len(s)-24:]

      ref_a = append(ref_a, s[:len(s)-24])
    }

    seq_a[0] = append(seq_a[0], trail_tag[0])
    seq_a[1] = append(seq_a[1], trail_tag[1])
    ref_a = append(ref_a, trail_ref_tag)

    seq0 := strings.Join(seq_a[0], "")
    seq1 := strings.Join(seq_a[1], "")
    refseq := strings.Join(ref_a, "")

    //fmt.Printf("seq0: %s\nseq1: %s\n ref: %s\n", seq0, seq1, refseq)

    first_step_in_path := false
    last_step_in_path := false

    chrom := glfd.TilepathToChrom[tilepath]
    ref_pos := 0
    if (anchor_tilestep + step_idx) > 0 {
      m := len(glfd.Assembly["hg19"][tilepath])
      ref_pos = glfd.Assembly["hg19"][tilepath][anchor_tilestep+step_idx-1]

      if m==(anchor_tilestep+step_idx) { last_step_in_path = true }
      first_step_in_path = false

    } else if tilepath>0 {
      m := len(glfd.Assembly["hg19"][tilepath-1])
      ref_pos = glfd.Assembly["hg19"][tilepath-1][m-1]

      if m==(anchor_tilestep+step_idx) { last_step_in_path = true }
      first_step_in_path = true
    }

    _ = last_step_in_path

    if first_step_in_path {
      EmitGVCF(outs, chrom, ref_pos, refseq, seq0, seq1)
    } else {
      EmitGVCF(outs, chrom, ref_pos, refseq[24:], seq0[24:], seq1[24:])
    }

    step_idx = z

  }

  return "", nil
}

// TESTING
//
func main() {
  local_debug := true


  //tglfd := GLFD{}
  //tglfd.StartSrv()
  //os.Exit(0)



  if local_debug {
    fmt.Printf(">>> loading...\n")
    t := time.Now()
    fmt.Print(t.Format(time.RFC3339))
    fmt.Printf("\n")
  }

  glfd,e := GLFDInit("/scratch/l7g/glf/glf.2bit.path",
    "/scratch/l7g/assembly/assembly.00.hg19.fw.gz",
    "/scratch/l7g/tagset.2bit/tagset.2bit",
    //"/scratch/l7g/glf/span/span.bin.gz")
    "/scratch/l7g/glf/span/span.gz")
  if e!=nil { panic(e) }

  if local_debug {
    fmt.Printf(">>> done\n")
    t := time.Now()
    fmt.Print(t.Format(time.RFC3339))
    fmt.Printf("\n")
  }

  if local_debug {
    fmt.Printf("starting web server...\n")
  }

  glfd.StartSrv()
  os.Exit(0)

  outs := bufio.NewWriter(os.Stdout)

  x,_ := glfd.TileSequence(3, 0, 13, 3)
  y,_ := glfd.TileSequence(3, 0, 13, 2)
  r,_ := glfd.TileSequence(3, 0, 13, 0)

  EmitGVCF(outs, "unk", 0, r, x, y)

  v,e := glfd.JSVMRun(` function f() { return 333; } ; var obj = { "x" : "y", "z" : 3 }; var abc = 2+2; console.log("...", abc); abc; f(); glfd_return(obj); `)
  if e!=nil { panic(e) }

  fmt.Printf("\n>>> %v\n", v)


  t := time.Now()
  fmt.Print(t.Format(time.RFC3339))
  fmt.Printf("\n")


  sampq,e := ioutil.ReadFile("js/sample-query.js")
  //sampq,e := ioutil.ReadFile("js/p2fb.js")
  //sampq,e := ioutil.ReadFile("js/p2fb_x.js")
  if e!=nil { panic(e) }



  //v,e = JSVM_run(string(sampq))
  v,e = glfd.JSVMRun(string(sampq))
  if e!=nil { panic(e); }
  fmt.Printf("sample-query.js: %v\n", v)

  t = time.Now()
  fmt.Print(t.Format(time.RFC3339))
  fmt.Printf("\n")

  glfd.StartSrv()

}
