#include "CommonIncludes.h"
#include "Entity.h"

class ExecutionContext;

class ExecutionTemplate : public Entity
{
protected:
	MSTRING s_StartVarName;
	LST_COMMANDPTR	lst_Commands;
	PENTITY p_Entity;
	MULONG ul_SpecialCommand;
	MSTRING s_CodeLine;
    
public:
	ExecutionTemplate();
	~ExecutionTemplate();
    
	void Destroy();
    
	ExecutionTemplate* GetCopy();	// override
    
	void SetStartVarName(MSTRING sName);
	void AddCommand(Command* pCommand);
	void SetEntity(PENTITY pEntity);
	void SetSpecialCommand(MULONG ulCmd);
	void SetCodeLine(MSTRING sLine);
	MSTRING GetStartVarName();
	PENTITY GetEntity();
	MULONG GetSpecialCommand();
	MSTRING GetCodeLine();
	PENTITY Execute(ExecutionContext* pContext);
	bool IsEmpty();
    
	friend std::ostream& operator << (ExecutionTemplate& et, std::ostream& x)
	{
		x << _MSTR(===ExecutionTemplate==\n);
		x << _MSTR(Start Var: ) << et.s_StartVarName.c_str();
		x << _MSTR(   Special Command: ) << et.ul_SpecialCommand;
		x << _MSTR(\nEntity: ) << et.p_Entity;
        return x;
	}
    
private:
    PENTITY ExecuteCommand(PENTITY entity, ExecutionContext* context, Command* cmd);
};