#include "Command.h"
#include "Entity.h"
#include "ExecutionContext.h"
#include "ExecutionTemplate.h"
#include "Node.h"
#include "Value.h"
#include "Int.h"
#include "DateTime.h"
#include "EntityList.h"
#include "MemMan.h"
#include "Null.h"
#include "Utils.h"
#include "AdditionalFunctions.h"
#include "Bool.h"
#include "MetaData.h"
#include "DateTimeOperations.h"
#include <set>
#include <algorithm>
#include <functional>
#include <math.h>

Command::Command()
:ul_CommandType(COMMAND_TYPE_INVALID), p_Arg(0), p_EntityArg(0), s_AdditionalFuncName(EMPTY_STRING)
{
    
}

Command::~Command()
{
    
}

void Command::Destroy()
{
	if(0 != p_Arg)
	{
		p_Arg->Destroy();
	}
	if(0 != p_EntityArg)
	{
		p_EntityArg->Destroy();
	}
}

Command* Command::GetCopy()
{
	Command* pCopy = 0;
	MemoryManager::Inst.CreateObject(&pCopy);
	pCopy->SetType(ul_CommandType);
	if(0 != p_Arg)
	{
		pCopy->SetArg(p_Arg->GetCopy());
	}
	if(0 != p_EntityArg)
	{
		if(ENTITY_TYPE_NODE != p_EntityArg->ul_Type)
		{
			pCopy->SetEntityArg(p_EntityArg->GetCopy());
		}
		else
		{
			pCopy->SetEntityArg(p_EntityArg);
		}
	}
	pCopy->SetAdditionalFuncName(s_AdditionalFuncName);
	return pCopy;
}

void Command::SetType(MULONG ulType)
{
	ul_CommandType = ulType;
}

void Command::SetArg(ExecutionTemplate* pArg)
{
	p_Arg = pArg;
}

void Command::SetEntityArg(PENTITY pArg)
{
	p_EntityArg = pArg;
}

void Command::SetAdditionalFuncName(MSTRING sFun)
{
	s_AdditionalFuncName = sFun;
}

MSTRING Command::GetAdditionalFuncName() {
    return s_AdditionalFuncName;
}

PENTITY Command::Execute(PENTITY pEntity, ExecutionContext* pContext)
{
    
	if(COMMAND_TYPE_ADDITIONAL_FUNCTION == ul_CommandType)
	{
		// Additional function can be defined either in the control code or inside the script
		// First check whether it is a function defined in the script
		MAP_STR_EXECTMPLIST::const_iterator iteFind2 = pContext->p_mapFunctions->find(s_AdditionalFuncName);
		if(pContext->p_mapFunctions->end() != iteFind2)
		{
			ExecutionContext ec;
			ec.p_mapFunctions = pContext->p_mapFunctions;
			ec.p_MD = pContext->p_MD;
			ec.map_Var[pContext->p_MD->s_FuncArg] = pEntity;
			((*iteFind2).second)->Execute(&ec);
			MAP_STR_ENTITYPTR::iterator iteFind3 = ec.map_Var.find(pContext->p_MD->s_FuncRet);
			if(ec.map_Var.end() == iteFind3)
			{
				return new String;  // a hack to return a dummy value when the return value is not expected
			}
			return (*iteFind3).second;
		}
		else
		{
			// Now try functions defined in control code
            AdditionalFunc fun = 0;
			MAP_STR_ADDITIONAL_FUNC::const_iterator iteFind = map_AdditionalFunctions.find(s_AdditionalFuncName);
			if(map_AdditionalFunctions.end() == iteFind)
			{
				return 0;
			}
			fun = (*iteFind).second;
			if(0 != p_Arg)
			{
				p_EntityArg = p_Arg->Execute(pContext);
			}
			return fun(p_EntityArg);
		}
	}
	else if(COMMAND_TYPE_STORE_AS_VARIABLE == ul_CommandType)
	{
        // This will change the execution context
		// Get a copy of the entity and add it as a new variable to the context
		// Variable name will be in the p_EntityArg
		PString pVarName = (PString)p_EntityArg;
		if(0 == pVarName)
		{
            return 0;
		}
		PENTITY pVar = pEntity;
        if(ENTITY_TYPE_NODE != pEntity->ul_Type)
		{
            pVar = pEntity->GetCopy();
		}
		MAP_STR_ENTITYPTR::iterator iteFind = pContext->map_Var.find(pVarName->GetValue());
		if(pContext->map_Var.end() != iteFind)
		{
            ((*iteFind).second)->Destroy();
		}
		pContext->map_Var[pVarName->GetValue()] = pVar;
		// Create a Null entity and return
		PNull pRet = 0;
		MemoryManager::Inst.CreateObject(&pRet);
		return pRet;
	}
	else
	{

        if (ENTITY_TYPE_LIST == pEntity->ul_Type) {
			if(0 != p_Arg)
			{
                p_EntityArg = p_Arg->Execute(pContext);
			}
            return ExecuteListCommand(ul_CommandType, pEntity, pContext, p_EntityArg);
        } else if (ENTITY_TYPE_NODE == pEntity->ul_Type) {
            return ExecuteNodeCommand(ul_CommandType, pEntity, pContext);
        } else {
            if(0 != p_Arg)
            {
                p_EntityArg = p_Arg->Execute(pContext);
            }
            return ExecuteEntityCommand(ul_CommandType, pEntity, p_EntityArg);
        }
	}
	return 0;
}

PENTITY Command::ExecuteEntityCommand(MULONG ulCommand, PENTITY pEntity, PENTITY pArg)
{
    // General functions in Entity level
	if(COMMAND_TYPE_IS_NULL == ulCommand)
	{
		PBool pBool = 0;
		MemoryManager::Inst.CreateObject(&pBool);
		pBool->SetValue(pEntity->IsNull());
		return pBool;
	}
    
	if(COMMAND_TYPE_IS_NOT_NULL == ulCommand)
	{
		PBool pBool = 0;
		MemoryManager::Inst.CreateObject(&pBool);
		pBool->SetValue(!pEntity->IsNull());
		return pBool;
	}
    
	switch (pEntity->ul_Type)
	{
        case ENTITY_TYPE_INT:
		{
			return ExecuteIntCommand(ulCommand, pEntity, pArg);
		}
        case ENTITY_TYPE_STRING:
		{
			return ExecuteStringCommand(ulCommand, pEntity, pArg);
		}
        case ENTITY_TYPE_BOOL:
		{
			return ExecuteBoolCommand(ulCommand, pEntity, pArg);
		}
        case ENTITY_TYPE_DATETIME:
        {
            return ExecuteDateTimeCommand(ulCommand, pEntity, pArg);
        }
	}
	return 0;
}

