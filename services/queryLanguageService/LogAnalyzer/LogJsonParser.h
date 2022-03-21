//
// Created by AfazD on 31-Oct-19.
//

#ifndef FLEXIBLECOMPUTERLANGUAGE_LOGJSONPARSER_H
#define FLEXIBLECOMPUTERLANGUAGE_LOGJSONPARSER_H

#include "CommonIncludes.h"
#include "document.h"
#include "writer.h"
#include "stringbuffer.h"
class LogJsonParser {
    public:
    static Node *LogJSONToNodeTree(std::string otpsString);
    static Node *LOGJSONToNodeTreeRecursively(rapidjson::Value& j,Node* parent);
    static void LogNodeTreetoJson(Node* parent);
    static void PrintNodeToFile(std::ofstream &newjsonfile,Node* node,int count);
};


#endif //FLEXIBLECOMPUTERLANGUAGE_LOGJSONPARSER_H
