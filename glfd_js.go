package main

//import "os"
import "fmt"
import "bytes"
import "bufio"
import "strings"
//import "strconv"
import "io/ioutil"
import "github.com/robertkrimen/otto"

//import "github.com/aebruno/twobit"
//import "github.com/abeconnelly/pasta"
//import "github.com/abeconnelly/pasta/gvcf"

import "github.com/abeconnelly/sloppyjson"

//import "reflect"

func status_otto(call otto.FunctionCall) otto.Value {
  v,e := otto.ToValue("ok status")
  if e!=nil { return otto.Value{} }
  return v
}

/*
func (glfd *GLFD) refinfo_otto(call otto.FunctionCall) otto.Value {
  otto_err,err := otto.ToValue("error")
  if err!=nil { return otto.Value{} }

  assembly_name,e := call.Argument(0).String()
  if e!=nil { return otto_err }
  assembly_pdh,e := call.Argument(1).String() ; _ = assembly_pdh
  if e!=nil { return otto_err }
  chrom,e := call.Argument(2).String()
  if e!=nil { return otto_err }
  beg_pos,e := call.Argument(3).ToInteger()
  if e!=nil { return otto_err }
  end_pos,e := call.Argument(4).ToInteger()
  if e!=nil { return otto_err }

  tilepath,tilestep,e := glfd.RefPosLookupTile(assembly_name, chrom, int(beg_pos), int(end_pos))
  if e!=nil { return otto_err }

  ref_varid := glfd.RefV[assembly_name][tilepath][tilestep]

  seq,e := glfd.TileSequence(tilepath, 0, tilestep, ref_varid)
  if e!=nil { return otto_err }

  //ret_s := fmt.Sprintf(`{"tile-id":"%04x.%02.%04x.%03x","seq":"%s","subseq":"%s"
}
*/

func (glfd *GLFD) tilespan_otto(call otto.FunctionCall) otto.Value {
  otto_err,err := otto.ToValue("error")
  if err!=nil { return otto.Value{} }

  tilepath,e := call.Argument(0).ToInteger()
  if e!=nil { return otto_err }
  libver,e := call.Argument(1).ToInteger()
  if e!=nil { return otto_err }
  tilestep,e := call.Argument(2).ToInteger()
  if e!=nil { return otto_err }
  tilevar,e := call.Argument(3).ToInteger()
  if e!=nil { return otto_err }

  span,e := glfd.TileSpan(int(tilepath), int(libver), int(tilestep), int(tilevar))
  if e!=nil { return otto_err }

  v,e := otto.ToValue(span)
  if e!=nil { return otto_err }

  return v
}

func (glfd *GLFD) tagend_seq_otto(call otto.FunctionCall) otto.Value {

  otto_err,err := otto.ToValue("error")
  if err!=nil { return otto.Value{} }

  tilepath,e := call.Argument(0).ToInteger()
  if e!=nil { return otto_err }
  libver,e := call.Argument(1).ToInteger()
  if e!=nil { return otto_err }
  tilestep,e := call.Argument(2).ToInteger()
  if e!=nil { return otto_err }


  tagseq,e := glfd.TagEnd(int(tilepath), int(libver), int(tilestep))
  if e!=nil { return otto_err }

  v,e := otto.ToValue(tagseq)
  if e!=nil { return otto_err }

  return v
}

func (glfd *GLFD) tilepos_info_otto(call otto.FunctionCall) otto.Value {
  tilepath,e := call.Argument(0).ToInteger()
  if e!=nil { return otto.Value{} }
  libver,e := call.Argument(1).ToInteger()
  if e!=nil { return otto.Value{} }
  tilestep,e := call.Argument(2).ToInteger()
  if e!=nil { return otto.Value{} }

  s,e := glfd.TileLibSequences(int(tilepath), int(libver), int(tilestep))
  if e!=nil { return otto.Value{} }

  v,e := otto.ToValue(s)
  return v
}

func (glfd *GLFD) tilesequence_otto(call otto.FunctionCall) otto.Value {
  tilepath,e := call.Argument(0).ToInteger()
  if e!=nil { return otto.Value{} }
  libver,e := call.Argument(1).ToInteger()
  if e!=nil { return otto.Value{} }
  tilestep,e := call.Argument(2).ToInteger()
  if e!=nil { return otto.Value{} }
  tilevarid,e := call.Argument(3).ToInteger()
  if e!=nil { return otto.Value{} }

  s,e := glfd.TileSequence(int(tilepath), int(libver), int(tilestep), int(tilevarid))
  if e!=nil { return otto.Value{} }

  v,e := otto.ToValue(s)
  return v
}

