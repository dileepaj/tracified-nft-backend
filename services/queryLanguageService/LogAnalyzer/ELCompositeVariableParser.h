//
//  ELCompositeVariableParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELCompositeVariableParser__
#define __LogAnalyzer__ELCompositeVariableParser__

#include <iostream>
#include "CommonIncludes.h"
#include "MetaData.h"

class ELCompositeVariableParser {
public:
    void ParseExpressions(MAP_STR_STR& expressions, MetaData *md, MAP_STR_ELVAR& basicVars, MAP_STR_ELVAR& compositeVars);
    
};

#endif /* defined(__LogAnalyzer__ELCompositeVariableParser__) */
