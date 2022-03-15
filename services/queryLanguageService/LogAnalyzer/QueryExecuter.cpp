// //
// // Created by AfazD on 14-Nov-19.
// //

// #include <string>
// #include <Node.h>
// #include <DefFileReader.h>
// #include <ScriptReader.h>
// #include <ExecutionContext.h>
// #include <TestCases/TestCaseBase.h>
// #include "QueryExecuter.h"
// #include "ResultGenerator.h"
// #include "EntityList.h"
// //#include "Debugger.h"

// // shared data
// int id = 0;
// MSTRING QueryExecuter::run(Node *root, MSTRING querycode)
// {

//     DefFileReader dfr;
//     // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
//     MetaData *pMD = dfr.Read("../FlexibleComputerLanguage/Defs.txt");
//     ScriptReader sr;
//     ScriptReaderOutput op;
//     //bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
//     bool bSucc = sr.ProcessScript(pMD, op, querycode);
//     if (!bSucc)
//     {
//         std::wcout << "\nFailed to read script\n";
//     }
//     ExecutionContext ec;
//     ec.p_mapFunctions = &op.map_Functions;
//     ec.p_MD = pMD;
//     Node *pY = MemoryManager::Inst.CreateNode(++id);
//     Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
//     Node *normal= MemoryManager::Inst.CreateNode(++id);
//     Node *object= MemoryManager::Inst.CreateNode(++id);
//     Node *array= MemoryManager::Inst.CreateNode(++id);
//     Node *Float=MemoryManager::Inst.CreateNode(++id);
//     Node *Integer= MemoryManager::Inst.CreateNode(++id);
//     Node *Datetime= MemoryManager::Inst.CreateNode(++id);
//     Node *String= MemoryManager::Inst.CreateNode(++id);
//     Node *cus = MemoryManager::Inst.CreateNode(++id);

//     normal->SetValue("normal");
//     object->SetValue("object");
//     array->SetValue("array");
//     Float->SetValue("1");
//     Integer->SetValue("2");
//     Datetime->SetValue("3");
//     String->SetValue("4");

//     //FLOAT IS 1
//     // INTEGER IS 2
//     // DATETIME IS 3
//     // STRING IS 4
//     std::string s = "52";
//     root->SetValue((char *)s.c_str());
//     ec.map_Var["X"] = root;
//     ec.map_Var["Y"] = pY;
//     ec.map_Var["Normal"]=normal;
//     ec.map_Var["Array"]=array;
//     ec.map_Var["Object"]=object;
//     ec.map_Var["Float"]=Float;
//     ec.map_Var["Integer"]=Integer;
//     ec.map_Var["Datetime"]=Datetime;
//     ec.map_Var["String"]=String;
//     ec.map_Var["RESULT"] = pRESULT;
//     ec.map_Var["CUS"] = cus;
//     op.p_ETL->Execute(&ec);
//     PENTITYLIST result = (PENTITYLIST)(ec.map_Var["Result"]);
//     std::cout<<pRESULT->GetAggregatedValue()<<"\n";
//     //Debugger db;
//      //db.DebugResult(ec)


//     //std::cout<<"result array size: "<<result->size()<<"\n";
//     TestCaseExecutionResult();
//     return ResultGenerator::CreateResult(pRESULT);
//    //return "Done";

// }
