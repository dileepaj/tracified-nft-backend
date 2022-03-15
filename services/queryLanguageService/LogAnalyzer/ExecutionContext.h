#ifndef _EXECUTIONCONTEXT_H
#define _EXECUTIONCONTEXT_H

#include "CommonIncludes.h"
#include "Entity.h"
#include "ExecutionTemplateList.h"

class MetaData;


class ExecutionContext
{
public:
	class Iteration
	{
	public:
		Iteration() : s_VarName(EMPTY_STRING)
		{

		}
		MSTRING s_VarName;
		PENTITY p_Entity;
		LST_ENTITYPTR::const_iterator ite_Curr;
		ExecutionTemplateList::const_iterator ite_FirstTemplate;
	};
	typedef std::stack<ExecutionContext::Iteration>		STK_ITERATION;
	MAP_STR_ENTITYPTR	map_Var;	// first = variable name		second = current value of variable
	STK_ITERATION		stk_Iteration;
	MAP_STR_EXECTMPLIST*	p_mapFunctions;
	MetaData* p_MD;

public:
	ExecutionContext()
		: p_mapFunctions(0), p_MD(0)
	{

	}
};

#endif