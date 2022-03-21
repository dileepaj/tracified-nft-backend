//
//  ELSpacesString.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELSpacesString__
#define __LogAnalyzer__ELSpacesString__

#include <iostream>
#include "ELFillerString.h"

class ELSpacesString : public ELFillerString {
public:
    ELSpacesString();
    virtual ~ELSpacesString();
    
    virtual bool EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos);
};

#endif /* defined(__LogAnalyzer__ELSpacesString__) */
