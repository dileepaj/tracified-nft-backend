//
//  NodeCustomValueShorthand.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 1/26/15.
//  Copyright (c) 2015 99x Eurocenter. All rights reserved.
//

#include "NodeCustomValueShorthand.h"
#include "Command.h"
#include "ExecutionTemplate.h"
#include "ExecutionContext.h"
#include "MetaData.h"
#include "Strings.h"

PENTITY NodeCustomValueShorthand::ExecuteSpecialCommand(PENTITY entity, ExecutionContext* context, Command* cmd) {
    MSTRING customstr = cmd->GetAdditionalFuncName();
    if (customstr.empty()) {
        return 0;
    }
    // Create a new command as follows
    // FilterSubtree($Item.GetCustomString.IsStringEqualTo(customstr))
    Command *newcmd = new Command;
    newcmd->SetType(COMMAND_TYPE_FILTER_SUBTREE);
    ExecutionTemplate* et = new ExecutionTemplate;
    et->SetStartVarName(context->p_MD->s_ListItemVar);
    Command *innercmd1 = new Command;
    innercmd1->SetType(COMMAND_TYPE_GET_CUSTOM_STRING);
    Command *innercmd2 = new Command;
    innercmd2->SetType(COMMAND_TYPE_IS_STRING_EQUAL_TO);
    ExecutionTemplate* argForIsStringEqualTo = new ExecutionTemplate;
    argForIsStringEqualTo->SetEntity(new String(customstr));
    innercmd2->SetArg(argForIsStringEqualTo);
    et->AddCommand(innercmd1);
    et->AddCommand(innercmd2);
    newcmd->SetArg(et);
    
    return newcmd->Execute(entity, context);

}
