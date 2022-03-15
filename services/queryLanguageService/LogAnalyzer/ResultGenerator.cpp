//
//  ResultGenerator.cpp
//  FlexibleComputerLanguage
//
//  Created by Murtaza Anverali on 7/5/18.
//  Copyright © 2018 Dileepa Jayathilaka. All rights reserved.
//

#include <stdio.h>
#include "ResultGenerator.h"
#include "easylogging++.h"

std::string ResultGenerator::TypeFormatting(std::string result, int type)
{
    switch (type)
    {
    case 0:
    {
        std::replace(result.begin(), result.end(), '"', ' ');

        if(result.length()==0){
            result = "\"\"" ;
        }
        return result;
    }
    case 1:
    {
        std::replace(result.begin(), result.end(), '"', ' ');
        return result;
    }
    case 2:
    {
        std::replace(result.begin(), result.end(), '"', ' ');
        return result;
    }
    case 3:
    case 4:
    default:
    {
        if (std::count(result.begin(), result.end(), '"') == 0)
        {

            return "\"" + result + "\"";
        }
        else
        {
            return "" + result + "";
        }
    }
    }
}

std::string ResultGenerator::CreateArray(Node *result, int type)
{
    std::string resultString = "";
    Node *child = result->GetFirstChild();
    while (child != 0)
    {
        type = child->GetRVal() != 0 ? atoi(child->GetRVal()) : type;
        std::string customString = child->GetCustomString();
        std::string value = "";
        if (customString.compare("array") == 0)
        {
            value = CreateArray(child, type);
        }
        else if (customString.compare("object") == 0)
        {
            value = CreateObject(child, type);
        }
        else if (customString.compare("normal") == 0)
        {
            value = TypeFormatting((child->GetValue() != 0) ? child->GetValue() : "", type);
        }
        if (resultString == "")
        {
            resultString += value;
        }
        else
        {
            resultString = resultString + "," + value;
        }
        child = child->GetRightSibling();
    }
    resultString = "[" + resultString + "]";
    return resultString;
}

std::string ResultGenerator::CreateObject(Node *result, int type)
{
    std::string resultString = "";
    Node *child = result->GetFirstChild();
    while (child != 0)
    {
        type = child->GetRVal() != 0 ? atoi(child->GetRVal()) : type;
        std::string customString = child->GetCustomString();
        std::string value = "";
        if (customString.compare("array") == 0)
        {
            value = CreateArray(child, type);
        }
        else if (customString.compare("object") == 0)
        {
            value = CreateObject(child, type);
        }
        else if (customString.compare("normal") == 0)
        {
            value = TypeFormatting((child->GetValue() != 0) ? child->GetValue() : "", type);
        }
        if (resultString == "")
        {
            resultString = child->GetLVal();
            resultString = "\"" + resultString + "\":";
            resultString = resultString + value;
        }
        else
        {
            resultString = resultString + ",\"" + child->GetLVal() + "\":" + value;
        }
        child = child->GetRightSibling();
    }
    resultString = "{" + resultString + "}";
    return resultString;
}

std::string ResultGenerator::ValueFormatting(Node *result)
{
    std::string stringResult = "\"\"";
    std::string typeString = result->GetCustomString();
    if (typeString.compare("array") == 0)
    {
        stringResult = CreateArray(result, atoi(result->GetRVal()));
    }
    else if (typeString.compare("object") == 0)
    {
        stringResult = CreateObject(result, atoi(result->GetRVal()));
    }
    else if (typeString.compare("normal") == 0)
    {
        stringResult = TypeFormatting((result->GetValue() != 0) ? result->GetValue() : "", atoi(result->GetRVal()));
    }
    return stringResult;
}

std::string ResultGenerator::CreateResult(Node *result)
{
    // BOOLEAN IS 0
    // FLOAT IS 1
    // INTEGER IS 2
    // DATETIME IS 3
    // STRING IS 4
    // PHOTO is an array of strings, ARTIFACT is an object, LIST is an array of strings
    // ARRAY OR OBJECT OR NORMAL is stated in CustomerString attr in Node
    std::string resultJSON = "";
    std::string value;
    //    resultJSON = "{\"type\": 0, \"val\": " + valueFormatting(result) + "}";

   if(result->GetRVal()){
       switch (atoi(result->GetRVal()))
       {
           case 0:
               resultJSON = "{\"type\": 0, \"val\": " + ValueFormatting(result) + "}";
               break;
           case 1:
               resultJSON = "{\"type\": 1, \"val\": " + ValueFormatting(result) + "}";
               break;
           case 2:
               resultJSON = "{\"type\": 2, \"val\": " + ValueFormatting(result) + "}";
               break;
           case 3:
               resultJSON = "{\"type\": 3, \"val\": " + ValueFormatting(result) + "}";
               break;
           case 4:
               resultJSON = "{\"type\": 4, \"val\": " + ValueFormatting(result) + "}";
               break;
               //     case 5:
               //         value = result->GetValue();
               //         resultJSON = "{\"type\": 5, \"val\": \"" + valueFormatting(result) + "\"}";
               //         break;
               //     case 6:
               //         // TODO -> HAVE TO KEEP OBJECT WITHIN TREE AND CREATE TREE
               //         value = result->GetValue();
               //         resultJSON = "{\"type\": 6, \"val\": \"" + valueFormatting(result) + "\"}";
               //         break;
               //     case 7:
               //         // TODO -> HAVE TO KEEP LIST WITHIN TREE AND CREATE TREE
               //         value = result->GetValue();
               //         resultJSON = "{\"type\": 7, \"val\":[" + valueFormatting(result) + "]}";
               //         break;
               //     default:
               //         value = result->GetValue();
               //         resultJSON = "{\"type\": 1, \"val\": " + valueFormatting(result) + "}";
       }

   }else{
       resultJSON = "{\"type\": 1, \"val\": \"Incomplete query\"}";
   }

    return resultJSON;
}
