/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: FCLWrapper.i

#define SWIGMODULE FCLlib

#ifdef __cplusplus
/* SwigValueWrapper is described in swig.swg */
template<typename T> class SwigValueWrapper {
  struct SwigMovePointer {
    T *ptr;
    SwigMovePointer(T *p) : ptr(p) { }
    ~SwigMovePointer() { delete ptr; }
    SwigMovePointer& operator=(SwigMovePointer& rhs) { T* oldptr = ptr; ptr = 0; delete oldptr; ptr = rhs.ptr; rhs.ptr = 0; return *this; }
  } pointer;
  SwigValueWrapper& operator=(const SwigValueWrapper<T>& rhs);
  SwigValueWrapper(const SwigValueWrapper<T>& rhs);
public:
  SwigValueWrapper() : pointer(0) { }
  SwigValueWrapper& operator=(const T& t) { SwigMovePointer tmp(new T(t)); pointer = tmp; return *this; }
  operator T&() const { return *pointer.ptr; }
  T *operator&() { return pointer.ptr; }
};

template <typename T> T SwigValueInit() {
  return T();
}
#endif

/* -----------------------------------------------------------------------------
 *  This section contains generic SWIG labels for method/variable
 *  declarations/attributes, and other compiler dependent labels.
 * ----------------------------------------------------------------------------- */

/* template workaround for compilers that cannot correctly implement the C++ standard */
#ifndef SWIGTEMPLATEDISAMBIGUATOR
# if defined(__SUNPRO_CC) && (__SUNPRO_CC <= 0x560)
#  define SWIGTEMPLATEDISAMBIGUATOR template
# elif defined(__HP_aCC)
/* Needed even with `aCC -AA' when `aCC -V' reports HP ANSI C++ B3910B A.03.55 */
/* If we find a maximum version that requires this, the test would be __HP_aCC <= 35500 for A.03.55 */
#  define SWIGTEMPLATEDISAMBIGUATOR template
# else
#  define SWIGTEMPLATEDISAMBIGUATOR
# endif
#endif

/* inline attribute */
#ifndef SWIGINLINE
# if defined(__cplusplus) || (defined(__GNUC__) && !defined(__STRICT_ANSI__))
#   define SWIGINLINE inline
# else
#   define SWIGINLINE
# endif
#endif

/* attribute recognised by some compilers to avoid 'unused' warnings */
#ifndef SWIGUNUSED
# if defined(__GNUC__)
#   if !(defined(__cplusplus)) || (__GNUC__ > 3 || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4))
#     define SWIGUNUSED __attribute__ ((__unused__))
#   else
#     define SWIGUNUSED
#   endif
# elif defined(__ICC)
#   define SWIGUNUSED __attribute__ ((__unused__))
# else
#   define SWIGUNUSED
# endif
#endif

#ifndef SWIG_MSC_UNSUPPRESS_4505
# if defined(_MSC_VER)
#   pragma warning(disable : 4505) /* unreferenced local function has been removed */
# endif
#endif

#ifndef SWIGUNUSEDPARM
# ifdef __cplusplus
#   define SWIGUNUSEDPARM(p)
# else
#   define SWIGUNUSEDPARM(p) p SWIGUNUSED
# endif
#endif

/* internal SWIG method */
#ifndef SWIGINTERN
# define SWIGINTERN static SWIGUNUSED
#endif

/* internal inline SWIG method */
#ifndef SWIGINTERNINLINE
# define SWIGINTERNINLINE SWIGINTERN SWIGINLINE
#endif

/* exporting methods */
#if defined(__GNUC__)
#  if (__GNUC__ >= 4) || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4)
#    ifndef GCC_HASCLASSVISIBILITY
#      define GCC_HASCLASSVISIBILITY
#    endif
#  endif
#endif