PENTITY Command::ExecuteIntCommand(MULONG ulCommand, PENTITY pEntity, PENTITY pArg)
{
	PInt pInt = (PInt)pEntity;
	if(0 == pInt)
	{
		return 0;
	}
    
	PBool pBoolRes = 0;
	PNull pNullRes = 0;
	PString pStrRes = 0;
    
	switch(ulCommand)
	{
        case COMMAND_TYPE_IS_INT_EQUAL_TO:
		{
			if(ENTITY_TYPE_INT == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PInt pIntArg = (PInt)pArg;
				pBoolRes->SetValue(pInt->GetValue() == pIntArg->GetValue());
				break;
			}
			break;
		}
        case COMMAND_TYPE_IS_INT_MEMBER_OF:
		{
			if(ENTITY_TYPE_LIST == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				pBoolRes->SetValue(false);
				PENTITYLIST pIntListArg = (PENTITYLIST)pArg;
				EntityList::const_iterator ite1 = pIntListArg->begin();
				EntityList::const_iterator iteEnd1 = pIntListArg->end();
				for( ; ite1 != iteEnd1; ++ite1)
				{
					if(((PInt)(*ite1))->GetValue() == pInt->GetValue())
					{
						pBoolRes->SetValue(true);
						break;
					}
				}
			}
			break;
		}
        case COMMAND_TYPE_IS_LESS_THAN:
		{
			if(ENTITY_TYPE_INT == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PInt pIntArg = (PInt)pArg;
				pBoolRes->SetValue(pInt->GetValue() < pIntArg->GetValue());
			}
			break;
		}
        case COMMAND_TYPE_IS_LESS_THAN_OR_EQUAL_TO:
		{
			if(ENTITY_TYPE_INT == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PInt pIntArg = (PInt)pArg;
				pBoolRes->SetValue(pInt->GetValue() <= pIntArg->GetValue());
			}
			break;
		}
        case COMMAND_TYPE_IS_GREATER_THAN:
		{
			if(ENTITY_TYPE_INT == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PInt pIntArg = (PInt)pArg;
				pBoolRes->SetValue(pInt->GetValue() > pIntArg->GetValue());
			}
			break;
		}
        case COMMAND_TYPE_IS_GREATER_THAN_OR_EQUAL_TO:
		{
			if(ENTITY_TYPE_INT == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PInt pIntArg = (PInt)pArg;
				pBoolRes->SetValue(pInt->GetValue() >= pIntArg->GetValue());
			}
			break;
		}
        case COMMAND_TYPE_ADD:
		{
			if(ENTITY_TYPE_INT == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pNullRes);
				PInt pIntArg = (PInt)pArg;
				MULONG ulVal = pInt->GetValue();
				ulVal += (pIntArg->GetValue());
				pInt->SetValue(ulVal);
			}
			break;
		}
        case COMMAND_TYPE_SUBTRACT:
		{
			if(ENTITY_TYPE_INT == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pNullRes);
				PInt pIntArg = (PInt)pArg;
				MULONG ulVal = pInt->GetValue();
				ulVal -= (pIntArg->GetValue());
				pInt->SetValue(ulVal);
			}
			break;
		}
        case COMMAND_TYPE_SET_INTEGER:
        {
            if(ENTITY_TYPE_INT == pArg->ul_Type)
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                PInt pIntArg = (PInt)pArg;
                MULONG ulVal = pIntArg->GetValue();
                pInt->SetValue(ulVal);
            }
            break;
        }
        case COMMAND_TYPE_TOSTRING:
		{
			MSTRINGSTREAM ss;
			ss<<pInt->GetValue();
			MemoryManager::Inst.CreateObject(&pStrRes);
			pStrRes->SetValue(ss.str());
			break;
		}
        case COMMAND_TYPE_PERCENTAGE:
        {
            if(ENTITY_TYPE_INT == pArg->ul_Type)
            {
                MemoryManager::Inst.CreateObject(&pStrRes);
                PInt pIntArg = (PInt)pArg;
                float ulVal = pInt->GetValue();
                ulVal = ulVal / pIntArg->GetValue();
                std::string floatString = std::to_string(roundf(ulVal * 10000 ) / 100);
                pStrRes->SetValue(floatString.substr(0, floatString.find(".") + 3) + "%");
            }
            break;
        }

	}
    
	if(0 != pBoolRes)
	{
		return pBoolRes;
	}
	if(0 != pNullRes)
	{
		return pNullRes;
	}
	if(0 != pStrRes)
	{
		return pStrRes;
	}
	return 0;
}

PENTITY Command::ExecuteBoolCommand(MULONG ulCommand, PENTITY pEntity, PENTITY pArg)
{
	PBool pBool = (PBool)pEntity;
	if(0 == pBool)
	{
		return 0;
	}
    
	PBool pBoolRes = 0;
	PString pStrRes = 0;
    
	switch(ulCommand)
	{
        case COMMAND_TYPE_BOOL_AND:
		{
			PBool pBoolArg = (PBool)pArg;
			if(0 != pBoolArg)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				pBoolRes->SetValue(pBool->And(pBoolArg)->GetValue());
			}
			break;
		}
        case COMMAND_TYPE_BOOL_OR:
		{
			PBool pBoolArg = (PBool)pArg;
			if(0 != pBoolArg)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				pBoolRes->SetValue(pBool->Or(pBoolArg)->GetValue());
			}
			break;
		}
        case COMMAND_TYPE_BOOLTOSTRING:
		{
			MemoryManager::Inst.CreateObject(&pStrRes);
			if (pBool->GetValue()) {
				pStrRes->SetValue("TRUE");
			}
			else {
				pStrRes->SetValue("FALSE");
			}
			break;
		}
		case COMMAND_TYPE_SET_BOOL:
		{
			if(ENTITY_TYPE_BOOL == pArg->ul_Type)
			{
				PBool pBoolArg = (PBool)pArg;
			}
//			MemoryManager::Inst.CreateObject(&pBoolRes);
//			pBoolRes->SetValue(pBoolArg);
//			pBool->SetValue(pBoolArg);
			break;
		}
		case COMMAND_TYPE_TO_FALSE:
		{
			pBool->SetValue(false);
			break;
		}
		case COMMAND_TYPE_TO_TRUE:
		{
			pBool->SetValue(true);
			break;
		}
	}
    
    
	if(0 != pBoolRes)
	{
		return pBoolRes;
	}
	if(0 != pStrRes)
	{
		return pStrRes;
	}
    
	return 0;
}

