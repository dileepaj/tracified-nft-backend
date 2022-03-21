//
//  ELLineParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELLineParser.h"
#include "MemMan.h"
#include "Utils.h"
#include "ELLineParserResult.h"

ELLineParser::ELLineParser() {
    
}

ELLineParser::~ELLineParser() {
    
}

ELLineParserResult* ELLineParser::ParseLine(MSTRING &line, MetaData *md) {
    ELLineParserResult *res = 0;
    MemoryManager::Inst.CreateObject(&res);
    
    MSTRING::size_type pos = line.find(md->s_ELAssignment);
    MSTRING::size_type posEnd = pos + md->s_ELAssignment.length();
    MSTRING::size_type len = line.length();
    if ((pos != MSTRING::npos) && (pos > 0) && (posEnd < len)) {
        MSTRING left = line.substr(0, pos);
        MSTRING right = line.substr(posEnd, len - posEnd);
        Utils::TrimLeft(left, _MSTR(\t \t));
        Utils::TrimRight(left, _MSTR(\t \t));
        Utils::TrimLeft(right, _MSTR(\t \t));
        Utils::TrimRight(right, _MSTR(\t \t));
        MSTRING::size_type leftLen = left.length();
        MSTRING::size_type rightLen = right.length();
        
        if ((leftLen > 0) && (rightLen > 0)) {
            MSTRING entityName = EMPTY_STRING;
            AssignmentType at = GetAssignmentType(left, md, entityName);
            res->entityName = entityName;
            res->entityType = at;
            res->expression = right;
        }
    }
    
    return res;
}

AssignmentType ELLineParser::GetAssignmentType(MSTRING &left, MetaData *md, MSTRING& entityName) {
    MSTRING::size_type len = left.length();
    MSTRING::size_type lenVarPrefix = md->s_ELVarPrefix.length();
    MSTRING::size_type lenLineTemplatePrefix = md->s_ELLineTemplatePrefix.length();
    MSTRING::size_type lenBlockTemplatePrefix = md->s_ELBlockTemplatePrefix.length();
    if ((len > lenBlockTemplatePrefix) && (left.substr(0, lenBlockTemplatePrefix) == md->s_ELBlockTemplatePrefix)) {
        entityName = left.substr(lenBlockTemplatePrefix, len - lenBlockTemplatePrefix);
        return AssignmentType_BlockTemplate;
    }
    if ((len > lenLineTemplatePrefix) && (left.substr(0, lenLineTemplatePrefix) == md->s_ELLineTemplatePrefix)) {
        entityName = left.substr(lenLineTemplatePrefix, len - lenLineTemplatePrefix);
        return AssignmentType_LineTemplate;
    }
    if ((len > lenVarPrefix) && (left.substr(0, lenVarPrefix) == md->s_ELVarPrefix)) {
        entityName = left.substr(lenVarPrefix, len - lenVarPrefix);
        return AssignmentType_Variable;
    }
    return AssignmentType_None;
}

