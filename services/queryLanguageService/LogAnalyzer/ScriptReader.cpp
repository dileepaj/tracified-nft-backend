#include <regex>
#include "ScriptReader.h"
#include "ExecutionTemplate.h"
#include "ExecutionTemplateList.h"
#include "MetaData.h"
#include "Utils.h"
#include "MemMan.h"
#include "Entity.h"
#include "Command.h"
#include "Value.h"
#include "Int.h"
#include "EntityList.h"
#include "Strings.h"
#include "DateTime.h"
#include "DateTime.h"
#include "BuildErrorHandler.h"


BuildErrorHandler errorHandler;

std::list<std::string> lst;
std::list<std::string> lstIf;
std::list<std::string> lstIfNot;
std::list<std::string> lstVar;
bool ScriptReader::ProcessScript(MSTRING sFile, MetaData* pMD, ScriptReaderOutput& op)
{
    p_MetaData = pMD;
    LST_STR lstLines;
    LST_INT lstLineNumbers;

    // if there's a code library, load all lines from code library first
    // these lines should be prepended to the lines read from the script file
    MSTRING sLoadFromCodeLibrary = pMD->s_LoadFromCodeLibrary;
    Utils::MakeUpper(sLoadFromCodeLibrary);
    if ((sLoadFromCodeLibrary == "TRUE") || (sLoadFromCodeLibrary == "YES")) {
        LST_INT lstLineNumbersDummy;    // line numbers in code library are disregarded
        ReadFileToLines(pMD->s_CodeLibraryFile, pMD->s_LineContinuation, pMD->s_CommentStart, lstLines, lstLineNumbersDummy);
    }

    ReadFileToLines(sFile, pMD->s_LineContinuation, pMD->s_CommentStart, lstLines, lstLineNumbers);
    if(lstLines.empty())
    {
        return false;
    }

    MemoryManager::Inst.CreateObject(&op.p_ETL);
    ExecutionTemplateList* pCurrFunction = 0;
    bool bInsideFunction = false;
    LST_STR::const_iterator ite1 = lstLines.begin();
    LST_STR::const_iterator iteEnd1 = lstLines.end();
    LST_INT::const_iterator ite2 = lstLineNumbers.begin();
    int val = 0;

    for( ; ite1 != iteEnd1; ++ite1, ++ite2)
    {
        val++;
        MSTRING line = *ite1;
        Utils::TrimLeft(line, _MSTR( \t\n));
        Utils::TrimRight(line, _MSTR( \t\n));
        if (line.empty()) {
            continue;
        }
        //######################################  ################################################
        //You can add a line here to check whether there is an error
        //Print the line number and the content of the line aswel


        ScriptReader::ProcessLineRetVal ret = ProcessLine(*ite1, pMD,*ite2,lstLines.size());
//        if(0 == ret.p_ET){
//            return false;
//		  }
        if(0 != ret.p_ET)
        {
            MSTRINGSTREAM sCodeLine;
            sCodeLine<<*ite2<<_MSTR(:)<<SPACE<<*ite1;
            ret.p_ET->SetCodeLine(sCodeLine.str());
        }
        if(SLT_FuncStart == ret.slt)
        {
            MemoryManager::Inst.CreateObject(&pCurrFunction);
            bInsideFunction = true;
            op.map_Functions[ret.s_Str] = pCurrFunction;
        }
        else if(SLT_FuncEnd == ret.slt)
        {
            pCurrFunction = 0;
            bInsideFunction = false;
        }
        else
        {
            if(bInsideFunction)
            {
                pCurrFunction->push_back(ret.p_ET);
            }
            else
            {
                op.p_ETL->push_back(ret.p_ET);
            }
        }
    }

    return true;
}

