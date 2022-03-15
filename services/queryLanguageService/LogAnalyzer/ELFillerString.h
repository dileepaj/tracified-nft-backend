//
//  ELFillerString.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELFillerString__
#define __LogAnalyzer__ELFillerString__

#include <iostream>
#include "ELVariable.h"

class ELFillerString : public ELVariable {
protected:
    SET_CHAR characters;
    
public:
    ELFillerString();
    virtual ~ELFillerString();
    void AddChar(MCHAR ch);
    virtual bool EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos);
};


#endif /* defined(__LogAnalyzer__ELFillerString__) */
