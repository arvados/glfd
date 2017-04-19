var assembly_name = "hg19";
var assembly_pdh = "";

var r_obj = {};
for (var tilestep = 0; tilestep<5; tilestep++) {
  var tileid = "00." + hexstr(0x35e) + "." + hexstr(tilestep, 4);
  var r = JSON.parse(api_locus_req(assembly_name, assembly_pdh, tileid));
  r_obj[tileid] = r;
}

glfd_return(r_obj, "  ");
