//
//  ELNumber.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELNumber.h"

ELNumber::ELNumber()
: ELFillerString() {
    AddChar('0');
    AddChar('1');
    AddChar('2');
    AddChar('3');
    AddChar('4');
    AddChar('5');
    AddChar('6');
    AddChar('7');
    AddChar('8');
    AddChar('9');
}

ELNumber::~ELNumber() {
    
}
