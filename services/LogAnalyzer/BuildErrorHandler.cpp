#include "BuildErrorHandler.h"
#include "ExecutionTemplate.h"
#include "ExecutionTemplateList.h"
#include "MetaData.h"
#include "Command.h"
#include "Int.h"
#include "EntityList.h"
#ifdef _WIN32
#include "windows.h"
#else
#define RESET   "\033[0m"
#define RED     "\033[31m"
#endif
#include "json.hpp"
#include "regex.h"




/**
 *
 * @param sCommand
 * @param vecCE
 * @param i
 * @return  bool
 * Check Basic Syntax Errors
 */

bool BuildErrorHandler::CheckForErrors(MSTRING sCommand, VEC_CE &vecCE, int value, MetaData *pData)
{


    for(int i=0;i<vecCE.size()-1;i++)
    {

           //Check  Function Exists
           if(vecCE[i].e_Type==5 && vecCE[i+1].e_Type==0 )
         {

          MAP_STR_MULONG::iterator iteFind = pData->map_FuncNamesReverse.find(vecCE[i+1].s_Str);
             if((*iteFind).first.size()>20000)
             {

                 return false;
             }

        }


        //Check String exists Immediately  After  Close Parentheses
        if(vecCE[i].e_Type==CET_ArgEnd && vecCE[i+1].e_Type==CET_String )
        {

            return false;

        }



        //Check variable is a function name
        if(vecCE[i].e_Type==CET_EqualSign && vecCE[i + 1].e_Type == CET_String)
        {

            Command* pCmd = 0;
            MemoryManager::Inst.CreateObject(&pCmd);
            MAP_STR_MULONG::iterator iteFind = pData->map_FuncNamesReverse.find(vecCE[i+1].s_Str);

            if((*iteFind).first.size()<20000)
            {
                    return false;
            }


        }



        //Check String exists after "."
        if((vecCE[i].e_Type==CET_FuncStart ) && vecCE[i + 1].e_Type != CET_String)
        {


            return false;

        }




        //check string is a function name
        if(vecCE[i].e_Type==0 )
        {

            Command* pCmd = 0;
            MemoryManager::Inst.CreateObject(&pCmd);
            MAP_STR_MULONG::iterator iteFind = pData->map_FuncNamesReverse.find(vecCE[i].s_Str);

            if((*iteFind).first.size()<20000)
            {
                if(vecCE[i-1].e_Type!=CET_FuncStart)
                {

                    return false;

                }
            }

        }



         //Check After Equal Sign (must be a string)
        if(vecCE[i].e_Type==CET_EqualSign && vecCE[i + 1].e_Type != CET_String)
        {

            return false;

        }



        if(vecCE[i].e_Type==CET_EqualSign && vecCE[i + 1].e_Type == CET_String)
        {


            for(auto i:vecCE[i + 1].s_Str)
            {
                if(!std::isalpha(i))
                {
                    return false;
                }
            }



        }




        if(vecCE[i].e_Type==CET_String && vecCE[i + 1].e_Type ==CET_ArgStart)
        {

            Command* pCmd = 0;
            MemoryManager::Inst.CreateObject(&pCmd);
            MAP_STR_MULONG::iterator iteFind = pData->map_FuncNamesReverse.find(vecCE[i].s_Str);

            if((*iteFind).first.size()>20000)
            {
                return false;
            }

        }



    }


    return true;


}



/**
 *
 * @param expr
 * @return bool
 * Checking Brackets Balanced
 */

