print = ((typeof(print)==="undefined") ? console.log : print);

var glf_info = {};

function setup_glf_info() {
  //glf_info.cgf = [];
  //glf_info.cgf.push({ "file":"../data/hu826751-GS03052-DNA_B01.cgf", "name":"hu826751-GS03052-DNA_B01", "id":0 });
  //glf_info.cgf.push({ "file":"../data/hu0211D6-GS01175-DNA_E02.cgf", "name":"hu0211D6-GS01175-DNA_E02", "id":1 });

  //glf_info.id = {};
  //for (var idx=0; idx<glf_info.cgf.length; idx++) {
  //  glf_info.id[ glf_info.cgf[idx].name ] = idx;
  //}

  glf_info["TagSetVersion"] = "xx";
  glf_info["CGFVersion"] = "0.1.0";
  glf_info["CGFLibVersion"] = "0.1.0";
  glf_info["PathCount"] = 863;
  glf_info["StepPerPath"] = [
     5433, 11585, 7112, 7550, 13094, 10061, 15111, 13212, 14838, 7361, 8565, 8238, 21058, 15318, 9982, 14543,
    20484, 11704, 9056, 29572, 3032, 58941, 13626, 13753, 10082, 19756, 9669, 18011, 17221, 16418, 6572, 10450,
    653, 1, 1, 43, 4603, 4524, 17225, 5245, 9951, 5416, 18877, 6467, 14301, 7627, 11539, 16593,
    21475, 19845, 11886, 19126, 30932, 16774, 11607, 37511, 1368, 9016, 14132, 15803, 6847, 26570, 19594, 17082,
    10529, 20354, 17716, 9931, 19189, 14703, 8418, 8231, 17045, 7804, 12459, 23570, 20025, 8246, 24611, 10263,
    17693, 11001, 7904, 5629, 32719, 19083, 565, 3431, 20757, 13319, 5383, 9608, 10026, 16921, 14381, 29377,
    6845, 8754, 6367, 21554, 7707, 18707, 4227, 2345, 16932, 19091, 15332, 23909, 32173, 10128, 9612, 24819,
    9782, 21619, 22599, 5851, 16177, 24645, 24453, 14657, 3551, 19209, 17178, 6784, 22677, 10729, 4764, 18388,
    11981, 5804, 12040, 29022, 9918, 17574, 4842, 16740, 11327, 16335, 1542, 416, 23880, 6126, 8255, 16187,
    20267, 23705, 17658, 21050, 14728, 14705, 2708, 9599, 1327, 17097, 6536, 3446, 7194, 13517, 6740, 12960,
    8454, 15276, 6666, 10736, 7497, 7113, 13394, 16658, 7897, 10893, 15843, 24193, 12589, 10989, 7735, 7704,
    6591, 26835, 12945, 19129, 12707, 14282, 6739, 5660, 7363, 17599, 20166, 15899, 5832, 18674, 15349, 10225,
    13863, 25249, 32580, 20511, 13259, 14135, 3468, 71, 25343, 27513, 14097, 21456, 9860, 13680, 6387, 10838,
    4120, 21815, 5451, 14460, 8533, 24975, 24610, 25300, 11590, 19404, 8688, 32414, 7729, 19437, 6621, 10118,
    17649, 24182, 10736, 21411, 6710, 17505, 4790, 22874, 15243, 14561, 17381, 7292, 13961, 20750, 12771, 17639,
    5133, 16978, 18906, 16519, 15821, 13209, 882, 4225, 31741, 15233, 1182, 13597, 6528, 11710, 13632, 16991,
    5455, 37078, 22890, 16898, 6764, 20266, 7277, 6180, 8009, 24144, 22877, 12483, 21662, 12287, 19473, 20872,
    11085, 11566, 16415, 34070, 16922, 13794, 14120, 8663, 7451, 11295, 13748, 3815, 7213, 7030, 38651, 6143,
    12781, 5883, 5178, 11753, 15562, 22214, 22047, 4132, 16117, 3941, 144, 4865, 400, 25489, 22288, 30139,
    3706, 11083, 19909, 24752, 4171, 19061, 35002, 14079, 794, 29730, 3892, 12776, 3515, 15587, 14919, 14827,
    11010, 13427, 13368, 11662, 21111, 13834, 24662, 10333, 6684, 8376, 25611, 10830, 17440, 17699, 9856, 3300,
    23551, 7908, 24000, 7739, 13746, 5876, 13653, 11619, 105, 227, 12689, 19087, 11490, 35461, 6928, 11137,
    6317, 19717, 18677, 2636, 10982, 28108, 11243, 14787, 10618, 12904, 7678, 4053, 8783, 21899, 18003, 16798,
    16058, 8543, 15728, 7511, 16071, 18591, 25102, 17085, 16227, 5457, 29901, 6958, 5306, 12761, 2290, 4222,
    15593, 1523, 10990, 23625, 2365, 14954, 7597, 9733, 12983, 17099, 7155, 17446, 7771, 24670, 22012, 9790,
    17944, 16958, 6352, 22341, 6025, 12803, 18803, 16509, 19724, 13970, 23963, 7842, 9501, 16725, 20807, 9222,
    7462, 5182, 22155, 9365, 20144, 11012, 8142, 1490, 180, 546, 1, 1, 15, 550, 4865, 7015,
    20266, 7250, 11850, 10403, 13346, 5036, 7311, 10212, 9994, 12206, 21611, 12006, 13925, 10860, 19459, 12846,
    17584, 11203, 1904, 7356, 5714, 14022, 11522, 3238, 10867, 22206, 19356, 3286, 381, 14758, 7681, 18901,
    6319, 11569, 13319, 2602, 1, 12601, 5388, 8544, 32551, 13246, 23124, 16676, 10420, 16083, 23002, 4756,
    13393, 4473, 10500, 8904, 9750, 4253, 7078, 3459, 24069, 12012, 16737, 10252, 5577, 17329, 11901, 19092,
    9991, 28650, 8063, 13688, 21339, 17049, 4291, 15046, 21055, 27571, 19581, 5339, 1, 2796, 15653, 6733,
    5702, 9463, 8431, 7485, 17429, 7445, 33236, 10017, 15088, 16390, 18985, 3047, 29163, 8290, 8000, 26700,
    10459, 15540, 11802, 16858, 12184, 8407, 15777, 9945, 7774, 20407, 5030, 20355, 4994, 11256, 9088, 5210,
    703, 31263, 9981, 8655, 12869, 6059, 5323, 19308, 6962, 10252, 14659, 16466, 18159, 25083, 8822, 14458,
    13654, 20804, 8472, 20356, 9936, 2048, 7595, 10099, 4973, 9834, 18782, 13534, 16861, 1, 1, 1,
    1, 449, 13648, 8140, 8894, 4307, 12796, 7164, 5979, 18211, 19843, 2279, 5677, 13654, 16553, 17021,
    10676, 13343, 11629, 19081, 8331, 7079, 7216, 33870, 9290, 20014, 12554, 4179, 9303, 12659, 8980, 13317,
    17551, 1, 1, 1, 1, 91, 16293, 33478, 7694, 4755, 4736, 21768, 13932, 14148, 12245, 5458,
    10017, 15321, 10317, 11761, 9101, 13816, 21162, 17182, 5312, 19338, 8096, 10791, 6468, 20877, 6861, 3000,
    11596, 1, 1, 1, 1, 650, 9508, 9670, 6240, 919, 7453, 25276, 10122, 2914, 3833, 18009,
    12803, 23978, 730, 17531, 13228, 383, 818, 20600, 9375, 4772, 6376, 13251, 9675, 14940, 19964, 17117,
    15125, 29145, 9758, 7794, 8325, 3452, 13738, 9153, 15025, 13310, 2344, 1, 1932, 22098, 16158, 2693,
    36810, 14074, 7919, 4845, 19451, 10051, 10058, 11572, 5454, 5493, 11041, 11843, 15854, 19846, 17827, 125,
    1653, 21451, 21850, 1084, 9274, 12463, 8800, 10895, 28728, 2071, 9705, 5530, 5548, 10683, 15160, 14696,
    1860, 22145, 10747, 16523, 5517, 9195, 13344, 12, 1771, 23326, 30739, 18023, 25450, 18584, 21768, 9509,
    10948, 10287, 21091, 7440, 17747, 19563, 23601, 23077, 347, 7918, 12998, 13442, 559, 2780, 15135, 11458,
    9677, 1431, 16004, 5610, 9780, 11468, 7764, 8969, 10185, 19284, 16238, 11893, 23036, 13336, 3819, 12729,
    1268, 1, 8908, 8382, 11966, 16626, 1577, 16554, 12901, 20235, 6003, 7836, 17926, 1, 1, 1214,
    534, 1, 4755, 30050, 11265, 18230, 16716, 7896, 7554, 11599, 21326, 1, 1, 1, 1, 2745,
    11905, 4602, 8007, 14401, 9459, 20828, 12737, 11643, 16587, 4104, 6858, 5235, 6576, 13137, 29283, 8472,
    9959, 11291, 16995, 8588, 23499, 18569, 14851, 10837, 14462, 10224, 1492, 3714, 5149, 10944, 13980, 6118,
    6368, 29977, 5799, 12454, 4748, 18033, 14477, 3916, 18518, 28427, 15228, 29028, 6516, 11944, 15846, 8098,
    6040, 18525, 26363, 1, 1045, 12490, 1, 361, 4758, 14711, 4019, 5647, 721, 181, 35
  ];

  return glf_info;
}
setup_glf_info();

