//
//  ELNodeWrapper.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/12/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELNodeWrapper__
#define __LogAnalyzer__ELNodeWrapper__

#include <iostream>
#include "CommonIncludes.h"
#include "Node.h"

const MBYTE ELNODE_TYPE_VARIABLE = 0;
const MBYTE ELNODE_TYPE_VARIABLE_SEQUENCE = 1;
const MBYTE ELNODE_TYPE_LINE = 2;
const MBYTE ELNODE_TYPE_LINE_SEQUENCE = 3;
const MBYTE ELNODE_TYPE_BLOCK = 4;
const MBYTE ELNODE_TYPE_BLOCK_SEQUENCE = 5;

class ELVariable;

class ELNodeWrapperInfo {
public:
    MBYTE type;
    MSTRING name;
    MSTRING value;
    PVOID parserElement;    // ELVariable, ELLineTemplate or ELBlockTemplate that defines this node.
};

class ELNodeWrapper {
public:
    static MAP_NODE_WRAPPER mapNodeToWrapper;
    ELNodeWrapper(ELNodeWrapperInfo& info);
    virtual ~ELNodeWrapper();
    PNODE GetNode();
    ELNodeWrapper* AddChild(ELNodeWrapperInfo& info);
    MSTRING PrintNode();
    void PrintNodeToFile(MOFSTREAM& file);
    void PrintNodeToJSONFile(MOFSTREAM& file, int count);
    
private:
    static MULONG nodeId;
    
    ELNodeWrapper();
    void PrintNodeToFile(MOFSTREAM& file, PNODE theNode, int tabCount);
    void PrintNodeToJSONFile(MOFSTREAM &jsonfile, PNODE theNode, int tabCount, int count);
    void StartNewLine(MOFSTREAM& file, int tabCount);
    
protected:
    Node* node;
};

#endif /* defined(__LogAnalyzer__ELNodeWrapper__) */
