//
//  ELInterpretedElement.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/11/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELInterpretedElement__
#define __LogAnalyzer__ELInterpretedElement__

#include <iostream>
#include "CommonIncludes.h"

class ELVariable;
class ELLineTemplate;
class ELBlockTemplate;

enum ELInterpretedElementType {
    ELInterpretedElementType_None = 0,
    ELInterpretedElementType_Variable,
    ELInterpretedElementType_Line,
    ELInterpretedElementType_Block
};

class ELInterpretedElement {
public:
    ELVariable *var;
    ELLineTemplate *line;
    ELBlockTemplate *block;
    ELInterpretedElementType type;
    
    ELInterpretedElement();
    virtual ~ELInterpretedElement();
};

#endif /* defined(__LogAnalyzer__ELInterpretedElement__) */
