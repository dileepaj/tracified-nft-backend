#include "EntityList.h"
#include "ExecutionTemplate.h"
#include "ExecutionContext.h"
#include "Utils.h"
#include "Int.h"

EntityList::EntityList()
: st_CurrElemPos(0)
{
	ul_Type = ENTITY_TYPE_LIST;
}

EntityList::~EntityList()
{
    
}

void EntityList::Destroy()
{
	// First destroy the children in the list
	DestroyCollection(*this);
    
	MemoryManager::Inst.DeleteObject(this);
}

Entity* EntityList::GetCopy()
{
	EntityList* pNew = 0;
	MemoryManager::Inst.CreateObject(&pNew);
	const_iterator ite1 = this->begin();
	const_iterator iteEnd1 = this->end();
	for( ; ite1 != iteEnd1; ++ite1)
	{
		if(ENTITY_TYPE_NODE != (*ite1)->ul_Type)
		{
			pNew->push_back((*ite1)->GetCopy());
		}
		else
		{
			pNew->push_back(*ite1);
		}
	}
	return pNew;
}

void EntityList::ExecuteElements(ExecutionContext* pContext)
{
	iterator ite1 = this->begin();
	const_iterator iteEnd1 = this->end();
	for( ; ite1 != iteEnd1; ++ite1)
	{
		if(ENTITY_TYPE_EXECUTION_TEMPLATE == (*ite1)->ul_Type)
		{
			ExecutionTemplate* pET = (ExecutionTemplate*)(*ite1);
			if(0 != pET)
			{
				PENTITY pRes = pET->Execute(pContext);
				(*ite1) = pRes;
				pET->Destroy();
			}
		}
	}
}

Entity* EntityList::GetCurrElem()
{
	EntityList::size_type stPos = 0;
	EntityList::const_iterator ite = begin();
	EntityList::const_iterator iteEnd = end();
	while(ite != iteEnd)
	{
		if(stPos == st_CurrElemPos)
		{
			return (*ite);
		}
		++stPos;
		++ite;
	}
	return 0;
}

void EntityList::SeekToBegin()
{
	st_CurrElemPos = 0;
}

void EntityList::SeekToEnd()
{
	st_CurrElemPos = size() - 1;
}

void EntityList::Seek(MULONG ulOffset, bool bIsNegative)
{
	if(bIsNegative)
	{
		st_CurrElemPos -= ulOffset;
	}
	else
	{
		st_CurrElemPos += ulOffset;
	}
}

EntityList* EntityList::GetInnerCount() {
    EntityList *pListRes = 0;
    MemoryManager::Inst.CreateObject(&pListRes);
    SeekToBegin();
    PENTITY ent = GetCurrElem();
    while (ent) {
        PInt count = 0;
        MemoryManager::Inst.CreateObject(&count);
        count->SetValue(((PENTITYLIST)ent)->size());
        pListRes->push_back(count);
        
        Seek(1, false);
        ent = GetCurrElem();
    }
    return pListRes;
}