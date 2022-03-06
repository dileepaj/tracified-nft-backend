//
//  ELLineTemplateParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELLineTemplateParser.h"
#include "ELCompositeVariableParser.h"
#include "ELLineTemplate.h"
#include "MemMan.h"
#include "ELSequenceVar.h"

void ELLineTemplateParser::BuildLineTemplatesFromExpressions(VEC_STR& vecLineTemplateNames, MAP_STR_STR &mapExpr, MAP_STR_ELVAR& vars, MetaData *md, VEC_ELLINETEMPLATE& lts) {
    ELCompositeVariableParser cvp;
    MAP_STR_ELVAR lines;
    cvp.ParseExpressions(mapExpr, md, vars, lines);
    VEC_STR::iterator ite = vecLineTemplateNames.begin();
    VEC_STR::iterator iteEnd = vecLineTemplateNames.end();
    for ( ; ite != iteEnd; ++ite) {
        ELLineTemplate *lt = 0;
        MemoryManager::Inst.CreateObject(&lt);
        lt->ch = ELLineTemplate::nextChar++;
        lt->var_Sequence = (ELSequenceVar *) (lines[*ite]);
        lt->name = lt->var_Sequence->name;
        lts.push_back(lt);
//        std::cout<<"\nLine Template Parsed -> "<<lt->name;
//        std::wcout<<L" = "<<lt->ch;
    }
}
