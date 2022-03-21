#ifndef _VALUE_H
#define _VALUE_H

#include "Entity.h"
#include "MemMan.h"

template <class T, int Type>
class Value : public Entity
{
protected:
	T val;
    
public:
	Value()
	{
		ul_Type = Type;
	}
	Value(T v)
    :val(v)
	{
		ul_Type = Type;
	}
    
	virtual ~Value()
	{
        
	}
    
	void SetValue(T v)
	{
		val = v;
	}
    
	T GetValue()
	{
		return val;
	}
    
	Entity* GetCopy()
	{
		Value<T, Type>* pNew = 0;
		MemoryManager::Inst.CreateObject(&pNew);
		pNew->SetValue(val);
		return pNew;
	}
};

#endif