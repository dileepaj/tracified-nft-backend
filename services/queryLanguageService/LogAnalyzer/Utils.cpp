#include "Utils.h"
#include <sys/timeb.h>

void Utils::TokenizeStringBasic(MSTRING& sStr, LST_STR& lstTokens, LST_STR& lstComponents, LST_INT& lstComponentTypes)
{
	// sStr is the string to tokenize
	// lstTokens contains tokens to be used in tokenization
	// lstComponents contains the splitted components of the string; both string values and seperators
	// lstComponentTypes contains the component type for each component in lstComponents 1 for value, 2 for seperator
	// e.g. Splitting the string <xx>(5) with tokens <, >, ( and ) outputs
	// lstComponents as <,xx,>,(,5,)	AND
	// lstComponentTypes as 2,1,2,2,1,2
	MSTRING::size_type stFindStart = 0;
	while(true)
	{
		MSTRING::size_type stTokenPos = MSTRING::npos;
		MSTRING sToken = EMPTY_STRING;
		LST_STR::const_iterator ite1 = lstTokens.begin();
		LST_STR::const_iterator iteEnd1 = lstTokens.end();
		for( ; ite1 != iteEnd1; ++ite1)
		{
			MSTRING::size_type stPos = sStr.find(*ite1, stFindStart);
			if(MSTRING::npos != stPos)
			{
				if(MSTRING::npos == stTokenPos)
				{
					stTokenPos = stPos;
					sToken = *ite1;
				}
				else
				{
					if(stPos < stTokenPos)
					{
						stTokenPos = stPos;
						sToken = *ite1;
					}
					else if(stPos == stTokenPos)
					{
						if(sToken.length() < (*ite1).length())
						{
							stTokenPos = stPos;
							sToken = *ite1;
						}
					}
				}
			}
		}

		if(MSTRING::npos == stTokenPos)
		{
			// Tokenization is finished
			lstComponents.push_back(sStr.substr(stFindStart, sStr.length() - stFindStart));
			lstComponentTypes.push_back(1);			
			break;
		}
		else
		{
			if(stTokenPos == stFindStart)
			{
				lstComponents.push_back(sToken);
				lstComponentTypes.push_back(2);
			}
			else
			{
				lstComponents.push_back(sStr.substr(stFindStart, stTokenPos - stFindStart));
				lstComponentTypes.push_back(1);
				lstComponents.push_back(sToken);
				lstComponentTypes.push_back(2);
			}
			if(stTokenPos + sToken.length() == sStr.length())
			{
				break;
			}
			stFindStart = stTokenPos + sToken.length();
		}
	}
}
void Utils::TokenizeString(MSTRING& sStr, LST_STR& lstTokens, LST_STR& lstSep, LST_STR& lstVal)
{
	// sStr is the string to tokenize
	// lstTokens contains tokens to be used in tokenization
	// lstSep is an output parameter which will contain the seperators (tokens) in sequence
	// lstVal is a output parameter which will contain the strings seperated by the tokens
	// size of lstSep should be one higher than size of lstVal
	LST_INT lstComponentTypes;
	LST_STR lstComponents;
	TokenizeStringBasic(sStr, lstTokens, lstComponents, lstComponentTypes);
	MSTRING sCurrToken = EMPTY_STRING;
	LST_STR::const_iterator ite1 = lstComponents.begin();
	LST_STR::const_iterator iteEnd1 = lstComponents.end();
	LST_INT::const_iterator ite2 = lstComponentTypes.begin();
	for( ; ite1 != iteEnd1; ++ite1, ++ite2)
	{
		if(1 == *ite2)
		{
			if((EMPTY_STRING == sCurrToken) && !lstSep.empty())
			{
				lstVal.push_back(*ite1);
			}
			else
			{
				lstSep.push_back(sCurrToken);
				sCurrToken = EMPTY_STRING;
				lstVal.push_back(*ite1);
			}
		}
		else if(2 == *ite2)
		{
			sCurrToken += *ite1;
		}
	}
	lstSep.push_back(sCurrToken);
}