function tile_concordance_slice(set_a, set_b, lvl) {
}

function help() {
  print("muduk server");
}

function query(q) {
  var qobj = {};
  var robj = {};
  try {
    qobj = JSON.parse(q);
  } catch(err) {
    return err;
  }

  if ("request" in qobj) {
    print("request: " + String(qobj["request"]));
  }

  return JSON.stringify(robj);
}

// Take the input and Stringify is if it's a JSON object.
// If it's a string or number, return it.
// Otherwise return an empty string.
//
function glfd_return(q, indent) {
  indent = ((typeof(indent)==="undefined") ? '' : indent);
  if (typeof(q)==="undefined") { return ""; }
  if (typeof(q)==="object") {
    var s = "";
    try {
      s = JSON.stringify(q, null, indent);
    } catch(err) {
    }
    return s;
  }
  if (typeof(q)==="string") { return q; }
  if (typeof(q)==="number") { return q; }
  return "";
}

function hexstr(x, sz) {
  sz = ((typeof sz==="undefined")?0:sz);
  var t = x.toString(16);
  if (t.length < sz) {
    t = Array(sz - t.length + 1).join("0") + t;
  }
  return t;
}

function info() {
  return glfd_return(glf_info);
}


function api_tile_variants(tilepos_str) {
  var tilepos_parts = tilepos_str.split(".");
  var tilever = parseInt(tilepos_parts[0], 16);
  var tilepath = parseInt(tilepos_parts[1], 16);
  var tilestep = parseInt(tilepos_parts[2], 16);

  var ret_a = [];

  var r = glfd_tilepos_info(tilepath, tilever, tilestep);
  var jj = JSON.parse(r);
  for (var idx=0; idx<jj.length; idx++) {
    var tileid_parts = jj[idx]["tile-id"].split(".");
    var m5 = jj[idx]["md5sum"];

    var varid = parseInt(tileid_parts[3], 16);

    ret_a.push({
      "tile-id":tileid_parts[1] + "." + tileid_parts[0] + "." + tileid_parts[2] + "." + m5,
      "md5sum":m5,
      "variant-id":varid
    });
  }

  return glfd_return(ret_a);
}

