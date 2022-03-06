//
//  OTPParser.h
//  FlexibleComputerLanguage
//
//  Created by Murtaza Anverali on 7/5/18.
//  Copyright © 2018 Dileepa Jayathilaka. All rights reserved.
//

#include "json.hpp"
#include "CommonIncludes.h"
#include "document.h"
#include "writer.h"
#include "stringbuffer.h"

// using namespace rapidjson;

class OTPParser
{
public:
    static void createTDTree(rapidjson::Value& j, Node *parent);
    static Node *OTPJSONToNodeTree(std::string otpsString);
    static Node *TDPJSONToNodeTree(std::string tdpsString);
};