bool  BuildErrorHandler::CheckBrackets(std::string expr )
{
    std::vector <char> leftArr;




    for(int i=0; i<expr.length(); i++) {

        if(expr[i] == '(' && expr[i+1] == ')' || expr[i] == '[' && expr[i+1] == ']' || expr[i]=='{' && expr[i+1] == '}')
        {
            return false;
        }

        if(expr[i] == '(' || expr[i] == '[' || expr[i]=='{') {
            leftArr.push_back(expr[i]);
        }

       if(leftArr.size() ==0 && (expr[i] == ')' || expr[i] == '}' || expr[i] == ']')){
            return false;
        }


        int leftArrLength = leftArr.size();

        if(expr[i] == ')' && leftArr[leftArrLength - 1] == '('){
            leftArr.pop_back();
        }else if(expr[i] == '}' && leftArr[leftArrLength - 1] == '{') {
            leftArr.pop_back();
        } else if(expr[i] == ']' && leftArr[leftArrLength - 1] == '[') {
            leftArr.pop_back();
        }

    }


    return leftArr.size() == 0;


}




/**
 *
 * @param sLine
 * @param set
 * Check Created Variables
 */

void BuildErrorHandler::CheckDeclearedVar(MSTRING sLine, std::set<std::string> &set)
{

    std::string delimiter = "=";
    set.insert("RESULT");
    set.insert("X");
    set.insert("Y");
    size_t pos = 0;
    std::string token;
    while ((pos = sLine.find(delimiter)) != std::string::npos) {
        token = sLine.substr(0, pos);
        sLine.erase(0, pos + delimiter.length());

    }

    std::string s = sLine;
    if (s.find('$') == std::string::npos)
    {
        set.insert(s);
    }



}


/**
 *
 * @param sLine
 * @param set
 * Check Used Variables
 */

void BuildErrorHandler::CheckUsedVar(MSTRING sLine, std::set<std::string> &set)
{

    std::string delimiter = ".";
    std::string delimiter2 = "$";



    unsigned first_del = sLine.find(delimiter2);
    unsigned end_first = first_del + delimiter2.length();
    unsigned last_del = sLine.find(delimiter);

    set.insert(sLine.substr(end_first,last_del - end_first));





}

/**
 *
 * @param s1
 * @param s2
 * @return  string
 * Check Variable Declared Before Use
 */
std::string BuildErrorHandler::CheckVar(std::set<std::string> &s1,std::set<std::string>  &s2)
{


    std::set<std::string> result;
    std::set_difference(s2.begin(), s2.end(), s1.begin(), s1.end(),
                        std::inserter(result, result.end()));


    if(!result.empty())
    {

        std::string res="Undefined variable  ";
        for(auto n:result)
        {

            res +="$"+n+" ";

        }
        return res;
    }
    else
    {
        return "";
    }

}




/**
 * Iterator for variable checking
 * @param list
 * @return
 */


std::string BuildErrorHandler::IterateVar(std::list<std::basic_string<char>> list)
{
    LST_STR::const_iterator iteStart = list.begin();
    LST_STR::const_iterator iteEnd = list.end();
    std::set<std::string> s1;
    std::set<std::string> s2;


    for( ; iteStart != iteEnd; ++iteStart)
    {

        CheckDeclearedVar(*iteStart, s1);
        CheckUsedVar(*iteStart, s2);
    }


   if(CheckVar(s1,s2) !="")
   {

      std::string result =CheckVar(s1,s2);


        return result;
   }

      return "";
}




/**
 *
 * @param lineNumber
 * @param line
 * Print error to console
 */
void BuildErrorHandler::PrintToConsole(int lineNumber,std::string line)
{
#ifdef _WIN32
    HANDLE hConsole;
    hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    std::cout << "\nError At line " << lineNumber << "\t" ;
    SetConsoleTextAttribute(hConsole, 192);
    std::cout<<line<< std::endl;
    SetConsoleTextAttribute(hConsole, 15);
#else
    std::cout<<"\nError At line "<<lineNumber<<"\t"<<RED<<line<< RESET << std::endl;
#endif
}










/**
 *  check line end with . or =
 * @param expr
 * @return
 */

bool  BuildErrorHandler::ChecklineEnd(std::string expr)
{

    if(expr.back() =='.' || expr.back() =='=')
    {
        return  false;
    }

    return true;


}


