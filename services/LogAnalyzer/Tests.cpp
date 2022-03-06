// // #include "Tests.h"
// // #include "DefFileReader.h"
// // #include "ScriptReader.h"
// // #include "MetaData.h"
// // #include "ExecutionTemplateList.h"
// // #include "ExecutionContext.h"
// // #include "Node.h"
// // #include "MemMan.h"
// // #include "Debugger.h"
// // #include "FlexibleComputerLanguage/LogJsonParser.h"
// // #include "FlexibleComputerLanguage/OTPParser.h"
// // #include "FlexibleComputerLanguage/LogJsonParser.h"
// // #include <iostream>
// // #include "FlexibleComputerLanguage/ResultGenerator.h"

// // void Tests::RunTest1()
// // {
// // 	DefFileReader dfr;
// //     MetaData* pMD = dfr.Read("../Core/TestCases/files/test1/Defs.txt");
// // 	ScriptReader sr;
// // 	ScriptReaderOutput op;
// // 	bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
// // 	if(!bSucc)
// // 	{
// // 		std::wcout<<"\nFailed to read script\n";
// // 	}
// // 	ExecutionContext ec;
// // 	ec.p_mapFunctions = &op.map_Functions;
// // 	ec.p_MD = pMD;
// // 	Node* pX = MemoryManager::Inst.CreateNode(1);
// // 	Node* pY = MemoryManager::Inst.CreateNode(2);
// // 	ec.map_Var["X"] = pX;
// // 	ec.map_Var["Y"] = pY;
// // 	op.p_ETL->Execute(&ec);
// // }

// // void Tests::RunTest2()
// // {
// // 	DefFileReader dfr;
// //     MetaData* pMD = dfr.Read("../Core/TestCases/files/test2/Defs.txt");
// // 	ScriptReader sr;
// // 	ScriptReaderOutput op;
// // 	bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
// // 	if(!bSucc)
// // 	{
// // 		std::wcout<<"\nFailed to read script\n";
// // 	}
// // 	ExecutionContext ec;
// // 	ec.p_mapFunctions = &op.map_Functions;
// // 	ec.p_MD = pMD;
// // 	Node* pX = MemoryManager::Inst.CreateNode(1);
// // 	Node* pY = MemoryManager::Inst.CreateNode(2);
// // 	Node* pZ = MemoryManager::Inst.CreateNode(3);
// // 	ec.map_Var["X"] = pX;
// // 	ec.map_Var["Y"] = pY;
// // 	ec.map_Var["Z"] = pZ;
// // 	op.p_ETL->Execute(&ec);
// // 	pX->DestroyWithSubTree();
// // 	pY->DestroyWithSubTree();
// // 	pZ->DestroyWithSubTree();
// // 	//op.p_ETL->Destroy();
// // }

// // void Tests::RunTest3()
// // {
// //     DefFileReader dfr;
// //     MetaData* pMD = dfr.Read("D:\\Tracified\\LogAnalyzer\\Leedl-backend\\LogAnalyzer\\TestCases\\files\\test3\\Defs.txt");
// //     ScriptReader sr;
// //     ScriptReaderOutput op;
// //     bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
// //     if(!bSucc)
// //     {
// //         std::wcout<<"\nFailed to read script\n";
// //     }
// //     ExecutionContext ec;
// //     ec.p_mapFunctions = &op.map_Functions;
// //     ec.p_MD = pMD;
// //     Node* pX = MemoryManager::Inst.CreateNode(1);
// //     Node* pY = MemoryManager::Inst.CreateNode(2);
// //     Node* pZ = MemoryManager::Inst.CreateNode(3);
// //     ec.map_Var["X"] = pX;
// //     ec.map_Var["Y"] = pY;
// //     ec.map_Var["Z"] = pZ;
// //     op.p_ETL->Execute(&ec);
// //     pX->DestroyWithSubTree();
// //     pY->DestroyWithSubTree();
// //     pZ->DestroyWithSubTree();
// // }



// // void Tests::RunTest4()
// // {
// //     int id = 0;
// //     DefFileReader dfr;
// //     // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
// //     MetaData *pMD = dfr.Read("C:\\Users\\thari\\Desktop\\Code2\\tests\\tdpTest\\Defs.txt");
// //     ScriptReader sr;
// //     ScriptReaderOutput op;
// //     bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
// //     //bool bSucc = sr.ProcessScript(pMD, op, querycode);
// //     if (!bSucc)
// //     {
// //         std::wcout << "\nFailed to read script\n";
// //     }
// //     ExecutionContext ec;
// //     ec.p_mapFunctions = &op.map_Functions;
// //     ec.p_MD = pMD;
// //     Node *pY = MemoryManager::Inst.CreateNode(++id);
// //     Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
// //     Node *normal= MemoryManager::Inst.CreateNode(++id);
// //     Node *object= MemoryManager::Inst.CreateNode(++id);
// //     Node *array= MemoryManager::Inst.CreateNode(++id);
// //     Node *Float=MemoryManager::Inst.CreateNode(++id);
// //     Node *Integer= MemoryManager::Inst.CreateNode(++id);
// //     Node *Datetime= MemoryManager::Inst.CreateNode(++id);
// //     Node *String= MemoryManager::Inst.CreateNode(++id);

