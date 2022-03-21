//
//  OTPParser.cpp
//  FlexibleComputerLanguage
//
//  Created by Murtaza Anverali on 7/5/18.
//  Copyright © 2018 Dileepa Jayathilaka. All rights reserved.
//

#include "document.h"
#include "json.hpp"
#include "OTPParser.h"
#include "MemMan.h"
#include "Node.h"
#include "MetaData.h"
#include "Strings.h"
#include "Entity.h"
#include "easylogging++.h"

#include "writer.h"
#include "stringbuffer.h"
#include "stdbool.h"

// using namespace rapidjson;

// This will iter add oobject to a tree and return tree at the end

void OTPParser::createTDTree(rapidjson::Value &j, Node *parent) {
    int id = 0;
    if (j.IsObject()) {
        for (rapidjson::Value::ConstMemberIterator data = j.MemberBegin(); data != j.MemberEnd(); ++data) {
            rapidjson::Value &jsonvalue = j[data->name.GetString()];
            if (jsonvalue.IsObject() || jsonvalue.IsArray()) {
                Node *datanode = MemoryManager::Inst.CreateNode(++id);
                //            std::cout << (char *)data.key().c_str();
                datanode->SetValue((char *) data->name.GetString());
                datanode->SetLValue((char *) "Placeholder");
                parent->AppendNode(datanode);
                createTDTree(jsonvalue, datanode);
            } else {
                PString pStr = 0;
                MemoryManager::Inst.CreateObject(&pStr);

                rapidjson::StringBuffer buffer;
                rapidjson::Writer<rapidjson::StringBuffer> writer(buffer);
                jsonvalue.Accept(writer);
                pStr->SetValue(buffer.GetString());

                Node *datanode = MemoryManager::Inst.CreateNode(++id);
                std::string val = buffer.GetString();
                //            std::replace(val.begin(), val.end(), '"', '\0');
                val.erase(std::remove(val.begin(), val.end(), '"'), val.end());
                datanode->SetEntityObj((PENTITY) pStr);
                datanode->SetValue((char *) data->name.GetString());
                datanode->SetLValue((char *) val.c_str());
                parent->AppendNode(datanode);
            }
        }
    } else if (j.IsArray()) {
        int iter = 0;
        for (rapidjson::Value::ConstValueIterator data = j.Begin(); data != j.End(); ++data) {
            rapidjson::Value &jsonvalue = (rapidjson::Value &) (*data);
            if (jsonvalue.IsObject() || jsonvalue.IsArray()) {
                Node *datanode = MemoryManager::Inst.CreateNode(++id);
                //            std::cout << (char *)data.key().c_str();
                datanode->SetValue((char *) std::to_string(iter).c_str());
                datanode->SetLValue((char *) "Placeholder");
                parent->AppendNode(datanode);
                createTDTree((rapidjson::Value &) jsonvalue, datanode);
            } else {
                PString pStr = 0;
                MemoryManager::Inst.CreateObject(&pStr);
                std::string val;
                if (jsonvalue.IsNumber()) {
                    pStr->SetValue(std::to_string(jsonvalue.GetDouble()));
                    val = std::to_string(jsonvalue.GetDouble());
                } else {
                    pStr->SetValue(jsonvalue.GetString());
                    val = jsonvalue.GetString();
                }
                Node *datanode = MemoryManager::Inst.CreateNode(++id);
//                std::string val = jsonvalue.GetString();
                //            std::replace(val.begin(), val.end(), '"', '\0');
                val.erase(std::remove(val.begin(), val.end(), '"'), val.end());
                datanode->SetEntityObj((PENTITY) pStr);
                datanode->SetValue((char *) std::to_string(iter).c_str());
                datanode->SetLValue((char *) val.c_str());
                parent->AppendNode(datanode);
            }
            ++iter;
        }
    }
}

