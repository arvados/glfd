#include <Python.h>
#include "asm_ukk.h"

static int g_debug=1;

static char module_docstring[] =
  "Ukkonen's algorithm for calculating edit distance of strings.";
static char asmukk_docstring[] =
  "Ukkonen's algorithm for calculating edit distance of strings.";

//static PyObject *asmukk_pyasmukk(PyObject *self, PyObject *args);
static PyObject *asmukk_pyscore(PyObject *self, PyObject *args);
static PyObject *asmukk_pyscore2(PyObject *self, PyObject *args);
//static PyObject *asmukk_pyscore3(PyObject *self, PyObject *args);
static PyObject *asmukk_pyalign(PyObject *self, PyObject *args);
static PyObject *asmukk_pyalign2(PyObject *self, PyObject *args);
//static PyObject *asmukk_pyalign3(PyObject *self, PyObject *args);
static PyObject *asmukk_pydebug(PyObject *self, PyObject *args);


static PyMethodDef module_methods[] = {
  {"debug", asmukk_pydebug, METH_VARARGS, asmukk_docstring},
  {"score", asmukk_pyscore, METH_VARARGS, asmukk_docstring},
  {"score2", asmukk_pyscore2, METH_VARARGS, asmukk_docstring},
  {"score3", asmukk_pyscore3, METH_VARARGS, asmukk_docstring},
  {"align", asmukk_pyalign, METH_VARARGS, asmukk_docstring},
  {"align2", asmukk_pyalign2, METH_VARARGS, asmukk_docstring},
  {"align3", asmukk_pyalign3, METH_VARARGS, asmukk_docstring},
  {NULL, NULL, 0, NULL}
};



//PyMODINIT_FUNC init_asmukk(void)
PyMODINIT_FUNC initasmukk(void)
{
  PyObject *m = Py_InitModule3("asmukk", module_methods, module_docstring);
  if (m == NULL)
    return;
}

static PyObject *asmukk_pyscore(PyObject *self, PyObject *args) {
  char *x, *y;
  int sc;

  if (!PyArg_ParseTuple(args, "ss", &x, &y))
    return NULL;

  sc = asm_ukk_score(x,y);

  if (g_debug) {
    printf("... %i %s %s\n", sc, x, y);
  }

  return Py_BuildValue("i", sc);
}

static PyObject *asmukk_pyscore2(PyObject *self, PyObject *args) {
  char *x, *y;
  int mismatch, gap;
  int sc;

  if (!PyArg_ParseTuple(args, "ssii", &x, &y, &mismatch, &gap))
    return NULL;

  sc = asm_ukk_score2(x,y,mismatch,gap);

  if (g_debug) {
    printf("... %d (m%i, g%i) %s %s\n", sc, mismatch, gap, x, y);
  }

  return Py_BuildValue("i", sc);
}

/*
static PyObject *asmukk_pyscore3(PyObject *self, PyObject *args) {
  char *x, *y;
  PyObject *score_cb;

  if (!PyArg_ParseTuple(args, "ssO", &x, &y, &score_cb))
    return NULL;

  if (!PyCallable_Check(score_cb)) {
    PyErr_SetString(PyExc_TypeError, "score function must be callable");
    return NULL;
  }

  int mismatch, gap;
  int sc;
  printf("... %s %s\n", x, y);

  Py_RETURN_NONE;
}
*/

//WIP
static PyObject *asmukk_pyalign(PyObject *self, PyObject *args) {
  char *x, *y;
  char *X, *Y;

  if (!PyArg_ParseTuple(args, "ss", &x, &y))
    return NULL;

  sc = asm_ukk_align(&X, &Y, x,y);

  if ((X==NULL) || (Y==NULL)) {
    if (X) free(X);
    if (Y) free(Y);
  }

  if (g_debug) {
    printf("... %s %s (%s %s)\n", x, y, X, Y);
  }

  if (X) free(X);
  if (Y) free(Y);

  Py_RETURN_NONE;
}

static PyObject *asmukk_pyalign2(PyObject *self, PyObject *args) {
  const char *x, *y;

  if (!PyArg_ParseTuple(args, "ss", &x, &y))
    return NULL;

  printf("... %s %s\n", x, y);

  Py_RETURN_NONE;
}

/*
static PyObject *asmukk_pyalign3(PyObject *self, PyObject *args) {
  const char *x, *y;

  if (!PyArg_ParseTuple(args, "ss", &x, &y))
    return NULL;

  printf("... %s %s\n", x, y);

  Py_RETURN_NONE;
}
*/

static PyObject *asmukk_pydebug(PyObject *self, PyObject *args) {
  const char *x, *y;

  if (!PyArg_ParseTuple(args, "ss", &x, &y))
    return NULL;

  printf("... %s %s\n", x, y);

  Py_RETURN_NONE;
}
