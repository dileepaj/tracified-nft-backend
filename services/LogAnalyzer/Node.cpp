#include "Node.h"
#include "EntityList.h"
#include "MemMan.h"
#include "StringOperations.h"
#include "Utils.h"
#include <limits>

Node::Node(MULONG ulID)
: ul_ID(ulID), us_Type(NODE_TYPE_GENERAL), z_Value(0), z_LValue(0), z_RValue(0), z_CustomStr(0), ul_MinChildWeight(std::numeric_limits<MULONG>::min()), ul_MaxChildWeight(std::numeric_limits<MULONG>::max())
,by_Nature(NODE_NATURE_GENERAL), ul_Weight(0), p_Next(0), p_Prev(0), p_Parent(0), p_FirstChild(0), ul_ChildCount(0), p_CustomObj(0)
{
	ul_Type = ENTITY_TYPE_NODE;
}

Node::~Node()
{
    
}

void Node::Destroy()
{
	
}

void Node::DestroyNodeAlone()
{
	MemoryManager::Inst.DeleteObject(this);
}

void Node::DestroyWithSubTree()
{
	// first destroy the child subtrees
	PNODE pChild = p_FirstChild;
	while(0 != pChild)
	{
		PNODE pChildCpy = pChild;
		pChild = pChild->GetRightSibling();
		pChildCpy->DestroyWithSubTree();
	}
	this->DestroyNodeAlone();
}

MSTRING Node::ToString() {
#ifdef ISWIDECHAR
    wchar_t buff[100];
    wsprintf(buff, L"%x", this);
    return MSTRING(buff);
#else
    char buff[100];
    sprintf(buff, "%p", this);
    return MSTRING(buff);
#endif
}

PNODE Node::GetLeftSibling()
{
	return p_Prev;
}

PNODE Node::GetRightSibling()
{
	return p_Next;
}

PNODE Node::GetParent()
{
	return p_Parent;
}

PNODE Node::GetFirstChild()
{
	return p_FirstChild;
}

PNODE Node::GetLastChild()
{
	return p_LastChild;
}

MULONG Node::GetChildCount()
{
	return ul_ChildCount;
}

PMCHAR Node::GetValue()
{
	return z_Value;
}

PMCHAR Node::GetLVal()
{
	return z_LValue;
}

PMCHAR Node::GetRVal()
{
	return z_RValue;
}

PMCHAR Node::GetCustomString()
{
	return z_CustomStr;
}

MULONG Node::GetID()
{
	return ul_ID;
}

MUSHORT Node::GetType()
{
	return us_Type;
}

MULONG Node::GetWeight()
{
	return ul_Weight;
}

MBYTE Node::GetNature()
{
	return by_Nature;
}

PVOID Node::GetCustomObj() {
    return p_CustomObj;
}

PENTITY Node::GetEntityObj() {
    return p_EntityObj;
}

/*!
 * Creates and outputs a string that is constructed by concatenating all the string values of the nodes
 * in the subtree
 */
MSTRING Node::GetAggregatedValue()
{
	MSTRING sAggVal = EMPTY_STRING;
	AddAggregatedValue(sAggVal);
	return sAggVal;
}

void Node::AddAggregatedValue(MSTRING& sStr)
{
	if(0 != z_Value)
	{
		sStr += z_Value;
	}
	PNODE pChild = p_FirstChild;
	while(0 != pChild)
	{
		if((pChild == p_FirstChild) && (pChild->GetLVal() != 0))
		{
			sStr += pChild->GetLVal();
		}
        
		pChild->AddAggregatedValue(sStr);
        
		if(pChild->GetRVal() != 0)
		{
			sStr += pChild->GetRVal();
		}
        
		pChild = pChild->GetRightSibling();
	}
}

PNODE Node::GetCopy()
{
	// Only the node properties are copied, not the hierarchical relationships in the tree (e.g. parent, left sibling, etc)
	PNODE pNewNode = MemoryManager::Inst.CreateNode(ul_ID);
	pNewNode->SetValue(z_Value);
	pNewNode->SetLValue(z_LValue);
	pNewNode->SetRValue(z_RValue);
	pNewNode->SetCustomString(z_CustomStr);
	pNewNode->SetType(us_Type);
	pNewNode->SetNature(by_Nature);
	pNewNode->SetWeight(ul_Weight);
	return pNewNode;
}

PNODE Node::GetCopyWithSubTree()
{
	// All the properties plus subtree are copied.
	// This copy does not include parent, left sibling and right sibling.
	PNODE pNewNode = GetCopy();
	PNODE pChild = p_FirstChild;
	while(0 != pChild)
	{
		PNODE pChildCopy = pChild->GetCopyWithSubTree();
		pNewNode->AddNode(pChildCopy, false);
		pChild = pChild->GetRightSibling();
	}
	return pNewNode;
}

MULONG Node::GetMinimumChildWeight()
{
	return ul_MinChildWeight;
}

