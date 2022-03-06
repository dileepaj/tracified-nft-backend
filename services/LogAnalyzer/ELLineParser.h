//
//  ELLineParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELLineParser__
#define __LogAnalyzer__ELLineParser__

#include <iostream>
#include "CommonIncludes.h"
#include "MetaData.h"

class ELLineParserResult;

typedef enum _AssignmentType {
    AssignmentType_None,
    AssignmentType_Variable,
    AssignmentType_LineTemplate,
    AssignmentType_BlockTemplate
} AssignmentType;

class ELLineParser {
public:
    ELLineParser();
    virtual ~ELLineParser();
    
    ELLineParserResult *ParseLine(MSTRING& line, MetaData *md);
    
private:
    AssignmentType GetAssignmentType(MSTRING& left, MetaData *md, MSTRING& entityName);
};

#endif /* defined(__LogAnalyzer__ELLineParser__) */
