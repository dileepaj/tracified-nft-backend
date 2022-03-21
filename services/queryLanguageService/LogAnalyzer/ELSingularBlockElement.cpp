//
//  SingularBlockElement.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/5/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#include "ELSingularBlockElement.h"

ELSingularBlockElement::ELSingularBlockElement(bool isLine) {
    if (isLine) {
        type = BlockElementType_Line;
    } else {
        type = BlockElementType_Block;
    }
}