func align2pasta_otto(call otto.FunctionCall) otto.Value {
  refseq := call.Argument(0).String()
  altseq := call.Argument(1).String()

  s,e := AlignToPasta(refseq, altseq)
  if e!=nil { return otto.Value{} }

  v,e := otto.ToValue(s)
  if e!=nil { return otto.Value{} }
  return v
}

func align_otto(call otto.FunctionCall) otto.Value {
  refseq := call.Argument(0).String()
  altseq := call.Argument(1).String()

  ref_align, alt_align, score := align(refseq, altseq) ; _ = score

  //sa := []string{ref_align, alt_align}
  //v,e := otto.ToValue(sa)

  v,e := otto.ToValue(ref_align + "\n" + alt_align)
  if e!=nil { return otto.Value{} }
  return v
}


/*
func emitgvcf_otto(call otto.FunctionCall) otto.Value {
  refseq := call.Argument(0).String()
  alt0seq := call.Argument(1).String()
  alt1seq := call.Argument(2).String()

  //DEBUG
  outs := bufio.NewWriter(os.Stdout)

  //EmitGVCF(refseq, alt0seq, alt1seq)
  EmitGVCF(outs, "unk", 0, refseq, alt0seq, alt1seq)

  //v,e := otto.ToValue("ok")
  v,e := otto.ToValue(outs.Bytes)
  if e!=nil { return otto.Value{} }
  return v
}
*/

//func (glfd *GLFD) tiletogvcf_x_otto(call otto.FunctionCall) otto.Value { }

func (glfd *GLFD) assembly_end_pos_otto(call otto.FunctionCall) otto.Value {
  assembly_name := call.Argument(0).String()
  assembly_pdh := call.Argument(1).String() ; _= assembly_pdh

  otto_err,e := otto.ToValue("error")
  if e!=nil { return otto.Value{} }


  tilepath,e := call.Argument(2).ToInteger()
  if e!=nil { return otto_err }

  tilever,e := call.Argument(3).ToInteger() ; _ = tilever
  if e!=nil { return otto_err }

  tilestep,e := call.Argument(4).ToInteger()
  if e!=nil { return otto_err }

  if _,ok := glfd.Assembly[(assembly_name)] ; ok {
    if _,okp := glfd.Assembly[(assembly_name)][int(tilepath)] ; okp {
      if end_ref,oks := glfd.Assembly[(assembly_name)][int(tilepath)][int(tilestep)] ; oks {
        //end_ref := glfd.Assembly[(assembly_name)][int(tilepath)][int(tilestep)]
        v,e := otto.ToValue(end_ref)
        if e!=nil { return otto_err }
        return v
      }
    }
  }

  return otto_err
}

func (glfd *GLFD) assembly_chrom_otto(call otto.FunctionCall) otto.Value {
  otto_err,e := otto.ToValue("error")
  if e!=nil { return otto.Value{} }

  assembly_name := call.Argument(0).String() ; _ = assembly_name
  assembly_pdh := call.Argument(1).String() ; _ = assembly_pdh

  tilepath,e := call.Argument(2).ToInteger()
  if e!=nil { return otto_err }

  if chromstr,ok := glfd.TilepathToChrom[int(tilepath)] ; ok {
    v,e := otto.ToValue(chromstr)
    if e!=nil { return otto_err }
    return v
  }

  return otto_err
}


func (glfd *GLFD) tiletogvcf_x_otto(call otto.FunctionCall) otto.Value {
  otto_err,e := otto.ToValue("error")
  if e!=nil { return otto.Value{} }

  str := call.Argument(0).String()
  sec_arg := call.Argument(1)

  json_out := true
  if (sec_arg.IsDefined()) {
    var e error
    json_out,e = sec_arg.ToBoolean()
    if e!=nil { return otto_err }
  }



  jso,e := sloppyjson.Loads(str)
  //if e!=nil {  panic(e) }
  if e!=nil {
    v,e := otto.ToValue("input parse error")
    if e!=nil { return otto.Value{} }
    return v
  }

  tilepath := int(jso.O["tilepath"].P)
  start_tilestep := int(jso.O["start_tilestep"].P)

  allele := [][]int{}
  allele = append(allele, []int{})
  allele = append(allele, []int{})

  n:=len(jso.O["allele"].L[0].L)
  for i:=0; i<n; i++ {
    allele[0] = append(allele[0], int(jso.O["allele"].L[0].L[i].P))
    allele[1] = append(allele[1], int(jso.O["allele"].L[1].L[i].P))
  }

  nocall := [][][]int{}
  nocall = append(nocall, [][]int{})
  nocall = append(nocall, [][]int{})
  for i:=0; i<n; i++ {
    nocall[0] = append(nocall[0], []int{})
    m:=len(jso.O["loq_info"].L[0].L[i].L)
    for j:=0; j<m; j++ {
      nocall[0][i] = append(nocall[0][i], int(jso.O["loq_info"].L[0].L[i].L[j].P))
    }

    nocall[1] = append(nocall[1], []int{})
    m=len(jso.O["loq_info"].L[1].L[i].L)
    for j:=0; j<m; j++ {
      nocall[1][i] = append(nocall[1][i], int(jso.O["loq_info"].L[1].L[i].L[j].P))
    }

  }

  ref_varid := []int{}
  for i:=0; i<len(allele[0]); i++ {
    //ref_varid = append(ref_varid, 0)
    ref_varid = append(ref_varid, glfd.RefV["hg19"][tilepath][start_tilestep+i])
  }

  _ = start_tilestep
  _ = allele
  _ = nocall

  bb := new(bytes.Buffer)
  outs := bufio.NewWriter(bb)

  s,e := glfd.TileToGVCF(outs, tilepath, 0, start_tilestep, allele, nocall, ref_varid) ; _ = s
  if e!=nil {
    //panic(e)
    v,err := otto.ToValue( fmt.Sprintf("%v", e) )
    if err!=nil { return otto.Value{} }
    return v
  }

  out_str := string(bb.Bytes())
  if json_out { out_str = _to_json_gvcf(string(bb.Bytes()), "unk") }
  v,e := otto.ToValue(out_str)

  //json_gvcf_str := _to_json_gvcf(string(bb.Bytes()), "unk")
  //v,e := otto.ToValue(json_gvcf_str)
  if e!=nil { return otto.Value{} }
  return v
}