#ifndef SWIGEXPORT
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   if defined(STATIC_LINKED)
#     define SWIGEXPORT
#   else
#     define SWIGEXPORT __declspec(dllexport)
#   endif
# else
#   if defined(__GNUC__) && defined(GCC_HASCLASSVISIBILITY)
#     define SWIGEXPORT __attribute__ ((visibility("default")))
#   else
#     define SWIGEXPORT
#   endif
# endif
#endif

/* calling conventions for Windows */
#ifndef SWIGSTDCALL
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   define SWIGSTDCALL __stdcall
# else
#   define SWIGSTDCALL
# endif
#endif

/* Deal with Microsoft's attempt at deprecating C standard runtime functions */
#if !defined(SWIG_NO_CRT_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_CRT_SECURE_NO_DEPRECATE)
# define _CRT_SECURE_NO_DEPRECATE
#endif

/* Deal with Microsoft's attempt at deprecating methods in the standard C++ library */
#if !defined(SWIG_NO_SCL_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_SCL_SECURE_NO_DEPRECATE)
# define _SCL_SECURE_NO_DEPRECATE
#endif

/* Deal with Apple's deprecated 'AssertMacros.h' from Carbon-framework */
#if defined(__APPLE__) && !defined(__ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES)
# define __ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES 0
#endif

/* Intel's compiler complains if a variable which was never initialised is
 * cast to void, which is a common idiom which we use to indicate that we
 * are aware a variable isn't used.  So we just silence that warning.
 * See: https://github.com/swig/swig/issues/192 for more discussion.
 */
#ifdef __INTEL_COMPILER
# pragma warning disable 592
#endif


#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>



typedef long long intgo;
typedef unsigned long long uintgo;


# if !defined(__clang__) && (defined(__i386__) || defined(__x86_64__))
#   define SWIGSTRUCTPACKED __attribute__((__packed__, __gcc_struct__))
# else
#   define SWIGSTRUCTPACKED __attribute__((__packed__))
# endif



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;




#define swiggo_size_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];
#define swiggo_size_assert(t, n) swiggo_size_assert_eq(sizeof(t), n, swiggo_sizeof_##t##_is_not_##n)

swiggo_size_assert(char, 1)
swiggo_size_assert(short, 2)
swiggo_size_assert(int, 4)
typedef long long swiggo_long_long;
swiggo_size_assert(swiggo_long_long, 8)
swiggo_size_assert(float, 4)
swiggo_size_assert(double, 8)

#ifdef __cplusplus
extern "C" {
#endif
extern void crosscall2(void (*fn)(void *, int), void *, int);
extern char* _cgo_topofstack(void) __attribute__ ((weak));
extern void _cgo_allocate(void *, int);
extern void _cgo_panic(void *, int);
#ifdef __cplusplus
}
#endif

static char *_swig_topofstack() {
  if (_cgo_topofstack) {
    return _cgo_topofstack();
  } else {
    return 0;
  }
}

static void _swig_gopanic(const char *p) {
  struct {
    const char *p;
  } SWIGSTRUCTPACKED a;
  a.p = p;
  crosscall2(_cgo_panic, &a, (int) sizeof a);
}




#define SWIG_contract_assert(expr, msg) \
  if (!(expr)) { _swig_gopanic(msg); } else


static _gostring_ Swig_AllocateString(const char *p, size_t l) {
  _gostring_ ret;
  ret.p = (char*)malloc(l);
  memcpy(ret.p, p, l);
  ret.n = l;
  return ret;
}


static void Swig_free(void* p) {
  free(p);
}

static void* Swig_malloc(int c) {
  return malloc(c);
}



#include "FCLWrapper.h"



#include <string>


#include <vector>
#include <stdexcept>

SWIGINTERN std::vector< std::string >::const_reference std_vector_Sl_std_string_Sg__get(std::vector< std::string > *self,int i){
                int size = int(self->size());
                if (i>=0 && i<size)
                    return (*self)[i];
                else
                    throw std::out_of_range("vector index out of range");
            }
