//
//  ELLineTemplate.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELLineTemplate.h"
#include "ELSequenceVar.h"

WIDECHAR ELLineTemplate::nextChar = LINE_TEMPLATE_CHAR_START;
WIDECHAR ELLineTemplate::defaultChar = UNIDENTIFIED_LINE_TEMPLATE_CHAR;

ELLineTemplate::ELLineTemplate() {
    
}

ELLineTemplate::~ELLineTemplate() {
    
}

bool ELLineTemplate::ParseLine(MSTRING& line) {
    if (!var_Sequence) {
        return false;
    }
    MSTRING::size_type pos;
    return var_Sequence->EvaluateString(line, 0, pos);
}

void ELLineTemplate::CreateNodesForLineElements(PNODE parent) {
    var_Sequence->AddNodesForEvaluatedStrings(parent);
}
