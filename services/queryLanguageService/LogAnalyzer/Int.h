#include "Value.h"

class Int : public Value<MULONG, ENTITY_TYPE_INT>
{
public:
	bool b_IsNegative;
    
	Int()
    : b_IsNegative(false), Value<MULONG, ENTITY_TYPE_INT>()
	{
        
	}
    
	virtual ~Int()
	{
        
	}
    
    virtual MSTRING ToString()
    {
#ifdef ISWIDECHAR
        wchar_t buff[100];
        swprintf(buff, L"%d", val);
        return MSTRING(buff);
#else
        char buff[100];
        sprintf(buff, "%lu", val);
        return MSTRING(buff);
#endif
    }
};