func _to_json_gvcf(s, samp_name string) string {
  tot_res_a := []string{}
  lines := strings.Split(s,"\n")
  for i:=0; i<len(lines); i++ {
    parts := strings.Split(lines[i], "\t")

    res_a := []string{}

    if len(parts)<6 { continue; }

    t := fmt.Sprintf(`{ "chrom":"%s","pos":%s,"ref":"%s","alt":[`, parts[0], parts[1], parts[3])
    res_a = append(res_a, t)

    alt_parts := strings.Split(parts[4], ",")
    for j:=0; j<len(alt_parts); j++ {
      if j>0 {
        t = fmt.Sprintf(`,"%s"`, alt_parts[j])
      } else {
        t = fmt.Sprintf(`"%s"`, alt_parts[j])
      }
      res_a = append(res_a, t)
    }
    res_a = append(res_a, "],")


    res_a = append(res_a, `"format":[{`)
    res_a = append(res_a, fmt.Sprintf(`"sample-name":"%s","GT":"%s"`, samp_name, parts[9]))
    res_a = append(res_a, `}],`)

    res_a = append(res_a, `"info":{`)
    x_parts := strings.Split(parts[7], ";")
    for j:=0; j<len(x_parts); j++ {
      kv := strings.Split(x_parts[j], "=")
      if len(kv)==2 {
        if j>0 { res_a = append(res_a, ",") }

        if kv[0] == "END" {
          res_a = append(res_a, `"` + kv[0] + `":[` + kv[1] + `]`)
        } else {
          res_a = append(res_a, `"` + kv[0] + `":"` + kv[1] + `"`)
        }
      }
    }
    res_a = append(res_a, `}`)
    res_a = append(res_a, `}`)

    tot_res_a = append(tot_res_a, strings.Join(res_a, ""))
  }

  return `[` + strings.Join(tot_res_a, ",") + `]`
}

func (glfd *GLFD) JSVMRun(src string) (rstr string, e error) {
  js_vm := otto.New()

  fmt.Printf("JSVM_run:\n\n")

  init_js,err := ioutil.ReadFile("js/init.js")
  if err!=nil { e = err; return }
  js_vm.Run(init_js)

  js_vm.Set("status", status_otto)
  js_vm.Set("tilesequence", glfd.tilesequence_otto)
  js_vm.Set("aligntopasta", align2pasta_otto)
  js_vm.Set("align", align_otto)
  //js_vm.Set("emitgvcf", emitgvcf_otto)
  //js_vm.Set("tiletogvcf", glfd.tiletogvcf_otto)

  js_vm.Set("tiletogvcf", glfd.tiletogvcf_x_otto)
  js_vm.Set("tiletogvcf_x", glfd.tiletogvcf_x_otto)

  js_vm.Set("glfd_assembly_end_pos", glfd.assembly_end_pos_otto)
  js_vm.Set("glfd_assembly_chrom", glfd.assembly_chrom_otto)

  js_vm.Set("glfd_tilepos_info", glfd.tilepos_info_otto)

  js_vm.Set("glfd_tagend_seq", glfd.tagend_seq_otto)

  js_vm.Set("glfd_tilespan", glfd.tilespan_otto)

  v,err := js_vm.Run(src)
  if err!=nil {
    e = err
    return
  }

  rstr,e = v.ToString()
  return
}