//######################################  ################################################
std::string ScriptReader::ProcessScript(MetaData* pMD, ScriptReaderOutput& op, MSTRING code)
{
    p_MetaData = pMD;
    LST_STR lstLines;
    LST_INT lstLineNumbers;

    MSTRING sLoadFromCodeLibrary = pMD->s_LoadFromCodeLibrary;
    Utils::MakeUpper(sLoadFromCodeLibrary);
    if ((sLoadFromCodeLibrary == "TRUE") || (sLoadFromCodeLibrary == "YES"))
    {
        LST_INT lstLineNumbersDummy;    // line numbers in code library are disregarded
        ReadFileToLines(pMD->s_CodeLibraryFile, pMD->s_LineContinuation, pMD->s_CommentStart, lstLines, lstLineNumbersDummy);
        std::cout<<"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n";
    }

    ReadStringToLines(code, pMD->s_LineContinuation, pMD->s_CommentStart, lstLines, lstLineNumbers);
    if(lstLines.empty())
    {
        return "empty query";
    }

    MemoryManager::Inst.CreateObject(&op.p_ETL);
    ExecutionTemplateList* pCurrFunction = 0;
    bool bInsideFunction = false;
    LST_STR::const_iterator ite1 = lstLines.begin();
    LST_STR::const_iterator iteEnd1 = lstLines.end();
    LST_INT::const_iterator ite2 = lstLineNumbers.begin();
    int lineNo =0;



     std::string result= errorHandler.IterateVar(lstLines);


    for( ; ite1 != iteEnd1; ++ite1, ++ite2)
    {
         std::string line = std::regex_replace(*ite1, std::regex("^ +"), "");
        ScriptReader::ProcessLineRetVal ret = ProcessLine(line, pMD,*ite2,lstLines.size());

         if(0 != ret.p_ET)
         {
                     MSTRINGSTREAM sCodeLine;
                     sCodeLine<<*ite2<<_MSTR(:)<<SPACE<<*ite1;
                     ret.p_ET->SetCodeLine(sCodeLine.str());
         }else{


                  if(ret.s_Str!="")
                  {
                                  errorHandler.PrintToConsole(*ite2,*ite1+"\t"+ret.s_Str);
                                  return "\nError At line "+ std::to_string(*ite2) +"\t" +*ite1+"\t"+ret.s_Str;
                  }


                              errorHandler.PrintToConsole(*ite2,*ite1);
                              return "\nError At line "+ std::to_string(*ite2) +"\t" +*ite1;
         }




        //try to return false when it captures a syntax error

        if(SLT_FuncStart == ret.slt)
        {
            MemoryManager::Inst.CreateObject(&pCurrFunction);
            bInsideFunction = true;
            op.map_Functions[ret.s_Str] = pCurrFunction;
        }
        else if(SLT_FuncEnd == ret.slt)
        {
            pCurrFunction = 0;
            bInsideFunction = false;
        }
        else
        {
            if(bInsideFunction)
            {
                pCurrFunction->push_back(ret.p_ET);
            }
            else
            {

                op.p_ETL->push_back(ret.p_ET);
            }
        }
    }



    if(result !="")
       {

           return result;
       }


    return "";
}


void ScriptReader::ReadStringToLines(MSTRING code, MSTRING sLineContinuation, MSTRING sCommentStart, LST_STR& lstLines, LST_INT& lstLineNumbers)
{
    MSTRING sCurr = EMPTY_STRING;
    MINT iLineNo = 0;
    std::istringstream iss(code);
    for (MSTRING sLine; std::getline(iss, sLine); ) {
        ++iLineNo;
        Utils::TrimLeft(sLine, _MSTR(\t));
        Utils::TrimRight(sLine, _MSTR(\t));
        if ((sLine.empty()) || (sCommentStart == sLine.substr(0, sCommentStart.length()))) {
            continue;
        }
        sCurr += sLine;
        if ((sCurr.length() >= sLineContinuation.length()) && (sLineContinuation ==
                                                               sCurr.substr(sCurr.length() - sLineContinuation.length(),
                                                                            sLineContinuation.length()))) {
            sCurr = sCurr.substr(0, sCurr.length() - sLineContinuation.length());
        } else {
            lstLines.push_back(sCurr);
            lstLineNumbers.push_back(iLineNo);
            sCurr = EMPTY_STRING;
        }
    }

}