PENTITY Command::ExecuteStringCommand(MULONG ulCommand, PENTITY pEntity, PENTITY pArg)
{
    PNODE pNode = (PNODE)pEntity;
//    std::cout<<pNode->GetValue()<<"\n";
    if(0 == pNode)
    {
        return 0;
    }
    PString pString = (PString)pEntity;
	if(0 == pString)
	{
		return 0;
	}

    PNODE pNodeRes = 0;
	PInt pIntRes = 0;
	PNull pNullRes = 0;
	PBool pBoolRes = 0;
    PString pStrRes = 0;
    
	switch(ulCommand)
	{
        case COMMAND_TYPE_IS_STRING_EQUAL_TO:
		{
			if(ENTITY_TYPE_STRING == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PString pStrArg = (PString)pArg;
				pBoolRes->SetValue(pString->GetValue() == pStrArg->GetValue());
			}
			break;
		}
        case COMMAND_TYPE_IS_STRING_MEMBER_OF:
		{
			if(ENTITY_TYPE_LIST == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PENTITYLIST pStrListArg = (PENTITYLIST)pArg;
				MSTRING sVal = pString->GetValue();
				pBoolRes->SetValue(false);
				EntityList::const_iterator ite1 = pStrListArg->begin();
				EntityList::const_iterator iteEnd1 = pStrListArg->end();
				for( ; ite1 != iteEnd1; ++ite1)
				{
					if(((PString)(*ite1))->GetValue() == sVal)
					{
						pBoolRes->SetValue(true);
						break;
					}
				}
			}
			break;
		}
        case COMMAND_TYPE_IS_HAVING_SUBSTRING:
		{
			if(ENTITY_TYPE_STRING == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PString pStrArg = (PString)pArg;
				pBoolRes->SetValue(pString->GetValue().find(pStrArg->GetValue()) != MSTRING::npos);
			}
			break;
		}
        case COMMAND_TYPE_IS_HAVING_LEFT_SUBSTRING:
		{
			if(ENTITY_TYPE_STRING == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PString pStrArg = (PString)pArg;
				MSTRING sArg = pStrArg->GetValue();
				pBoolRes->SetValue(pString->GetValue().substr(0, sArg.length()) == sArg);
			}
			break;
		}
        case COMMAND_TYPE_IS_HAVING_RIGHT_SUBSTRING:
		{
			if(ENTITY_TYPE_STRING == pArg->ul_Type)
			{
				MemoryManager::Inst.CreateObject(&pBoolRes);
				PString pStrArg = (PString)pArg;
				MSTRING sArg = pStrArg->GetValue();
				pBoolRes->SetValue(pString->GetValue().substr(pString->GetValue().length() - sArg.length(), sArg.length()) == sArg);
			}
			break;
		}
        case COMMAND_TYPE_ADD_PREFIX:
		{
			MemoryManager::Inst.CreateObject(&pNullRes);
			if(ENTITY_TYPE_STRING == pArg->ul_Type)
			{
				PString pStrArg = (PString)pArg;
				MSTRING sVal = pString->GetValue();
				sVal = pStrArg->GetValue() + sVal;
				pString->SetValue(sVal);
			}
			break;
		}
        case COMMAND_TYPE_ADD_POSTFIX:
		{
			MemoryManager::Inst.CreateObject(&pNullRes);
			if(ENTITY_TYPE_STRING == pArg->ul_Type)
			{
                PString pStrArg = (PString)pArg;
				MSTRING sVal = pString->GetValue();
				sVal += pStrArg->GetValue();
				pString->SetValue(sVal);
			}
			break;
		}
        case COMMAND_TYPE_TRIM_LEFT:
		{
			MemoryManager::Inst.CreateObject(&pNullRes);
			MSTRING sVal = pString->GetValue();
			Utils::TrimLeft(sVal, _MSTR( \t\n));
			pString->SetValue(sVal);
			break;
		}
        case COMMAND_TYPE_TRIM_RIGHT:
		{
			MemoryManager::Inst.CreateObject(&pNullRes);
			MSTRING sVal = pString->GetValue();
			Utils::TrimRight(sVal, _MSTR( \t\n));
			pString->SetValue(sVal);
			break;
		}
        case COMMAND_TYPE_WRITE_TO_FILE:
		{
			MemoryManager::Inst.CreateObject(&pNullRes);
			PString pStrArg = (PString)pArg;
			if(0 != pStrArg)
			{
				MOFSTREAM file;
				file.open(pStrArg->GetValue().c_str(), std::ios::out | std::ios::trunc);
				file<<(pString->GetValue().c_str());
				file.close();
			}
			break;
		}
        case COMMAND_TYPE_GET_LENGTH:
		{
			MemoryManager::Inst.CreateObject(&pIntRes);
			pIntRes->SetValue(pString->GetValue().length());
			break;
		}
        case COMMAND_TYPE_SECONDS_TO_MONTHS:
        {
            MemoryManager::Inst.CreateObject(&pStrRes);
            if (pString->GetValue() != "")
            {
                pStrRes->SetValue(std::to_string(DateTimeOperations::SecondsToMonths(stol(pString->ToString()))));
            }
            break;
        }
        case COMMAND_TYPE_SECONDS_TO_DAYS:
        {
            MemoryManager::Inst.CreateObject(&pStrRes);
            if (pString->GetValue() != "")
            {
                pStrRes->SetValue(std::to_string(DateTimeOperations::SecondsToDays(stol(pString->ToString()))));
            }
            break;
        }
        case COMMAND_TYPE_SECONDS_TO_YEARS:
        {
            MemoryManager::Inst.CreateObject(&pStrRes);
            if (pString->GetValue() != "")
            {
                pStrRes->SetValue(std::to_string(DateTimeOperations::SecondsToYears(stol(pString->ToString()))));
            }
            break;
        }
        case COMMAND_TYPE_GET_DIFFERENCE_BY_STRING:
        {
            MemoryManager::Inst.CreateObject(&pStrRes);
            if (pString->GetValue() != "")
            {
                String* pStrArg = (String*)pArg;
                pStrRes->SetValue(std::to_string(DateTimeOperations::GetDifferenceByString(pString->ToString(), pStrArg->ToString())));
            }
            break;
        }
        case COMMAND_TYPE_STRING_TO_READABLE_DATETIME:
        {
            MemoryManager::Inst.CreateObject(&pStrRes);
            if (pString->GetValue() != "")
            {
                pStrRes->SetValue(DateTimeOperations::StringToReadable(pString->GetValue()));
            }
            break;
        }
		case COMMAND_TYPE_GET_DAY_OF_THE_WEEK_SHORT_STRING:
		{
			MemoryManager::Inst.CreateObject(&pStrRes);
			if (pString->GetValue() != "")
			{
				pStrRes->SetValue(DateTimeOperations::GetDayOfTheWeekShortString(pString->GetValue()));
			}
			break;
		}
		case COMMAND_TYPE_GET_DAY_STRING:
		{
			MemoryManager::Inst.CreateObject(&pStrRes);
			if (pString->GetValue() != "")
			{
				pStrRes->SetValue(DateTimeOperations::GetDayString(pString->GetValue()));
			}
			break;
		}
		case COMMAND_TYPE_GET_MONTH_SHORT_STRING:
		{
			MemoryManager::Inst.CreateObject(&pStrRes);
			if (pString->GetValue() != "")
			{
				pStrRes->SetValue(DateTimeOperations::GetMonthShortString(pString->GetValue()));
			}
			break;
		}
		case COMMAND_TYPE_GET_TIME_24_HOUR_FORMAT:
		{
			MemoryManager::Inst.CreateObject(&pStrRes);
			if (pString->GetValue() != "")
			{
				pStrRes->SetValue(DateTimeOperations::GetTime24HourFormat(pString->GetValue()));
			}
			break;
		}
		case COMMAND_TYPE_GET_YEAR:
		{
			MemoryManager::Inst.CreateObject(&pStrRes);
			if (pString->GetValue() != "")
			{
				pStrRes->SetValue(DateTimeOperations::GetYear(pString->GetValue()));
			}
			break;
		}
        case COMMAND_TYPE_DATE_NOW:
        {
            MemoryManager::Inst.CreateObject(&pStrRes);
            if (pString->GetValue() != "")
            {
                pStrRes->SetValue(std::to_string(DateTimeOperations::GetDateNow()));
            }
            break;
        }
        case COMMAND_TYPE_STRING_TO_UNIX_TIME:
        {
            MemoryManager::Inst.CreateObject(&pStrRes);
            if (pString->GetValue() != "")
            {
                pStrRes->SetValue(std::to_string(DateTimeOperations::StringToUnix(pString->ToString())));
            }
            break;
        }
        case COMMAND_TYPE_STRINGTOINTEGER:
		{
			MemoryManager::Inst.CreateObject(&pIntRes);
			if (pString->GetValue() != "")
			{
				if (pString->GetValue().find_first_not_of("0123456789") == std::string::npos)
				{
					try
					{
						pIntRes->SetValue(std::atoi(pString->GetValue().c_str()));
					}
					catch (...)
					{
						pIntRes->SetValue(0);
					}
				}
				else
				{
					pIntRes->SetValue(0);
				}
			}
			else
			{
				pIntRes->SetValue(0);
			}
			break;
        }
        case COMMAND_TYPE_STRINGTOBOOLEAN:
        {
            MemoryManager::Inst.CreateObject(&pBoolRes);
            if (pString->GetValue() != "")
            {
                std::string val = pString->GetValue();
                if (val.compare("true") == 0)
                {
                    pBoolRes->SetValue(1);
                }
                else
                {
                    pBoolRes->SetValue(0);
                }
            }
            break;
        }
        case COMMAND_TYPE_STRINGTOBOOL:
        {
            MemoryManager::Inst.CreateObject(&pBoolRes);
            PString pStrArg = (PString)pEntity;
            MSTRING sArg = pStrArg->GetValue();
            Utils::MakeLower(sArg);
            pBoolRes->SetValue(sArg == "true");
            break;
        }
		case COMMAND_TYPE_CONVERT_TO_SENTENCE_CASE:
		{
			MemoryManager::Inst.CreateObject(&pStrRes);
			PString pStrArg = (PString)pEntity;
			std::string str = pStrArg->GetValue();
			str[0] = std::toupper(str[0]);
			pStrRes->SetValue(str);
			break;
		}
        case COMMAND_TYPE_ADD_PERIOD:
        {
            MemoryManager::Inst.CreateObject(&pNullRes);
            PString pStrArg = (PString)pArg;
            MSTRING sVal = pString->GetValue();
            sVal += ".";
            pString->SetValue(sVal);
        }
            break;
	}
    if(0 != pNodeRes)
    {
        return pNodeRes;
    }
    
	if(0 != pIntRes)
	{
		return pIntRes;
	}
	if(0 != pNullRes)
	{
		return pNullRes;
	}
	if(0 != pBoolRes)
	{
		return pBoolRes;
    }
    if(0 != pStrRes)
    {
        return pStrRes;
    }
    
	return 0;
}

