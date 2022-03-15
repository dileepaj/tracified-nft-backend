//
//  ELSpacesString.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELSpacesString.h"

ELSpacesString::ELSpacesString()
: ELFillerString() {
    AddChar(' ');
    AddChar('\t');
}

ELSpacesString::~ELSpacesString() {
    
}

bool ELSpacesString::EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos) {
    MSTRING::size_type len = str.length();
    // Spaces string will evaluate to true even if there are no matches
    if ((len == 0) || (startPos >= len)) {
        newPos = startPos;
        return true;
    }
    MCHAR ch = str.at(startPos);
    if (characters.find(ch) == characters.end()) {
        newPos = startPos;
        return true;
    }
    
    return ELFillerString::EvaluateString(str, startPos, newPos);
}