SWIGINTERN void std_vector_Sl_std_string_Sg__set(std::vector< std::string > *self,int i,std::vector< std::string >::value_type const &val){
                int size = int(self->size());
                if (i>=0 && i<size)
                    (*self)[i] = val;
                else
                    throw std::out_of_range("vector index out of range");
            }
SWIGINTERN std::vector< char >::const_reference std_vector_Sl_char_Sg__get(std::vector< char > *self,int i){
                int size = int(self->size());
                if (i>=0 && i<size)
                    return (*self)[i];
                else
                    throw std::out_of_range("vector index out of range");
            }
SWIGINTERN void std_vector_Sl_char_Sg__set(std::vector< char > *self,int i,std::vector< char >::value_type const &val){
                int size = int(self->size());
                if (i>=0 && i<size)
                    (*self)[i] = val;
                else
                    throw std::out_of_range("vector index out of range");
            }
#ifdef __cplusplus
extern "C" {
#endif

void _wrap_Swig_free_FCLlib_1ee54ad49562e6e8(void *_swig_go_0) {
  void *arg1 = (void *) 0 ;
  
  arg1 = *(void **)&_swig_go_0; 
  
  Swig_free(arg1);
  
}


void *_wrap_Swig_malloc_FCLlib_1ee54ad49562e6e8(intgo _swig_go_0) {
  int arg1 ;
  void *result = 0 ;
  void *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = (void *)Swig_malloc(arg1);
  *(void **)&_swig_go_result = (void *)result; 
  return _swig_go_result;
}


std::vector< std::string > *_wrap_new_StringVector__SWIG_0_FCLlib_1ee54ad49562e6e8() {
  std::vector< std::string > *result = 0 ;
  std::vector< std::string > *_swig_go_result;
  
  
  result = (std::vector< std::string > *)new std::vector< std::string >();
  *(std::vector< std::string > **)&_swig_go_result = (std::vector< std::string > *)result; 
  return _swig_go_result;
}


std::vector< std::string > *_wrap_new_StringVector__SWIG_1_FCLlib_1ee54ad49562e6e8(long long _swig_go_0) {
  std::vector< std::string >::size_type arg1 ;
  std::vector< std::string > *result = 0 ;
  std::vector< std::string > *_swig_go_result;
  
  arg1 = (size_t)_swig_go_0; 
  
  result = (std::vector< std::string > *)new std::vector< std::string >(arg1);
  *(std::vector< std::string > **)&_swig_go_result = (std::vector< std::string > *)result; 
  return _swig_go_result;
}


long long _wrap_StringVector_size_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  std::vector< std::string >::size_type result;
  long long _swig_go_result;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  
  result = ((std::vector< std::string > const *)arg1)->size();
  _swig_go_result = result; 
  return _swig_go_result;
}


long long _wrap_StringVector_capacity_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  std::vector< std::string >::size_type result;
  long long _swig_go_result;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  
  result = ((std::vector< std::string > const *)arg1)->capacity();
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_StringVector_reserve_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0, long long _swig_go_1) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  std::vector< std::string >::size_type arg2 ;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  arg2 = (size_t)_swig_go_1; 
  
  (arg1)->reserve(arg2);
  
}


bool _wrap_StringVector_isEmpty_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  
  result = (bool)((std::vector< std::string > const *)arg1)->empty();
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_StringVector_clear_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  
  (arg1)->clear();
  
}


void _wrap_StringVector_add_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0, _gostring_ _swig_go_1) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  std::vector< std::string >::value_type *arg2 = 0 ;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  
  std::vector< std::string >::value_type arg2_str(_swig_go_1.p, _swig_go_1.n);
  arg2 = &arg2_str;
  
  
  (arg1)->push_back((std::vector< std::string >::value_type const &)*arg2);
  
}


_gostring_ _wrap_StringVector_get_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0, intgo _swig_go_1) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  int arg2 ;
  std::vector< std::string >::value_type *result = 0 ;
  _gostring_ _swig_go_result;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  try {
    result = (std::vector< std::string >::value_type *) &std_vector_Sl_std_string_Sg__get(arg1,arg2);
  }
  catch(std::out_of_range &_e) {
    (void)_e;
    _swig_gopanic("C++ std::out_of_range exception thrown");
    
  }
  
  _swig_go_result = Swig_AllocateString((*result).data(), (*result).length()); 
  return _swig_go_result;
}


