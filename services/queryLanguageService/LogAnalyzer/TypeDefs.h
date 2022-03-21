// New type definitions

#ifndef _TYPEDEFS_H
#define _TYPEDEFS_H

#include <string>
#include <sstream>
#include <list>
#include <stack>
#include "Defs.h"

class Node;
class Entity;
class Command;
template <class T, int Type> class Value;
class EntityList;
class Null;
class Int;
class Bool;
class DateTime;
class ExecutionTemplateList;
class ELVariable;
class ELLineTemplate;
class ELBlockTemplate;
class ELListParserResult;
class ELBlockElement;
class ELNodeWrapper;
class ELLineAnnotationElement;
class ELLineAnnotation;
class ELCachedEvaluationResult;
class ELBlockDetectiveResultElement;
class String;

#ifndef WIDECHAR
typedef char				MCHAR;
typedef std::string			MSTRING;
typedef std::fstream		MFSTREAM;
typedef std::ifstream		MIFSTREAM;
typedef std::ofstream		MOFSTREAM;
typedef std::ostream        MOSTREAM;
typedef std::istream        MISTREAM;
typedef std::stringstream	MSTRINGSTREAM;
#else
typedef wchar_t				MCHAR;
typedef std::wstring		MSTRING;
typedef std::wfstream		MFSTREAM;
typedef std::wifstream		MIFSTREAM;
typedef std::wofstream		MOFSTREAM;
typedef std::wostream       MOSTREAM;
typedef std::wistream       MISTREAM;
typedef std::wstringstream	MSTRINGSTREAM;
#endif

#ifdef PRODUCTION
typedef std::wstring        WIDESTRING;
typedef wchar_t             WIDECHAR;
#else
typedef std::string         WIDESTRING;
typedef char                WIDECHAR;
#endif

typedef const MCHAR			CMCHAR;
typedef MCHAR*				PMCHAR;
typedef PMCHAR*				PPMCHAR;
typedef const MCHAR*		CPMCHAR;

typedef bool				MBOOL;
typedef unsigned char		MBYTE;
typedef int					MINT;
typedef unsigned int		MUINT;
typedef long				MLONG;
typedef unsigned long		MULONG;
typedef short				MSHORT;
typedef unsigned short		MUSHORT;

typedef MBYTE*				PMBYTE;
typedef void*               PVOID;

typedef Node*				PNODE;
typedef Entity*				PENTITY;
typedef Command*			PCOMMAND;
typedef std::list<MSTRING>	LST_STR;
typedef std::vector<MSTRING> VEC_STR;
typedef std::set<MSTRING> SET_STR;
typedef std::list<int>		LST_INT;
typedef std::set<int> SET_INT;
typedef std::set<WIDECHAR> SET_WIDECHAR;
typedef std::map<int, MSTRING>	MAP_INT_STR;
typedef std::map<int, int> MAP_INT_INT;
typedef std::map<MSTRING, MULONG>	MAP_STR_MULONG;
typedef std::map<MULONG, MSTRING>	MAP_MULONG_STR;
typedef std::map<MSTRING, PENTITY>		MAP_STR_ENTITYPTR;
typedef std::list<PENTITY>	LST_ENTITYPTR;
typedef std::list<PCOMMAND> LST_COMMANDPTR;
typedef std::map<MSTRING, ExecutionTemplateList*>	MAP_STR_EXECTMPLIST;
typedef Int* PInt;
typedef Bool* PBool;
typedef String* PString;
typedef Null* PNull;
typedef EntityList*	PENTITYLIST;
typedef DateTime* PDateTime;
typedef std::set<MCHAR>     SET_CHAR;
typedef std::map<MSTRING, MSTRING>  MAP_STR_STR;
typedef std::map<MSTRING, ELVariable*>  MAP_STR_ELVAR;
typedef std::map<MSTRING, WIDECHAR> MAP_STR_WIDECHAR;
typedef std::vector<int> VEC_INT;
typedef std::map<WIDESTRING, int> MAP_WIDESTRING_INT;
typedef std::map<int, std::map<WIDESTRING, int> > MAP_INT_MAP_WIDESTRING_INT;

typedef std::vector<ELVariable*>        VEC_ELVARIABLE;
typedef std::vector<ELLineTemplate*>    VEC_ELLINETEMPLATE;
typedef std::vector<ELBlockTemplate*>   VEC_ELBLOCKTEMPLATE;
typedef std::vector<ELListParserResult*>    VEC_ELLIST_PARSER_RESULT;
typedef std::map<MSTRING, VEC_ELLIST_PARSER_RESULT > MAP_STR_VECELLISTPARSERRESULT;
typedef std::vector<ELBlockElement*> VEC_BLOCKELEMENT;
typedef std::map<PNODE, ELNodeWrapper*> MAP_NODE_WRAPPER;
typedef std::map<WIDECHAR, ELBlockTemplate*> MAP_WIDECHAR_ELBLOCKTEMPLATE;;
typedef std::vector<ELLineAnnotationElement*> VEC_ELLINEANNOTATIONELEMENT;
typedef std::vector<ELLineAnnotation*> VEC_ELLINEANNOTATION;
typedef std::map<WIDECHAR, unsigned long> MAP_CHAR_ULONG;
typedef std::map<unsigned long, ELCachedEvaluationResult*> MAP_ULONG_CACHED_EVALUATION_RESULT;
typedef std::vector<ELBlockDetectiveResultElement*> VEC_ELBLOCKDETECTIVE_RESULT_ELEMENT;

#endif

