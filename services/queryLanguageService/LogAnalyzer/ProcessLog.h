#ifndef _PROCESSLOG_H
#define _PROCESSLOG_H

#include "CommonIncludes.h"
#include <time.h>

class ProcessLog
{
public:
	static void Write(MSTRING& sStr)
	{
		time_t rawtime;
		struct tm * timeinfo;
		time(&rawtime);
		timeinfo = gmtime(&rawtime);
		MOFSTREAM file;
		file.open(_MSTR(ProcessLog.txt), std::ios::out | std::ios::app);
		file<<_MSTR([)<<timeinfo->tm_year<<_MSTR(-)<<timeinfo->tm_mon<<_MSTR(-)<<timeinfo->tm_mday<<SPACE<<timeinfo->tm_hour<<_MSTR(:)<<timeinfo->tm_min<<_MSTR(:)<<timeinfo->tm_sec<<_MSTR(])<<SPACE<<sStr.c_str()<<_MSTR(\n);
		file.close();
	}
};

#endif