//
//  ELBlockTemplateParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELBlockTemplateParser.h"
#include "ELLineTemplate.h"
#include "ELBlockElement.h"
#include "Utils.h"
#include "ELSingularBlockElement.h"
#include "ELSequenceBlockElement.h"
#include "MemMan.h"

ELBlockTemplate* ELBlockTemplateParser::ParseExpression(MSTRING& expr) {
    ELBlockTemplate *ret = 0;
    
    return ret;
}

void ELBlockTemplateParser::BuildBlockTemplatesFromExpressions(VEC_STR& vecBlockTemplateNames, MAP_STR_STR &mapExpr, MetaData *md, VEC_ELLINETEMPLATE& lts, VEC_ELBLOCKTEMPLATE& bts) {
    // first, assign characters to each block template
    MAP_STR_WIDECHAR charMap;
    MAP_STR_STR::iterator ite = mapExpr.begin();
    MAP_STR_STR::iterator iteEnd = mapExpr.end();
    for ( ; ite != iteEnd; ++ite) {
        charMap[(*ite).first] = ELLineTemplate::nextChar++;
        std::cout<<"\nBlock Template -> "<<(*ite).first;
        std::wcout<<L" = "<<charMap[(*ite).first];
    }
    
    VEC_ELLINETEMPLATE::iterator ite2 = lts.begin();
    VEC_ELLINETEMPLATE::iterator iteEnd2 = lts.end();
    for ( ; ite2 != iteEnd2; ++ite2) {
        ELLineTemplate *lt = (*ite2);
        charMap[lt->name] = lt->ch;
        std::cout<<"\nLine Template -> "<<lt->name;
        std::wcout<<L" = "<<lt->ch;
    }
    std::cout<<"\n\n";
    
    MSTRING::size_type lenVarSeqStart = md->s_ELVarSequenceStart.length();
    MSTRING::size_type lenVarSeqEnd = md->s_ELVarSequenceEnd.length();
    MSTRING::size_type lenVarSeq = lenVarSeqStart + lenVarSeqEnd;
    
    MSTRING::size_type lenSetStart = md->s_ELSetStart.length();
    MSTRING::size_type lenSetEnd = md->s_ELSetEnd.length();
    MSTRING::size_type lenSet = lenSetStart + lenSetEnd;
    
    VEC_STR::iterator ite3 = vecBlockTemplateNames.begin();
    VEC_STR::iterator iteEnd3 = vecBlockTemplateNames.end();
    for ( ; ite3 != iteEnd3; ++ite3) {
        MSTRING expr = mapExpr[*ite3];
        MSTRING::size_type len = expr.length();
        if ((len > lenVarSeq) && (expr.substr(0, lenVarSeqStart) == md->s_ELVarSequenceStart) && (expr.substr(len - lenVarSeqEnd, lenVarSeqEnd) == md->s_ELVarSequenceEnd)) {
            // This is a sequence   e.g. [$$TEST_START, $$NONEMPTY_LINE_SEQUENCE, $$TEST_FAILED]
            ELBlockTemplate *bt = 0;
            MemoryManager::Inst.CreateObject(&bt);
            bt->name = (*ite3);
            bt->ch = charMap[(*ite3)];
            bt->isUnion = false;
            expr = expr.substr(lenVarSeqStart, len - lenVarSeq);
            ProcessCollection(expr, md, charMap, bt->elements);
            bts.push_back(bt);
        } else if ((len > lenSet) && (expr.substr(0, lenSetStart) == md->s_ELSetStart) && (expr.substr(len - lenSetEnd, lenSetEnd) == md->s_ELSetEnd)) {
            // This is a set     e.g. {$$$TEST_EXECUTION_SUCCESS, $$$TEST_EXECUTION_FAILURE}
            ELBlockTemplate *bt = 0;
            MemoryManager::Inst.CreateObject(&bt);
            bt->name = (*ite3);
            bt->ch = charMap[(*ite3)];
            bt->isUnion = true;
            expr = expr.substr(lenSetStart, len - lenSet);
            ProcessCollection(expr, md, charMap, bt->elements);
            bts.push_back(bt);
        }
    }
}

