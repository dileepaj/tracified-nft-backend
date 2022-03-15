//
// Created by Tharindu-Balasuriya on 9/7/2021.
//


#include <iomanip>
#include "Debugger.h"
#include "ExecutionContext.h"
#include "Null.h"
#include "string"
#include "json.hpp"
#include "Int.h"
#include "Node.h"
#include "StringOperations.h"
#include "Bool.h"
#include "EntityList.h"
#include "Strings.h"

using json = nlohmann::json;

void Debugger::DebugResult(MAP_STR_ENTITYPTR *ecVarMap,MetaData *pMD) {


    json j;
    json varibleObj;
    json childObj;


    for (auto const &x : *ecVarMap) {
        switch (x.second->ul_Type) {
            case ENTITY_TYPE_NULL:{
                    //ep:entity pointer
                    PNull ep = (PNull)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "NULL ENTITY";
                    MSTRING val;
                    if(ep->IsNull()){
                        val="TRUE";
                    }
                    varibleObj["value"] =  "Is Null : " +val ;
                    j["variables"] += varibleObj;
                    break;
                }
                case ENTITY_TYPE_INT:{
                    PInt ep = (PInt)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "INTEGER";
                    if(ep->IsNull()){
                        varibleObj["value"] = "NULL" ;
                    }else{
                        varibleObj["value"] =  std::to_string(ep->GetValue());
                    }
                    j["variables"] += varibleObj;

                    break;
                }
                case ENTITY_TYPE_STRING:{
                    PString ep = (PString)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "STRING";
                    varibleObj["value"] = ep->GetValue();
                    j["variables"] += varibleObj;
                    break;
                }
                case ENTITY_TYPE_NODE:{
                    PNODE ep = (PNODE)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "NODE";
                    json  nodeDetailsObj;
                    if(ep->IsNull()){
                        varibleObj["value"] = "NULL NODE" ;
                    }else{
                         varibleObj["value"]=Debugger::nodeToJSON(ep);
                    }
                    j["variables"] += varibleObj;
                    break;
                }
                case ENTITY_TYPE_LIST:{
                    PENTITYLIST ep = (PENTITYLIST)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "LIST";
                    if(ep->IsNull()){
                        varibleObj["value"] = "NULL" ;
                    }else{
                        varibleObj["list_size"]=std::to_string(ep->size());
                        EntityList::const_iterator ite1 = ep->begin();
                        EntityList::const_iterator iteEnd1 = ep->end();
                        auto ItemsArray = json::array();
                        for(;ite1 != iteEnd1; ++ite1){
                            MULONG currElemType = ((PENTITY)(*ite1))->ul_Type;
                            if(currElemType == ENTITY_TYPE_NODE){
                                PNODE curElm = ((PNODE)(*ite1));
                                json curElmJSON;
                                curElmJSON=Debugger::nodeToJSON(curElm);
                                ItemsArray.push_back(curElmJSON);
                            }
                        }
                        varibleObj["value"]=ItemsArray;
                    }
                    j["variables"] += varibleObj;
                    break;
                }
            }
        }
    std::ofstream o(pMD->s_DebugJSON_File);
    o << std::setw(4) << j <<"\n";
}

json  Debugger::nodeToJSON(PNODE pnode){
    json nodeJSON;
    if(pnode->GetValue()){
        nodeJSON["z_Value"]=pnode->GetValue();
    }
    if(pnode->GetLVal()){
        nodeJSON["l_Value"]=pnode->GetLVal();
    }
    if(pnode->GetRVal()){
        nodeJSON["r_Value"]=pnode->GetRVal();
    }
    if(pnode->GetCustomString()){
        nodeJSON["custom_string"]=pnode->GetCustomString();
    }
    nodeJSON["dataType"]="NODE";
    nodeJSON["child_count"]=pnode->GetChildCount();
    if(pnode->GetFirstChild()){
        PNODE firstChild = pnode->GetFirstChild();
        nodeJSON["first_child"]=Debugger::nodeToJSON(firstChild);
    }
    return nodeJSON;

}