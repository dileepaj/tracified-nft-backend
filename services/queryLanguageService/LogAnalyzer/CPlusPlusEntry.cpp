// //
// //  CPlusPlusEntry.cpp
// //  LogAnalyzer
// //
// //  Created by Dileepa Jayathilake on 7/4/13.
// //  Copyright (c) 2013 99x Eurocenter. All rights reserved.
// //

// #include "CPlusPlusEntry.h"
// #include "CommonIncludes.h"
// #include "Utils.h"
// #include "Tests.h"

// void RunMenu();

// typedef int (*FuncType)();
// typedef std::map<int, FuncType>	MAP_INT_FUNC;
// class Menu;
// typedef std::map<int, Menu*>	MAP_INT_MENUPTR;

// class Menu
// {
// 	class MenuItem
// 	{
// 	public:
// 		FuncType func;
// 		Menu* pMenu;
        
// 		MenuItem()
//         : func(0), pMenu(0)
// 		{
            
// 		}
        
// 		void Destroy()
// 		{
// 			if(0 != pMenu)
// 			{
// 				pMenu->Destroy();
// 			}
// 			delete this;
// 		}
// 		int Execute()
// 		{
// 			if(0 != pMenu)
// 			{
// 				return pMenu->Execute();
// 			}
// 			else if(0 != func)
// 			{
// 				return func();
// 			}
// 			return -1;
// 		}
// 	};
// 	typedef std::map<int, MenuItem*> MAP_INT_MENUITEMPTR;
// public:
    
// 	void Destroy()
// 	{
// 		MAP_INT_MENUITEMPTR::const_iterator ite1 = map_Items.begin();
// 		MAP_INT_MENUITEMPTR::const_iterator iteEnd1 = map_Items.end();
// 		for( ; ite1 != iteEnd1; ++ite1)
// 		{
// 			((*ite1).second)->Destroy();
// 		}
// 		delete this;
// 	}
    
// 	void AddFunction(MSTRING sDisplayName, FuncType func)
// 	{
// 		int iNewIndex = map_Items.size() + 1;
// 		MenuItem* pNewItem = new MenuItem;
// 		pNewItem->func = func;
// 		map_Items[iNewIndex] = pNewItem;
// 		map_ItemNames[iNewIndex] = sDisplayName;
// 	}
    
// 	Menu* AddMenu(MSTRING sDisplayName)
// 	{
// 		int iNewIndex = map_Items.size() + 1;
// 		MenuItem* pItem = new MenuItem;
// 		Menu* pNewMenu = new Menu;
// 		pItem->pMenu = pNewMenu;
// 		map_Items[iNewIndex] = pItem;
// 		map_ItemNames[iNewIndex] = sDisplayName;
// 		return pNewMenu;
// 	}
	
// 	int Execute()
// 	{
// 		std::wcout<<"\n";
// 		MAP_INT_STR::const_iterator ite1 = map_ItemNames.begin();
// 		MAP_INT_STR::const_iterator iteEnd1 = map_ItemNames.end();
// 		for( ; ite1 != iteEnd1; ++ite1)
// 		{
// 			std::wcout<<(*ite1).first<<". "<<(*ite1).second.c_str()<<"\n";
// 		}
// 		std::wcout<<"Select Option:";
// 		int iSel = -1;
// 		std::wcin>>iSel;
// 		MAP_INT_MENUITEMPTR::const_iterator iteFind = map_Items.find(iSel);
// 		if(map_Items.end() != iteFind)
// 		{
// 			return ((*iteFind).second)->Execute();
// 		}
// 		return -1;
// 	}
// 	MAP_INT_MENUITEMPTR map_Items;
// 	MAP_INT_STR	 map_ItemNames;
// };

// int Quit()
// {
// 	return 0;
// }

// int TestTokenize()
// {
// 	std::wcout<<"Testing Utils::TokenizeStringBasic\n";
// 	MSTRING sStr = "<xx>(5)";
// 	LST_STR lstTokens, lstStr, lstSep, lstComponents;
// 	LST_INT lstComponentTypes;
// 	lstTokens.push_back("<");
// 	lstTokens.push_back(">");
// 	lstTokens.push_back("(");
// 	lstTokens.push_back(")");
// 	Utils::TokenizeStringBasic(sStr, lstTokens, lstComponents, lstComponentTypes);
// 	std::wcout<<"String to tokenize: "<<&sStr<<"\n\n";
// 	std::wcout<<"Components\n";
// 	LST_STR::const_iterator ite1 = lstComponents.begin();
// 	LST_STR::const_iterator iteEnd1 = lstComponents.end();
// 	for( ; ite1 != iteEnd1; ++ite1)
// 	{
// 		std::wcout<<&*ite1<<"\n";
// 	}
// 	std::wcout<<"\nComponent types\n";
// 	LST_INT::const_iterator ite2 = lstComponentTypes.begin();
// 	LST_INT::const_iterator iteEnd2 = lstComponentTypes.end();
// 	for( ; ite2 != iteEnd2; ++ite2)
// 	{
// 		std::wcout<<*ite2<<"\n";
// 	}
    