Node *OTPParser::OTPJSONToNodeTree(std::string otpsString) {
    int id = 0;
    rapidjson::Document otps;
    otps.Parse(otpsString.c_str());
    Node *root = MemoryManager::Inst.CreateNode(++id);
    int i = 0, j = 0, k = 0;
    int otpCount = 0;

//    for (rapidjson::Value::ConstMemberIterator tp = otps[0].MemberBegin(); tp != otps[0].MemberEnd(); ++tp)
    for (rapidjson::Value::ConstValueIterator itr = otps.Begin(); itr != otps.End(); ++itr) {
        Node *otpNode = MemoryManager::Inst.CreateNode(++id);
        otpCount++;
        otpNode->SetValue("otp");
        otpNode->SetLValue("otp");
        otpNode->SetRValue("otp");
        root->AppendNode(otpNode);
        for (rapidjson::Value::ConstMemberIterator tps = itr->MemberBegin(); tps != itr->MemberEnd(); ++tps) {
            rapidjson::Value &tpsArray = (rapidjson::Value &) (itr->GetObjectA()[tps->name.GetString()]);
            for (rapidjson::Value::ConstValueIterator tp = tpsArray.Begin(); tp != tpsArray.End(); ++tp) {
                rapidjson::Value &tpjson = (rapidjson::Value &) (*tp);
                Node *tpnode = MemoryManager::Inst.CreateNode(++id);
                Node *itemnode = MemoryManager::Inst.CreateNode(++id);
                tpnode->SetLValue((char *) tpjson["id"].GetString());
                tpnode->SetValue((char *) tpjson["stageID"].GetString());
                tpnode->SetRValue((char *) tpjson["tenantID"].GetString());
                if (tpjson.HasMember("identifier")) {
                    if (!tpjson["identifier"].IsNull()) {
                        if (tpjson["identifier"].IsObject()) {
                            if (tpjson["identifier"].HasMember("identifier")) {
                                tpnode->SetCustomString((char *) tpjson["identifier"]["identifier"].GetString());
                            } else {
                                tpnode->SetCustomString("placeholder");
                            }
                        } else {
                            tpnode->SetCustomString("placeholder");
                        }
                    } else {
                        tpnode->SetCustomString("placeholder");
                    }
                } else {
                    tpnode->SetCustomString("placeholder");
                }
//                try {
//                    bool check = tpjson["identifier"].IsNull();
//                    bool val = tpjson["identifier"].HasMember("identifier");
//                    if (val) {
//                        bool val2 = tpjson["identifier"].IsNull();
//                        std::string s = "";
//                    }
//                } catch(int e) {
//                }
                itemnode->SetRValue((char *) tpjson["item"]["itemID"].GetString());
                itemnode->SetLValue((char *) tpjson["item"]["itemName"].GetString());
                tpnode->SetCustomObj(itemnode);
                otpNode->SetCustomObj(itemnode);
                otpNode->AppendNode(tpnode);
                for (rapidjson::Value::ConstValueIterator tdp = tpjson["traceabilityDataPackets"].Begin();
                     tdp != tpjson["traceabilityDataPackets"].End(); ++tdp) {
                    rapidjson::Value &tdpjson = (rapidjson::Value &) (*tdp);
                    Node *tdpnode = MemoryManager::Inst.CreateNode(++id);
                    Node *tdpidnode = MemoryManager::Inst.CreateNode(++id);
                    tdpnode->SetValue((char *) tdpjson["userID"].GetString());
                    tdpidnode->SetValue((char *) tdpjson["id"].GetString());
                    tpnode->AppendNode(tdpnode);
                    tdpnode->AppendNode(tdpidnode);
                    for (rapidjson::Value::ConstValueIterator td = tdpjson["traceabilityData"].Begin();
                         td != tdpjson["traceabilityData"].End(); ++td) {
                        rapidjson::Value &tdjson = (rapidjson::Value &) (*td);
                        Node *tdnode = MemoryManager::Inst.CreateNode(++id);
                        tdpidnode->AppendNode(tdnode);
                        tdnode->SetValue((char *) tdjson["key"].GetString());
                        //                tdnode->SetValue((char *)"something is better");
                        if (tdjson["val"].IsObject() || tdjson["val"].IsArray()) {
                            rapidjson::Value &val = (rapidjson::Value &) tdjson["val"];
                            createTDTree(val, tdnode);
                        } else if (tdjson["val"].IsBool()) {
                            PString pStr = 0;
                            MemoryManager::Inst.CreateObject(&pStr);
                            bool val = tdjson["val"].GetBool();
                            std::string val_bool = (val ? "true" : "false");
                            pStr->SetValue(val ? "true" : "false");
                            tdnode->SetEntityObj((PENTITY) pStr);
                            tdnode->SetLValue((char *) val_bool.c_str());
                        } else if (tdjson["val"].IsInt()) {
                            PString pStr = 0;
                            MemoryManager::Inst.CreateObject(&pStr);
                            int val = tdjson["val"].GetInt();
                            pStr->SetValue(std::to_string(val));
                            tdnode->SetEntityObj((PENTITY) pStr);
                            tdnode->SetLValue((char *) std::to_string(val).c_str());
                        } else if (tdjson["val"].IsFloat()) {
                            PString pStr = 0;
                            MemoryManager::Inst.CreateObject(&pStr);
                            float val = tdjson["val"].GetFloat();
                            pStr->SetValue(std::to_string(val));
                            tdnode->SetEntityObj((PENTITY) pStr);
                            tdnode->SetLValue((char *) std::to_string(val).c_str());
                        } else {
                            PString pStr = 0;
                            MemoryManager::Inst.CreateObject(&pStr);
                            std::string val = tdjson["val"].GetString();
                            //                    std::replace(val.begin(), val.end(), '"', '\0');
                            val.erase(std::remove(val.begin(), val.end(), '"'), val.end());
                            pStr->SetValue(tdjson["val"].GetString());
                            tdnode->SetEntityObj((PENTITY) pStr);
                            tdnode->SetLValue((char *) val.c_str());
                            //                    std::cout << ((PENTITY)tdnode->GetEntityObj())->ul_Type;

                            //                    tdnode->SetValue((char *)tdjson["val"].dump().c_str());
                        }
                        //                std::cout << (char *)std::to_string(tdjson["type"].get<int>()).c_str();
                        //                std::cout << tdjson.HasMember("type");
                        if (tdjson.HasMember("type")) {
                            tdnode->SetRValue((char *) std::to_string(tdjson["type"].GetInt()).c_str());
                        } else {
                            tdnode->SetRValue("90");
                        }
                    }
                }
            }
        }
    }
    return root;
}

