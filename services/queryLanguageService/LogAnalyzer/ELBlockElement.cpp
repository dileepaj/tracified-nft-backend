//
//  BlockElement.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/5/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#include "ELBlockElement.h"

bool ELBlockElement::IsSequence() {
    return ((type == BlockElementType_BlockSequence) || (type == BlockElementType_LineSequence));
}
