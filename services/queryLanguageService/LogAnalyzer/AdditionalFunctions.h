#include "CommonIncludes.h"
#include "MemMan.h"
#include <time.h>
#include "Strings.h"

PENTITY TimeStrToInt(PENTITY pStr)
{
	// converts a time in the format YYYY-MM-DD : hh:mm:ss into a long value
	PString pString = (PString)pStr;
	if(0 == pString)
	{
		return 0;
	}
	struct tm t1;
	MSTRING sStr = pString->GetValue();
	if(sStr.length() != 21)
	{
		return 0;
	}
	t1.tm_year = _MATOI(sStr.substr(0, 4).c_str()) - 1900;
	t1.tm_mon = _MATOI(sStr.substr(5, 2).c_str()) - 1;
	t1.tm_mday = _MATOI(sStr.substr(8, 2).c_str());
	t1.tm_hour = _MATOI(sStr.substr(13,2).c_str());
	t1.tm_min = _MATOI(sStr.substr(16, 2).c_str());
	t1.tm_sec = _MATOI(sStr.substr(19, 2).c_str());

	time_t tt = mktime(&t1);
	PInt pInt = 0;
	MemoryManager::Inst.CreateObject(&pInt);
	pInt->SetValue(tt);
	return pInt;
}

typedef PENTITY (*AdditionalFunc) (PENTITY);

typedef std::map<MSTRING, AdditionalFunc>	MAP_STR_ADDITIONAL_FUNC;

MAP_STR_ADDITIONAL_FUNC CreateMap()
{
	MAP_STR_ADDITIONAL_FUNC m;
	m[_MSTR(TimeStrToInt)] = TimeStrToInt;
	return m;
}

MAP_STR_ADDITIONAL_FUNC map_AdditionalFunctions = CreateMap();

