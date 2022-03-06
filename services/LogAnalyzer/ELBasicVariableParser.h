//
//  ELBasicVariableParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELBasicVariableParser__
#define __LogAnalyzer__ELBasicVariableParser__

#include <iostream>
#include "CommonIncludes.h"
#include "MetaData.h"

class ELBasicVariableParser {
public:
    void ParseExpressions(MAP_STR_STR& expressions, MetaData *md, MAP_STR_ELVAR& basicVars);
};

#endif /* defined(__LogAnalyzer__ELBasicVariableParser__) */
