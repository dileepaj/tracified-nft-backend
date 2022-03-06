//
//  ELNodeWrapper.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/12/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#include "ELNodeWrapper.h"
#include "MemMan.h"
#include "ELVariable.h"
#include "json.hpp"

using json = nlohmann::json;

MAP_NODE_WRAPPER ELNodeWrapper::mapNodeToWrapper;
MULONG ELNodeWrapper::nodeId = 1;

ELNodeWrapper::ELNodeWrapper()
{
    node = NULL;
}

ELNodeWrapper::ELNodeWrapper(ELNodeWrapperInfo &info)
{
    node = MemoryManager::Inst.CreateNode(nodeId++);
    node->SetNature(info.type);
    node->SetValue((PMCHAR)info.value.c_str());
    node->SetCustomString((PMCHAR)info.name.c_str());
    node->SetCustomObj(info.parserElement);
    ELNodeWrapper::mapNodeToWrapper[node] = this;
}

ELNodeWrapper::~ELNodeWrapper()
{
}

PNODE ELNodeWrapper::GetNode()
{
    return node;
}

ELNodeWrapper *ELNodeWrapper::AddChild(ELNodeWrapperInfo &info)
{
    ELNodeWrapper *wrapper = new ELNodeWrapper(info);
    PNODE childNode = wrapper->GetNode();
    node->AppendNode(childNode);
    return wrapper;
}

// yet to be implemented (if needed)
MSTRING ELNodeWrapper::PrintNode()
{
    return EMPTY_STRING;
}

void ELNodeWrapper::PrintNodeToFile(MOFSTREAM &file)
{
    PrintNodeToFile(file, node, 0);
}

void ELNodeWrapper::PrintNodeToJSONFile(MOFSTREAM &file,int count)
{
    PrintNodeToJSONFile(file, node, 0,count);
}

void ELNodeWrapper::PrintNodeToFile(MOFSTREAM &file, PNODE theNode, int tabCount)
{
    MSTRING nodeName = theNode->GetCustomString();
    MSTRING nodeValue = theNode->GetValue();
    MBYTE nodeType = theNode->GetNature();
    MSTRING nodeTypeString = EMPTY_STRING;
    if (nodeType == ELNODE_TYPE_VARIABLE)
    {
        nodeTypeString = _MSTR(VARIABLE);
    }
    else if (nodeType == ELNODE_TYPE_VARIABLE_SEQUENCE)
    {
        nodeTypeString = _MSTR(VARIABLE SEQUENCE);
    }
    else if (nodeType == ELNODE_TYPE_LINE)
    {
        nodeTypeString = _MSTR(LINE);
    }
    else if (nodeType == ELNODE_TYPE_LINE_SEQUENCE)
    {
        nodeTypeString = _MSTR(LINE SEQUENCE);
    }
    else if (nodeType == ELNODE_TYPE_BLOCK)
    {
        nodeTypeString = _MSTR(BLOCK);
    }
    else if (nodeType == ELNODE_TYPE_BLOCK_SEQUENCE)
    {
        nodeTypeString = _MSTR(BLOCK SEQUENCE);
    }

    StartNewLine(file, tabCount);
    file << nodeTypeString << _MSTR(\t
                                    :\t)
         << nodeName;
    StartNewLine(file, tabCount);
    file << _MSTR(
        {);
            StartNewLine(file, tabCount);
            file << _MSTR(\tValue =) << SPACE << QUOTE << nodeValue << QUOTE;
            StartNewLine(file, tabCount);

            PNODE child = theNode->GetFirstChild();
            bool hasChildren = false;
            while (child != NULL)
            {
                hasChildren = true;
                PrintNodeToFile(file, child, tabCount + 1);
                child = child->GetRightSibling();
            }
            if (hasChildren)
            {
                StartNewLine(file, tabCount);
            }

    file << _MSTR(
        });
}

void ELNodeWrapper::PrintNodeToJSONFile(MOFSTREAM &jsonfile, PNODE theNode, int tabCount ,int count)
{
    std::string str = std::to_string(count);
    MSTRING nodeName = theNode->GetCustomString();
    MSTRING nodeValue = theNode->GetValue();
    MBYTE nodeType = theNode->GetNature();
    MSTRING nodeTypeString = EMPTY_STRING;
    if (nodeType == ELNODE_TYPE_VARIABLE) {
        nodeTypeString = _MSTR(VARIABLE);
    } else if (nodeType == ELNODE_TYPE_VARIABLE_SEQUENCE) {
        nodeTypeString = _MSTR(VARIABLE SEQUENCE);
    } else if (nodeType == ELNODE_TYPE_LINE) {
        nodeTypeString = _MSTR(LINE);
    } else if (nodeType == ELNODE_TYPE_LINE_SEQUENCE) {
        nodeTypeString = _MSTR(LINE SEQUENCE);
    } else if (nodeType == ELNODE_TYPE_BLOCK) {
        nodeTypeString = _MSTR(BLOCK);
    } else if (nodeType == ELNODE_TYPE_BLOCK_SEQUENCE) {
        nodeTypeString = _MSTR(BLOCK SEQUENCE);
    }

    json j1 =nodeTypeString.append(str) ;
    json j2 = nodeName;
    json j3 = _MSTR(Value);
    json j4 = nodeValue;

    //StartNewLine(jsonfile, tabCount);
    jsonfile <<j1 << _MSTR(:)<<_MSTR({)<<j2<< _MSTR(:);
                                             //StartNewLine(jsonfile, tabCount);
                                            jsonfile<< _MSTR({)<<j3 << _MSTR(:)<<j4;
                jsonfile<<_MSTR(});
                                             PNODE child = theNode->GetFirstChild();
                                             bool hasChildren = false;
                                             while (child != NULL) {
            hasChildren = true;
            jsonfile << "," ;
            PrintNodeToJSONFile(jsonfile, child, tabCount + 1,count);
            child = child->GetRightSibling();
            count++;
        }
                                             jsonfile << _MSTR(});
}

void ELNodeWrapper::StartNewLine(MOFSTREAM &file, int tabCount)
{
    file << _MSTR(\n);
    for (int i = 0; i < tabCount; ++i)
    {
        file << _MSTR(\t);
    }
}