void ScriptReader::ReadFileToLines(MSTRING sFile, MSTRING sLineContinuation, MSTRING sCommentStart, LST_STR& lstLines, LST_INT& lstLineNumbers)
{
    MIFSTREAM file(sFile.c_str());
    MSTRING sLine;
    MSTRING sCurr = EMPTY_STRING;
    if(file.is_open())
    {
        MINT iLineNo = 0;
        while(!file.eof())
        {
            ++iLineNo;
            getline(file, sLine);
            Utils::TrimLeft(sLine, _MSTR( \t));
            Utils::TrimRight(sLine, _MSTR( \t));
            if((sLine.empty()) || (sCommentStart == sLine.substr(0, sCommentStart.length())))
            {
                continue;
            }
            sCurr += sLine;
            if((sCurr.length() >= sLineContinuation.length()) && (sLineContinuation == sCurr.substr(sCurr.length() - sLineContinuation.length(), sLineContinuation.length())))
            {
                sCurr = sCurr.substr(0, sCurr.length() - sLineContinuation.length());
            }
            else
            {
                lstLines.push_back(sCurr);
                lstLineNumbers.push_back(iLineNo);
                sCurr = EMPTY_STRING;
            }
        }
        file.close();
    }
}

ScriptReader::ProcessLineRetVal ScriptReader::ProcessLine(MSTRING sLine, MetaData* pMD,int i,int total)
{
    // First, parse the string with the following as tokens
    // {, }, (, ), ,, =, .
    VEC_CE vecCE;

        if( !errorHandler.CheckVarConv(sLine))
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = "Wrong variable convention";
            return  ret;
        }


        if(!errorHandler.CheckBrackets(sLine))
        {

            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = "Brackets not closed or empty";
            return  ret;
        }

        if(!errorHandler.CheckTags(sLine))
        {

            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = "Tags not closed or empty";
            return  ret;
        }


        if(!errorHandler.ChecklineEnd(sLine))
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str="Line not ended correctly";
            return  ret;
        }



        std::string res =errorHandler.CheckIfCondition(lstIf,sLine);
        if( res !="")
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = "EndIf without IF condition";
            return  ret;
        }


        std::string res2 =errorHandler.CheckIfNotCondition(lstIfNot,sLine);
        if( res2 !="")
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = " IfNot without If condition";
            return  ret;
        }


        if( !errorHandler.CheckCondition(sLine))
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = " does not contain a condition ";
            return  ret;
        }


        if( !errorHandler.CheckEndLine(sLine,i,total))
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = "loop without usage ";
            return  ret;
        }


        if( !errorHandler.ChecklineStart(sLine))
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = "missing $ or wrong  line start";
            return  ret;
        }



        if( !errorHandler.CheckEqual(sLine))
        {
            ScriptReader::ProcessLineRetVal ret;
            ret.s_Str = "missing =";
            return  ret;
        }



    GetCommandElements(sLine, vecCE, pMD);
    // Now this command element list needs to be unified with one of the following
    // 1. Entity
    // 2. Entity=String
    // 3. If(Entity)
    // 4. IfNot(Entity)
    // 5. EndIf
    // 6. While
    // 7. While(Entity)
    // 8. Do
    // 9. Break
    // 10. Continue
    // 11. Function=FuncName
    // 12. EndFunction



     if(errorHandler.CheckForErrors(sLine, vecCE, i, pMD) == 0)
        {


            ScriptReader::ProcessLineRetVal ret;
            return  ret;
        }

    ScriptReader::ProcessLineRetVal ret;



    // case 12
    if((vecCE.size() == 1) && (vecCE.at(0).e_Type == CET_FunctionEnd))
    {
        ret.slt = SLT_FuncEnd;
    }
        // cases 5, 6, 8, 9 & 10
    else if(1 == vecCE.size())
    {
        CommandElementType cet = vecCE.front().e_Type;

        switch(cet)
        {
            case CET_EndIf:
            case CET_While:
            case CET_Do:
            case CET_Break:
            case CET_Continue:
            {
                ExecutionTemplate* pET = 0;
                MemoryManager::Inst.CreateObject(&pET);
                if(CET_EndIf == cet)
                {
                    pET->SetSpecialCommand(COMMAND_TYPE_ENDIF);
                }
                else if(CET_While == cet)
                {
                    pET->SetSpecialCommand(COMMAND_TYPE_WHILE);
                }
                else if(CET_Do == cet)
                {
                    pET->SetSpecialCommand(COMMAND_TYPE_DO);
                }
                else if(CET_Break == cet)
                {
                    pET->SetSpecialCommand(COMMAND_TYPE_BREAK);
                }
                else if(CET_Continue == cet)
                {
                    pET->SetSpecialCommand(COMMAND_TYPE_CONTINUE);
                }
                ret.p_ET = pET;
            }
                break;

        }




    }
        // case 11
    else if((vecCE.size() == 3) && (vecCE.at(0).e_Type == CET_FunctionStart) && (vecCE.at(1).e_Type == CET_EqualSign) && (vecCE.at(2).e_Type ==  CET_String))
    {
        ret.slt = SLT_FuncStart;
        ret.s_Str = vecCE.at(2).s_Str;
    }
        // case 3
    else if((vecCE.size() >= 4) && (vecCE.at(0).e_Type == CET_If) && (vecCE.at(1).e_Type == CET_ArgStart) && (vecCE.at(vecCE.size() - 1).e_Type == CET_ArgEnd))
    {
        ExecutionTemplate* pET = GetEntity(vecCE, 2, vecCE.size() - 2);
        if(0 != pET)
        {
            pET->SetSpecialCommand(COMMAND_TYPE_IF);
        }
        ret.p_ET = pET;
    }
        // case 4
    else if((vecCE.size() >= 4) && (vecCE.at(0).e_Type == CET_IfNot) && (vecCE.at(1).e_Type == CET_ArgStart) && (vecCE.at(vecCE.size() - 1).e_Type == CET_ArgEnd))
    {
        ExecutionTemplate* pET = GetEntity(vecCE, 2, vecCE.size() - 2);
        if(0 != pET)
        {
            pET->SetSpecialCommand(COMMAND_TYPE_IFNOT);
        }
        ret.p_ET = pET;
    }
        // case 7
    else if((vecCE.size() > 3) && (vecCE.at(0).e_Type == CET_While) && (vecCE.at(1).e_Type == CET_ArgStart) && (vecCE.at(vecCE.size() - 1).e_Type == CET_ArgEnd))
    {
        ExecutionTemplate* pET = GetEntity(vecCE, 2, vecCE.size() - 2);
        if(0 != pET)
        {
            pET->SetSpecialCommand(COMMAND_TYPE_WHILE);
        }
        ret.p_ET = pET;
    }
        // case 2
    else if((vecCE.size() >= 3) && (vecCE.at(vecCE.size() - 2).e_Type == CET_EqualSign) && (vecCE.at(vecCE.size() - 1).e_Type == CET_String))
    {
        ExecutionTemplate* pET = GetEntity(vecCE, 0, vecCE.size() - 3);
        if(0 != pET)
        {
            Command* pCmd = 0;
            MemoryManager::Inst.CreateObject(&pCmd);
            pCmd->SetType(COMMAND_TYPE_STORE_AS_VARIABLE);
            PString pString = 0;
            MemoryManager::Inst.CreateObject(&pString);
            pString->SetValue(vecCE.at(vecCE.size() - 1).s_Str);
            pCmd->SetEntityArg(pString);
            pET->AddCommand(pCmd);
        }
        ret.p_ET = pET;
    }
        // case 1
    else
    {


        ExecutionTemplate* pET = GetEntity(vecCE, 0, vecCE.size() - 1);
        ret.p_ET = pET;
//        ExecutionTemplate* pET = 0;
//        ret.p_ET = pET;
        if(pET == 0){
//            return ret;

        }
        //return ret;
    }

    return ret;
}