/**
 * Check line start
 * @param expr
 * @return
 */


bool  BuildErrorHandler::ChecklineStart(std::string expr)
{




    if(expr[0] =='$' )
    {
        return  true;
    }
    else if(!std::isalpha(expr[0]))
    {
        return false;
    }

    return true;


}



/**
 * Check EndIf exists without If
 * @param list
 * @return
 */

std::string BuildErrorHandler::CheckIfCondition(std::list<std::string> &list,std::string sLine)
{


    std::string delimiter = "(";
    std::string token = sLine.substr(0, sLine.find(delimiter));

    if(token=="If")
    {
        list.push_back(token);
    }
    if(token=="EndIf" && list.size()!=0)
    {
        list.pop_back();
    }
    else if(token=="EndIf" &&  list.empty())
    {
        return "EndIf without If";
    }


    return "";

}


/**
 *Check IfNot exists without If
 * @param list
 * @return
 */


std::string BuildErrorHandler::CheckIfNotCondition(std::list<std::string> &list,std::string sLine)
{


    std::string delimiter = "(";
    std::string token = sLine.substr(0, sLine.find(delimiter));

    if(token=="If")
    {
        list.push_back(token);
    }
    if(token=="IfNot" && list.size()!=0)
    {
        list.push_back(token);
    }
    else if(token=="IfNot" &&  list.empty())
    {
        return "IfNot without If";
    }

    return "";
}




/**
 * Check <> tags are closed or empty
 * @param expr
 * @return
 */




bool BuildErrorHandler::CheckTags(std::string expr)
{

    std::vector <char> leftArr;
    for(int i=0; i<expr.length(); i++) {

        if(expr[i] == '<' && expr[i+1] == '>')
        {
            return false;
        }

        if(expr[i] == '<' ) {
            leftArr.push_back(expr[i]);
        }

        if(leftArr.size() ==0 && (expr[i] == '>')){
            return false;
        }


        int leftArrLength = leftArr.size();

        if(expr[i] == '>' && leftArr[leftArrLength - 1] == '<'){
            leftArr.pop_back();
        }
    }


    return leftArr.size() == 0;

}



/**
 * Check variable convention
 * @param list
 * @return
 */


bool BuildErrorHandler::CheckVarConv(std::string sLine)
{
    std::string delimiter = "=";
    std::string::size_type n;
    n = sLine.find('=');

    if (n != std::string::npos) {


        size_t pos = 0;
        std::string token;
        while ((pos = sLine.find(delimiter)) != std::string::npos) {
            token = sLine.substr(0, pos);
            sLine.erase(0, pos + delimiter.length());
        }


        if (!isalpha( sLine[0]))
        {
            return false;
        }


        for(int i=0;i<sLine.length();i++)
        {
            if(!std::isalnum(sLine[i]))
            {
                return false;
            }
        }


    }

    return true;

}


/**
 *
 * @param line
 * @param lineNumber
 * @param totalLine
 * @return
 */
bool BuildErrorHandler::CheckEndLine(std::string line,int lineNumber,int totalLine)
{

    if(lineNumber >= totalLine)
    {
        if(line=="Do" || (line[0]=='I'&& line[1]=='f') || (line[0]=='W' && line[1]=='h'))
        {
            return false;
        }
    }


    return true;


}

/**
 *
 * @param line
 * @return
 */


bool BuildErrorHandler::CheckCondition(std::string line)
{
   if(line=="IF" || line=="While")
   {
       return false;
   }
   return true;
}





/**
 *
 * @param line
 * @return
 */

bool  BuildErrorHandler::CheckEqual(std::string line)
{
    std::vector <char> leftArr;
    for(int i=0; i<line.length(); i++) {

        if(line[i] == ':' && line[i+1] != '=' )
        {
            std::cout<<line[i+1];
            return false;
        }


    }


    return true;
}
