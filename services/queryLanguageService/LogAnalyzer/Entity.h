#ifndef _ENTITY_H
#define _ENTITY_H

#include "CommonIncludes.h"
#include "MemMan.h"

class Entity
{
public:
	MULONG ul_Type;
    
	virtual ~Entity()	{}
	virtual Entity* GetCopy() = 0;
    virtual MSTRING ToString() {return EMPTY_STRING;}
	virtual bool IsNull()
	{
		return false;
	}
	virtual void Destroy()
	{
		MemoryManager::Inst.DeleteObject(this);
	}
};

#endif