void ScriptReader::GetCommandElements(MSTRING sCommand, VEC_CE& vecCE, MetaData* pMD)
{
    // First identify strings in the line
    LST_STR lstStringEnclosureSymbol, lstHighLevel;
    LST_INT lstTp;
    if (pMD->s_StringEnclosureSymbol.empty())
    {
        lstHighLevel.push_back(sCommand);
        lstTp.push_back(1);
    }
    else
    {
        lstStringEnclosureSymbol.push_back(pMD->s_StringEnclosureSymbol);
        Utils::TokenizeStringBasic(sCommand, lstStringEnclosureSymbol, lstHighLevel, lstTp);
    }

    LST_STR::const_iterator ite = lstHighLevel.begin();
    LST_STR::const_iterator iteEnd = lstHighLevel.end();
    LST_INT::const_iterator iteTypes = lstTp.begin();
    bool bStringOn = false;
    for ( ; ite != iteEnd; ++ite, ++iteTypes)
    {
        if (1 == *iteTypes)
        {
            if (bStringOn)
            {
                // Current component has to be taken as a string value
                CommandElement ce;
                ce.e_Type = CET_String;
                MSTRING str = *ite;
                Utils::ReplaceSpecialCharacters(str);
                ce.s_Str = str;
                vecCE.push_back(ce);
            } else {
                LST_STR lstTokens, lstRes;
                lstTokens.push_back(pMD->s_FuncSeperator);
                lstTokens.push_back(pMD->s_EqualSign);
                lstTokens.push_back(pMD->s_ArgumentStart);
                lstTokens.push_back(pMD->s_ArgumentEnd);
                lstTokens.push_back(pMD->s_ListStart);
                lstTokens.push_back(pMD->s_ListEnd);
                lstTokens.push_back(pMD->s_ListElementSeperator);
                LST_INT lstTypes;
                MSTRING sCommandPart = *ite;
                Utils::TokenizeStringBasic(sCommandPart, lstTokens, lstRes, lstTypes);
                LST_STR::const_iterator ite1 = lstRes.begin();
                LST_STR::const_iterator iteEnd1 = lstRes.end();
                LST_INT::const_iterator ite2 = lstTypes.begin();
                for( ; ite1 != iteEnd1; ++ite1, ++ite2)
                {
                    CommandElement ce;
                    if(1 == *ite2)
                    {
                        MSTRING sStr = *ite1;
                        if(Utils::IsStringPrefix(sStr, pMD->s_VarNamePrefix))
                        {
                            ce.e_Type = CET_VarName;
                            ce.s_Str = sStr.substr(pMD->s_VarNamePrefix.length(), sStr.length() - pMD->s_VarNamePrefix.length());
                        }
                        else if(Utils::IsStringPrefix(sStr, pMD->s_IntPrefix))
                        {
                            ce.e_Type = CET_Int;
                            ce.s_Str = sStr.substr(pMD->s_IntPrefix.length(), sStr.length() - pMD->s_IntPrefix.length());
                        }
                        else if(sStr == pMD->s_BoolTrue)
                        {
                            ce.e_Type = CET_BoolTrue;
                        }
                        else if(sStr == pMD->s_BoolFalse)
                        {
                            ce.e_Type = CET_BoolFalse;
                        }
                        else if(sStr == pMD->s_If)
                        {
                            ce.e_Type = CET_If;
                        }
                        else if(sStr == pMD->s_IfNot)
                        {
                            ce.e_Type = CET_IfNot;
                        }
                        else if(sStr == pMD->s_EndIf)
                        {
                            ce.e_Type = CET_EndIf;
                        }
                        else if(sStr == pMD->s_While)
                        {
                            ce.e_Type = CET_While;
                        }
                        else if(sStr == pMD->s_Do)
                        {
                            ce.e_Type = CET_Do;
                        }
                        else if(sStr == pMD->s_Break)
                        {
                            ce.e_Type = CET_Break;
                        }
                        else if(sStr == pMD->s_Continue)
                        {
                            ce.e_Type = CET_Continue;
                        }
                        else if(sStr == pMD->s_FuncStart)
                        {
                            ce.e_Type = CET_FunctionStart;
                        }
                        else if(sStr == pMD->s_FuncEnd)
                        {
                            ce.e_Type = CET_FunctionEnd;
                        }
                        else
                        {
                            ce.e_Type = CET_String;
                            Utils::ReplaceSpecialCharacters(sStr);
                            ce.s_Str = sStr;
                        }
                    }
                    else if(2 == *ite2)
                    {
                        if(*ite1 == pMD->s_FuncSeperator)
                        {
                            ce.e_Type = CET_FuncStart;
                        }
                        else if(*ite1 == pMD->s_EqualSign)
                        {
                            ce.e_Type = CET_EqualSign;
                        }
                        else if(*ite1 == pMD->s_ArgumentStart)
                        {
                            ce.e_Type = CET_ArgStart;
                        }
                        else if(*ite1 == pMD->s_ArgumentEnd)
                        {
                            ce.e_Type = CET_ArgEnd;
                        }
                        else if(*ite1 == pMD->s_ListStart)
                        {
                            ce.e_Type = CET_ListStart;
                        }
                        else if(*ite1 == pMD->s_ListEnd)
                        {
                            ce.e_Type = CET_ListEnd;
                        }
                        else if(*ite1 == pMD->s_ListElementSeperator)
                        {
                            ce.e_Type = CET_ListElemSep;
                        }
                    }
                    vecCE.push_back(ce);
                }
            }
        }
        else if(2 == *iteTypes)
        {
            bStringOn = !bStringOn;
        }
    }
}

