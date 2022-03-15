#ifndef _MEMORYMANAGER_H
#define _MEMORYMANAGER_H

#include "CommonIncludes.h"

class MemoryManager
{
public:
	static MemoryManager	Inst;

	template <class T> void CreateObject(T** ppBuff)
	{
		(*ppBuff) = new T;
	}

	template <class T> void DeleteObject(T* pBuff)
	{
		delete pBuff;
	}

	template <class T> void	CreateBuffer(T** ppBuff, MULONG ulSize)
	{
		(*ppBuff) = new T[ulSize];
	}

	template <class T> void DestroyBuffer(T* pBuff)
	{
		delete [] pBuff;
	}

	PNODE CreateNode(MULONG ulID);
	void DestroyNode(PNODE pNode);

//private:
	MemoryManager();
	~MemoryManager();	
};

#endif