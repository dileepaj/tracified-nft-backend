//
//  ELFloat.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELFloat.h"

ELFloat::ELFloat()
: ELVariable() {
    
}

ELFloat::~ELFloat() {
    
}

bool ELFloat::EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos) {
    MSTRING::size_type pos = startPos;
    bool succ = false;
    bool decimalPointFound = false;
    MSTRING nums = "0123456789";
    MSTRING stringEvaluatedSoFar = EMPTY_STRING;
    while (true) {
        if (pos == MSTRING::npos) {
            break;
        }
        MCHAR ch = str.at(pos);
        if (nums.find(ch) != MSTRING::npos) {
            ++pos;
            stringEvaluatedSoFar += ch;
            if (decimalPointFound) {
                succ = true;
                newPos = pos;
            }
        } else if (ch == '.'){
            if (decimalPointFound) {
                break;
            } else {
                stringEvaluatedSoFar += ch;
                decimalPointFound = true;
                ++pos;
            }
        } else {
            break;
        }
    }
    return succ;
}
