//
//  ELConstantString.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELConstantString__
#define __LogAnalyzer__ELConstantString__

#include <iostream>
#include "ELVariable.h"

class ELConstantString : public ELVariable {
protected:
    MSTRING s_Const;
    
public:
    ELConstantString();
    virtual ~ELConstantString();
    void SetString(MSTRING str);
    
    // overwrites ELVariable
    bool EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos);
    bool IsConstant();
};

#endif /* defined(__LogAnalyzer__ELConstantString__) */
