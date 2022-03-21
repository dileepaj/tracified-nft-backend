//
//  Strings.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/25/15.
//  Copyright (c) 2015 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__String__
#define __LogAnalyzer__String__

#include "Value.h"

class String : public Value<MSTRING, ENTITY_TYPE_STRING>
{
public:
    String()
    : Value<MSTRING, ENTITY_TYPE_STRING>() {
        
    }
    
    String(MSTRING str)
    : Value<MSTRING, ENTITY_TYPE_STRING>(str) {
        
    }
    
	virtual MSTRING ToString()
    {
        return val;
    }
};

#endif /* defined(__LogAnalyzer__String__) */
