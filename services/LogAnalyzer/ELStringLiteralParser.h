//
//  ELStringLiteralParser.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/26/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELStringLiteralParser__
#define __LogAnalyzer__ELStringLiteralParser__

#include <iostream>
#include "CommonIncludes.h"
#include "MetaData.h"

class ELStringLiteralParser {
public:
    bool ParseStringForLiteral(MSTRING str, MSTRING::size_type startPos, MetaData *md, MSTRING& parsedStr, MSTRING::size_type& endPos);
};

#endif /* defined(__LogAnalyzer__ELStringLiteralParser__) */
