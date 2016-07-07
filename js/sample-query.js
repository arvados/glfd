print("info: " + info());

var tilepath = 0x2fb;
var libver = 0;
var tilestep = 0x30;
var tilevarid = 0x1;
var seq = tilesequence(tilepath, libver, tilestep, tilevarid);

print("tilesequence (" +  String(tilepath) + ", " + String(libver) + ", " + String(tilestep) + ", " + String(tilevarid) + "):\n");
print(seq);

var canon_seq = tilesequence(tilepath, libver, tilestep, 0);
var alt0seq = tilesequence(tilepath, libver, tilestep, 1);
var alt1seq = tilesequence(tilepath, libver, tilestep, 2);

var al = align(canon_seq, alt0seq);
print("align:\n" + al + "\n");

//print("aligntopasta(0,1,2):\n" + aligntopasta(canon_seq, alt0seq, alt1seq) + "\n");

print("emitgvcf..\n");
emitgvcf(canon_seq, alt0seq, alt1seq);

print("\n\n---\n\n");

var start_step = 3;
var end_step = 6;
for (var tilestep=start_step; tilestep<=end_step; tilestep++) {

  //print(">>>>>" + String(tilestep))

  var c = tilesequence(tilepath, libver, tilestep, 0);
  var a0 = tilesequence(tilepath, libver, tilestep, 1);

  var a1 = "";
  if (tilestep==9) {
    a1 = tilesequence(tilepath, libver, tilestep, 1);
  } else if (tilestep==6) {
    a0 = tilesequence(tilepath, libver, tilestep, 2);
    a1 = tilesequence(tilepath, libver, tilestep, 3);
  } else {
    a1 = tilesequence(tilepath, libver, tilestep, 2);
  }

  //print(" c: " + c + "\na0: " + a0 + "\na1: " + a1 + "\n");
  //print(" c: " + c.slice(0,-24) + "\na0: " + a0.slice(0,-24) + "\na1: " +  a1.slice(0,-24) + "\n");

  emitgvcf(c.slice(0,-24), a0.slice(0,-24), a1.slice(0,-24));

}

print("\n\ntiletogvcf:\n");
var q = {"tilepath":763, "allele": [[0,1,0],[0,0,0]], "loq_info":[[[],[],[]],[[],[],[]]], "start_tilestep":0};
tiletogvcf(JSON.stringify(q));


print("\ntesting span:\n");
var s0 = glfd_tilespan(0x2fb, 0, 0x84, 0);
var s1 = glfd_tilespan(0x2fb, 0, 0x84, 3);

print("  got: " + s0.toString() + " " + s1.toString() + "\n");
