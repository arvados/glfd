var tilepath = 0x2fb;
var libver = 0;
var tilestep = 0x30;
var tilevarid = 0x1;

var a = "hg19";
var apdh = "x";

var span = glfd_tilespan(tilepath, libver, tilestep, tilevarid);

var chrom = glfd_assembly_chrom(a, apdh, tilepath);
var alt_chrom = chrom;

var ref_start = 0;
if (tilestep>0) { ref_start = glfd_assembly_end_pos(a, apdh, tilepath, libver, tilestep-1); }
else if (tilepath>0) { 
    alt_chrom = glfd_assembly_chrom(a, apdh, tilepath-1);
      var end_step = glf_info.StepPerPath[tilepath-1];
        ref_start = glfd_assembly_end_pos(a, apdh, tilepath-1, libver, end_step-1);
}

var ref_end = glfd_assembly_end_pos(a, apdh, tilepath, libver, tilestep+span-1);
if (alt_chrom!=chrom) { ref_start = 0; }
var tilepos_str = [ hexstr(tilepath, 4), hexstr(libver, 2), hexstr(tilestep, 4) ].join(".");

var ret_obj = {"assembly-name":a, "assembly-pdh":apdh, "chromosome-name":chrom, "indexing":0, "start-position":ref_start, "end-position":ref_end };
glfd_return(ret_obj, "  ");