bool ELBlockTemplateParser::ProcessCollection(MSTRING expr, MetaData *md, MAP_STR_WIDECHAR& charMap, VEC_BLOCKELEMENT &blockElements) {
    Utils::TrimLeft(expr, _MSTR(\t \t));
    Utils::TrimRight(expr, _MSTR(\t \t));
    
    MSTRING::size_type lenBlockTempletePrefix = md->s_ELBlockTemplatePrefix.length();
    MSTRING::size_type lenLineTemplatePrefix = md->s_ELLineTemplatePrefix.length();
    MSTRING seqVarSuffix = md->s_ELSequenceVarStartIndicator + md->s_ELSequenceVarSuffix;
    MSTRING::size_type lenSeqVarSuffix = seqVarSuffix.length();
    MSTRING::size_type lenVarSequenceSeperator = md->s_ELVarSequenceSeperator.length();
    
    while (true) {
        MSTRING::size_type len = expr.length();
        bool isLineTemplate = false;
        bool isBlockTemplate = false;
        MSTRING::size_type lenPrefix = 0;
        if ((len > lenBlockTempletePrefix) && (expr.substr(0, lenBlockTempletePrefix) == md->s_ELBlockTemplatePrefix)) {
            // block template
            isBlockTemplate = true;
            lenPrefix = lenBlockTempletePrefix;
        } else if ((len > lenLineTemplatePrefix) && (expr.substr(0, lenLineTemplatePrefix) == md->s_ELLineTemplatePrefix)) {
            // line template
            isLineTemplate = true;
            lenPrefix = lenLineTemplatePrefix;
        }
        
        if (isLineTemplate || isBlockTemplate) {
            LST_STR lstTerminatingStrings;
            lstTerminatingStrings.push_back(md->s_ELVarSequenceSeperator);
            lstTerminatingStrings.push_back(SPACE);
            lstTerminatingStrings.push_back(_MSTR(\t));
            MSTRING matchedSubstring = EMPTY_STRING;
            MSTRING::size_type pos = Utils::GetFirstPositionOfSubstring(expr, lenBlockTempletePrefix, lstTerminatingStrings, matchedSubstring);
            MSTRING::size_type p = pos;
            if (pos == MSTRING::npos) {
                p = len;
            }
            
            if ((p > (lenPrefix + lenSeqVarSuffix)) && (expr.substr(p - lenSeqVarSuffix, lenSeqVarSuffix) == seqVarSuffix)) {
                // this is a sequence
                MSTRING name = expr.substr(lenPrefix, p - lenPrefix - lenSeqVarSuffix);
                ELSequenceBlockElement *be = new ELSequenceBlockElement(isLineTemplate);
                be->ch = charMap[name];
                blockElements.push_back(be);
            } else if (p > lenPrefix) {
                MSTRING name = expr.substr(lenPrefix, p - lenPrefix);
                ELSingularBlockElement *be = new ELSingularBlockElement(isLineTemplate);
                be->ch = charMap[name];
                blockElements.push_back(be);
            }
            
            if (pos == MSTRING::npos) {
                break;
            }
            
            if (matchedSubstring != md->s_ELVarSequenceSeperator) {
                expr = expr.substr(pos, len - pos);
                Utils::TrimLeft(expr, _MSTR(\t \t));
                if (expr.empty()) {
                    break;
                }
                if (expr.substr(0, lenVarSequenceSeperator) != md->s_ELVarSequenceSeperator) {
                    throw "Block template parsing failed.";
                }
            } else {
                expr = expr.substr(pos, len - pos);
            }
            len = expr.length();
            if (len > (lenVarSequenceSeperator)) {
                expr = expr.substr(lenVarSequenceSeperator, len - lenVarSequenceSeperator);
                Utils::TrimLeft(expr, _MSTR(\t \t));
                if (expr.empty()) {
                    break;
                }
            }
        }
    }
    
    return true;
}