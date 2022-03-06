//
//  ELBasicVariableParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELBasicVariableParser.h"
#include "MemMan.h"
#include "ELNumber.h"
#include "ELString.h"
#include "ELSpacesString.h"
#include "ELFloat.h"
#include "ELText.h"

void ELBasicVariableParser::ParseExpressions(MAP_STR_STR &expressions, MetaData *md, MAP_STR_ELVAR &basicVars) {
    MAP_STR_STR::iterator ite = expressions.begin();
    MAP_STR_STR::iterator iteEnd = expressions.end();
    for ( ; ite != iteEnd; ++ite) {
        MSTRING varName = (*ite).first;
        MSTRING expr = (*ite).second;
        ELVariable *ret = 0;
        if (!expr.empty()) {
            MSTRING::size_type len = expr.length();
            MSTRING::size_type lenVarPrefix = md->s_ELVarPrefix.length();
            
            if ((len > lenVarPrefix) && (expr.substr(0, lenVarPrefix) == md->s_ELVarPrefix)) {
                // e.g. $X = $(NUMBER)
                MSTRING var = expr.substr(lenVarPrefix, len - lenVarPrefix);
                if (var == md->s_ELNumber) {
                    ELNumber *num = 0;
                    MemoryManager::Inst.CreateObject(&num);
                    num->name = varName;
                    ret = num;
                } else if (var == md->s_ELString) {
                    ELString *str = 0;
                    MemoryManager::Inst.CreateObject(&str);
                    str->name = varName;
                    ret = str;
                }
                else if (var == md->s_ELText) {
                    ELText *txt = 0;
                    MemoryManager::Inst.CreateObject(&txt);
                    txt->name = varName;
                    ret = txt;
                } else if (var == md->s_ELSpacesString) {
                    ELSpacesString *ss = 0;
                    MemoryManager::Inst.CreateObject(&ss);
                    ss->name = varName;
                    ret = ss;
                } else if (var == md->s_ELFloat) {
                    ELFloat *f = 0;
                    MemoryManager::Inst.CreateObject(&f);
                    f->name = varName;
                    ret = f;
                }
            }
        }
        
        if (ret) {
            basicVars[varName] = ret;
        }
    }
    
    MAP_STR_ELVAR::iterator ite2 = basicVars.begin();
    MAP_STR_ELVAR::iterator iteEnd2 = basicVars.end();
    for ( ; ite2 != iteEnd2; ++ite2) {
        expressions.erase((*ite2).first);
    }
}
