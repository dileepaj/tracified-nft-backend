//
//  ELCompositeVariableParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELCompositeVariableParser.h"
#include "ELListParser.h"
#include "ELConstantString.h"
#include "ELVariable.h"
#include "MemMan.h"
#include "ELSequenceVar.h"
#include "ELBasicVariableParser.h"

void ELCompositeVariableParser::ParseExpressions(MAP_STR_STR &expressions, MetaData *md, MAP_STR_ELVAR &basicVars, MAP_STR_ELVAR &compositeVars) {
    MAP_STR_ELVAR currentVars(basicVars);
    MAP_STR_VECELLISTPARSERRESULT intermediates;
    MAP_STR_STR::iterator ite = expressions.begin();
    MAP_STR_STR::iterator iteEnd = expressions.end();
    for ( ; ite != iteEnd; ++ite) {
        MSTRING expr = (*ite).second;
        MSTRING::size_type len = expr.length();
        MSTRING::size_type lenSeqStart = md->s_ELVarSequenceStart.length();
        MSTRING::size_type lenSeqEnd = md->s_ELVarSequenceEnd.length();
        if ((len > (lenSeqStart + lenSeqEnd)) && (expr.substr(0, lenSeqStart) == md->s_ELVarSequenceStart) && (expr.substr(len - lenSeqEnd) == md->s_ELVarSequenceEnd)) {
            // expr is in the form [X, Y, Z]
            expr = expr.substr(lenSeqStart, (len - lenSeqStart - lenSeqEnd));
            ELListParser lp;
            VEC_ELLIST_PARSER_RESULT lpr;
            if (!lp.ParseString(expr, md, lpr)) {
                continue;
            }
            intermediates[(*ite).first] = lpr;
        }
    }
    
    while (true) {
        if (intermediates.size() == 0) {
            break;
        }
        LST_STR varsResolvedDuringThisPass;
        MAP_STR_VECELLISTPARSERRESULT::iterator ite1 = intermediates.begin();
        MAP_STR_VECELLISTPARSERRESULT::iterator iteEnd1 = intermediates.end();
        for ( ; ite1 != iteEnd1; ++ite1) {
            bool resolvable = true;
            VEC_ELLIST_PARSER_RESULT::iterator ite2 = (*ite1).second.begin();
            VEC_ELLIST_PARSER_RESULT::iterator iteEnd2 = (*ite1).second.end();
            for ( ; ite2 != iteEnd2; ++ite2) {
                ELListParserResult *lpr = *ite2;
                if (lpr->type == ELListElementType_Variable) {
                    // check whether it is an anonymous variable   e.g. $(NUMBER)
                    ELBasicVariableParser bvp;
                    MAP_STR_STR mapForSingleVar;
                    mapForSingleVar["Anonymous"] = (md->s_ELVarPrefix + lpr->content);  // remember to add the var prefix again since we removed it early
                    MAP_STR_ELVAR mapHolder;
                    bvp.ParseExpressions(mapForSingleVar, md, mapHolder);
                    if ((mapHolder.size() == 0) && (currentVars.find(lpr->content) == currentVars.end())) {
                        resolvable = false;
                        break;
                    }
                }
            }
            if (resolvable) {
                ELSequenceVar *cv = 0;
                MemoryManager::Inst.CreateObject(&cv);
                cv->name = (*ite1).first;
                cv->expression = expressions[(*ite1).first];
                ite2 = (*ite1).second.begin();
                iteEnd2 = (*ite1).second.end();
                for ( ; ite2 != iteEnd2; ++ite2) {
                    ELListParserResult *lpr = *ite2;
                    if (lpr->type == ELListElementType_String) {
                        ELConstantString *cs = 0;
                        MemoryManager::Inst.CreateObject(&cs);
                        cs->SetString(lpr->content);
                        cv->AddVar(cs);
                    } else if (lpr->type == ELListElementType_Variable) {
                        // first check whether it is an anonymous variable   e.g. $(NUMBER)
                        ELBasicVariableParser bvp;
                        MAP_STR_STR mapForSingleVar;
                        mapForSingleVar["Anonymous"] = (md->s_ELVarPrefix + lpr->content);  // remember to add the var prefix again since we removed it early
                        MAP_STR_ELVAR mapHolder;
                        bvp.ParseExpressions(mapForSingleVar, md, mapHolder);
                        if (mapHolder.size() > 0) {
                            cv->AddVar(mapHolder.begin()->second);
                        } else {
                            cv->AddVar(currentVars.find(lpr->content)->second);
                        }
                    }
                }
                compositeVars[cv->name] = cv;
                currentVars[cv->name] = cv;
                varsResolvedDuringThisPass.push_back(cv->name);
            }
        }
        
        LST_STR::iterator ite3 = varsResolvedDuringThisPass.begin();
        LST_STR::iterator iteEnd3 = varsResolvedDuringThisPass.end();
        for ( ; ite3 != iteEnd3; ++ite3) {
            intermediates.erase(*ite3);
        }
    }
}