// 	std::wcout<<"\nTesting Utils::TokenizeString\n";
// 	sStr = "<root>\n<errors>\t\t<error level=1 desc=\"error1\"/>\t\t<error level=2 desc=\"error2\"/>\t</errors></root>";
// 	lstTokens.clear();
// 	lstTokens.push_back(" ");
// 	lstTokens.push_back("\t");
// 	lstTokens.push_back("<");
// 	lstTokens.push_back(">");
// 	lstTokens.push_back("\n");
// 	Utils::TokenizeString(sStr, lstTokens, lstSep, lstStr);
// 	std::wcout<<"String to tokenize: "<<&sStr<<"\n\n";
// 	std::wcout<<"Seperators\n";
// 	ite1 = lstSep.begin();
// 	iteEnd1 = lstSep.end();
// 	for( ; ite1 != iteEnd1; ++ite1)
// 	{
// 		std::wcout<<&*ite1<<"\n";
// 	}
// 	std::wcout<<"\nStrings\n";
// 	ite1 = lstStr.begin();
// 	iteEnd1 = lstStr.end();
// 	for( ; ite1 != iteEnd1; ++ite1)
// 	{
// 		std::wcout<<&*ite1<<"\n";
// 	}
    
// 	// Another string
// 	sStr = "$X.LeftSibling.GetValue(cc)=Y";
// 	lstTokens.clear();
// 	lstStr.clear();
// 	lstSep.clear();
// 	lstTokens.push_back(".");
// 	lstTokens.push_back("(");
// 	lstTokens.push_back(")");
// 	lstTokens.push_back("=");
// 	Utils::TokenizeString(sStr, lstTokens, lstSep, lstStr);
// 	std::wcout<<"String to tokenize: "<<&sStr<<"\n\n";
// 	std::wcout<<"Seperators\n";
// 	ite1 = lstSep.begin();
// 	iteEnd1 = lstSep.end();
// 	for( ; ite1 != iteEnd1; ++ite1)
// 	{
// 		std::wcout<<&*ite1<<"\n";
// 	}
// 	std::wcout<<"\nStrings\n";
// 	ite1 = lstStr.begin();
// 	iteEnd1 = lstStr.end();
// 	for( ; ite1 != iteEnd1; ++ite1)
// 	{
// 		std::wcout<<&*ite1<<"\n";
// 	}
    
// 	return 1;
// }

// int RunUnitTest1()
// {
// 	Tests tests;
// 	tests.RunTest1();
// 	return 1;
// }

// int RunUnitTest2()
// {
// 	Tests tests;
// 	tests.RunTest2();
// 	return 1;
// }

// int RunUnitTest3()
// {
// 	Tests tests;
// 	tests.RunTest3();
// 	return 1;
// }

// void CPlusPlusEntry::RunMenu()
// {
// 	Menu* pMain = new Menu;
// 	{
// 		Menu* pRunTestsMenu = pMain->AddMenu("Run Tests");
// 		{
// 			Menu* pTestUtilsMenu = pRunTestsMenu->AddMenu("Test Utils");
// 			{
// 				pTestUtilsMenu->AddFunction("Test Tokenize", TestTokenize);
// 			}
// 			Menu* pTestScriptMenu = pMain->AddMenu("Test Script");
// 			{
// 				pTestScriptMenu->AddFunction("test 1", RunUnitTest1);
// 				pTestScriptMenu->AddFunction("test 2", RunUnitTest2);
// 				pTestScriptMenu->AddFunction("test 3", RunUnitTest3);
// 			}
// 		}
// 		pMain->AddFunction("Quit", Quit);
// 	}
// 	while(true)
// 	{
// 		if(0 == pMain->Execute())
// 		{
// 			break;
// 		}
// 	}
// 	pMain->Destroy();	
// }

// void CPlusPlusEntry::RunDefault() {
//     //RunUnitTest3();
// }


