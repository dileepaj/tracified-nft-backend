//
//  SequenceBlockElement.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/5/14.
//  Copyright (c) 2014 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__SequenceBlockElement__
#define __LogAnalyzer__SequenceBlockElement__

#include <iostream>
#include "ELBlockElement.h"

class ELSequenceBlockElement : public ELBlockElement {
public:
    ELSequenceBlockElement(bool isLine);
};

#endif /* defined(__LogAnalyzer__SequenceBlockElement__) */