ExecutionTemplate* ScriptReader::GetEntity(VEC_CE& vecCE, VEC_CE::size_type stStart, VEC_CE::size_type stEnd)
{
    // An entity is in one of the following forms
    // 1. Integer
    // 2. String
    // 3. List
    // 4. Variable.<Any number of functions>
    // 5. Variable
    // 6. FuncName(Entity)
    if(stEnd < stStart)
    {
        return 0;
    }
    ExecutionTemplate* pET = 0;
    // case 1, 2 & 5
    if(stEnd == stStart)
    {
        if(vecCE.at(stStart).e_Type == CET_Int)
        {
            Int* pInt = 0;
            MemoryManager::Inst.CreateObject(&pInt);
            long lVal = _MATOI(vecCE.at(stStart).s_Str.c_str());
            pInt->SetValue((lVal < 0) ? -1 * lVal : lVal);
            pInt->b_IsNegative = (lVal < 0);
            MemoryManager::Inst.CreateObject(&pET);
            pET->SetEntity(pInt);
            return pET;
        }
        if(vecCE.at(stStart).e_Type == CET_String)
        {
            String* pString = 0;
            MemoryManager::Inst.CreateObject(&pString);
            pString->SetValue(vecCE.at(stStart).s_Str);
            MemoryManager::Inst.CreateObject(&pET);
            pET->SetEntity(pString);
            return pET;
        }
        if(vecCE.at(stStart).e_Type == CET_VarName)
        {
            MemoryManager::Inst.CreateObject(&pET);
            pET->SetStartVarName(vecCE.at(stStart).s_Str);
            return pET;
        }
    }
        // case 3
    else if((vecCE.at(stStart).e_Type == CET_ListStart) && (vecCE.at(stEnd).e_Type == CET_ListEnd))
    {
        if(stEnd == (stStart + 1))
        {
            // empty list
            EntityList* pEL = 0;
            MemoryManager::Inst.CreateObject(&pEL);
            MemoryManager::Inst.CreateObject(&pET);
            pET->SetEntity(pEL);
            return pET;
        }
        EntityList* pEL = GetList(vecCE, stStart + 1, stEnd - 1);
        MemoryManager::Inst.CreateObject(&pET);
        pET->SetEntity(pEL);
        return pET;
    }
        // case 4
    else if((vecCE.at(stStart).e_Type == CET_VarName) && (vecCE.at(stStart + 1).e_Type == CET_FuncStart))
    {
        MemoryManager::Inst.CreateObject(&pET);
        pET->SetStartVarName(vecCE.at(stStart).s_Str);
        // Now go through the function list and fill Command objects
        VEC_CE::size_type stPos = stStart + 2;
        std::map<CommandElementType, CommandElementType> mapContextChangeElements;
        mapContextChangeElements[CET_ArgStart] = CET_ArgEnd;
        mapContextChangeElements[CET_ListStart] = CET_ListEnd;
        while(true)
        {
            VEC_CE::size_type stNext;
            GetNextFirstLevelCommandElementPos(vecCE, stPos, stEnd, CET_FuncStart, mapContextChangeElements, stNext);
            if(stNext > stEnd)
            {
                // No more symbols
                Command* pCmd = GetFunction(vecCE, stPos, stEnd);
                pET->AddCommand(pCmd);
                break;
            }
            else
            {
                Command* pCmd = GetFunction(vecCE, stPos, stNext - 1);
                pET->AddCommand(pCmd);
                if(stNext < stEnd)
                {
                    stPos = stNext + 1;
                }
                else
                {
                    // stNext = stEnd
                    break;
                }
            }
        }
        return pET;
    }
        // case 6
    else if((stEnd  - stStart >= 3) && (vecCE.at(stStart).e_Type == CET_String) && (vecCE.at(stStart + 1).e_Type == CET_ArgStart) && (vecCE.at(stEnd).e_Type == CET_ArgEnd))
    {
        pET = GetEntity(vecCE, stStart + 2, stEnd - 1);
        Command* pCmd = GetFunction(vecCE, stStart, stStart);
        pET->AddCommand(pCmd);
        return pET;
    }
    return 0;
}