void Utils::TrimLeft(MSTRING& str, CPMCHAR chars2remove)
{
	if (!str.empty())
	{
		MSTRING::size_type pos = str.find_first_not_of(chars2remove);

		if (pos != MSTRING::npos)
		{
			str.erase(0,pos);
		}
		else
		{
			str.erase( str.begin() , str.end() ); // make empty
		}
	}
}

void Utils::TrimRight(MSTRING& str, CPMCHAR chars2remove)
{
	if (!str.empty())
	{
		MSTRING::size_type pos = str.find_last_not_of(chars2remove);

		if (pos != MSTRING::npos)
		{
			str.erase(pos + 1);
		}
		else
		{
			str.erase( str.begin() , str.end() ); // make empty
		}
	}
}

void Utils::MakeUpper(MSTRING &str)
{
	std::transform(str.begin(),str.end(),str.begin(),toupper);
}

void Utils::MakeLower(MSTRING &str)
{
	std::transform(str.begin(),str.end(),str.begin(),tolower);
}

bool Utils::IsStringPrefix(MSTRING sFullString, MSTRING sPotencialPrefix)
{
	if(sFullString.length() < sPotencialPrefix.length())
	{
		return false;
	}
	return (sFullString.substr(0, sPotencialPrefix.length()) == sPotencialPrefix);
}

void Utils::ReplaceSpecialCharacters(MSTRING& sStr)
{
	std::vector<MSTRING> vecStringsToReplace;
	std::vector<MSTRING> vecStringsToReplaceWith;

	vecStringsToReplace.push_back(_MSTR(@NEWLINE));
	vecStringsToReplace.push_back(_MSTR(@TAB));
	
	vecStringsToReplaceWith.push_back(_MSTR(\n));
	vecStringsToReplaceWith.push_back(_MSTR(\t));

	Utils::ReplaceSubstrings(sStr, vecStringsToReplace, vecStringsToReplaceWith);
}

void Utils::ReplaceSubstrings(MSTRING& sStr, std::vector<MSTRING>& vecStringsToReplace, std::vector<MSTRING>& vecStringsToReplaceWith)
{
	std::vector<MSTRING>::iterator ite1 = vecStringsToReplace.begin();
	std::vector<MSTRING>::iterator iteEnd1 = vecStringsToReplace.end();
	std::vector<MSTRING>::iterator ite2 = vecStringsToReplaceWith.begin();
	for( ; ite1 != iteEnd1; ++ite1, ++ite2)
	{
		MSTRING sStr1 = (*ite1);
		MSTRING sStr2 = (*ite2);
		MSTRING::size_type stPos = 0;
		while(true)
		{
			MSTRING::size_type stCurr = sStr.find(sStr1, stPos);
			if(MSTRING::npos != stCurr)
			{
				sStr.replace(stCurr, sStr1.length(), sStr2);
				if((stCurr + sStr1.length()) == sStr.length())
				{
					break;
				}
				stPos = stCurr + sStr2.length();
			}
			else
			{
				break;
			}
		}
	}
}

bool Utils::IsStringContainsSubstringAtPosition(MSTRING& str, MSTRING::size_type pos, MSTRING& substr) {
    MSTRING::size_type len = substr.length();
    return ((str.length() > len) && (str.substr(pos, len) == substr));
}

MSTRING::size_type Utils::GetFirstPositionOfSubstring(MSTRING& str, MSTRING::size_type startPos, LST_STR& substrings, MSTRING& matchedSubstring) {
    MSTRING::size_type ret = MSTRING::npos;
    LST_STR::iterator ite = substrings.begin();
    LST_STR::iterator iteEnd = substrings.end();
    for ( ; ite != iteEnd; ++ite) {
        MSTRING::size_type pos = str.find(*ite, startPos);
        if (pos != MSTRING::npos) {
            if ((ret == MSTRING::npos) || (pos < ret)) {
                ret = pos;
                matchedSubstring = *ite;
            }
        }
    }
    return ret;
}

int Utils::getMilliCount(){
	timeb tb;
	ftime(&tb);
	int nCount = tb.millitm + (tb.time & 0xfffff) * 1000;
	return nCount;
}

int Utils::getMilliSpan(int nTimeStart){
	int nSpan = getMilliCount() - nTimeStart;
	if(nSpan < 0)
		nSpan += 0x100000 * 1000;
	return nSpan;
}