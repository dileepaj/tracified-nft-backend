//
// Created by tharindu on 8/27/2021.
//

#include <string>
#include "LDAL_Wrapper.h"
#include "DefFileReader.h"
#include "ScriptReader.h"
#include "MetaData.h"
#include "ExecutionTemplateList.h"
#include "ExecutionContext.h"
#include "Node.h"
#include "MemMan.h"
#include "ELInterpretter.h"
#include "CommonIncludes.h"
#include "Debugger.h"
#include "OTPParser.h"
#include "LogJsonParser.h"
#include "ELNodeWrapper.h"
#include "ResultGenerator.h"


std::string LDAL_Wrapper::GetLDALResult(std::string defFilePath)
{

    DefFileReader dfr;
    MetaData *pMD = dfr.Read(defFilePath);

    std::cout << defFilePath;
    ScriptReader sr;
    ScriptReaderOutput op;
    bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
    if (!bSucc)
    {
        std::wcout << "\nFailed to read script\n";
        return "Failed to read script";
    }
    ExecutionContext ec;
    ec.p_mapFunctions = &op.map_Functions;
    ec.p_MD = pMD;
    Node *pLog = MemoryManager::Inst.CreateNode(1);
    Node *pY = MemoryManager::Inst.CreateNode(2);
    Node *pOut = MemoryManager::Inst.CreateNode(3);

    ec.map_Var["LOG"] = pLog;
    ec.map_Var["OUTPUT"] = pOut;
    ec.map_Var["Y"] = pY;

    std::string Location = pMD->s_TREELocation;
    pLog->ReadValueFromFile(Location.c_str());

    op.p_ETL->Execute(&ec);

    std::string result = pOut->GetAggregatedValue();
    Debugger b;
    b.DebugResult(&ec.map_Var, pMD);

    pLog->DestroyWithSubTree();
    pY->DestroyWithSubTree();
    pOut->DestroyWithSubTree();

    return result;
}

std::string LDAL_Wrapper::GetTDPResult(std::string defFilePath)
{

    int id = 0;
    DefFileReader dfr;
    // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
    MetaData *pMD = dfr.Read(defFilePath);

    //Read Query to string
    std::ifstream queryFile(pMD->s_RuleFileName);
    std::string query = "";
    std::string qline = "";
    while (getline(queryFile, qline))
    {
        query += qline;
        query += "\n";
    }
    //std::cout<<query;
    ScriptReader sr;
    ScriptReaderOutput op;
   std::string bSucc = sr.ProcessScript(pMD, op, query);
       if (bSucc!="")
       {

           std::wcout << "\nFailed to read script\n";
           return bSucc;
       }

    //Parse text to TDPNodeTree
    std::string line;
    std::string jsonline;
    std::ifstream jsonfile(pMD->s_TREELocation);
    if (jsonfile.is_open())
    {
        getline(jsonfile, line);
        jsonline = line;
        jsonfile.close();
    }
    OTPParser otp;
    Node *root = otp.TDPJSONToNodeTree(jsonline);
    //Node *root= LogJsonParser::LogJSONToNodeTree(jsonline);
    ExecutionContext ec;
    ec.p_mapFunctions = &op.map_Functions;
    ec.p_MD = pMD;
    Node *pY = MemoryManager::Inst.CreateNode(++id);
    Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
    std::string s = "52";
    root->SetValue((char *)s.c_str());
    ec.map_Var["X"] = root;
    ec.map_Var["Y"] = pY;
    ec.map_Var["RESULT"] = pRESULT;
    op.p_ETL->Execute(&ec);
    Debugger db;
    db.DebugResult(&ec.map_Var,pMD);
    //std::cout << pRESULT->GetAggregatedValue();
    return ResultGenerator::CreateResult(pRESULT);
}

std::string LDAL_Wrapper::GetOTPResult(std::string defFilePath){
        int id = 0;
    DefFileReader dfr;
    // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
    MetaData *pMD = dfr.Read(defFilePath);

    //Read Query to string
    std::ifstream queryFile(pMD->s_RuleFileName);
    std::string query = "";
    std::string qline = "";
    while (getline(queryFile, qline))
    {
        query += qline;
        query += "\n";
    }
    //std::cout<<query;
    ScriptReader sr;
    ScriptReaderOutput op;
   std::string bSucc = sr.ProcessScript(pMD, op, query);
       if (bSucc!="")
       {

           std::wcout << "\nFailed to read script\n";
           return bSucc;
       }

    //Parse text to TDPNodeTree
    std::string line;
    std::string jsonline;
    std::ifstream jsonfile(pMD->s_TREELocation);
    if (jsonfile.is_open())
    {
        getline(jsonfile, line);
        jsonline = line;
        jsonfile.close();
    }
    OTPParser otp;
    Node *root = otp.OTPJSONToNodeTree(jsonline);
    //Node *root= LogJsonParser::LogJSONToNodeTree(jsonline);
    ExecutionContext ec;
    ec.p_mapFunctions = &op.map_Functions;
    ec.p_MD = pMD;
    Node *pY = MemoryManager::Inst.CreateNode(++id);
    Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
    std::string s = "52";
    root->SetValue((char *)s.c_str());
    ec.map_Var["X"] = root;
    ec.map_Var["Y"] = pY;
    ec.map_Var["RESULT"] = pRESULT;
    op.p_ETL->Execute(&ec);
    Debugger db;
    db.DebugResult(&ec.map_Var,pMD);
    //std::cout << pRESULT->GetAggregatedValue();
    return ResultGenerator::CreateResult(pRESULT);

}


