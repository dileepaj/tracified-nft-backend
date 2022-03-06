//
//  ELSequenceVar.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELSequenceVar.h"
#include "Node.h"
#include "ELNodeWrapper.h"

ELSequenceVar::ELSequenceVar()
: ELVariable() {
    
}

ELSequenceVar::~ELSequenceVar() {
    
}

void ELSequenceVar::AddVar(ELVariable *var) {
    vars.push_back(var);
}

bool ELSequenceVar::EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos) {
    std::vector<ELVariable*>::iterator ite = vars.begin();
    std::vector<ELVariable*>::iterator iteEnd = vars.end();
    MSTRING::size_type pos = startPos;
    for ( ; ite != iteEnd; ++ite) {
        MSTRING::size_type temp;
        if ((*ite)->EvaluateString(str, pos, temp)) {
            pos = temp;
        } else {
            return false;
        }
    }
    newPos = pos;
    evaluatedStr = str.substr(startPos, pos - startPos);
    return true;
}

void ELSequenceVar::AddNodesForEvaluatedStrings(PNODE parent) {
    ELNodeWrapperInfo info;
    info.type = ELNODE_TYPE_VARIABLE_SEQUENCE;
    info.name = name;
    info.value = evaluatedStr;
    info.parserElement = this;
    evaluatedStr = EMPTY_STRING;
    ELNodeWrapper *wrapper = new ELNodeWrapper(info);
    PNODE seqNode = wrapper->GetNode();
    parent->AppendNode(seqNode);
    
    std::vector<ELVariable*>::iterator ite = vars.begin();
    std::vector<ELVariable*>::iterator iteEnd = vars.end();
    for ( ; ite != iteEnd; ++ite) {
        ELVariable *var = (*ite);
        var->AddNodesForEvaluatedStrings(seqNode);
    }
}
