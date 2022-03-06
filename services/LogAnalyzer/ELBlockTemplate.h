//
//  ELBlockTemplate.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__ELBlockTemplate__
#define __LogAnalyzer__ELBlockTemplate__

#include <iostream>
#include "CommonIncludes.h"

class ELBlockTemplate {
public:
    VEC_BLOCKELEMENT elements;
    bool isUnion;
    MSTRING name;
    WIDECHAR ch;
    
    bool IsReadyToProcess(WIDESTRING& alreadyResolvedEntities);     // the string alreadyResolvedEntities contains all characters corresponding to line / block templates that are already resolved
    bool IsReadyToProcessRecursively(WIDESTRING& alreadyResolvedEntities, MAP_WIDECHAR_ELBLOCKTEMPLATE& mapBlocks);
    bool IsReadyToProcessIfDependencyIsGiven(WIDESTRING& alreadyResolvedEntities, WIDECHAR chGivenDependency, WIDESTRING unresolvedAncestorDependencies, MAP_WIDECHAR_ELBLOCKTEMPLATE& mapBlocks);
    bool IsDefinite();
    bool TryUnify(WIDESTRING str, WIDESTRING::size_type startPos, WIDESTRING::size_type& newPos);
    
private:
    bool TryUnifyForSequence(WIDESTRING& str, WIDESTRING::size_type startPos, WIDESTRING::size_type& newPos);
    bool TryUnifyForUnion(WIDESTRING& str, WIDESTRING::size_type startPos, WIDESTRING::size_type& newPos);
};

#endif /* defined(__LogAnalyzer__ELBlockTemplate__) */