function api_tagseq_end(tilepos_str) {
  var tilepos_parts = tilepos_str.split(".");
  var tilever = parseInt(tilepos_parts[0], 16);
  var tilepath = parseInt(tilepos_parts[1], 16);
  var tilestep = parseInt(tilepos_parts[2], 16);

  var seq = glfd_tagseq_end(tilepath, tilever, tilestep);

  return glfd_return(seq);
}

function api_tile_variant_info(tile_md5_var_str) {
  var tilepos_parts = tilepos_str.split(".");
  var tilever = parseInt(tilepos_parts[0], 16);
  var tilepath = parseInt(tilepos_parts[1], 16);
  var tilestep = parseInt(tilepos_parts[2], 16);
  var md5var_str = tilepos_parts[3];


  var x = api_locus_req("hg19", "xx", hexstr(tilever, 2) + "." + hexstr(tilepath, 4) + "." + hexstr(tilestep, 4));
  var jj = JSON.parse(x);

  var is_start_tile = false;
  var is_end_tile = false;
  if (tilestep == 0) { is_start_tile = true; }
  if (tilestep == (glf_info.StepPerPath[tilepath]-1)) { is_end_tile = true; }

  var y = api_tile_variants(hexstr(tilever, 2) + "." + hexstr(tilepath, 4) + "." + hexstr(tilestep, 4));
  var yy = JSON.parse(y);

  var ret_obj = {};
  ret_obj["tile-variant"] = tile_md5_var_str;
  ret_obj["tag-length"] = 24;
  ret_obj["start-tag"] = "";
  ret_obj["end-tag"] = "";
  ret_obj["is-start-of-path"] = is_start_tile;
  ret_obj["is-end-of-path"] = is_end_tile;
  ret_obj["sequence"] = "";
  ret_obj["md5sum"] = md5var_str;
  ret_obj["length"] = -1;
  ret_obj["number-of-positions-spanned"] = -1;

  for (var idx=0; idx<yy.length; idx++) {
    if (yy[idx].md5sum == md5var_str) {

      var tileid = yy[idx]["tile-id"].split(".");
      var varid = yy[idx]["variant-id"];

      var seq = yy[idx].seq;
      if (!is_start_tile) {
        ret_obj["start-tag"] = seq.slice(0,24);
      }
      if (!is_end_tile) {
        ret_obj["end-tag"] = seq.slice(seq.length-24, seq.length);
      }
      ret_obj["sequence"] = seq;
      ret_obj["length"] = seq.length;

      var span = glfd_tilespan(tilepath, tilever, tilestep, varid);
      ret_obj["number-of-positions-spanned"] = span;
    }
  }

  return glfd_return(ret_obj);
}

