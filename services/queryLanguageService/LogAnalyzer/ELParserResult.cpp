//
//  ELParserResult.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELParserResult.h"
#include "ELVariable.h"
#include "ELLineTemplate.h"
#include "ELBlockTemplate.h"

ELParserResult::ELParserResult() {
    
}

ELParserResult::~ELParserResult() {
    
}

void ELParserResult::Destroy() {
    MAP_STR_ELVAR::iterator ite = map_Variables.begin();
    MAP_STR_ELVAR::iterator iteEnd = map_Variables.end();
    for ( ; ite != iteEnd; ++ite) {
        delete (*ite).second;
    }
    
    VEC_ELLINETEMPLATE::iterator ite2 = vec_LineTemplates.begin();
    VEC_ELLINETEMPLATE::iterator iteEnd2 = vec_LineTemplates.end();
    for ( ; ite2 != iteEnd2; ++ite2) {
        delete (*ite2);
    }
    
    VEC_ELBLOCKTEMPLATE::iterator ite3 = vec_BlockTemplates.begin();
    VEC_ELBLOCKTEMPLATE::iterator iteEnd3 = vec_BlockTemplates.end();
    for ( ; ite3 != iteEnd3; ++ite3) {
        delete (*ite3);
    }
    
    delete this;
}
