//
//  ELInterpretter.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/29/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELInterpretter__
#define __LogAnalyzer__ELInterpretter__

#include <iostream>
#include "CommonIncludes.h"
#include "Node.h"

class ELInterpretterResult {
public:
    PNODE startNode;
    WIDESTRING finalString;
    int millisecondsForParsing;
    int millisecondsForInterpreting;
    ELInterpretterResult();
    
};

class ELLineAnnotationElement {
public:
    MSTRING::size_type start;
    MSTRING::size_type len;
    MSTRING name;
};

class ELLineAnnotation {
public:
    VEC_ELLINEANNOTATIONELEMENT elements;
};

class ELInterpretter {
public:
    ELInterpretterResult* EvaluateCase(MSTRING sDefFile);
    VEC_ELLINEANNOTATION* AnnotateAgainstLineTemplate(MSTRING sDefFile, MSTRING sLineTemplateName);
    
private:
    WIDESTRING ProcessLinesInFile(MSTRING sLogFile, VEC_ELLINETEMPLATE& vecLineTemplates, ELInterpretterResult *res);
    WIDESTRING ProcessBlocks(WIDESTRING str, VEC_ELBLOCKTEMPLATE& blockTemplates, VEC_ELLINETEMPLATE& lineTemplates, ELInterpretterResult *res);
    void UnifyBlock(WIDESTRING& str, ELBlockTemplate *block, ELInterpretterResult *res);
    void PrintInterpretterResult(ELInterpretterResult *ir);
    void PrintInterpretterResultInJSON(ELInterpretterResult *ir);
    VEC_ELBLOCKTEMPLATE GetBlockTemplatesPreparedToHandleRecursiveDefs(VEC_ELBLOCKTEMPLATE& blockTemplates);
    void FillAnnotationElements(PNODE node, ELLineAnnotation *an, MSTRING::size_type& startPos);
};

#endif /* defined(__LogAnalyzer__ELInterpretter__) */
