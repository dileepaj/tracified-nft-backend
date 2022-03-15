//
//  SpecialCommandExecuter.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/26/15.
//  Copyright (c) 2015 99x Eurocenter. All rights reserved.
//

#include "SpecialCommandExecuter.h"
#include "ISpecialCommand.h"
#include "NodeCustomValueShorthand.h"

SpecialCommandExecuter SpecialCommandExecuter::inst;

SpecialCommandExecuter::SpecialCommandExecuter() {
    RegisterSpecialCommands();
}

SpecialCommandExecuter::~SpecialCommandExecuter() {
    
}

void SpecialCommandExecuter::RegisterSpecialCommands() {
    commands.push_back(new NodeCustomValueShorthand);
}

PENTITY SpecialCommandExecuter::ExecuteSpecialCommand(PENTITY entity, ExecutionContext* context, Command* cmd) {
    std::vector<ISpecialCommand*>::iterator ite = commands.begin();
    std::vector<ISpecialCommand*>::iterator end = commands.end();
    for ( ; ite != end; ++ite) {
        PENTITY res = (*ite)->ExecuteSpecialCommand(entity, context, cmd);
        if (res) {
            return res;
        }
    }
    return 0;
}
