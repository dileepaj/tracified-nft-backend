#ifndef _BOOL_H
#define _BOOL_H

#include "Value.h"

class Bool : public Value<bool, ENTITY_TYPE_BOOL>
{
public:
	Bool();
	virtual ~Bool();
	PBool And(PBool pArg);
	PBool Or(PBool pArg);
};


#endif


