#ifndef _EXECUTIONTEMPLATELIST_H
#define _EXECUTIONTEMPLATELIST_H

#include "CommonIncludes.h"

class ExecutionTemplate;
class ExecutionContext;

class ExecutionTemplateList : public std::vector<ExecutionTemplate*>
{
	typedef std::map<ExecutionTemplateList::const_iterator, ExecutionTemplateList::const_iterator> MAP_ITER_ITER;
	typedef std::stack<ExecutionTemplateList::const_iterator> STK_ITER;
    
	class StackFrame
	{
	public:
		ExecutionTemplateList::const_iterator ite_LoopStart;
		ExecutionTemplateList::const_iterator ite_LoopEnd;
	};
    
	typedef std::stack<StackFrame>	STK_FRAMES;
public:
	void Destroy();
	void Execute(ExecutionContext* pContext);
};

#endif