PENTITY Command::ExecuteDateTimeCommand(MULONG ulCommand, PENTITY pEntity, PENTITY pArg)
{
    PDateTime pDateTime = (PDateTime)pEntity;
    if(0 == pDateTime)
    {
        return 0;
    }
    
    PInt pIntRes = 0;
    PNull pNullRes = 0;
    PBool pBoolRes = 0;
    PString pStrRes = 0;
    PDateTime pDateTimeRes = 0;
    
//    switch(ulCommand)
//    {
//        case COMMAND_TYPE_DATETOSTRING:
//        {
//            MemoryManager::Inst.CreateObject(&pStrRes);
//            if (pDateTime->GetValue() != "")
//            {
//                //TODO
//            }
//        }
//        case COMMAND_TYPE_STRINGTODATE:
//        {
//            MemoryManager::Inst.CreateObject(&pStrRes);
//            if (pDateTime->GetValue() != "")
//            {
//                //TODO
//            }
//        }
//        case COMMAND_TYPE_DATEDIFFERENCE:
//        {
//            MemoryManager::Inst.CreateObject(&pStrRes);
//            if (pDateTime->GetValue() != "")
//            {
//                //TODO
//            }
//        }
//        case COMMAND_TYPE_DATENOW:
//        {
//            MemoryManager::Inst.CreateObject(&pStrRes);
//            if (pDateTime->GetValue() != "")
//            {
//                //TODO
//            }
//        }
//    }
    
    if(0 != pIntRes)
    {
        return pIntRes;
    }
    if(0 != pNullRes)
    {
        return pNullRes;
    }
    if(0 != pBoolRes)
    {
        return pBoolRes;
    }
    if(0 != pStrRes)
    {
        return pStrRes;
    }
    if(0 != pDateTimeRes)
    {
        return pDateTimeRes;
    }
    return 0;
}

