//
//  ELFillerString.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELFillerString.h"

ELFillerString::ELFillerString()
: ELVariable() {
    
}

ELFillerString::~ELFillerString() {
    
}

void ELFillerString::AddChar(MCHAR ch) {
    characters.insert(ch);
}

bool ELFillerString::EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos) {
    MSTRING::size_type pos = startPos;
    MSTRING::size_type strlen = str.length();
    MSTRING stringEvaluatedSoFar = EMPTY_STRING;
    bool succ = false;
    while (true) {
        if (pos == MSTRING::npos) {
            break;
        }
        if (pos >= strlen) {
            break;
        }
        MCHAR ch = str.at(pos);
        if (characters.find(ch) != characters.end()) {
            succ = true;
            newPos = (++pos);
            stringEvaluatedSoFar += ch;
        } else {
            break;
        }
    }
    if (succ) {
        evaluatedStr = stringEvaluatedSoFar;
    }
    return succ;
}
