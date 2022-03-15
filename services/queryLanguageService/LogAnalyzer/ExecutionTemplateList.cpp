#include "ExecutionTemplateList.h"
#include "ExecutionTemplate.h"
#include "ExecutionContext.h"
#include "Value.h"
#include "Bool.h"
#include "Utils.h"
#include "MemMan.h"
#include "ProcessLog.h"

void ExecutionTemplateList::Destroy()
{
	DestroyCollection(*this);
    
	MemoryManager::Inst.DeleteObject(this);
}

void ExecutionTemplateList::Execute(ExecutionContext* pContext)
{
	STK_ITER stkLoopStart;
	MAP_ITER_ITER mapLoopStartEnd;
	MAP_ITER_ITER mapLoopEndStart;
	ExecutionTemplateList::const_iterator ite1 = begin();
	ExecutionTemplateList::const_iterator iteEnd1 = end();
    for( ; ite1 != iteEnd1; ++ite1)
	{

        switch((*ite1)->GetSpecialCommand())
		{
            case COMMAND_TYPE_IF:
            case COMMAND_TYPE_IFNOT:
            case COMMAND_TYPE_WHILE:
			{
				stkLoopStart.push(ite1);
				break;
			}
            case COMMAND_TYPE_ENDIF:
            case COMMAND_TYPE_DO:
			{
				ExecutionTemplateList::const_iterator iteStart = stkLoopStart.top();
				mapLoopStartEnd[iteStart] = ite1;
				mapLoopEndStart[ite1] = iteStart;
				stkLoopStart.pop();
				break;
			}
		}
	}
    
	STK_FRAMES stkFrames;
	ite1 = begin();
	iteEnd1 = end();
	while(ite1 != iteEnd1)
	{
#ifdef LOG
		MSTRING sLog = MSTRING(_MSTR(Executing)) + MSTRING(SPACE) + (*ite1)->GetCodeLine();
		ProcessLog::Write(sLog);
#endif
		switch((*ite1)->GetSpecialCommand())
		{
            case COMMAND_TYPE_IF:
			{
				PENTITY pRes = (*ite1)->Execute(pContext);
				if((0 != (PBool)pRes) && (((PBool)pRes)->GetValue()))
				{
					++ite1;
				}
				else
				{
					// If evaluated to false
					// start from the statement that just follows the corresponding EndIf
					ite1 = (mapLoopStartEnd.find(ite1)->second) + 1;
				}
				break;
			}
            case COMMAND_TYPE_IFNOT:
			{
				PENTITY pRes = (*ite1)->Execute(pContext);
				if((0 != (PBool)pRes) && (((PBool)pRes)->GetValue()))
				{
					// If evaluated to false
					// start from the statement that just follows the corresponding EndIf
					ite1 = (mapLoopStartEnd.find(ite1)->second) + 1;
				}
				else
				{
					++ite1;
				}
				break;
			}
            case COMMAND_TYPE_ENDIF:
			{
				++ite1;
				break;
			}
            case COMMAND_TYPE_WHILE:
			{
				bool bGoInsideWhileLoop = false;
				if((*ite1)->IsEmpty())
				{
					// While(true)
					bGoInsideWhileLoop = true;
				}
				else
				{
					PENTITY pRes = (*ite1)->Execute(pContext);
					if((0 != (PBool)pRes) && (((PBool)pRes)->GetValue()))
					{
						bGoInsideWhileLoop = true;
					}
				}
				if(bGoInsideWhileLoop)
				{
					// First check whether this While is in the stack already (it happens when the execution hits this point after the first time)
					if((stkFrames.size() == 0) || (stkFrames.top().ite_LoopStart != ite1))
					{
						// This While is reached for the first time
						// Add a frame to the stack corresponding to this While-Do loop
						StackFrame sf;
						sf.ite_LoopStart = ite1;	// While
						sf.ite_LoopEnd = mapLoopStartEnd.find(ite1)->second;	// Do
						stkFrames.push(sf);
					}
					++ite1;
				}
				else
				{
					if((stkFrames.size() > 0) && (stkFrames.top().ite_LoopStart == ite1))
					{
						// Frame stack contains a frame for this While loop
						// Remove the frame and go out of the While loop
						stkFrames.pop();
					}
					ite1 = (mapLoopStartEnd.find(ite1)->second) + 1;
				}
				
				break;
			}
            case COMMAND_TYPE_DO:
			{
				// Find the corresponding While statement and move to it
				ite1 = mapLoopEndStart.find(ite1)->second;
				break;
			}
            case COMMAND_TYPE_BREAK:
			{
				// Pop the stack
				// Find the Do element from that frame and move to the statement that just follows it
				ite1 = stkFrames.top().ite_LoopEnd + 1;
				stkFrames.pop();
				break;
			}
            case COMMAND_TYPE_CONTINUE:
			{
				// Read the top frame and move to the corresponding While statement
				ite1 = stkFrames.top().ite_LoopStart;
				break;
			}
            default:
			{
				(*ite1)->Execute(pContext);
				++ite1;
			}
		}
	}
}