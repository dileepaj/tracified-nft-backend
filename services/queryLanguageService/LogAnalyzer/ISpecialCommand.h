//
//  ISpecialCommand.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/26/15.
//  Copyright (c) 2015 99x Eurocenter. All rights reserved.
//

#ifndef LogAnalyzer_ISpecialCommand_h
#define LogAnalyzer_ISpecialCommand_h

#include "CommonIncludes.h"

class Entity;
class Command;
class ExecutionContext;

class ISpecialCommand {
public:
    virtual PENTITY ExecuteSpecialCommand(PENTITY entity, ExecutionContext* context, Command* cmd) = 0;
};


#endif