function api_locus_req(assembly_name, assembly_pdh, tilepos_str) {
  var r = [];

  var tilepos_parts = tilepos_str.split(".");
  var tilever = parseInt(tilepos_parts[0], 16);
  var tilepath = parseInt(tilepos_parts[1], 16);
  var tilestep = parseInt(tilepos_parts[2], 16);

  var chrom = glfd_assembly_chrom(assembly_name, assembly_pdh, tilepath);

  var end_pos = glfd_assembly_end_pos(assembly_name, assembly_pdh, tilepath, tilever, tilestep);
  var beg_pos = 0;
  if (tilestep>0) {
    beg_pos = glfd_assembly_end_pos(assembly_name, assembly_pdh, tilepath, tilever, tilestep-1);
  }
  else if (tilepath>0) {
    var end_step = glf_info.StepPerPath[tilepath-1]-1;

    var prev_chrom = glfd_assembly_chrom(assembly_name, assembly_pdh, tilepath-1);
    if (prev_chrom.toString() == chrom.toString()) {
      beg_pos = glfd_assembly_end_pos(assembly_name, assembly_pdh, tilepath-1, tilever, end_step);
    } else {
      beg_pos = 0;
    }

  }

  r.push({
    "assembly-name":assembly_name,
    "assembly-pdh":assembly_pdh,
    "chromosome-name":chrom.toString(),
    "start-position":beg_pos,
    "end-position":end_pos,
    "indexing":0
  });

  return glfd_return(r);
}

// This all needs to be fixed to take in whatever
// the `rest_req` variable has stored.
//
function api_query(rest_req) {
  var parts = rest_req.split("/");
  var n = parts.length;
  var f = parts[n-1];

  if (f=="status") {
    return glfd_return({"api-version":glf_info.GGFLibVersion});
  }

  else if (f=="tag-sets") {
    return glfd_return([ glf_info.TagSetVersion ]);
  }

  else if (f=="tag-set-identifier") {
    return glfd_return({"tag-set-identifier": glf_info.TagSetVersion, "tag-set-integer":0 });
  }

  else if (f=="paths") {
    var n_path = glf_info.StepPerPath.length;
    var path_a = [];
    for (var i=0; i<n_path; i++) {
      path_a.append(i);
    }
    return glfd_return(path_a);
  }

  else if (f=="path-int") {
    var tilepath = 0x247;
    return glfd_return({"path":tilepath, "num-positions":glf_info.StepPerPath[tilepath]});
  }

  else if (f=="tile-positions") {
    var tilepath = 0x247;
    var tilepos_a = [];
    var libver = 0;
    var m = glf_info.StepPerPath[tilepath];
    for (var i=0; i<m; i++) {
      tilepos_a.push( hexstr(libver,2) + "." + hexstr(tilepath, 3) + "." + hexstr(i, 4) );
    }
    return glfd_return(tilepos_a);
  }

  // outside of glfd scope
  //
  else if (f=="tile-position-id") {
    return "";
  }

  else if (f=="locus") {
    var assembly_name = "hg19";
    var assembly_pdh = "xxx";
    var tilepos_str = "00.247.0000";

    return api_locus_req(assembly_name, assembly_pdh, tilepos_str);
  }

  // unreasonable, not implementing as specified.
  // only consider tile positions
  //
  else if (f=="tile-variants") {
    var tilepos_str = "00.247.0000";
    return api_tile_variants(tilepos_str);
  }

  else if (f=="tile-variant-id") {
    return api_tile_variant_info(tile_md5_var_str);
  }

}