void _wrap_StringVector_set_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0, intgo _swig_go_1, _gostring_ _swig_go_2) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  int arg2 ;
  std::vector< std::string >::value_type *arg3 = 0 ;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  std::vector< std::string >::value_type arg3_str(_swig_go_2.p, _swig_go_2.n);
  arg3 = &arg3_str;
  
  
  try {
    std_vector_Sl_std_string_Sg__set(arg1,arg2,(std::string const &)*arg3);
  }
  catch(std::out_of_range &_e) {
    (void)_e;
    _swig_gopanic("C++ std::out_of_range exception thrown");
    
  }
  
  
}


void _wrap_delete_StringVector_FCLlib_1ee54ad49562e6e8(std::vector< std::string > *_swig_go_0) {
  std::vector< std::string > *arg1 = (std::vector< std::string > *) 0 ;
  
  arg1 = *(std::vector< std::string > **)&_swig_go_0; 
  
  delete arg1;
  
}


std::vector< char > *_wrap_new_ByteVector__SWIG_0_FCLlib_1ee54ad49562e6e8() {
  std::vector< char > *result = 0 ;
  std::vector< char > *_swig_go_result;
  
  
  result = (std::vector< char > *)new std::vector< char >();
  *(std::vector< char > **)&_swig_go_result = (std::vector< char > *)result; 
  return _swig_go_result;
}


std::vector< char > *_wrap_new_ByteVector__SWIG_1_FCLlib_1ee54ad49562e6e8(long long _swig_go_0) {
  std::vector< char >::size_type arg1 ;
  std::vector< char > *result = 0 ;
  std::vector< char > *_swig_go_result;
  
  arg1 = (size_t)_swig_go_0; 
  
  result = (std::vector< char > *)new std::vector< char >(arg1);
  *(std::vector< char > **)&_swig_go_result = (std::vector< char > *)result; 
  return _swig_go_result;
}


long long _wrap_ByteVector_size_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  std::vector< char >::size_type result;
  long long _swig_go_result;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  
  result = ((std::vector< char > const *)arg1)->size();
  _swig_go_result = result; 
  return _swig_go_result;
}


long long _wrap_ByteVector_capacity_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  std::vector< char >::size_type result;
  long long _swig_go_result;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  
  result = ((std::vector< char > const *)arg1)->capacity();
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_ByteVector_reserve_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0, long long _swig_go_1) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  std::vector< char >::size_type arg2 ;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  arg2 = (size_t)_swig_go_1; 
  
  (arg1)->reserve(arg2);
  
}


bool _wrap_ByteVector_isEmpty_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  
  result = (bool)((std::vector< char > const *)arg1)->empty();
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_ByteVector_clear_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  
  (arg1)->clear();
  
}


void _wrap_ByteVector_add_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0, char _swig_go_1) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  std::vector< char >::value_type *arg2 = 0 ;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  arg2 = (std::vector< char >::value_type *)&_swig_go_1; 
  
  (arg1)->push_back((std::vector< char >::value_type const &)*arg2);
  
}


char _wrap_ByteVector_get_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0, intgo _swig_go_1) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  int arg2 ;
  std::vector< char >::value_type *result = 0 ;
  char _swig_go_result;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  try {
    result = (std::vector< char >::value_type *) &std_vector_Sl_char_Sg__get(arg1,arg2);
  }
  catch(std::out_of_range &_e) {
    (void)_e;
    _swig_gopanic("C++ std::out_of_range exception thrown");
    
  }
  
  _swig_go_result = (char)*result; 
  return _swig_go_result;
}