MULONG Node::GetMaximumChildWeight()
{
	return ul_MaxChildWeight;
}

void Node::SetValue(PMCHAR pVal)
{
	StringOperations::SetString(&z_Value, pVal);
}

void Node::SetLValue(PMCHAR pLVal)
{
	StringOperations::SetString(&z_LValue, pLVal);
}

void Node::SetRValue(PMCHAR pRVal)
{
	StringOperations::SetString(&z_RValue, pRVal);
}

void Node::SetCustomString(PMCHAR pCustStr)
{
	StringOperations::SetString(&z_CustomStr, pCustStr);
}

void Node::SetID(MULONG ulID)
{
	ul_ID = ulID;
}

void Node::SetWeight(MULONG ulWeight)
{
	ul_Weight = ulWeight;
}

void Node::SetNature(MBYTE byNature)
{
	by_Nature = byNature;
}

void Node::SetType(MUSHORT usType)
{
	us_Type = usType;
}

void Node::SetRightSibling(PNODE pNext)
{
	p_Next = pNext;
}

void Node::SetLeftSibling(PNODE pPrev)
{
	p_Prev = pPrev;
}

void Node::SetParent(PNODE pParent)
{
	p_Parent = pParent;
}

void Node::SetFirstChild(PNODE pFirstChild) {
	p_FirstChild = pFirstChild;
}

void Node::SetLastChild(PNODE pLastChild) {
	p_LastChild = pLastChild;
}

void Node::SetMinimumChildWeight(MULONG ulMinChildWeight)
{
	ul_MinChildWeight = ulMinChildWeight;
}

void Node::SetMaximumChildWeight(MULONG ulMaxChildWeight)
{
	ul_MaxChildWeight = ulMaxChildWeight;
}

void Node::SetCustomObj(PVOID obj) {
    p_CustomObj = obj;
}

void Node::SetEntityObj(PENTITY obj) {
    p_EntityObj = obj;
}

void Node::Expand(LST_STR& lstTokens)
{
	if(0 == z_Value)
	{
		return;
	}
	LST_STR lstSep;
	LST_STR lstVal;
    MSTRING sValue(z_Value);
	Utils::TokenizeString(sValue, lstTokens, lstSep, lstVal);
	assert(lstSep.size() == (lstVal.size() + 1));
	LST_STR::const_iterator iteSep = lstSep.begin();
	LST_STR::const_iterator iteVal = lstVal.begin();
	LST_STR::const_iterator iteValEnd = lstVal.end();
	while(true)
	{
		if(iteVal == iteValEnd)
		{
			break;
		}
		MSTRING sLeftString = *iteSep;
		++iteSep;
		MSTRING sRightString = *iteSep;
		MSTRING sVal = *iteVal;
		++iteVal;
		PNODE pNewChild = AddNode();
		pNewChild->SetValue((PMCHAR)sVal.c_str());
		pNewChild->SetLValue((PMCHAR)sLeftString.c_str());
		pNewChild->SetRValue((PMCHAR)sRightString.c_str());
	}
	SetValue(0);
}

PNODE Node::AddNode()
{
	++ul_ChildCount;
	PNODE pNewChild = MemoryManager::Inst.CreateNode(ul_ChildCount + NODE_ID_OFFSET);
	if(0 == p_FirstChild)
	{
		p_FirstChild = pNewChild;
		p_LastChild = pNewChild;
	}
	else
	{
		p_LastChild->SetRightSibling(pNewChild);
		pNewChild->SetLeftSibling(p_LastChild);
		p_LastChild = pNewChild;
	}
	pNewChild->SetParent(this);
    
	return pNewChild;
}

void Node::AppendNode(PNODE pNewChild) {
    ++ul_ChildCount;
	if(0 == p_FirstChild)
	{
		p_FirstChild = pNewChild;
		p_LastChild = pNewChild;
	}
	else
	{
		p_LastChild->SetRightSibling(pNewChild);
		pNewChild->SetLeftSibling(p_LastChild);
		p_LastChild = pNewChild;
	}
	pNewChild->SetParent(this);
}

