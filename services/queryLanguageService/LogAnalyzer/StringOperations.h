#include "CommonIncludes.h"

class StringOperations
{
public:
	static void SetString(PPMCHAR ppDest, CPMCHAR pSrc);
	static bool Copy(PMCHAR pDest, CPMCHAR pSrc, MULONG ulSize);
	static MULONG Len(CPMCHAR pStr);

private:
	void SetStr(PMCHAR pDest, CPMCHAR pSrc);
};