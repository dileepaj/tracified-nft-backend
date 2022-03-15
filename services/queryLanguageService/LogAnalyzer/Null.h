#ifndef _NULL_H
#define _NULL_H

#include "CommonIncludes.h"
#include "Entity.h"
#include "MemMan.h"

class Null : public Entity
{
public:
	Null()
	{
		ul_Type = ENTITY_TYPE_NULL;
	}
    
	Entity* GetCopy()
	{
		Null* pNew = 0;
		MemoryManager::Inst.CreateObject(&pNew);
		return pNew;
	}
    
	bool IsNull()
	{
		return true;
	}
};

#endif