std::string LDAL_Wrapper::GetLOGLDALResult(std::string defFilePath) {

    int id = 0;
    DefFileReader dfr;
    // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
    MetaData *pMD = dfr.Read(defFilePath);

    //Read Query to string
    std::ifstream queryFile(pMD->s_RuleFileName);
    std::string query = "";
    std::string qline = "";
    while (getline(queryFile, qline))
    {
        query += qline;
        query += "\n";
    }
    //std::cout<<query<<"\n";
    ScriptReader sr;
    ScriptReaderOutput op;
    std::string bSucc = sr.ProcessScript(pMD, op, query);
        if (bSucc!="")
        {

            std::wcout << "\nFailed to read script\n";
            return bSucc;
        }

    //Parse text to TDPNodeTree
    std::string line;
    std::string jsonline;
    std::ifstream jsonfile(pMD->s_TREELocation);
    if (jsonfile.is_open())
    {
        getline(jsonfile, line);
        jsonline = line;
        jsonfile.close();
    }

    //std::cout<<jsonline<<"\n";
    Node *root= LogJsonParser::LogJSONToNodeTree(jsonline);
    ExecutionContext ec;
    ec.p_mapFunctions = &op.map_Functions;
    ec.p_MD = pMD;
    Node *pY = MemoryManager::Inst.CreateNode(++id);
    Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
    std::string s = "52";
    root->SetValue((char *)s.c_str());
    ec.map_Var["X"] = root;
    ec.map_Var["Y"] = pY;
    ec.map_Var["RESULT"] = pRESULT;
    op.p_ETL->Execute(&ec);
    Debugger db;
    db.DebugResult(&ec.map_Var,pMD);
    //std::cout << pRESULT->GetAggregatedValue();
    std::string result="";

      result = ResultGenerator::CreateResult(pRESULT);

    return result;

}


//Build method
std::string  LDAL_Wrapper::GetBuildResult(std::string defFilePath) {
    nlohmann::json j;
    int id = 0;
    DefFileReader dfr;
    // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
    MetaData *pMD = dfr.Read(defFilePath);

    //Read Query to string
    std::ifstream queryFile(pMD->s_RuleFileName);
    std::string query = "";
    std::string qline = "";
    while (getline(queryFile, qline))
    {
        query += qline;
        query += "\n";
    }

    ScriptReader sr;
    ScriptReaderOutput op;
    std::string bSucc = sr.ProcessScript(pMD, op, query);
        if (bSucc!="")
        {
         std::wcout << "\nFailed to read script\n";

        j["Result"] = { {"status", false}, {"message",bSucc } };
        return j.dump();

        }
        else{

        j["Result"] = { {"status", true}, {"message","" } };
         return j.dump();
        }



}
std::string LDAL_Wrapper::GetLOGLDALResultV2(std::string defFilePath,std::string queryString,std::string jsonString) {

    int id = 0;
    DefFileReader dfr;
    // CAUTION: This file path is hardcoded and can cause crashes. You have been warned!
    MetaData *pMD = dfr.Read(defFilePath);
    ScriptReader sr;
    ScriptReaderOutput op;
    queryString.erase(std::remove(queryString.begin(),  queryString.end(), ' '),queryString.end());

    std::string bSucc = sr.ProcessScript(pMD, op, queryString);
        if (bSucc!="")
        {

            std::wcout << "\nFailed to read script\n";
            return bSucc;
        }

    //Parse text to TDPNodeTree

    std::string json ="";
    std::string jsonline = "";


    MSTRINGSTREAM jsonStringStream(jsonString);

    while (std::getline(jsonStringStream,jsonline)){

        json+=jsonline;
    }
    jsonStringStream.clear();

    //std::cout<<jsonline<<"\n";
    Node *root= LogJsonParser::LogJSONToNodeTree(json);
    ExecutionContext ec;
    ec.p_mapFunctions = &op.map_Functions;
    ec.p_MD = pMD;
    Node *pY = MemoryManager::Inst.CreateNode(++id);
    Node *pRESULT = MemoryManager::Inst.CreateNode(++id);
    std::string s = "52";
    root->SetValue((char *)s.c_str());
    ec.map_Var["X"] = root;
    ec.map_Var["Y"] = pY;
    ec.map_Var["RESULT"] = pRESULT;
    op.p_ETL->Execute(&ec);
    Debugger db;
    db.DebugResult(&ec.map_Var,pMD);
    //std::cout << pRESULT->GetAggregatedValue();
    std::string result="";

      result = ResultGenerator::CreateResult(pRESULT);

    return result;

}
