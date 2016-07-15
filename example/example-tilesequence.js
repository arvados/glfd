var tilepath = 0x2fb;
var libver = 0;
var tilestep = 0x30;
var tilevarid = 0x1;
var span = glfd_tilespan(tilepath, libver, tilestep, tilevarid);
var seq = tilesequence(tilepath, libver, tilestep, tilevarid);
var tileid = [ hexstr(tilepath, 4), hexstr(libver, 2), hexstr(tilestep, 4), hexstr(tilevarid) ].join(".") + "+" + span.toString();
glfd_return({"tileid":tileid, "seq":seq}, "  ");
