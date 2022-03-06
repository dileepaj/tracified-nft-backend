//
//  SequenceBlockElement.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/5/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#include "ELSequenceBlockElement.h"

ELSequenceBlockElement::ELSequenceBlockElement(bool isLine) {
    if (isLine) {
        type = BlockElementType_LineSequence;
    } else {
        type = BlockElementType_BlockSequence;
    }
}
