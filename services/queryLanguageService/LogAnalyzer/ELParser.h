//
//  ELParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELParser__
#define __LogAnalyzer__ELParser__

#include <iostream>
#include "CommonIncludes.h"
#include "MetaData.h"
#include "ELParserResult.h"

class ELParser {
public:
    ELParser();
    virtual ~ELParser();
    
    bool ProcessScript(MSTRING& sFile, MetaData *md, ELParserResult& result);
    
private:
    void ReadFileToLines(MSTRING sFile, MSTRING sLineContinuation, MSTRING sCommentStart, LST_STR& lstLines, LST_INT& lstLineNumbers);
};

#endif /* defined(__LogAnalyzer__ELParser__) */
