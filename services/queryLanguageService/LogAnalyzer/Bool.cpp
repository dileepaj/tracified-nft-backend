#include "Bool.h"
#include "MemMan.h"

Bool::Bool()
: Value<bool, ENTITY_TYPE_BOOL>()
{
    
}

Bool::~Bool()
{
    
}

PBool Bool::And(PBool pArg)
{
	bool bRes = val && (pArg->GetValue());
	PBool pRes = 0;
	MemoryManager::Inst.CreateObject(&pRes);
	pRes->SetValue(bRes);
	return pRes;
}
PBool Bool::Or(PBool pArg)
{
	bool bRes = val || (pArg->GetValue());
	PBool pRes = 0;
	MemoryManager::Inst.CreateObject(&pRes);
	pRes->SetValue(bRes);
	return pRes;
}