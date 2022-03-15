//
//  ELListParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/26/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELListParser__
#define __LogAnalyzer__ELListParser__

#include <iostream>
#include "CommonIncludes.h"
#include "MetaData.h"

enum ELListElementType {
    ELListElementType_None,
    ELListElementType_String,
    ELListElementType_Variable,
    ELListElementType_LineTemplate,
    ELListElementType_BlockTemplate
};

class ELListParserResult {
public:
    ELListElementType type;
    MSTRING content;    // will contain the string if the element is a string, otherwise the name of the entity (e.g. line template name)
    
    ELListParserResult() {
        type = ELListElementType_None;
        content = EMPTY_STRING;
    }
};

class ELListParser {
public:
    bool ParseString(MSTRING& str, MetaData* md, VEC_ELLIST_PARSER_RESULT& results);
};

#endif /* defined(__LogAnalyzer__ELListParser__) */
