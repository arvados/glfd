#include <stdio.h>
#include <stdlib.h>

#include <vector>
#include <string>
#include <map>

int main(int argc, char **argv) {
  int i, j, k;
  size_t n;
  char buf[1024];
  std::string s;
  std::string assembly_name, chrom, tilepath_str;
  std::string tilestep_str, end_pos_str;
  std::map<std::string, std::map<int, std::map<int, int> > > assembly;

  int tilepath, tilestep, end_pos;


  while (!feof(stdin)) {
    if (fgets(buf, 1024, stdin)==NULL) { continue; }
    s = buf;
    if (s.size()==0) { continue; }

    if (s[0]=='>') {

      assembly_name.clear();
      chrom.clear();
      tilepath_str.clear();

      n = s.size();
      for (k=0, i=1; i<n; i++) {
        if (s[i]==':') { k++; continue; }
        if (k==0) { assembly_name += s[i]; }
        else if (k==1) { chrom += s[i]; }
        else { tilepath_str += s[i]; }
      }

      tilepath = (int)strtol(tilepath_str.c_str(), NULL, 16);

      printf(">>> %s %s %i\n", assembly_name.c_str(), chrom.c_str(), tilepath);
      continue;
    }

    tilestep_str.clear();
    end_pos_str.clear();

    n = s.size();
    for (i=0; i<n; i++) {
      if ((s[i]==' ') || (s[i]=='\t')) { k++; continue; }
      if (k==0) { tilestep_str += s[i]; }
      else { end_pos_str += s[i]; }
    }

    tilestep = (int)strtol(tilestep_str.c_str(), NULL, 16);
    end_pos = atoi(end_pos_str.c_str());

    assembly[tilepath_str][tilepath][tilestep] = end_pos;

  }


  printf("ok\n");

}