EntityList* ScriptReader::GetList(VEC_CE& vecCE, VEC_CE::size_type stStart, VEC_CE::size_type stEnd)
{
    EntityList* pEL = 0;
    MemoryManager::Inst.CreateObject(&pEL);
    std::map<CommandElementType, CommandElementType> mapContextChangeElements;
    mapContextChangeElements[CET_ArgStart] = CET_ArgEnd;
    mapContextChangeElements[CET_ListStart] = CET_ListEnd;
    VEC_CE::size_type stPos = stStart;
    while(true)
    {
        VEC_CE::size_type stNext;
        GetNextFirstLevelCommandElementPos(vecCE, stPos, stEnd, CET_ListElemSep, mapContextChangeElements, stNext);
        if(stNext > stEnd)
        {
            // No symbols found
            ExecutionTemplate* pEntity = GetEntity(vecCE, stPos, stEnd);
            pEL->push_back(pEntity);
            break;
        }
        else
        {
            ExecutionTemplate* pEntity = GetEntity(vecCE, stPos, stNext - 1);
            pEL->push_back(pEntity);
            if(stNext < stEnd)
            {
                stPos = stNext + 1;
            }
            else
            {
                //stNext = stEnd
                break;
            }
        }
    }
    return pEL;
}

Command* ScriptReader::GetFunction(VEC_CE& vecCE, VEC_CE::size_type stStart, VEC_CE::size_type stEnd)
{
    // First check whether this is a function taking an argument
    if((vecCE.at(stEnd).e_Type == CET_ArgEnd) && (vecCE.at(stStart + 1).e_Type == CET_ArgStart))
    {
        // Function taking an argument
        ExecutionTemplate* pET = GetEntity(vecCE, stStart + 2, stEnd - 1);
        Command* pCmd = 0;
        MemoryManager::Inst.CreateObject(&pCmd);
        pCmd->SetArg(pET);
        MAP_STR_MULONG::iterator iteFind = p_MetaData->map_FuncNamesReverse.find(vecCE.at(stStart).s_Str);
        if(p_MetaData->map_FuncNamesReverse.end() == iteFind)
        {
            // This is an additional function (user added function)
            pCmd->SetAdditionalFuncName(vecCE.at(stStart).s_Str);
            pCmd->SetType(COMMAND_TYPE_ADDITIONAL_FUNCTION);
        }
        else
        {
            pCmd->SetType((*iteFind).second);
        }
        return pCmd;
    }
    else
    {
        // Function that does not take any argument
        // There should be only one element i.e. stStart should be equal to stEnd
        if(vecCE.at(stStart).e_Type == CET_String)
        {
            Command* pCmd = 0;
            MemoryManager::Inst.CreateObject(&pCmd);
            MAP_STR_MULONG::iterator iteFind = p_MetaData->map_FuncNamesReverse.find(vecCE.at(stStart).s_Str);
            if(p_MetaData->map_FuncNamesReverse.end() == iteFind)
            {
                // This is an additional function (user added function)
                pCmd->SetAdditionalFuncName(vecCE.at(stStart).s_Str);
                pCmd->SetType(COMMAND_TYPE_ADDITIONAL_FUNCTION);
            }
            else
            {
                pCmd->SetType((*iteFind).second);
            }
            return pCmd;
        }
    }
    return 0;
}

