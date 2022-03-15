//
//  ELInterpretedElement.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/11/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#include "ELInterpretedElement.h"

ELInterpretedElement::ELInterpretedElement() {
    var = NULL;
    line = NULL;
    block = NULL;
    type = ELInterpretedElementType_None;
}

ELInterpretedElement::~ELInterpretedElement() {
    
}

