//
//  ELBlockTemplate.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/24/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELBlockTemplate.h"
#include "ELBlockElement.h"

bool ELBlockTemplate::TryUnify(WIDESTRING str, WIDESTRING::size_type startPos, WIDESTRING::size_type& newPos) {
    if (isUnion) {
        return TryUnifyForUnion(str, startPos, newPos);
    } else {
        return TryUnifyForSequence(str, startPos, newPos);
    }
}

bool ELBlockTemplate::IsReadyToProcess(WIDESTRING& alreadyResolvedEntities) {
    VEC_BLOCKELEMENT::iterator ite = elements.begin();
    VEC_BLOCKELEMENT::iterator iteEnd = elements.end();
    for ( ; ite != iteEnd; ++ite) {
        WIDECHAR elementRep = (*ite)->ch;
        if (alreadyResolvedEntities.find(elementRep) == WIDESTRING::npos) {
            return false;
        }
    }
    return true;
}

bool ELBlockTemplate::IsReadyToProcessRecursively(WIDESTRING &alreadyResolvedEntities, MAP_WIDECHAR_ELBLOCKTEMPLATE& mapBlocks) {
    VEC_BLOCKELEMENT::iterator ite = elements.begin();
    VEC_BLOCKELEMENT::iterator iteEnd = elements.end();
    for ( ; ite != iteEnd; ++ite) {
        WIDECHAR elementRep = (*ite)->ch;
        if (alreadyResolvedEntities.find(elementRep) != WIDESTRING::npos) {
            continue;
        } else {
            if (!IsReadyToProcessIfDependencyIsGiven(alreadyResolvedEntities, ch, EMPTY_WIDESTRING, mapBlocks)) {
                return false;
            }
        }
    }
    return true;
}

bool ELBlockTemplate::IsReadyToProcessIfDependencyIsGiven(WIDESTRING &alreadyResolvedEntities, WIDECHAR chGivenDependency, WIDESTRING unresolvedAncestorDependencies, MAP_WIDECHAR_ELBLOCKTEMPLATE& mapBlocks) {
    WIDESTRING unresolvedAncestorDependenciesNew = unresolvedAncestorDependencies + ch;
    VEC_BLOCKELEMENT::iterator ite = elements.begin();
    VEC_BLOCKELEMENT::iterator iteEnd = elements.end();
    for ( ; ite != iteEnd; ++ite) {
        WIDECHAR elementRep = (*ite)->ch;
        if (elementRep == chGivenDependency) {
            continue;
        }
        if (alreadyResolvedEntities.find(elementRep) != WIDESTRING::npos) {
            continue;
        }
        if (unresolvedAncestorDependencies.find(elementRep) != WIDESTRING::npos) {
            return false;
        }
        
        if (!mapBlocks[elementRep]->IsReadyToProcessIfDependencyIsGiven(alreadyResolvedEntities, chGivenDependency, unresolvedAncestorDependenciesNew, mapBlocks)) {
            return false;
        }
    }
    return true;
}

bool ELBlockTemplate::IsDefinite() {
    // if first and last elements are not sequence types return true, otherwise return false
    VEC_BLOCKELEMENT::size_type len = elements.size();
    if (len == 0) {
        throw "Block cannot be empty!!";
    }
    ELBlockElement* b = elements.front();
    if (b->IsSequence()) {
        return false;
    }
    if (len > 1) {
        b = elements.at(len - 1);
        if (b->IsSequence()) {
            return false;
        }
    }
    return true;
}

bool ELBlockTemplate::TryUnifyForSequence(WIDESTRING &str, WIDESTRING::size_type startPos, WIDESTRING::size_type &newPos) {
    WIDESTRING::size_type strlen = str.length();
    newPos = startPos;
    WIDESTRING::size_type pos = startPos;
    VEC_BLOCKELEMENT::iterator ite = elements.begin();
    VEC_BLOCKELEMENT::iterator iteEnd = elements.end();
    while (true) {
        ELBlockElement *b = (*ite);
        WIDECHAR ch = str.at(pos);
        if (!b->IsSequence()) {
            if (ch != b->ch) {
                return false;
            }
            ++ite;
            if (ite == iteEnd) {
                newPos = (pos + 1);
                return true;
            }
        } else {
            if (ch != b->ch) {
                ++ite;
                if (ite == iteEnd) {
                    newPos = (pos + 1);
                    return true;
                }
                continue;
            }
        }
        ++pos;
        if (pos >= strlen) {
            if ((b->IsSequence()) && ((ite + 1) == iteEnd)) {
                newPos = pos;
                return true;
            }
            return false;
        }
    }
}

bool ELBlockTemplate::TryUnifyForUnion(WIDESTRING &str, WIDESTRING::size_type startPos, WIDESTRING::size_type &newPos) {
    newPos = startPos;
    WIDECHAR ch = str.at(startPos);
    VEC_BLOCKELEMENT::iterator ite = elements.begin();
    VEC_BLOCKELEMENT::iterator iteEnd = elements.end();
    for ( ; ite != iteEnd; ++ite) {
        ELBlockElement *b = (*ite);
        if (b->IsSequence()) {
            if (ch == b->ch) {
                newPos = str.find_first_not_of(str, startPos);
                return true;
            }
        } else {
            if (ch == b->ch) {
                newPos = startPos + 1;
                return true;
            }
        }
    }
    return false;
}