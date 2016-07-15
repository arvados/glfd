var tilepath = 0x2fb;
var libver = 0;
var tilestep = 0x30;

var tilepos_str = [ hexstr(tilepath, 4), hexstr(libver, 2), hexstr(tilestep, 4) ].join(".");

var r = JSON.parse(glfd_tilepos_info(tilepath, libver, tilestep));
glfd_return(r, "  ");
