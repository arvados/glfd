var tilepath = 0x2fb;
var libver = 0;
var tilestep = 0x30;
var tilevarid = 0x1;

var span = glfd_tilespan(tilepath, libver, tilestep, tilevarid);

var tileid_str = [ hexstr(tilepath, 4), hexstr(libver, 2), hexstr(tilestep, 4), hexstr(tilevarid, 4) ].join(".") + "+" + span.toString();

var base_seq = tilesequence(tilepath, libver, tilestep, 0);
var alt0seq = tilesequence(tilepath, libver, tilestep, 1);
var alt1seq = tilesequence(tilepath, libver, tilestep, 2);

var align_arr0 = align(base_seq, alt0seq).split("\n");
var align_arr1 = align(base_seq, alt0seq).split("\n");
glfd_return({"tile-id":tileid_str,"align":[align_arr0, align_arr1], "base_seq":base_seq, "alt_seq":[alt0seq, alt1seq]}, '  ');