// //     normal->SetValue("normal");
// //     object->SetValue("object");
// //     array->SetValue("array");
// //     Float->SetValue("1");
// //     Integer->SetValue("2");
// //     Datetime->SetValue("3");
// //     String->SetValue("4");

// //     std::cout << "Started\n";
// //     //Tests tt;
// //     // tt.RunTest6();
// //     std::string line;
// //     std::string jsonline;
// //     std::ifstream jsonfile (pMD->s_TREELocation);
// //     if (jsonfile.is_open())
// //     {
// //         getline (jsonfile,line);
// //         jsonline = line;
// //         jsonfile.close();
// //         std::cout<<line;
// //     }
// //     Node* jsonroot = LogJsonParser::LogJSONToNodeTree(jsonline);

// //     //FLOAT IS 1
// //     // INTEGER IS 2
// //     // DATETIME IS 3
// //     // STRING IS 4
// //     std::string s = "52";
// //     jsonroot->SetValue((char *)s.c_str());
// //     ec.map_Var["X"] = jsonroot;
// //     ec.map_Var["Y"] = pY;
// //     ec.map_Var["Normal"]=normal;
// //     ec.map_Var["Array"]=array;
// //     ec.map_Var["Object"]=object;
// //     ec.map_Var["Float"]=Float;
// //     ec.map_Var["Integer"]=Integer;
// //     ec.map_Var["Datetime"]=Datetime;
// //     ec.map_Var["String"]=String;
// //     ec.map_Var["RESULT"] = pRESULT;
// //     //op.p_ETL->Execute(&ec);


// // }

// // void Tests::RunTest5()
// // {
// //     std::wcout<<"Test 5 Started\n";
// // 	DefFileReader dfr;
// //     MetaData* pMD = dfr.Read("../Core/TestCases/files/test5/Defs.txt");
// // 	ScriptReader sr;
// // 	ScriptReaderOutput op;
// // 	bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
// // 	if(!bSucc)
// // 	{
// // 		std::wcout<<"\nFailed to read script\n";
// // 	}
// // 	ExecutionContext ec;
// // 	ec.p_mapFunctions = &op.map_Functions;
// // 	ec.p_MD = pMD;
// // 	Node* pX = MemoryManager::Inst.CreateNode(1);
// // 	Node* pY = MemoryManager::Inst.CreateNode(2);
// // 	Node* pZ = MemoryManager::Inst.CreateNode(3);
// // 	ec.map_Var["X"] = pX;
// // 	ec.map_Var["Y"] = pY;
// // 	ec.map_Var["Z"] = pZ;
// // 	op.p_ETL->Execute(&ec);
// //     std::cout <<"X : "<< pX->GetValue()<<"\n";
// //     std::cout <<"Y : "<< pY->GetValue()<<"\n";
// // 	pX->DestroyWithSubTree();
// // 	pY->DestroyWithSubTree();
// // 	pZ->DestroyWithSubTree();
// // }
// // // void Tests::RunTest6()
// // // {
// // //     DefFileReader dfr;
// // //     MetaData* pMD = dfr.Read("D:\\Tracified\\LogAnalyzer\\Leedl-backend\\LogAnalyzer\\TestCases\\files\\test3\\Defs.txt");
// // //     ScriptReader sr;
// // //     ScriptReaderOutput op;
// // //     std::cout<<pMD->s_RuleFileName;
// // //     bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
// // //     if(!bSucc)
// // //     {
// // //         std::wcout<<"\nFailed to read script\n";
// // //     }
// // //     ExecutionContext ec;
// // //     ec.p_mapFunctions = &op.map_Functions;
// // //     ec.p_MD = pMD;
// // //     Node* pX = MemoryManager::Inst.CreateNode(1);
// // //     Node* pY = MemoryManager::Inst.CreateNode(2);
// // //     Node* pZ = MemoryManager::Inst.CreateNode(3);
// // //     ec.map_Var["X"] = pX;
// // //     ec.map_Var["Y"] = pY;
// // //     ec.map_Var["Z"] = pZ;
// // //     op.p_ETL->Execute(&ec);
// // //     //std::cout <<"X : "<< pX->GetValue()<<"\n";
// // //     // std::cout <<"Y : "<< pY->GetValue()<<"\n";
// // //     std::cout <<"Z : "<< pY->GetAggregatedValue();
// // //     pX->DestroyWithSubTree();
// // //     pY->DestroyWithSubTree();
// // //     pZ->DestroyWithSubTree();
// // // }


// void Tests::TDPTest(){
//     int id = 0;
//     DefFileReader dfr;
//     // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
//     MetaData *pMD = dfr.Read("C:\\Users\\thari\\Desktop\\Code2\\tests\\newTdpTest\\Defs.txt");



