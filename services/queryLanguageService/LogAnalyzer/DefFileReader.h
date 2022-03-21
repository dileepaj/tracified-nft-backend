#ifndef _DEFFILEREADER_H
#define _DEFFILEREADER_H

#include "CommonIncludes.h"

class MetaData;

class DefFileReader
{
public:
	MetaData* Read(MSTRING sFile);
    
private:
	void ProcessLine(MSTRING& sLine, MetaData* pMD);
	void AddKeyAndValue(MetaData* pMD, MSTRING sKey, MSTRING sVal);
    void AddFuncNames(MetaData* pMD, MSTRING sKey, MSTRING sVal);
    void ModifyFilePathsIfNeeded(MetaData *md, MSTRING sDefFilePath);
    void PrependFolderIfNeeded(MSTRING folder, MSTRING& file);
    void PrependFolderToLogFilesIfNeeded(MSTRING folder, MSTRING& logFiles);
};

#endif