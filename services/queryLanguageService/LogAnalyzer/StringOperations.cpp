#include "StringOperations.h"
#include "MemMan.h"

bool StringOperations::Copy(PMCHAR pDest, CPMCHAR pSrc, MULONG ulSize)
{
#ifndef WIDECHAR
	strncpy(pDest, pSrc, ulSize);	
#else
	wcsncpy(pDest, pSrc, ulSize);
#endif

	return true;
}

MULONG StringOperations::Len(CPMCHAR pStr)
{
#ifndef WIDECHAR
	return strlen(pStr);	
#else
	return wcslen(pStr);
#endif
}

void StringOperations::SetString(PPMCHAR ppDest, CPMCHAR pSrc)
{
	MemoryManager::Inst.DestroyBuffer(*ppDest);
	if(0 == pSrc)
	{
		*ppDest = 0;
		return;
	}
	MULONG ulLen = StringOperations::Len(pSrc);
	MemoryManager::Inst.CreateBuffer(ppDest, ulLen + 1);
	StringOperations::Copy(*ppDest, pSrc, ulLen);
	(*ppDest)[ulLen] = STRING_END_CHAR;
}