Node *OTPParser::TDPJSONToNodeTree(std::string tdpsString) {
    int id = 0;
    rapidjson::Document tdps;
    tdps.Parse(tdpsString.c_str());
    Node *root = MemoryManager::Inst.CreateNode(++id);
    int i = 0, j = 0, k = 0;
    int tdpCount = 0;
    root->SetValue("tdp");
    root->SetLValue("tdp");
    root->SetRValue("tdp");
    for (rapidjson::Value::ConstValueIterator tdpItr = tdps.Begin(); tdpItr != tdps.End(); ++tdpItr) {
        Node *tdpNode = MemoryManager::Inst.CreateNode(++id);
        tdpCount++;
        rapidjson::Value &tdpjson = (rapidjson::Value &) (*tdpItr);
        Node *tdpidnode = MemoryManager::Inst.CreateNode(++id);
        tdpNode->SetLValue((char *) tdpjson["userID"].GetString());
        tdpNode->SetRValue((char *) tdpjson["id"].GetString());
        tdpNode->SetValue((char *) tdpjson["stageID"].GetString());
        Node *tenant = MemoryManager::Inst.CreateNode(++id);
        tenant->SetRValue((char *) tdpjson["tenantID"]["tenantId"].GetString());
        tenant->SetLValue((char *) tdpjson["tenantID"]["name"].GetString());
        tenant->SetValue((char *) tdpjson["tenantID"]["itemName"].GetString());
        Node *tenantCustomObj = MemoryManager::Inst.CreateNode(++id);
        tenantCustomObj->SetValue((char *) tdpjson["tenantID"]["itemID"].GetString());
        tenant->SetCustomObj(tenantCustomObj);
        tenant->SetCustomString((char *) tdpjson["tenantID"]["identifier"].GetString());
        tdpNode->SetCustomObj(tenant);
        root->AppendNode(tdpNode);
        for (rapidjson::Value::ConstValueIterator td = tdpjson["traceabilityData"].Begin();
             td != tdpjson["traceabilityData"].End(); ++td) {
            rapidjson::Value &tdjson = (rapidjson::Value &) (*td);
            Node *tdnode = MemoryManager::Inst.CreateNode(++id);
            tdpNode->AppendNode(tdnode);
            tdnode->SetValue((char *) tdjson["key"].GetString());
            tdnode->SetLValue((char *) "Placeholder");
            if (tdjson["val"].IsObject() || tdjson["val"].IsArray()) {
                rapidjson::Value &val = (rapidjson::Value &) tdjson["val"];
                createTDTree(val, tdnode);
            } else if (tdjson["val"].IsBool()) {
                PString pStr = 0;
                MemoryManager::Inst.CreateObject(&pStr);
                bool val = tdjson["val"].GetBool();
                std::string val_bool = (val ? "true" : "false");
                pStr->SetValue(val ? "true" : "false");
                tdnode->SetEntityObj((PENTITY) pStr);
                tdnode->SetLValue((char *) val_bool.c_str());
            } else if (tdjson["val"].IsInt()) {
                PString pStr = 0;
                MemoryManager::Inst.CreateObject(&pStr);
                int val = tdjson["val"].GetInt();
                pStr->SetValue(std::to_string(val));
                tdnode->SetEntityObj((PENTITY) pStr);
                tdnode->SetLValue((char *) std::to_string(val).c_str());
            } else if (tdjson["val"].IsFloat()) {
                PString pStr = 0;
                MemoryManager::Inst.CreateObject(&pStr);
                float val = tdjson["val"].GetFloat();
                pStr->SetValue(std::to_string(val));
                tdnode->SetEntityObj((PENTITY) pStr);
                tdnode->SetLValue((char *) std::to_string(val).c_str());
            } else {
                PString pStr = 0;
                MemoryManager::Inst.CreateObject(&pStr);
                std::string val = tdjson["val"].GetString();
                //                    std::replace(val.begin(), val.end(), '"', '\0');
                val.erase(std::remove(val.begin(), val.end(), '"'), val.end());
                pStr->SetValue(tdjson["val"].GetString());
                tdnode->SetEntityObj((PENTITY) pStr);
                tdnode->SetLValue((char *) val.c_str());
            }
            if (tdjson.HasMember("type")) {
                tdnode->SetRValue((char *) std::to_string(tdjson["type"].GetInt()).c_str());
            } else {
                tdnode->SetRValue("90");
            }
        }
    }
    return root;
}