//
//  ELVariable.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELVariable.h"
#include "Node.h"
#include "ELNodeWrapper.h"

ELVariable::ELVariable() {
    name = EMPTY_STRING;
    evaluatedStr = EMPTY_STRING;
}

ELVariable::~ELVariable() {
    
}

void ELVariable::AddNodesForEvaluatedStrings(PNODE parent) {
    ELNodeWrapperInfo info;
    info.type = ELNODE_TYPE_VARIABLE;
    info.name = name;
    info.value = evaluatedStr;
    info.parserElement = this;
    evaluatedStr = EMPTY_STRING;
    ELNodeWrapper *wrapper = new ELNodeWrapper(info);
    PNODE newNode = wrapper->GetNode();
    parent->AppendNode(newNode);
}

bool ELVariable::IsConstant() {
    return false;
}