//
//  NodeCustomValueShorthand.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/26/15.
//  Copyright (c) 2015 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__NodeCustomValueShorthand__
#define __LogAnalyzer__NodeCustomValueShorthand__

#include "CommonIncludes.h"
#include "ISpecialCommand.h"

class NodeCustomValueShorthand : public ISpecialCommand {
public:
    virtual PENTITY ExecuteSpecialCommand(PENTITY entity, ExecutionContext* context, Command* cmd);
};

#endif /* defined(__LogAnalyzer__NodeCustomValueShorthand__) */
