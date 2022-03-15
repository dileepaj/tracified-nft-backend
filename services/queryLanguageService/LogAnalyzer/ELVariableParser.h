//
//  ELVariableParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELVariableParser__
#define __LogAnalyzer__ELVariableParser__

#include <iostream>
#include "CommonIncludes.h"
#include "ELVariable.h"
#include "MetaData.h"

class ELVariableParser {
public:
    void BuildVariablesFromExpressions(MAP_STR_STR& mapExpr, MetaData *md, MAP_STR_ELVAR& vars);       // in mapExpr  first = variable name        second = expression for variable

};

#endif /* defined(__LogAnalyzer__ELVariableParser__) */
