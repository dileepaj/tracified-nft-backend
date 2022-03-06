#ifndef _UTILS_H
#define _UTILS_H

#include "CommonIncludes.h"

// Global functions
// Only put template functions here
template <class Collection>
void DestroyCollection(Collection& col)
{
	typename Collection::const_iterator ite1 = col.begin();
	typename Collection::const_iterator iteEnd1 = col.end();
	for( ; ite1 != iteEnd1; ++ite1)
	{
		(*ite1)->Destroy();
	}
}

class Utils
{
public:
	static void TokenizeString(MSTRING& sStr, LST_STR& lstTokens, LST_STR& lstSep, LST_STR& lstVal);
	static void TokenizeStringBasic(MSTRING& sStr, LST_STR& lstTokens, LST_STR& lstComponents, LST_INT& lstComponentTypes);
	static void TrimLeft(MSTRING& str, CPMCHAR chars2remove);
	static void TrimRight(MSTRING& str, CPMCHAR chars2remove);
	static void MakeUpper(MSTRING &str);
	static void MakeLower(MSTRING &str);
	static bool IsStringPrefix(MSTRING sFullString, MSTRING sPotencialPrefix);
	static void ReplaceSpecialCharacters(MSTRING& sStr);
	static void ReplaceSubstrings(MSTRING& sStr, std::vector<MSTRING>& vecStringsToReplace, std::vector<MSTRING>& vecStringsToReplaceWith);
    static bool IsStringContainsSubstringAtPosition(MSTRING& str, MSTRING::size_type pos, MSTRING& substr);
    static MSTRING::size_type GetFirstPositionOfSubstring(MSTRING& str, MSTRING::size_type startPos, LST_STR& substrings, MSTRING& matchedSubstring);
    
    // Time functions - for performance testing
    static int getMilliCount();
    static int getMilliSpan(int nTimeStart);
};

#endif