void _wrap_ByteVector_set_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0, intgo _swig_go_1, char _swig_go_2) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  int arg2 ;
  std::vector< char >::value_type *arg3 = 0 ;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  arg3 = (std::vector< char >::value_type *)&_swig_go_2; 
  
  try {
    std_vector_Sl_char_Sg__set(arg1,arg2,(char const &)*arg3);
  }
  catch(std::out_of_range &_e) {
    (void)_e;
    _swig_gopanic("C++ std::out_of_range exception thrown");
    
  }
  
  
}


void _wrap_delete_ByteVector_FCLlib_1ee54ad49562e6e8(std::vector< char > *_swig_go_0) {
  std::vector< char > *arg1 = (std::vector< char > *) 0 ;
  
  arg1 = *(std::vector< char > **)&_swig_go_0; 
  
  delete arg1;
  
}


void _wrap_FCLWrapper_RunELInterpretter_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0, _gostring_ _swig_go_1) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  char *arg2 = (char *) 0 ;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  (arg1)->RunELInterpretter((char const *)arg2);
  
  free(arg2); 
}


_gostring_ _wrap_FCLWrapper_GetLDALResult_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0, _gostring_ _swig_go_1) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  char *arg2 = (char *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  result = (arg1)->GetLDALResult((char const *)arg2);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  free(arg2); 
  return _swig_go_result;
}


_gostring_ _wrap_FCLWrapper_GetTDPResult_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0, _gostring_ _swig_go_1) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  char *arg2 = (char *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  result = (arg1)->GetTDPResult((char const *)arg2);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  free(arg2); 
  return _swig_go_result;
}


_gostring_ _wrap_FCLWrapper_GetLogLDALResult_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0, _gostring_ _swig_go_1) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  char *arg2 = (char *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  result = (arg1)->GetLogLDALResult((char const *)arg2);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  free(arg2); 
  return _swig_go_result;
}


_gostring_ _wrap_FCLWrapper_GetOTPResult_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0, _gostring_ _swig_go_1) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  char *arg2 = (char *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  result = (arg1)->GetOTPResult((char const *)arg2);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  free(arg2); 
  return _swig_go_result;
}


_gostring_ _wrap_FCLWrapper_GetBuildResult_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0, _gostring_ _swig_go_1) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  char *arg2 = (char *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  result = (arg1)->GetBuildResult((char const *)arg2);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  free(arg2); 
  return _swig_go_result;
}


_gostring_ _wrap_FCLWrapper_GetLogLDALResultV2_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0, _gostring_ _swig_go_1, _gostring_ _swig_go_2, _gostring_ _swig_go_3) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  char *arg2 = (char *) 0 ;
  char *arg3 = (char *) 0 ;
  char *arg4 = (char *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  arg3 = (char *)malloc(_swig_go_2.n + 1);
  memcpy(arg3, _swig_go_2.p, _swig_go_2.n);
  arg3[_swig_go_2.n] = '\0';
  
  
  arg4 = (char *)malloc(_swig_go_3.n + 1);
  memcpy(arg4, _swig_go_3.p, _swig_go_3.n);
  arg4[_swig_go_3.n] = '\0';
  
  
  result = (arg1)->GetLogLDALResultV2((char const *)arg2,(char const *)arg3,(char const *)arg4);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  free(arg2); 
  free(arg3); 
  free(arg4); 
  return _swig_go_result;
}


FCLWrapper *_wrap_new_FCLWrapper_FCLlib_1ee54ad49562e6e8() {
  FCLWrapper *result = 0 ;
  FCLWrapper *_swig_go_result;
  
  
  result = (FCLWrapper *)new FCLWrapper();
  *(FCLWrapper **)&_swig_go_result = (FCLWrapper *)result; 
  return _swig_go_result;
}


void _wrap_delete_FCLWrapper_FCLlib_1ee54ad49562e6e8(FCLWrapper *_swig_go_0) {
  FCLWrapper *arg1 = (FCLWrapper *) 0 ;
  
  arg1 = *(FCLWrapper **)&_swig_go_0; 
  
  delete arg1;
  
}


#ifdef __cplusplus
}
#endif

