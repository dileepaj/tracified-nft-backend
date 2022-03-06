//
//  ELSequencevVar.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELSequencevVar__
#define __LogAnalyzer__ELSequencevVar__

#include <iostream>
#include "ELVariable.h"

class ELSequenceVar : public ELVariable {
protected:
    std::vector<ELVariable*> vars;
public:
    MSTRING expression;     // added for debugging purposes
    ELSequenceVar();
    virtual ~ELSequenceVar();
    void AddVar(ELVariable *var);
    bool EvaluateString (MSTRING& str, MSTRING::size_type startPos, MSTRING::size_type& newPos);
    virtual void AddNodesForEvaluatedStrings(PNODE parent);
};

#endif /* defined(__LogAnalyzer__ELSequencevVar__) */
