//
//  ELParser.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/25/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELParser.h"
#include "Utils.h"
#include "ELLineParser.h"
#include "ELLineParserResult.h"
#include "ELVariableParser.h"
#include "ELLineTemplateParser.h"
#include "ELBlockTemplateParser.h"

ELParser::ELParser() {
    
}

ELParser::~ELParser() {
    
}

bool ELParser::ProcessScript(MSTRING &sFile, MetaData *md, ELParserResult &result) {
    LST_STR lstLines;
	LST_INT lstLineNumbers;
	ReadFileToLines(sFile, md->s_LineContinuation, md->s_CommentStart, lstLines, lstLineNumbers);
    if(lstLines.empty())
	{
		return false;
	}
    
    ELLineParser lp;
    MAP_STR_STR mapVars;
    MAP_STR_STR mapLineTemplates;
    MAP_STR_STR mapBlockTemplates;
    VEC_STR vecLineTemplates;
    VEC_STR vecBlockTemplates;
    
    LST_STR::const_iterator ite1 = lstLines.begin();
	LST_STR::const_iterator iteEnd1 = lstLines.end();
	LST_INT::const_iterator ite2 = lstLineNumbers.begin();
	for( ; ite1 != iteEnd1; ++ite1, ++ite2)
	{
        MSTRING line = *ite1;
        ELLineParserResult *lpr = lp.ParseLine(line, md);
        if (lpr->entityType == AssignmentType_Variable) {
            mapVars[lpr->entityName] = lpr->expression;
        } else if (lpr->entityType == AssignmentType_LineTemplate) {
            vecLineTemplates.push_back(lpr->entityName);
            mapLineTemplates[lpr->entityName] = lpr->expression;
        } else if (lpr->entityType == AssignmentType_BlockTemplate) {
            vecBlockTemplates.push_back(lpr->entityName);
            mapBlockTemplates[lpr->entityName] = lpr->expression;
        }
    }
    
    ELVariableParser vp;
    vp.BuildVariablesFromExpressions(mapVars, md, result.map_Variables);
    
    ELLineTemplateParser ltp;
    ltp.BuildLineTemplatesFromExpressions(vecLineTemplates, mapLineTemplates, result.map_Variables, md, result.vec_LineTemplates);
    
    ELBlockTemplateParser btp;
    btp.BuildBlockTemplatesFromExpressions(vecBlockTemplates, mapBlockTemplates, md, result.vec_LineTemplates, result.vec_BlockTemplates);
    
    return true;
}

void ELParser::ReadFileToLines(MSTRING sFile, MSTRING sLineContinuation, MSTRING sCommentStart, LST_STR& lstLines, LST_INT& lstLineNumbers)
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
