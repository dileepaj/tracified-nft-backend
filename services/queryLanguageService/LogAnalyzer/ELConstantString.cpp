//
//  ELConstantString.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELConstantString.h"

ELConstantString::ELConstantString()
: ELVariable() {
    name = _MSTR(CONSTANT STRING);
}

ELConstantString::~ELConstantString() {
    
}

void ELConstantString::SetString(MSTRING str) {
    s_Const = str;
}

bool ELConstantString::EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos) {
    MSTRING::size_type len = s_Const.length();
    if (len == 0) {
        return false;
    }
    if ((startPos + len) > str.length()) {
        return false;
    }
    if (str.substr(startPos, len) == s_Const) {
        newPos = startPos + len;
        evaluatedStr = s_Const;
        return true;
    }
    return false;
}

bool ELConstantString::IsConstant() {
    return true;
}