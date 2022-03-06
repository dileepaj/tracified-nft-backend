//
//  ELLineTemplate.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELLineTemplate__
#define __LogAnalyzer__ELLineTemplate__

#include <iostream>
#include "CommonIncludes.h"

class ELSequenceVar;

class ELLineTemplate {
public:
    static WIDECHAR nextChar;
    static WIDECHAR defaultChar;
    ELSequenceVar *var_Sequence;
    WIDECHAR ch;
    MSTRING name;
    
    ELLineTemplate();
    virtual ~ELLineTemplate();
    bool ParseLine(MSTRING& line);
    void CreateNodesForLineElements(PNODE parent);
};

#endif /* defined(__LogAnalyzer__ELLineTemplate__) */
