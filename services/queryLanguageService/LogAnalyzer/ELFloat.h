//
//  ELFloat.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELFloat__
#define __LogAnalyzer__ELFloat__

#include <iostream>
#include "ELVariable.h"

class ELFloat : public ELVariable {
public:
    ELFloat();
    virtual ~ELFloat();
    bool EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos);
};

#endif /* defined(__LogAnalyzer__ELFloat__) */
