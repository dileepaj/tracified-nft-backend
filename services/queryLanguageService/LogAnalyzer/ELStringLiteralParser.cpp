//
//  ELStringLiteralParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/26/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELStringLiteralParser.h"
#include "Utils.h"

bool ELStringLiteralParser::ParseStringForLiteral(MSTRING str, MSTRING::size_type startPos, MetaData *md, MSTRING &parsedStr, MSTRING::size_type& endPos) {
    if (!Utils::IsStringContainsSubstringAtPosition(str, startPos, md->s_ELStringLiteralStart)) {
        return false;
    }
    MSTRING::size_type len = md->s_ELStringLiteralStart.length();
    str = str.substr(startPos + len, str.length() - (startPos + len));
    endPos = startPos + len;
    parsedStr = EMPTY_STRING;
    while (true) {
        LST_STR lstSep;
        lstSep.push_back(md->s_ELStringLiteralEscape);
        lstSep.push_back(md->s_ELStringLiteralEnd);
        MSTRING matched = EMPTY_STRING;
        MSTRING::size_type pivot = Utils::GetFirstPositionOfSubstring(str, 0, lstSep, matched);
        if (pivot == MSTRING::npos) {
            return false;
        }
        if (matched == md->s_ELStringLiteralEnd) {
            parsedStr += str.substr(0, pivot);
            endPos += (pivot + md->s_ELStringLiteralEnd.length());
            break;
        } else if (matched == md->s_ELStringLiteralEscape) {
            parsedStr += str.substr(0, pivot);
            MSTRING::size_type escapedCharacterPos = pivot + md->s_ELStringLiteralEscape.length();
            parsedStr += str.substr(escapedCharacterPos, 1);
            endPos += (escapedCharacterPos + 1);
            str = str.substr(escapedCharacterPos + 1, str.length() - (escapedCharacterPos + 1));
        }
    }
    return true;
}
