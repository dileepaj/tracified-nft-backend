//
//  SpecialCommandExecuter.h
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/26/15.
//  Copyright (c) 2015 99x Eurocenter. All rights reserved.
//

#ifndef __LogAnalyzer__SpecialCommandExecuter__
#define __LogAnalyzer__SpecialCommandExecuter__

#include "CommonIncludes.h"

class ExecutionContext;
class Command;
class ISpecialCommand;

class SpecialCommandExecuter {
public:
    static SpecialCommandExecuter inst;
    
    SpecialCommandExecuter();
    ~SpecialCommandExecuter();
    PENTITY ExecuteSpecialCommand(PENTITY entity, ExecutionContext* context, Command* cmd);
    
private:
    std::vector<ISpecialCommand*> commands;
    
    void RegisterSpecialCommands();
};

#endif /* defined(__LogAnalyzer__SpecialCommandExecuter__) */
