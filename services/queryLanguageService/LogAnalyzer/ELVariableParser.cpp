//
//  ELVariableParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELVariableParser.h"
#include "MemMan.h"
#include "ELBasicVariableParser.h"
#include "ELCompositeVariableParser.h"

void ELVariableParser::BuildVariablesFromExpressions(MAP_STR_STR &mapExpr, MetaData *md, MAP_STR_ELVAR& vars) {
    ELBasicVariableParser bvp;
    MAP_STR_ELVAR mapBasicVars;
    bvp.ParseExpressions(mapExpr, md, mapBasicVars);
    
    ELCompositeVariableParser cvp;
    MAP_STR_ELVAR mapCompositeVars;
    cvp.ParseExpressions(mapExpr, md, mapBasicVars, mapCompositeVars);
  
    vars.insert(mapBasicVars.begin(), mapBasicVars.end());
    vars.insert(mapCompositeVars.begin(), mapCompositeVars.end());
}
