//
//  ELParserResult.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELParserResult__
#define __LogAnalyzer__ELParserResult__

#include <iostream>
#include "CommonIncludes.h"

class ELParserResult {
public:
    MAP_STR_ELVAR map_Variables;
    VEC_ELLINETEMPLATE  vec_LineTemplates;
    VEC_ELBLOCKTEMPLATE vec_BlockTemplates;
    
    ELParserResult();
    virtual ~ELParserResult();
    void Destroy();
};

#endif /* defined(__LogAnalyzer__ELParserResult__) */
