from distutils.core import setup, Extension
import numpy.distutils.misc_util

setup(
    ext_modules=[Extension("asmukk", ["pyasmukk.c", "asm_ukk.c", "asm_ukk3.c"])],
    include_dirs=numpy.distutils.misc_util.get_numpy_include_dirs(),
)


