//
//  ELLineParserResult.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELLineParserResult__
#define __LogAnalyzer__ELLineParserResult__

#include <iostream>
#include "CommonIncludes.h"
#include "ELVariable.h"
#include "ELLineTemplate.h"
#include "ELBlockTemplate.h"
#include "ELLineParser.h"

class ELLineParserResult {
public:
    MSTRING entityName;
    MSTRING expression;
    AssignmentType  entityType;
    
    ELLineParserResult();
    virtual ~ELLineParserResult();
};

#endif /* defined(__LogAnalyzer__ELLineParserResult__) */
