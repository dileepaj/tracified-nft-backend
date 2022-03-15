//
//  ELBlockTemplateParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELBlockTemplateParser__
#define __LogAnalyzer__ELBlockTemplateParser__

#include <iostream>
#include "CommonIncludes.h"
#include "ELBlockTemplate.h"
#include "ELLineTemplate.h"
#include "MetaData.h"

class ELBlockElement;

class ELBlockTemplateParser {
public:
    void BuildBlockTemplatesFromExpressions(VEC_STR& vecBlockTemplateNames, MAP_STR_STR& mapExpr, MetaData *md, VEC_ELLINETEMPLATE& lts, VEC_ELBLOCKTEMPLATE& bts);       // in mapExpr  first = block template name        second = expression for block template
    ELBlockTemplate *ParseExpression(MSTRING& expr);
    
private:
    bool ProcessCollection(MSTRING expr, MetaData *md, MAP_STR_WIDECHAR& charMap, VEC_BLOCKELEMENT& blockElements);
};

#endif /* defined(__LogAnalyzer__ELBlockTemplateParser__) */
