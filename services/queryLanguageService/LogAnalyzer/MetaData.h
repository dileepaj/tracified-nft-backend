#ifndef _METADATA_H
#define _METADATA_H

#include "CommonIncludes.h"

class MetaData
{
public:
	MSTRING s_RuleFileName;
	MSTRING	s_FuncSeperator;
	MSTRING s_EqualSign;
	MSTRING s_LineContinuation;
	MSTRING s_CommentStart;
	MSTRING s_ArgumentStart;
	MSTRING s_ArgumentEnd;
	MSTRING s_ListStart;
	MSTRING s_ListEnd;
	MSTRING s_ListElementSeperator;
	MSTRING s_FuncStart;
	MSTRING s_FuncEnd;
	MSTRING s_FuncArg;
	MSTRING s_FuncRet;
	MSTRING s_VarNamePrefix;
	MSTRING s_IntPrefix;
    MSTRING s_StringEnclosureSymbol;
	MSTRING s_BoolTrue;
	MSTRING s_BoolFalse;
	MSTRING s_If;
	MSTRING s_IfNot;
	MSTRING s_EndIf;
	MSTRING s_While;
	MSTRING s_Do;
	MSTRING s_Break;
	MSTRING s_Continue;
    MSTRING s_ListItemVar;
    
	MAP_MULONG_STR	map_FuncNames;
	MAP_STR_MULONG	map_FuncNamesReverse;
    MSTRING s_ResultJSONFile;
    MSTRING s_TREELocation;
    MSTRING s_DebugJSON_File;
    MSTRING s_GetNodeObj;
    // Added for LDEL
    MSTRING s_ScriptFile;
    MSTRING s_LogFile;
    MSTRING s_ResultFile;
    MSTRING s_CodeLibraryFile;
    MSTRING s_LoadFromCodeLibrary;
    MSTRING s_ELAssignment;
    MSTRING s_ELVarPrefix;
    MSTRING s_ELLineTemplatePrefix;
    MSTRING s_ELBlockTemplatePrefix;
    MSTRING s_ELNumber;
    MSTRING S_ELFormattedNumber;
    MSTRING s_ELString;
    MSTRING s_ELText;
    MSTRING s_ELTrimmedText;
    MSTRING s_ELFilePath;
    MSTRING s_ELSpacesString;
    MSTRING s_ELFloat;
    MSTRING s_ELFormattedFloat;
    MSTRING s_ELTimestamp;
    MSTRING s_ELVarSequenceStart;
    MSTRING s_ELVarSequenceEnd;
    MSTRING s_ELVarSuperFlexiSequenceStart;
    MSTRING s_ELVarSuperFlexiSequenceEnd;
    MSTRING s_ELVarFlexiSequenceStart;
    MSTRING s_ELVarFlexiSequenceEnd;
    MSTRING s_ELVarSequenceSeperator;
    MSTRING s_ELStringLiteralStart;
    MSTRING s_ELStringLiteralEnd;
    MSTRING s_ELStringLiteralEscape;
    MSTRING s_ELSetStart;
    MSTRING s_ELSetEnd;
    MSTRING s_ELSetSeperator;
    MSTRING s_ELSequenceVarStartIndicator;
    MSTRING s_ELSequenceVarSuffix;
    MSTRING s_ELImportLineStart;
    MSTRING s_ELImportFileStart;
    MSTRING s_ELImportFileEnd;
    MSTRING s_ELIgnoreEmptyLines;
    MSTRING s_ELIsNumberFormatEuropean;
    MSTRING s_ELIgnoreText;
    MSTRING s_ELAnyText;
    MSTRING s_ELDescriptiveVarEncloserStart;
    MSTRING s_ELDescriptiveVarEncloderEnd;
    MSTRING s_ELDescriptiveVarPropertySeparator;
    MSTRING s_ELDescriptiveVarPropertyAssignment;
    MSTRING s_ELDescriptiveVarPropertyFormat;
};

#endif