PENTITY Command::ExecuteNodeCommand(MULONG ulCommand, PENTITY pEntity, ExecutionContext* pContext)
{
    PNODE pNode = (PNODE)pEntity;
	if(0 == pNode)
	{
		return 0;
	}
    
	PNODE pNodeRes = 0;
	PInt pIntRes = 0;
	PString pStrRes = 0;
	PENTITYLIST pNodeListRes = 0;
	PNull pNullRes = 0;
	PBool pBoolRes = 0;
    PENTITY pEntityRes = 0;
    
    // first handle the commands that would need to access the execution context
    if (COMMAND_TYPE_FILTER_SUBTREE == ulCommand) {
        MemoryManager::Inst.CreateObject(&pNodeListRes);
        FilterSubTree(pNode, p_Arg, pContext, pNodeListRes);

    } else {
        // now handle commands that would not explicitly need the execution context
        // for these command, for the sake of simplicity, we first evaluate the command argument and use it subsequently
        PENTITY pArg = 0;
        if(0 != p_Arg)
		{
			p_EntityArg = p_Arg->Execute(pContext);
            pArg = p_EntityArg;
		}
        
        switch(ulCommand)
        {
            case COMMAND_TYPE_ADD_INNER_OBJ:
            {
                    if(ENTITY_TYPE_STRING == pArg->ul_Type){
                        String* pStrArg = (String*)pArg;
                        if(0!=pStrArg){
                            MemoryManager::Inst.CreateObject(&pNullRes);
                            pNodeRes = pNode->AddNode();
                            pNodeRes->SetCustomString("object");
                            pNodeRes->SetLValue((PMCHAR)pStrArg->GetValue().c_str());
                        }
                    }
                break;
            }
            case COMMAND_TYPE_GET_NODE_OBJ:
            {
//                MemoryManager::Inst.CreateObject(&pNodeRes);
                pNodeRes = MemoryManager::Inst.CreateNode(7777);
                pNodeRes->SetValue("");
                pNodeRes->SetLValue("");
                pNodeRes->SetRValue("");
                break;
            }
            case COMMAND_TYPE_LEFT_SIBLING:
            {
                pNodeRes = pNode->GetLeftSibling();
                break;
            }
            case COMMAND_TYPE_RIGHT_SIBLING:
            {
                pNodeRes = pNode->GetRightSibling();
                break;
            }
            case COMMAND_TYPE_PARENT:
            {
                pNodeRes = pNode->GetParent();
                break;
            }
            case COMMAND_TYPE_FIRST_CHILD:
            {
                pNodeRes = pNode->GetFirstChild();
                break;
            }
            case COMMAND_TYPE_CHILDREN:
            {
                MemoryManager::Inst.CreateObject(&pNodeListRes);
                PNODE pChild = pNode->GetFirstChild();
                while(0 != pChild)
                {
                    pNodeListRes->push_back((PENTITY)pChild);
                    pChild = pChild->GetRightSibling();
                }
                break;
            }
            case COMMAND_TYPE_CHILD_COUNT:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(pNode->GetChildCount());
                break;
            }
            case COMMAND_TYPE_GET_VALUE:
            {
                MemoryManager::Inst.CreateObject(&pStrRes);
                pStrRes->SetValue(MSTRING(pNode->GetValue()));
                break;
            }
            case COMMAND_TYPE_GET_LVALUE:
            {
                MemoryManager::Inst.CreateObject(&pStrRes);
                pStrRes->SetValue(MSTRING(pNode->GetLVal()));
                break;
            }
            case COMMAND_TYPE_GET_RVALUE:
            {
                MemoryManager::Inst.CreateObject(&pStrRes);
                pStrRes->SetValue(MSTRING(pNode->GetRVal()));
                break;
            }
            case COMMAND_TYPE_GET_CUSTOM_STRING:
            {
                MemoryManager::Inst.CreateObject(&pStrRes);
                pStrRes->SetValue(MSTRING(pNode->GetCustomString()));
                break;
            }
            case COMMAND_TYPE_GET_ID:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(pNode->GetID());
                break;
            }
            case COMMAND_TYPE_GET_TYPE:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(pNode->GetType());
                break;
            }
            case COMMAND_TYPE_GET_NATURE:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(pNode->GetNature());
                break;
            }
            case COMMAND_TYPE_GET_WEIGHT:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(pNode->GetWeight());
                break;
            }
            case COMMAND_TYPE_GET_MIN_CHILD_WEIGHT:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(pNode->GetMinimumChildWeight());
                break;
            }
            case COMMAND_TYPE_GET_MAX_CHILD_WEIGHT:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(pNode->GetMaximumChildWeight());
                break;
            }
            case COMMAND_TYPE_SET_VALUE:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    String* pStrArg = (String*)pArg;
                    if(0 != pStrArg)
                    {
                        pNode->SetValue((PMCHAR)pStrArg->GetValue().c_str());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_ATTRIBUTES:
            {
                if(ENTITY_TYPE_LIST==pArg->ul_Type){
                    PENTITYLIST pStrListArg = (PENTITYLIST)pArg;
                    if(0!=pStrListArg){
                        LST_STR lstTokens;
                        EntityList::const_iterator ite1 = pStrListArg->begin();
                        EntityList::const_iterator iteEnd1 = pStrListArg->end();
                         for( ; ite1 != iteEnd1; ++ite1){
                            MSTRING val=((PNODE)(*ite1))->GetValue();
                            if(val=="normal"){
                                pNode->SetCustomString("normal");
                            }
                            if(val=="array"){
                                pNode->SetCustomString("array");
                            }
                            if(val=="object"){
                                pNode->SetCustomString("object");
                            }
                            if(val=="1"){
                                pNode->SetRValue("1");
                                pNode->SetLValue("1");
                            }
                            if(val=="2"){
                                pNode->SetRValue("2");
                                pNode->SetLValue("2");
                            }
                            if(val=="3"){
                                pNode->SetRValue("3");
                                pNode->SetLValue("3");
                            }
                            if(val=="4"){
                                pNode->SetRValue("4");
                                pNode->SetLValue("4");
                            }
                        }
                    }
                }
                break;
            }

            case COMMAND_TYPE_SET_LVALUE:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    String* pStrArg = (String*)pArg;
                    if(0 != pStrArg)
                    {
                        pNode->SetLValue((PMCHAR)pStrArg->GetValue().c_str());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_RVALUE:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    String* pStrArg = (String*)pArg;
                    if(0 != pStrArg)
                    {
                        pNode->SetRValue((PMCHAR)pStrArg->GetValue().c_str());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_TYPE:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    Int* pIntArg = (Int*)pArg;
                    if(0 != pIntArg)
                    {
                        pNode->SetType((MUSHORT)pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_NATURE:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    Int* pIntArg = (Int*)pArg;
                    if(0 != pIntArg)
                    {
                        pNode->SetNature((MBYTE)pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_CUSTOM_STRING:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    String* pStrArg = (String*)pArg;
                    if(0 != pStrArg)
                    {
                        pNode->SetCustomString((PMCHAR)pStrArg->GetValue().c_str());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_MIN_CHILD_WEIGHT:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    Int* pIntArg = (Int*)pArg;
                    if(0 != pIntArg)
                    {
                        pNode->SetMinimumChildWeight(pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_MAX_CHILD_WEIGHT:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    Int* pIntArg = (Int*)pArg;
                    if(0 != pIntArg)
                    {
                        pNode->SetMaximumChildWeight(pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_SET_WEIGHT:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    Int* pIntArg = (Int*)pArg;
                    if(0 != pIntArg)
                    {
                        pNode->SetWeight(pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_EXPAND:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_LIST == pArg->ul_Type)
                {
                    PENTITYLIST pStrListArg = (PENTITYLIST)pArg;
                    if(0 != pStrListArg)
                    {
                        LST_STR lstTokens;
                        EntityList::const_iterator ite1 = pStrListArg->begin();
                        EntityList::const_iterator iteEnd1 = pStrListArg->end();
                        for( ; ite1 != iteEnd1; ++ite1)
                        {
                            lstTokens.push_back(((PString)(*ite1))->GetValue());
                        }
                        pNode->Expand(lstTokens);
                    }
                }
                break;
            }
            case COMMAND_TYPE_ADD_NODE:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(0 == pArg)
                {
                    pNodeRes = pNode->AddNode();
                }
                else if(ENTITY_TYPE_NODE == pArg->ul_Type)
                {
                    pNodeRes = pNode->AddNode((PNODE)pArg, true);
                }
                break;
            }
            case COMMAND_TYPE_ADD_NODE_WITH_WEIGHT:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    Int* pIntArg = (Int*)pArg;
                    if(0 != pIntArg)
                    {
                        pNodeRes = pNode->AddNodeWithWeight(pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_READ_FROM_FILE:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    PString pStrArg = (PString)pArg;
                    if(0 != pStrArg)
                    {
                        pNode->ReadValueFromFile(pStrArg->GetValue().c_str());
                    }
                }
                break;
            }
            case COMMAND_TYPE_GET_AGGREGATED_VALUE:
            {
                MemoryManager::Inst.CreateObject(&pStrRes);
                pStrRes->SetValue(pNode->GetAggregatedValue());
                break;
            }
            case COMMAND_TYPE_GET_ENTITY_OBJECT:
            {
                pEntityRes = (pNode->GetEntityObj());
                break;
            }
            case COMMAND_TYPE_SET_ENTITY_OBJECT:
            {
                if(ENTITY_TYPE == pArg->ul_Type)
                {
                    Entity* pEntityArg = (Entity*)pArg;
                    if(0 != pEntityArg)
                    {
                        pNode->SetEntityObj(pEntityArg);
                    }
                }
                break;
            }
            case COMMAND_TYPE_GET_SUBTREE:
            {
                MemoryManager::Inst.CreateObject(&pNodeListRes);
                AddSubtreeToNodeList(pNodeListRes, pNode);
                break;
            }
            case COMMAND_TYPE_IS_TYPE:
            {
                MemoryManager::Inst.CreateObject(&pBoolRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    PInt pIntArg = (PInt)pArg;
                    if(0 != pIntArg)
                    {
                        pBoolRes->SetValue(pNode->GetType() == pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_IS_VALUE:
            {
                MemoryManager::Inst.CreateObject(&pBoolRes);
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    PString pStrArg = (PString)pArg;
                    if(0 != pStrArg)
                    {
                        pBoolRes->SetValue(MSTRING(pNode->GetValue()) == pStrArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_GET_CHILD_OF_TYPE:
            {
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    PInt pIntArg = (PInt)pArg;
                    PNODE pChild = pNode->GetFirstChild();
                    while(0 != pChild)
                    {
                        if(pChild->GetType() == pIntArg->GetValue())
                        {
                            pNodeRes = pChild;
                            break;
                        }
                        pChild = pChild->GetRightSibling();
                    }
                }
                break;
            }
//                pNodeRes = pNode->GetRightSibling();
//                break;
            case COMMAND_TYPE_LAST_CHILD:
            {
                pNodeRes = pNode->GetLastChild();
                break;
            }
            case COMMAND_TYPE_IS_HAVING_CUSTOM_STRING:
            {
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    PString pStringArg = (PString)pArg;
                    pNodeRes = pNode->IsHavingCustomString(pStringArg->GetValue());
                }
                break;
            }
            case COMMAND_TYPE_GET_CHILD_NODE_BY_CUSTOM_STRING:
            {
                if(ENTITY_TYPE_STRING == pArg->ul_Type)
                {
                    PString pStringArg = (PString)pArg;
                    pNodeRes = pNode->GetChildNodeByCustomString(pStringArg->GetValue());
                }
                break;
            }
            case COMMAND_TYPE_SET_ID:
            {
                MemoryManager::Inst.CreateObject(&pNullRes);
                if(ENTITY_TYPE_INT == pArg->ul_Type)
                {
                    PInt pIntArg = (PInt)pArg;
                    if(0 != pIntArg)
                    {
                        pNode->SetID(pIntArg->GetValue());
                    }
                }
                break;
            }
            case COMMAND_TYPE_CHECK_NOT_NULL:
            {
                MemoryManager::Inst.CreateObject(&pBoolRes);
                pBoolRes->SetValue(!pNode->IsNull());
                break;
            }
            case COMMAND_TYPE_GET_STRING:
            {
                MemoryManager::Inst.CreateObject(&pStrRes);
                pStrRes->SetValue("");
                break;
            }
            case COMMAND_TYPE_GET_INTEGER:
            {
                MemoryManager::Inst.CreateObject(&pIntRes);
                pIntRes->SetValue(0);
                break;
            }
            case COMMAND_TYPE_GET_BOOLEAN:
            {
                MemoryManager::Inst.CreateObject(&pBoolRes);
				pBoolRes->SetValue(true);
                break;
            }
			case COMMAND_TYPE_GET_COMMA:
			{
				MemoryManager::Inst.CreateObject(&pStrRes);
				PString pStr = (PString) ",";
				pStrRes->SetValue(",");
				break;
			}
            case COMMAND_TYPE_NEXT_SIBLING:
            {
				pNodeRes = pNode->GetRightSibling();
				if (pNodeRes == 0)
				{
					MemoryManager::Inst.CreateObject(&pNullRes);
					pNodeRes == 0;
				}
                break;
            }
            case COMMAND_TYPE_GET_CUSTOM_OBJ:
            {
                pNodeRes = (PNODE)pNode->GetCustomObj();
                break;
            }
        }
    }
    
	
	
	if(0 != pNodeRes)
	{
		return pNodeRes;
	}
    if(0 != pEntityRes)
    {
        return pEntityRes;
    }
	if(0 != pNodeListRes)
	{
		return pNodeListRes;
	}
	if(0 != pStrRes)
	{
		return pStrRes;
	}
	if(0 != pIntRes)
	{
		return pIntRes;
	}
	if(0 != pNullRes)
	{
		return pNullRes;
	}
	if(0 != pBoolRes)
	{
		return pBoolRes;
	}
	return 0;
}

PENTITY Command::ExecuteListCommand(MULONG ulCommand, PENTITY pEntity, ExecutionContext* pContext, PENTITY pArg)
{
	PENTITYLIST pEntityList = (PENTITYLIST)pEntity;
	if(0 == pEntityList)
	{
		return 0;
	}
    
	PInt pIntRes = 0;
	PENTITYLIST pListRes = 0;
	PNull pNullRes = 0;
	PENTITY pEntityRes = 0;
	PNODE pNodeRes = 0;
	PString pStrRes = 0;
    
	if(COMMAND_TYPE_GET_ITEM_COUNT == ulCommand)
	{
		MemoryManager::Inst.CreateObject(&pIntRes);
		pIntRes->SetValue(pEntityList->size());
	}
    else if(COMMAND_TYPE_GET_INNER_ITEM_COUNT == ulCommand)
    {
        pListRes = pEntityList->GetInnerCount();
        //        MemoryManager::Inst.CreateObject(&pListRes);
        //        pEntityList->SeekToBegin();
        //        PENTITY ent = pEntityList->GetCurrElem();
        //        while (ent) {
        //            PInt count = 0;
        //            MemoryManager::Inst.CreateObject(&count);
        //            count->SetValue(((PENTITYLIST)ent)->size());
        //            pListRes->push_back(count);
        //
        //            pEntityList->Seek(1, false);
        //            ent = pEntityList->GetCurrElem();
        //        }
    }
    else if(COMMAND_TYPE_LIST_FILTER == ulCommand)
    {
        MemoryManager::Inst.CreateObject(&pListRes);
		EntityList::const_iterator ite1 = pEntityList->begin();
		EntityList::const_iterator iteEnd1 = pEntityList->end();
		for( ; ite1 != iteEnd1; ++ite1)
		{
            pContext->map_Var[pContext->p_MD->s_ListItemVar] = *ite1;
            if(0 != p_Arg)
            {
                PBool pRes = (PBool)p_Arg->Execute(pContext);
                if (pRes->GetValue()) {
                    pListRes->push_back(*ite1);
                }
            }
        }
    }
    else if(COMMAND_TYPE_LIST_GROUPBY == ulCommand)
    {
        MemoryManager::Inst.CreateObject(&pListRes);
        std::map<MSTRING, PENTITYLIST> groupedLists;
		EntityList::const_iterator ite1 = pEntityList->begin();
		EntityList::const_iterator iteEnd1 = pEntityList->end();
		for( ; ite1 != iteEnd1; ++ite1)
		{
            pContext->map_Var[pContext->p_MD->s_ListItemVar] = *ite1;
            if(0 != p_Arg)
            {
                PENTITY pRes = p_Arg->Execute(pContext);
                MSTRING key = pRes->ToString();
                std::map<MSTRING, PENTITYLIST>::iterator ite = groupedLists.find(key);
                if (ite == groupedLists.end()) {
                    PENTITYLIST newlist = 0;
                    MemoryManager::Inst.CreateObject(&newlist);
                    newlist->push_back(*ite1);
                    groupedLists[key] = newlist;
                } else {
                    (*ite).second->push_back(*ite1);
                }
            }
        }
        
        std::map<MSTRING, PENTITYLIST>::iterator ite2 = groupedLists.begin();
        std::map<MSTRING, PENTITYLIST>::iterator end2 = groupedLists.end();
        for ( ; ite2 != end2; ++ite2) {
            pListRes->push_back(ite2->second);
        }
    }
    else if(COMMAND_TYPE_GET_UNIQUE_NODE_LIST_WITH_COUNT == ulCommand)
    {
        // ONLY FOR NODE LIST
        MemoryManager::Inst.CreateObject(&pListRes);
        pEntityList->SeekToBegin();
        PNODE currNode = (PNODE)pEntityList->GetCurrElem();
        std::map<std::string, int> uniqueMap;
        String* pStrArg = (String*)pArg;
        while(currNode != 0) {
            std::string str;
            if (pStrArg != 0 && std::strcmp((char *)pStrArg, "LValue") && currNode->GetLVal() != 0) {
                str.assign(currNode->GetLVal());
            } else if (pStrArg != 0 && std::strcmp((char *)pStrArg, "RValue") && currNode->GetRVal() != 0) {
                str.assign(currNode->GetRVal());
            } else {
                str.assign(currNode->GetValue());
            }
            if (uniqueMap[str] == 0)
            {
                uniqueMap[str] = 1;
            }
            else
            {
                uniqueMap[str] = uniqueMap[str] + 1;
            }
            pEntityList->Seek(1, false);
            currNode = (PNODE)pEntityList->GetCurrElem();
        }

        for (auto const& x : uniqueMap)
        {
            PNODE item = MemoryManager::Inst.CreateNode(999);
            item->SetValue((char *)x.first.c_str());
            item->SetLValue((char *)std::to_string(x.second).c_str());
            pListRes->push_back(item);
        }
    }
        else if(COMMAND_TYPE_GET_UNIQUE_NODE_LIST_WITH_NODE_REF == ulCommand)
    {
        // ONLY FOR NODE LIST
        MemoryManager::Inst.CreateObject(&pListRes);
        pEntityList->SeekToBegin();
        PNODE currNode = (PNODE)pEntityList->GetCurrElem();
        std::map<std::string, PNODE> uniqueMap;
        String* pStrArg = (String*)pArg;
        while(currNode != 0) {
            std::string str;
            if (pStrArg != 0 && std::strcmp((char *)pStrArg, "LValue") && currNode->GetLVal() != 0) {
                str.assign(currNode->GetLVal());
            } else if (pStrArg != 0 && std::strcmp((char *)pStrArg, "RValue") && currNode->GetRVal() != 0) {
                str.assign(currNode->GetRVal());
            } else {
                str.assign(currNode->GetValue());
            }
            if (uniqueMap[str] == 0)
            {
                uniqueMap[str] = currNode;
            }
//            else
//            {
//                uniqueMap[str] = uniqueMap[str] + 1;
//            }
            pEntityList->Seek(1, false);
            currNode = (PNODE)pEntityList->GetCurrElem();
        }

        for (auto const& x : uniqueMap)
        {
            PNODE item = MemoryManager::Inst.CreateNode(999);
            item->SetValue((char *)x.first.c_str());
            item->SetCustomObj((x.second));
            pListRes->push_back(item);
        }
    }
    else if(COMMAND_TYPE_SORT_NODE_LIST == ulCommand)
    {
        // ONLY FOR NODE LIST
        MemoryManager::Inst.CreateObject(&pListRes);
		String* pStrArg = (String*)pArg;
        pEntityList->SeekToBegin();
        PNODE currNode = (PNODE)pEntityList->GetCurrElem();
        std::map<std::string, int> uniqueMap;
        while(currNode != 0)
        {
            std::string str;
            str.assign(currNode->GetValue());
            uniqueMap[str] =  std::stoi(currNode->GetLVal());
            pEntityList->Seek(1, false);
            currNode = (PNODE)pEntityList->GetCurrElem();
        }

//        for (auto const& x : uniqueMap)
//        {
//            std::cout << x.first  // string (key)
//                      << ':'
//                      << x.second // string's value
//                      << std::endl ;
//        }

        // Declaring the type of Predicate that accepts 2 pairs and return a bool
        typedef std::function<bool(std::pair<std::string, int>, std::pair<std::string, int>)> Comparator;

        // Defining a lambda function to compare two pairs. It will compare two pairs using second field
        Comparator compFunctor;
		if (pStrArg != 0)
		{
			int pInt = atoi(pStrArg->GetValue().c_str());
			if (pInt >= 0)
			{
				compFunctor = [](std::pair<std::string, int> elem1 ,std::pair<std::string, int> elem2)
				{
					return elem1.second >= elem2.second;
				};
			}
			else if (pInt <= 0)
			{
				compFunctor = [](std::pair<std::string, int> elem1 ,std::pair<std::string, int> elem2)
				{
					return elem1.second <= elem2.second;
				};
			}
		} else {
			compFunctor = [](std::pair<std::string, int> elem1 ,std::pair<std::string, int> elem2)
					{
						return elem1.second >= elem2.second;
					};
		}
        // Declaring a set that will store the pairs using above comparision logic
        std::set<std::pair<std::string, int>, Comparator> setOfSorted(uniqueMap.begin(), uniqueMap.end(), compFunctor);
        for (auto const& x : setOfSorted)
        {
            PNODE item = MemoryManager::Inst.CreateNode(999);
            item->SetValue((char *)x.first.c_str());
            item->SetLValue((char *)std::to_string(x.second).c_str());
            pListRes->push_back(item);
        }
    }
    else if(COMMAND_TYPE_EXTRACT_NODE_LIST_TOP == ulCommand)
    {
        // ONLY FOR NODE LIST
        MemoryManager::Inst.CreateObject(&pListRes);
		String* pStrArg = (String*)pArg;
		int pInt = atoi(pStrArg->GetValue().c_str());
        pEntityList->SeekToBegin();
        PNODE currNode = (PNODE)pEntityList->GetCurrElem();
        int entitySize = pEntityList->size();
        for(int i = 0; i < (entitySize < pInt ? entitySize : pInt); i++)
        {
            pListRes->push_back(currNode->GetCopy());
            pEntityList->Seek(1, false);
            currNode = (PNODE)pEntityList->GetCurrElem();
        }
    }
    else if(COMMAND_TYPE_LIST_GROUP_SEQUENCE_BY == ulCommand)
    {
        bool firstKeyDetected = false;
        MSTRING currentkey;
        PENTITYLIST currentlist = 0;
        EntityList::const_iterator ite1 = pEntityList->begin();
		EntityList::const_iterator iteEnd1 = pEntityList->end();
		for( ; ite1 != iteEnd1; ++ite1)
		{
            pContext->map_Var[pContext->p_MD->s_ListItemVar] = *ite1;
            if(0 != p_Arg)
            {
                PENTITY pRes = p_Arg->Execute(pContext);
                MSTRING key = pRes->ToString();
                if (firstKeyDetected && (key == currentkey)) {
                    currentlist->push_back(*ite1);
                }
                else {
                    PENTITYLIST newlist = 0;
                    MemoryManager::Inst.CreateObject(&newlist);
                    newlist->push_back(*ite1);
                    currentlist = newlist;
                    pListRes->push_back(newlist);
                }
                currentkey = key;
                firstKeyDetected = true;
            }
        }
    }
	else if(COMMAND_TYPE_SEEK == ulCommand)
	{
		PInt pInt = (PInt)p_Arg->GetEntity();
		if(0 != pInt)
		{
			pEntityList->Seek(pInt->GetValue(), pInt->b_IsNegative);
		}
		MemoryManager::Inst.CreateObject(&pNullRes);
	}
	else if(COMMAND_TYPE_SEEK_TO_BEGIN == ulCommand)
	{
		pEntityList->SeekToBegin();
		MemoryManager::Inst.CreateObject(&pNullRes);
	}
	else if(COMMAND_TYPE_SEEK_TO_END == ulCommand)
	{
		pEntityList->SeekToEnd();
		MemoryManager::Inst.CreateObject(&pNullRes);
	}
	else if(COMMAND_TYPE_GET_CURR_ELEM == ulCommand)
	{
		PENTITY pEntity = pEntityList->GetCurrElem();
		if(0 != pEntity)
		{
			// If the entity is not a node then it will be deleted immediately after use.
			// Therefore we need to get a copy.
			if(ENTITY_TYPE_NODE != pEntity->ul_Type)
			{
				pEntityRes = pEntity->GetCopy();
			}
			else
			{
				pEntityRes = pEntity;
			}
		}
		else
		{
			MemoryManager::Inst.CreateObject(&pNullRes);
		}
	}
    else if (COMMAND_TYPE_GET_NEXT_ELEM == ulCommand)
    {
        pEntityList->Seek(1, false);
        PENTITY pEntity = pEntityList->GetCurrElem();
        if(0 != pEntity)
        {
            // If the entity is not a node then it will be deleted immediately after use.
            // Therefore we need to get a copy.
            if(ENTITY_TYPE_NODE != pEntity->ul_Type)
            {
                pEntityRes = pEntity->GetCopy();
            }
            else
            {
                pEntityRes = pEntity;
            }
        }
        else
        {
            MemoryManager::Inst.CreateObject(&pNullRes);
        }
    }
	else if (COMMAND_TYPE_GET_OLDEST_DATE == ulCommand)
	{
		MemoryManager::Inst.CreateObject(&pStrRes);
		pEntityList->SeekToBegin();
		PNODE currNode = (PNODE)pEntityList->GetCurrElem();
		String* pStrArg = (String*)pArg;
		std::vector<std::string> dateList;
		while(currNode != 0) {
			if (currNode->GetLVal() != 0)
			{
				dateList.push_back(std::string(currNode->GetLVal()));
			}
			pEntityList->Seek(1, false);
			currNode = (PNODE)pEntityList->GetCurrElem();
		}
		std::string oldestDate = (DateTimeOperations::GetOldestDate(dateList));
		pStrRes->SetValue(DateTimeOperations::GetOldestDate(dateList));
	}
	else if (COMMAND_TYPE_GET_LATEST_DATE == ulCommand)
	{
		MemoryManager::Inst.CreateObject(&pStrRes);
		pEntityList->SeekToBegin();
		PNODE currNode = (PNODE)pEntityList->GetCurrElem();
		String* pStrArg = (String*)pArg;
		std::vector<std::string> dateList;
		while(currNode != 0) {
			if (currNode->GetLVal() != 0)
			{
				dateList.push_back(currNode->GetLVal());
			}
			pEntityList->Seek(1, false);
			currNode = (PNODE)pEntityList->GetCurrElem();
		}
		std::string latestDate = (DateTimeOperations::GetLatestDate(dateList));
		pStrRes->SetValue(DateTimeOperations::GetLatestDate(dateList));
	}
	// first handle the commands that would need to access the execution context
	else if (COMMAND_TYPE_FILTER_SUBTREE == ulCommand) {
        MemoryManager::Inst.CreateObject(&pListRes);
        pEntityList->SeekToBegin();
        PNODE currNode = (PNODE)pEntityList->GetCurrElem();
        while(currNode != 0)
        {
            PENTITYLIST pNodeList = 0;
            MemoryManager::Inst.CreateObject(&pNodeList);
            FilterSubTree(currNode, p_Arg, pContext, pNodeList);
            pListRes->SeekToBegin();
            PNODE internalNode = (PNODE)pNodeList->GetCurrElem();
            while(internalNode != 0)
            {
                pListRes->push_back(internalNode->GetCopy());
                pNodeList->Seek(1, false);
                internalNode = (PNODE)pNodeList->GetCurrElem();
            }
            pEntityList->Seek(1, false);
            currNode = (PNODE)pEntityList->GetCurrElem();
        }
	}
	else
	{
        if(0 != p_Arg)
		{
			p_EntityArg = p_Arg->Execute(pContext);
		}
		MemoryManager::Inst.CreateObject(&pListRes);
		EntityList::const_iterator ite1 = pEntityList->begin();
		EntityList::const_iterator iteEnd1 = pEntityList->end();
		for( ; ite1 != iteEnd1; ++ite1)
		{
			PENTITY pRes = 0;
            if ((*ite1)->ul_Type == ENTITY_TYPE_NODE) {
                pRes = ExecuteNodeCommand(ulCommand, *ite1, pContext);
            } else {
                pRes = ExecuteEntityCommand(ulCommand, *ite1, p_EntityArg);
            }
            
			switch(pRes->ul_Type)
			{
                case ENTITY_TYPE_NULL:
				{
					MemoryManager::Inst.DeleteObject(pRes);
				}
                case ENTITY_TYPE_INT:
                case ENTITY_TYPE_NODE:
                case ENTITY_TYPE_STRING:
                case ENTITY_TYPE_LIST:
				{
					pListRes->push_back(pRes);
					break;
				}
                case ENTITY_TYPE_BOOL:
				{
					if(((PBool)pRes)->GetValue())
					{
						if(ENTITY_TYPE_NODE != (*ite1)->ul_Type)
						{
							pListRes->push_back((*ite1)->GetCopy());
						}
						else
						{
							pListRes->push_back(*ite1);
						}
					}
					MemoryManager::Inst.DeleteObject(pRes);
					break;
				}
			}
		}
	}
    
	if(0 != pIntRes)
	{
		return pIntRes;
	}
	if(0 != pListRes)
	{
		return pListRes;
	}
	if(0 != pNullRes)
	{
		return pNullRes;
	}
	if(0 != pEntityRes)
	{
		return pEntityRes;
	}
	if(0 != pNodeRes)
	{
		return pNodeRes;
	}
    if(0 != pStrRes)
    {
        return pStrRes;
    }
	return 0;
}

void Command::AddSubtreeToNodeList(PENTITYLIST pList, PNODE pRoot)
{
	pList->push_back(pRoot);
	PNODE pChild = pRoot->GetFirstChild();
	while(0 != pChild)
	{
		AddSubtreeToNodeList(pList, pChild);
		pChild = pChild->GetRightSibling();
	}
}

void Command::FilterSubTree(PNODE root, ExecutionTemplate* arg, ExecutionContext* context, PENTITYLIST resultList)
{
    context->map_Var[context->p_MD->s_ListItemVar] = root;
    PBool res = (PBool)arg->Execute(context);
    if (res->GetValue()) {
        resultList->push_back(root);
    }
    PNODE pChild = root->GetFirstChild();
	while(0 != pChild)
	{
		FilterSubTree(pChild, arg, context, resultList);
		pChild = pChild->GetRightSibling();
	}
}
