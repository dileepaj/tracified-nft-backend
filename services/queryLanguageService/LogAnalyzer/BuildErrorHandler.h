
#ifndef LOGANALYZER_BUILDERRORHANDLER_H
#define LOGANALYZER_BUILDERRORHANDLER_H
#include "CommonIncludes.h"
#include "ScriptReader.h"
class ExecutionTemplate;
class ExecutionTemplateList;
class MetaData;
class Entity;
class Command;

class BuildErrorHandler:public ScriptReader {


public:
          bool CheckForErrors(MSTRING sCommand, VEC_CE &vecCE, int i, MetaData *pData);
          bool CheckBrackets(std::string expr);
          void CheckUsedVar(MSTRING sLine, std::set<std::string> &set);
          void CheckDeclearedVar(MSTRING sLine, std::set<std::string> &set);
          std::string CheckVar(std::set<std::string> &s1, std::set<std::string> &s2);
          std::string  IterateVar(std::list<std::basic_string<char>> list);
          void PrintToConsole(int lineNumber, std::string line);
          bool ChecklineEnd(std::string expr);
          std::string CheckIfNotCondition(std::list<std::string> &list, std::string sLine);
          std::string CheckIfCondition(std::list<std::string> &list, std::string sLine);
          bool CheckTags(std::string expr);
          bool CheckVarConv(std::string sLine);
          bool CheckCondition(std::string line);
          bool CheckEndLine(std::string line, int lineNumber, int totalLine);
          bool ChecklineStart(std::string expr);
          bool CheckEqual(std::string line);
};


#endif //LOGANALYZER_BUILDERRORHANDLER_H
