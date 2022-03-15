//
//  ELListParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/26/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELListParser.h"
#include "Utils.h"
#include "MemMan.h"
#include "ELStringLiteralParser.h"

// string passed to this method should be the list string
// it should not contain the list start and end identifiers ([ and ])
// e.g. ‘Test Case \’-[“’, $TEST_SUITE, ‘ ‘, $TEST, ‘“]\’ started.’
// it is ok to have spaces at both sides and between elements - they will be removed during the parsing
bool ELListParser::ParseString(MSTRING &str, MetaData *md, VEC_ELLIST_PARSER_RESULT &results) {
    Utils::TrimLeft(str, _MSTR( \t));
    MSTRING::size_type pos = 0;
    while (true) {
        bool isAName = false;
        ELListElementType let = ELListElementType_None;
        MSTRING::size_type prefixLen = 0;
        if (Utils::IsStringContainsSubstringAtPosition(str, pos, md->s_ELBlockTemplatePrefix)) {
            isAName = true;
            let = ELListElementType_BlockTemplate;
            prefixLen = md->s_ELBlockTemplatePrefix.length();
        } else if (Utils::IsStringContainsSubstringAtPosition(str, pos, md->s_ELLineTemplatePrefix)) {
            isAName = true;
            let = ELListElementType_LineTemplate;
            prefixLen = md->s_ELLineTemplatePrefix.length();
        } else if (Utils::IsStringContainsSubstringAtPosition(str, pos, md->s_ELVarPrefix)) {
            isAName = true;
            let = ELListElementType_Variable;
            prefixLen = md->s_ELVarPrefix.length();
        }
        
        if (isAName) {
            ELListParserResult *res = 0;
            MemoryManager::Inst.CreateObject(&res);
            res->type = let;
            results.push_back(res);
            
            MSTRING::size_type sp = pos + prefixLen;
            LST_STR lstTerminatingStrings;
            lstTerminatingStrings.push_back(md->s_ELVarSequenceSeperator);
            lstTerminatingStrings.push_back(SPACE);
            lstTerminatingStrings.push_back(_MSTR(\t));
            MSTRING matchedSubstring = EMPTY_STRING;
            pos = Utils::GetFirstPositionOfSubstring(str, sp, lstTerminatingStrings, matchedSubstring);
            if (pos == MSTRING::npos) {
                res->content = str.substr(sp, str.length() - sp);
                break;
            } else {
                res->content = str.substr(sp, pos - sp);
                str = str.substr(pos, str.length() - pos);
                Utils::TrimLeft(str, _MSTR(\t \t));
                if (str.empty() || (!Utils::IsStringContainsSubstringAtPosition(str, 0, md->s_ELVarSequenceSeperator))) {
                    break;
                }
                MSTRING::size_type len = md->s_ELVarSequenceSeperator.length();
                str = str.substr(len, str.length() - len);
                Utils::TrimLeft(str, _MSTR(\t \t));
                pos = 0;
            }
        } else {
            ELStringLiteralParser slp;
            MSTRING parsedStr = EMPTY_STRING;
            MSTRING::size_type newStartPos = 0;
            if (!slp.ParseStringForLiteral(str, 0, md, parsedStr, newStartPos)) {
                return false;
            }
            ELListParserResult *res = 0;
            MemoryManager::Inst.CreateObject(&res);
            res->type = ELListElementType_String;
            res->content = parsedStr;
            results.push_back(res);
            if (newStartPos == str.length()) {
                break;
            }
            str = str.substr(newStartPos, str.length() - newStartPos);
            Utils::TrimLeft(str, _MSTR( \t));
            if (str.empty()) {
                break;
            }
            
            if (str.empty() || (!Utils::IsStringContainsSubstringAtPosition(str, 0, md->s_ELVarSequenceSeperator))) {
                break;
            }
            MSTRING::size_type len = md->s_ELVarSequenceSeperator.length();
            str = str.substr(len, str.length() - len);
            Utils::TrimLeft(str, _MSTR(\t \t));
            pos = 0;
        }
    }
    return true;
}
