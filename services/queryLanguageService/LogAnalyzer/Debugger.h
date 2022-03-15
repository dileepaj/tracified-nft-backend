//
// Created by tharindu on 9/7/2021.
//

#ifndef CODE2_DEBUGGER_H
#define CODE2_DEBUGGER_H
#include "CommonIncludes.h"
#include "MetaData.h"
#include "json.hpp"
class ExecutionTemplate;
class ExecutionContext;
class Debugger {
public:
    void DebugResult(MAP_STR_ENTITYPTR *ecVarMap,MetaData *pMD);
    void appendNode(nlohmann::json &j ,nlohmann::json &varibleObj,const std::pair<const std::basic_string<char>, PENTITY> &x);
    nlohmann::json nodeToJSON(PNODE pnode);

};


#endif //CODE2_DEBUGGER_H
