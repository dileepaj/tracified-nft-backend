//
//  ELLineTemplateParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELLineTemplateParser__
#define __LogAnalyzer__ELLineTemplateParser__

#include <iostream>
#include "CommonIncludes.h"
#include "ELLineTemplate.h"
#include "MetaData.h"

class ELLineTemplateParser {
public:
    void BuildLineTemplatesFromExpressions(VEC_STR& vecLineTemplateNames, MAP_STR_STR& mapExpr, MAP_STR_ELVAR& vars, MetaData *md, VEC_ELLINETEMPLATE& lts);       // in mapExpr  first = line template name        second = expression for line template

};

#endif /* defined(__LogAnalyzer__ELLineTemplateParser__) */
