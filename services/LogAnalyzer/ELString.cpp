//
//  ELString.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELString.h"

ELString::ELString()
: ELFillerString() {
    for (char ch = 'a'; ch <= 'z'; ++ch) {
        AddChar(ch);
    }
    for (char ch = 'A'; ch <= 'Z'; ++ch) {
        AddChar(ch);
    }
    for (char ch = '0'; ch <= '9'; ++ch) {
        AddChar(ch);
    }
    AddChar('_');
}

ELString::~ELString() {
    
}
