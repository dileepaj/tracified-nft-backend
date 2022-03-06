#include "MemMan.h"
#include "Node.h"

MemoryManager MemoryManager::Inst;

MemoryManager::MemoryManager()
{

}

MemoryManager::~MemoryManager()
{

}

PNODE MemoryManager::CreateNode(MULONG ulID)
{
	return new Node(ulID);
}

void MemoryManager::DestroyNode(PNODE pNode)
{
	delete pNode;
}