var tilepath = 0x2fb;
var libver = 0;
var tilestep = 0x30;
var tilevarid = 0x1;

var span = glfd_tilespan(tilepath, libver, tilestep, tilevarid);

var tilepos_str = [ hexstr(tilepath, 4), hexstr(libver, 2), hexstr(tilestep, 4) ].join(".");

var end_tagseq = glfd_tagend_seq(tilepath, libver, tilestep);
var beg_tagseq = "";
if (tilestep>0) { beg_tagseq = glfd_tagend_seq(tilepath, libver, tilestep-1); }
var beg_tile=false, end_tile=false
if (tilestep==0) { beg_tile = true; } 
if ((tilestep+span)==(glf_info.StepPerPath[tilepath])) { end_tile = true; } 

var ret_obj = {"beg_tagseq":beg_tagseq, "end_tagseq":end_tagseq, "is-start-of-path":beg_tile, "is-end-of-path":end_tile, "tilepos":tilepos_str };
glfd_return(ret_obj, "  ");