PNODE Node::AddNode(PNODE pNode, bool bMakeCopy)
{
	PNODE pNodeToAdd = pNode;
	if(bMakeCopy)
	{
		pNodeToAdd = pNode->GetCopyWithSubTree();
	}
	// Check whether the ID of the new node already exist
	// If so we need to merge the new node with the existing matching node
	MULONG ulID = pNodeToAdd->GetID();
	if(ulID <= (ul_ChildCount + NODE_ID_OFFSET))
	{
		// Matching Node exists
		// Now traverse the child list and find the matching node
		PNODE pMatchingNode = 0;
		// Should we traverse forward or backward?
		if((ulID - NODE_ID_OFFSET) < (ul_ChildCount / 2))
		{
			// Matching node is within the first half of the child list and hence we should traverse forward
			PNODE pChild = p_FirstChild;
			while(0 != pChild)
			{
				if(pChild->GetID() == ulID)
				{
					pMatchingNode = pChild;
					break;
				}
				pChild = pChild->GetRightSibling();
			}
		}
		else
		{
			// Matching node is within the second half of the child list and hence we should traverse backward
			PNODE pChild = p_LastChild;
			while(0 != pChild)
			{
				if(pChild->GetID() == ulID)
				{
					pMatchingNode = pChild;
					break;
				}
				pChild = pChild->GetLeftSibling();
			}
		}
        
		if(0 == pMatchingNode)
		{
			// We should never arrive here
			pNodeToAdd->DestroyWithSubTree();
			return 0;
		}
        
		// The node being added will lose all its properties
		// The new node's children will be added to the existing node
		PNODE pChild = pNodeToAdd->GetFirstChild();
		while(0 != pChild)
		{
			PNODE pChildCpy = pChild;
			pChild = pChild->GetRightSibling();
			pMatchingNode->AddNode(pChildCpy, false);
		}
        
		// Now we can destroy NodeToAdd because it is redundant
		pNodeToAdd->DestroyNodeAlone();
        
		return pMatchingNode;
	}
	else
	{
		// No matching node
		// Add the node as a completely new child node
		// However, the new node is going to get a new id
		++ul_ChildCount;
		pNodeToAdd->SetID(ul_ChildCount + NODE_ID_OFFSET);
		if(0 == p_FirstChild)
		{
			p_FirstChild = pNodeToAdd;
			p_LastChild = pNodeToAdd;
			pNodeToAdd->SetLeftSibling(0);
		}
		else
		{
			p_LastChild->SetRightSibling(pNodeToAdd);
			pNodeToAdd->SetLeftSibling(p_LastChild);
			p_LastChild = pNodeToAdd;
		}
		pNodeToAdd->SetParent(this);
		pNodeToAdd->SetRightSibling(0);
		return pNodeToAdd;
	}
}

PNODE Node::AddNodeWithWeight(MULONG ulWeight)
{
	if(NODE_NATURE_ORGANIZING != by_Nature)
	{
		return 0;
	}
    
	if((ulWeight < ul_MinChildWeight) || (ulWeight > ul_MaxChildWeight))
	{
		return 0;
	}
    
	PNODE pChild = p_FirstChild;
	PNODE pNodeInsertAfter = 0;
	while(0 != pChild)
	{
		if(pChild->GetWeight() > ulWeight)
		{
			break;
		}
		pNodeInsertAfter = pChild;
		pChild = pChild->GetRightSibling();
	}
    
	++ul_ChildCount;
	PNODE pNewNode = MemoryManager::Inst.CreateNode(ul_ChildCount + NODE_ID_OFFSET);
	pNewNode->SetParent(this);
	if(0 == pNodeInsertAfter)
	{
		// New node will be the first child
		if(0 != p_FirstChild)
		{
			p_FirstChild->SetLeftSibling(pNewNode);
		}
		pNewNode->SetRightSibling(p_FirstChild);
		p_FirstChild = pNewNode;
	}
	else
	{
		PNODE pNodeInsertBefore = pNodeInsertAfter->GetRightSibling();
		if(0 == pNodeInsertBefore)
		{
			// New node is going to be the last child node
			p_LastChild = pNewNode;
		}
		else
		{
			pNodeInsertBefore->SetLeftSibling(pNewNode);
			pNewNode->SetRightSibling(pNodeInsertBefore);
		}
		pNodeInsertAfter->SetRightSibling(pNewNode);
		pNewNode->SetLeftSibling(pNodeInsertAfter);
	}
    
	return pNewNode;
}

void Node::ReadValueFromFile(CPMCHAR zFilename)
{
	//MIFSTREAM file;
	std::ifstream file;
	file.open(zFilename, std::ios::in | std::ios::binary);
	if(!file.is_open())
	{
		return;
	}
	file.seekg(0, std::ios::end);
	MULONG ulLen = file.tellg();
	file.seekg(0, std::ios::beg);
	// ignore first two bytes if the file is unicode
	if(ulLen >= 2)
	{
		MBYTE abyTemp[2];
		file.read((char*)abyTemp, 2);
		if((0xff == abyTemp[0]) && (0xfe == abyTemp[1]))
		{
			file.seekg(2, std::ios::beg);
			ulLen -= 2;
		}
		else
		{
			file.seekg(0, std::ios::beg);
		}
	}
	PMBYTE byBuff = 0;
	MemoryManager::Inst.CreateBuffer(&byBuff, ulLen);
	file.read((char*)byBuff, ulLen);
	file.close();
	z_Value = (PMCHAR)byBuff;
}
PNODE Node::IsHavingCustomString(MSTRING customString)
{
	PNODE result = 0;
	if (z_CustomStr == customString)
	{
		result = this;
	}
	return result;
}

PNODE Node::GetChildNodeByCustomString(MSTRING customString)
{
	PNODE node = 0;
	return node;
}