//     //Read Query to string
//     std::ifstream queryFile (pMD->s_RuleFileName);
//     std::string query="";
//     std::string qline="";
//     while(getline(queryFile,qline))
//     {
//         query+=qline;
//         query+="\n";
//     }
//     //std::cout<<query;
//     ScriptReader sr;
//     ScriptReaderOutput op;
//     bool bSucc = sr.ProcessScript(pMD, op, query);
//     if (!bSucc)
//     {
//         std::wcout << "\nFailed to read script\n";
//     }

//     //Parse text to TDPNodeTree
//     std::string line;
//     std::string jsonline;
//     std::ifstream jsonfile (pMD->s_TREELocation);
//     if (jsonfile.is_open())
//     {
//         getline (jsonfile,line);
//         jsonline = line;
//         jsonfile.close();
//     }
//     Node* root = OTPParser::TDPJSONToNodeTree(jsonline);
//     ExecutionContext ec;
//     ec.p_mapFunctions = &op.map_Functions;
//     ec.p_MD = pMD;
//     Node *pY = MemoryManager::Inst.CreateNode(++id);
//     Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
//     std::string s = "52";
//     root->SetValue((char *)s.c_str());
//     ec.map_Var["X"] = root;
//     ec.map_Var["Y"] = pY;
//     ec.map_Var["RESULT"] = pRESULT;
//     op.p_ETL->Execute(&ec);
//     std::cout<<pRESULT->GetAggregatedValue();
//     //LogJsonParser::LogNodeTreetoJson(root);
// //    MOFSTREAM file;
// //    file.open("C:\\Users\\thari\\Desktop\\Code2\\tests\\logTest\\result.json");
// //    LogJsonParser::PrintNodeToFile(file,pRESULT,0);
//     //MSTRING result = ResultGenerator::CreateResult(pRESULT);
//     //std::cout<<result;
// }


// // // void Tests::LogTest() {
// // //     int id = 0;
// // //     DefFileReader dfr;
// // //     // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
// // //     MetaData *pMD = dfr.Read("C:\\Users\\thari\\Desktop\\Code2\\tests\\logTest\\Defs.txt");



// // //     //Read Query to string
// // //     std::ifstream queryFile (pMD->s_RuleFileName);
// // //     std::string query="";
// // //     std::string qline="";
// // //     while(getline(queryFile,qline))
// // //     {
// // //         query+=qline;
// // //         query+="\n";
// // //     }
// // //     std::cout<<query;
// // //     ScriptReader sr;
// // //     ScriptReaderOutput op;
// // //     bool bSucc = sr.ProcessScript(pMD, op, query);
// // //     if (!bSucc)
// // //     {
// // //         std::wcout << "\nFailed to read script\n";
// // //     }

// // //     //Parse text to LogNode
// // //     std::string line;
// // //     std::string jsonline;
// // //     std::ifstream jsonfile (pMD->s_TREELocation);
// // //     if (jsonfile.is_open())
// // //     {
// // //         getline (jsonfile,line);
// // //         jsonline = line;
// // //         jsonfile.close();
// // //     }
// // //     //std::cout<<jsonline;
// // //    Node* root = LogJsonParser::LogJSONToNodeTree(jsonline);
// // //     ExecutionContext ec;
// // //     ec.p_mapFunctions = &op.map_Functions;
// // //     ec.p_MD = pMD;
// // //     Node *pY = MemoryManager::Inst.CreateNode(++id);
// // //     Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
// // //     Node *normal= MemoryManager::Inst.CreateNode(++id);
// // //     Node *object= MemoryManager::Inst.CreateNode(++id);
// // //     Node *array= MemoryManager::Inst.CreateNode(++id);
// // //     Node *Float=MemoryManager::Inst.CreateNode(++id);
// // //     Node *Integer= MemoryManager::Inst.CreateNode(++id);
// // //     Node *Datetime= MemoryManager::Inst.CreateNode(++id);
// // //     Node *String= MemoryManager::Inst.CreateNode(++id);

// // //     normal->SetValue("normal");
// // //     object->SetValue("object");
// // //     array->SetValue("array");
// // //     Float->SetValue("1");
// // //     Integer->SetValue("2");
// // //     Datetime->SetValue("3");
// // //     String->SetValue("4");

// // //     //FLOAT IS 1
// // //     // INTEGER IS 2
// // //     // DATETIME IS 3
// // //     // STRING IS 4
// // //     std::string s = "52";
// // //     root->SetValue((char *)s.c_str());
// // //     ec.map_Var["X"] = root;
// // //     ec.map_Var["Y"] = pY;
// // //     ec.map_Var["Normal"]=normal;
// // //     ec.map_Var["Array"]=array;
// // //     ec.map_Var["Object"]=object;
// // //     ec.map_Var["Float"]=Float;
// // //     ec.map_Var["Integer"]=Integer;
// // //     ec.map_Var["Datetime"]=Datetime;
// // //     ec.map_Var["String"]=String;
// // //     ec.map_Var["RESULT"] = pRESULT;
// // //     op.p_ETL->Execute(&ec);

// // // }