void ScriptReader::GetNextFirstLevelCommandElementPos(VEC_CE& vecCE, VEC_CE::size_type stStart, VEC_CE::size_type stEnd, CommandElementType cet, std::map<CommandElementType, CommandElementType>& mapContextChangeElements, VEC_CE::size_type& stElemPos)
{
    // mapContextChangeElements contains first = <context change start element> e.g. (		second = <context change end element> e.g. )
    std::map<CommandElementType, CommandElementType> mapContextChangeElementsRev;
    std::map<CommandElementType, int> mapContextChanges;
    std::map<CommandElementType, CommandElementType>::iterator ite1 = mapContextChangeElements.begin();
    std::map<CommandElementType, CommandElementType>::iterator iteEnd1 = mapContextChangeElements.end();
    for( ; ite1 != iteEnd1; ++ite1)
    {
        mapContextChangeElementsRev[(*ite1).second] = (*ite1).first;
        mapContextChanges[(*ite1).first] = 0;
    }

    int iContextChangeCount = 0;
    VEC_CE::size_type stPos = stStart;
    while(stPos <= stEnd)
    {
        CommandElementType ce = vecCE.at(stPos).e_Type;
        if((0 == iContextChangeCount) && (ce == cet))
        {
            stElemPos = stPos;
            return;
        }
        std::map<CommandElementType, CommandElementType>::iterator iteFind = mapContextChangeElements.find(ce);
        if(iteFind != mapContextChangeElements.end())
        {
            iContextChangeCount++;
            mapContextChanges[ce]++;
        }
        iteFind = mapContextChangeElementsRev.find(ce);
        if(iteFind != mapContextChangeElementsRev.end())
        {
            iContextChangeCount--;
            mapContextChanges[(*iteFind).second]--;
        }
        stPos++;
    }

    stElemPos = stEnd + 1;
}
