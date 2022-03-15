//
//  BlockElement.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/5/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__BlockElement__
#define __LogAnalyzer__BlockElement__

#include <iostream>
#include "CommonIncludes.h"

enum BlockElementType {
    BlockElementType_Line = 0,
    BlockElementType_LineSequence,
    BlockElementType_Block,
    BlockElementType_BlockSequence
};

class ELBlockElement {
public:
    BlockElementType type;
    WIDECHAR ch;
    
    bool IsSequence();
    
protected:
    ELBlockElement() {}
};

#endif /* defined(__LogAnalyzer__BlockElement__) */
