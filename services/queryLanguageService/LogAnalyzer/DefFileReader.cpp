#include "DefFileReader.h"
#include "MetaData.h"
#include "MemMan.h"
#include "Utils.h"

MetaData* DefFileReader::Read(MSTRING sFile)
{
	MetaData* pMD = 0;
	MIFSTREAM file(sFile.c_str());
	MSTRING sLine;
	if(file.is_open())
	{
		MemoryManager::Inst.CreateObject(&pMD);
		while(!file.eof())
		{
			getline(file, sLine);
			ProcessLine(sLine, pMD);
		}
		file.close();
	}
    ModifyFilePathsIfNeeded(pMD, sFile);
	return pMD;
}

void DefFileReader::ModifyFilePathsIfNeeded(MetaData *md, MSTRING sDefFilePath)
{
    // This methid checks whether a path is specified for Rule file, script file, log file & result file
    // If a path is not given, this method prepends the def file folder name so that all other files are considered to be in the same folder as the def file
    
    // Extract the folder compoment from def file path
    MSTRING::size_type pos = sDefFilePath.find_last_of(_MSTR(/\\));
    if (pos == MSTRING::npos) {
        return;
    }
    MSTRING folder = sDefFilePath.substr(0, pos + 1);
    
    PrependFolderIfNeeded(folder, md->s_RuleFileName);
    PrependFolderIfNeeded(folder, md->s_ScriptFile);
    PrependFolderIfNeeded(folder, md->s_ResultFile);
    PrependFolderIfNeeded(folder, md->s_CodeLibraryFile);
    PrependFolderToLogFilesIfNeeded(folder, md->s_LogFile);
}

void DefFileReader::PrependFolderToLogFilesIfNeeded(MSTRING folder, MSTRING& logFiles) {
    MSTRING ret = EMPTY_STRING;
    LST_STR filenames;
    LST_STR sep;
    LST_STR sepout;
    sep.push_back(COMMA);
    Utils::TokenizeString(logFiles, sep, sepout, filenames);
    LST_STR::iterator filenamesEnd = filenames.end();
    for (LST_STR::iterator ite = filenames.begin(); ite != filenamesEnd; ++ite) {
        MSTRING filename = *ite;
        Utils::TrimLeft(filename, _MSTR(\t \t));
        Utils::TrimRight(filename, _MSTR(\t \t));
        PrependFolderIfNeeded(folder, filename);
        if (ret.length() > 0) {
            ret += COMMA;
        }
        ret += filename;
    }
    logFiles = ret;
}

void DefFileReader::PrependFolderIfNeeded(MSTRING folder, MSTRING &file)
{
    if (file.find_first_of(_MSTR(/\\)) != MSTRING::npos) {
        return;
    }
    file = folder + file;
}

void DefFileReader::ProcessLine(MSTRING& sLine, MetaData* pMD)
{
    // Example for Line format
	// DEF PARENT Parent	#This is a comment
	// Any line that does not start with DEF is a comment line
    
	Utils::TrimLeft(sLine, _MSTR( \t));
	Utils::TrimRight(sLine, _MSTR( \t));
    
    MSTRING::size_type len = sLine.length();
    if (len <= 3) {
        return;
    }
    MSTRING def = sLine.substr(0, 3);
    Utils::MakeUpper(def);
    if (def != _MSTR(DEF)) {
        return;
    }
    MSTRING keyval = sLine.substr(3, len - 3);
    Utils::TrimLeft(keyval, _MSTR(\t \t));
    MSTRING::size_type pos = keyval.find_first_of(_MSTR(\t \t));
    if (pos == MSTRING::npos) {
        return;
    }
    MSTRING sKey = keyval.substr(0, pos);
    MSTRING sVal = keyval.substr(pos, keyval.length() - pos);
    Utils::TrimLeft(sVal, _MSTR(\t \t));
    
    //	LST_STR lstTokens;
    //	lstTokens.push_back(SPACE);
    //	lstTokens.push_back(_MSTR(\t));
    //	LST_STR lstSep;
    //	LST_STR lstVal;
    //	Utils::TokenizeString(sLine, lstTokens, lstSep, lstVal);
    //	if(lstVal.size() < 3)
    //	{
    //		return;
    //	}
    //
    //	LST_STR::const_iterator ite = lstVal.begin();
    //	MSTRING sStr = *ite;
    //	Utils::MakeUpper(sStr);
    //	if(sStr != _MSTR(DEF))
    //	{
    //		return;
    //	}
    //	++ite;
    //	MSTRING sKey = *ite;
    //	++ite;
    //	MSTRING sVal = *ite;
	AddKeyAndValue(pMD, sKey, sVal);
}

// We use the following macro to ease typing of else if blocks for function names inside the function AddKeyAndValue
#define ADD_FUNC_NAME_FIRST(X) if(_MSTR(X) == sKey){pMD->map_FuncNames[COMMAND_TYPE_##X] = sVal;pMD->map_FuncNamesReverse[sVal] = COMMAND_TYPE_##X;}
#define ADD_FUNC_NAME(X) else if(_MSTR(X) == sKey){pMD->map_FuncNames[COMMAND_TYPE_##X] = sVal;pMD->map_FuncNamesReverse[sVal] = COMMAND_TYPE_##X;}

void DefFileReader::AddKeyAndValue(MetaData* pMD, MSTRING sKey, MSTRING sVal)
{
	if(_MSTR(RULE_FILE_NAME) == sKey)
	{
		pMD->s_RuleFileName = sVal;
	}
	else if(_MSTR(LINE_CONTINUATION_STRING) == sKey)
	{
		pMD->s_LineContinuation = sVal;
	}
	else if(_MSTR(FUNCTION_SEPERATOR) == sKey)
	{
		pMD->s_FuncSeperator = sVal;
	}
	else if(_MSTR(EQUAL_SIGN) == sKey)
	{
		pMD->s_EqualSign = sVal;
	}
	else if(_MSTR(COMMENT_START) == sKey)
	{
		pMD->s_CommentStart = sVal;
	}
	else if(_MSTR(ARGUMENT_START) == sKey)
	{
		pMD->s_ArgumentStart = sVal;
	}
	else if(_MSTR(ARGUMENT_END) == sKey)
	{
		pMD->s_ArgumentEnd = sVal;
	}
	else if(_MSTR(LIST_START) == sKey)
	{
		pMD->s_ListStart = sVal;
	}
	else if(_MSTR(LIST_END) == sKey)
	{
		pMD->s_ListEnd = sVal;
	}
	else if(_MSTR(LIST_ELEMENT_SEPERATOR) == sKey)
	{
		pMD->s_ListElementSeperator = sVal;
	}
	else if(_MSTR(FUNCTION_START) == sKey)
	{
		pMD->s_FuncStart = sVal;
	}
	else if(_MSTR(FUNCTION_END) == sKey)
	{
		pMD->s_FuncEnd = sVal;
	}
	else if(_MSTR(FUNCTION_ARGUMENT) == sKey)
	{
		pMD->s_FuncArg = sVal;
	}
	else if(_MSTR(FUNCTION_RETURN_VALUE) == sKey)
	{
		pMD->s_FuncRet = sVal;
	}
	else if(_MSTR(VARIABLE_NAME_PREFIX) == sKey)
	{
		pMD->s_VarNamePrefix = sVal;
	}
	else if(_MSTR(INTEGER_PREFIX) == sKey)
	{
		pMD->s_IntPrefix = sVal;
	}
    else if(_MSTR(STRING_ENCLOSURE_SYMBOL) == sKey)
    {
        pMD->s_StringEnclosureSymbol = sVal;
    }
	else if(_MSTR(BOOL_TRUE) == sKey)
	{
		pMD->s_BoolTrue = sVal;
	}
	else if(_MSTR(BOOL_FALSE) == sKey)
	{
		pMD->s_BoolFalse = sVal;
	}
	else if(_MSTR(IF) == sKey)
	{
		pMD->s_If = sVal;
	}
	else if(_MSTR(IFNOT) == sKey)
	{
		pMD->s_IfNot = sVal;
	}
	else if(_MSTR(ENDIF) == sKey)
	{
		pMD->s_EndIf = sVal;
	}
	else if(_MSTR(WHILE) == sKey)
	{
		pMD->s_While = sVal;
	}
	else if(_MSTR(DO) == sKey)
	{
		pMD->s_Do = sVal;
	}
	else if(_MSTR(BREAK) == sKey)
	{
		pMD->s_Break = sVal;
	}
	else if(_MSTR(CONTINUE) == sKey)
	{
		pMD->s_Continue = sVal;
	}
    else if(_MSTR(LIST_ITEM_VAR) == sKey)
    {
        pMD->s_ListItemVar = sVal;
    }
    else if(_MSTR(CODE_LIBRARY_FILE) == sKey)
    {
        pMD->s_CodeLibraryFile = sVal;
    }
    else if(_MSTR(LOAD_FROM_CODE_LIBRARY) == sKey)
    {
        pMD->s_LoadFromCodeLibrary = sVal;
    } else if(_MSTR(GET_NODE_OBJ) == sKey){
        pMD->s_GetNodeObj = sVal;
    }
    // LDEL
    else if(_MSTR(LDEL_SCRIPT_FILE) == sKey) {
        pMD->s_ScriptFile = sVal;
    }
    else if(_MSTR(LDEL_LOG_FILE) == sKey) {
        pMD->s_LogFile = sVal;
    }
     //LDAL TREE LOCATION
    else if(_MSTR(TREE_LOCATION)== sKey)
    {
        pMD->s_TREELocation = sVal;
    }

    //read result file
    else if(_MSTR(LDEL_RESULT_FILE) == sKey) {
        pMD->s_ResultFile = sVal;
    }

	 //read result file
    else if(_MSTR(LDEL_RESULT_JSONFILE) == sKey) {
        pMD->s_ResultJSONFile = sVal;
    }

    //Debug JSON location
    else if(_MSTR(DEBUG_JSON)==sKey){
        pMD->s_DebugJSON_File = sVal;
    }

    else if(_MSTR(LDEL_ASSIGNMENT) == sKey)
	{
		pMD->s_ELAssignment = sVal;
	}
    else if(_MSTR(LDEL_VARIABLE_PREFIX) == sKey)
	{
		pMD->s_ELVarPrefix = sVal;
	}
    else if(_MSTR(LDEL_LINE_TEMPLATE_PREFIX) == sKey)
	{
		pMD->s_ELLineTemplatePrefix = sVal;
	}
    else if(_MSTR(LDEL_BLOCK_TEMPLATE_PREFIX) == sKey)
	{
		pMD->s_ELBlockTemplatePrefix = sVal;
	}
    else if(_MSTR(LDEL_NUMBER) == sKey)
	{
		pMD->s_ELNumber = sVal;
	}
    else if(_MSTR(LDEL_FORMATTED_NUMBER) == sKey)
	{
		pMD->S_ELFormattedNumber = sVal;
	}
    else if(_MSTR(LDEL_STRING) == sKey)
	{
		pMD->s_ELString = sVal;
	}
    else if(_MSTR(LDEL_TEXT) == sKey)
	{
		pMD->s_ELText = sVal;
	}
    else if(_MSTR(LDEL_TRIMMEDTEXT) == sKey)
    {
        pMD->s_ELTrimmedText = sVal;
    }
    else if(_MSTR(LDEL_FILEPATH) == sKey)
    {
        pMD->s_ELFilePath = sVal;
    }
    else if(_MSTR(LDEL_SPACE_STRING) == sKey)
	{
		pMD->s_ELSpacesString = sVal;
	}
    else if(_MSTR(LDEL_FLOAT) == sKey)
	{
		pMD->s_ELFloat = sVal;
	}
    else if(_MSTR(LDEL_FORMATTED_FLOAT) == sKey)
	{
		pMD->s_ELFormattedFloat = sVal;
	}
    else if(_MSTR(LDEL_TIMESTAMP) == sKey)
	{
		pMD->s_ELTimestamp = sVal;
	}
    else if(_MSTR(LDEL_VAR_SEQUENCE_START) == sKey)
	{
		pMD->s_ELVarSequenceStart = sVal;
	}
    else if(_MSTR(LDEL_VAR_SEQUENCE_END) == sKey)
	{
		pMD->s_ELVarSequenceEnd = sVal;
	}
    else if(_MSTR(LDEL_VAR_FLEXISEQUENCE_START) == sKey)
	{
		pMD->s_ELVarFlexiSequenceStart = sVal;
	}
    else if(_MSTR(LDEL_VAR_FLEXISEQUENCE_END) == sKey)
	{
		pMD->s_ELVarFlexiSequenceEnd = sVal;
	}
    else if(_MSTR(LDEL_VAR_SUPERFLEXISEQUENCE_START) == sKey)
	{
		pMD->s_ELVarSuperFlexiSequenceStart = sVal;
	}
    else if(_MSTR(LDEL_VAR_SUPERFLEXISEQUENCE_END) == sKey)
	{
		pMD->s_ELVarSuperFlexiSequenceEnd = sVal;
	}
    else if(_MSTR(LDEL_VAR_SEQUENCE_SEPARATOR) == sKey)
	{
		pMD->s_ELVarSequenceSeperator = sVal;
	}
    else if(_MSTR(LDEL_STRING_LITERAL_START) == sKey)
	{
		pMD->s_ELStringLiteralStart = sVal;
	}
    else if(_MSTR(LDEL_STRING_LITERAL_END) == sKey)
	{
		pMD->s_ELStringLiteralEnd = sVal;
	}
    else if(_MSTR(LDEL_STRING_LITERAL_ESCAPE) == sKey)
	{
		pMD->s_ELStringLiteralEscape = sVal;
	}
    else if(_MSTR(LDEL_SET_START) == sKey)
	{
		pMD->s_ELSetStart = sVal;
	}
    else if(_MSTR(LDEL_SET_END) == sKey)
	{
		pMD->s_ELSetEnd = sVal;
	}
    else if(_MSTR(LDEL_SET_SEPARATOR) == sKey)
	{
		pMD->s_ELSetSeperator = sVal;
	}
    else if(_MSTR(LDEL_SEQUENCE_VAR_START_INDICATOR) == sKey)
	{
		pMD->s_ELSequenceVarStartIndicator = sVal;
	}
    else if(_MSTR(LDEL_SEQUENCE_VAR_SUFFIX) == sKey)
	{
		pMD->s_ELSequenceVarSuffix = sVal;
	}
    else if(_MSTR(LDEL_IMPORT_LINE_START) == sKey)
    {
        pMD->s_ELImportLineStart = sVal;
    }
    else if(_MSTR(LDEL_IMPORT_FILE_START) == sKey)
    {
        pMD->s_ELImportFileStart = sVal;
    }
    else if(_MSTR(LDEL_IMPORT_FILE_END) == sKey)
    {
        pMD->s_ELImportFileEnd = sVal;
    }
    else if(_MSTR(LDEL_IGNORE_EMPTY_LINES) == sKey)
    {
        Utils::MakeLower(sVal);
        pMD->s_ELIgnoreEmptyLines = sVal;
    }
    else if(_MSTR(LDEL_NUMBER_FORMAT_EUROPEAN) == sKey)
    {
        Utils::MakeLower(sVal);
        pMD->s_ELIsNumberFormatEuropean = sVal;
    }
    else if(_MSTR(LDEL_IGNORE_TEXT) == sKey)
    {
        pMD->s_ELIgnoreText = sVal;
    }
    else if(_MSTR(LDEL_ANY_TEXT) == sKey)
    {
        pMD->s_ELAnyText = sVal;
    }
    else if(_MSTR(LDEL_DESCRIPTIVE_VAR_ENCLOSER_START) == sKey)
    {
        pMD->s_ELDescriptiveVarEncloserStart = sVal;
    }
    else if(_MSTR(LDEL_DESCRIPTIVE_VAR_ENCLOSER_END) == sKey)
    {
        pMD->s_ELDescriptiveVarEncloderEnd = sVal;
    }
    else if(_MSTR(LDEL_DESCRIPTIVE_VAR_PROPERTY_SEPARATOR) == sKey)
    {
        pMD->s_ELDescriptiveVarPropertySeparator = sVal;
    }
    else if(_MSTR(LDEL_DESCRIPTIVE_VAR_PROPERTY_ASSIGNMENT) == sKey)
    {
        pMD->s_ELDescriptiveVarPropertyAssignment = sVal;
    }
    else if(_MSTR(LDEL_DESCRIPTIVE_VAR_PROPERTY_FORMAT) == sKey)
    {
        pMD->s_ELDescriptiveVarPropertyFormat = sVal;
    }
    // LDEL
    
    else
    {
        AddFuncNames(pMD, sKey, sVal);
    }
}

void DefFileReader::AddFuncNames(MetaData* pMD, MSTRING sKey, MSTRING sVal) {
    ADD_FUNC_NAME_FIRST(LEFT_SIBLING)
	ADD_FUNC_NAME(RIGHT_SIBLING)
	ADD_FUNC_NAME(PARENT)
	ADD_FUNC_NAME(FIRST_CHILD)
	ADD_FUNC_NAME(CHILDREN)
	ADD_FUNC_NAME(CHILD_COUNT)
	ADD_FUNC_NAME(GET_VALUE)
	ADD_FUNC_NAME(GET_LVALUE)
	ADD_FUNC_NAME(GET_RVALUE)
	ADD_FUNC_NAME(GET_CUSTOM_STRING)
	ADD_FUNC_NAME(GET_ID)
	ADD_FUNC_NAME(GET_TYPE)
	ADD_FUNC_NAME(GET_NATURE)
	ADD_FUNC_NAME(GET_WEIGHT)
	ADD_FUNC_NAME(GET_MIN_CHILD_WEIGHT)
	ADD_FUNC_NAME(GET_MAX_CHILD_WEIGHT)
	ADD_FUNC_NAME(SET_VALUE)
	ADD_FUNC_NAME(SET_LVALUE)
	ADD_FUNC_NAME(SET_RVALUE)
	ADD_FUNC_NAME(SET_TYPE)
	ADD_FUNC_NAME(SET_NATURE)
	ADD_FUNC_NAME(SET_CUSTOM_STRING)
	ADD_FUNC_NAME(SET_MIN_CHILD_WEIGHT)
	ADD_FUNC_NAME(SET_MAX_CHILD_WEIGHT)
	ADD_FUNC_NAME(SET_WEIGHT)
	ADD_FUNC_NAME(EXPAND)
	ADD_FUNC_NAME(ADD_NODE)
	ADD_FUNC_NAME(ADD_NODE_WITH_WEIGHT)
	ADD_FUNC_NAME(READ_FROM_FILE)
	ADD_FUNC_NAME(GET_AGGREGATED_VALUE)
	ADD_FUNC_NAME(GET_SUBTREE)
    ADD_FUNC_NAME(FILTER_SUBTREE)
	ADD_FUNC_NAME(IS_TYPE)
	ADD_FUNC_NAME(IS_VALUE)
	ADD_FUNC_NAME(GET_CHILD_OF_TYPE)
	ADD_FUNC_NAME(IS_STRING_EQUAL_TO)
	ADD_FUNC_NAME(IS_STRING_MEMBER_OF)
	ADD_FUNC_NAME(IS_HAVING_SUBSTRING)
	ADD_FUNC_NAME(IS_HAVING_LEFT_SUBSTRING)
	ADD_FUNC_NAME(IS_HAVING_RIGHT_SUBSTRING)
	ADD_FUNC_NAME(ADD_PREFIX)
	ADD_FUNC_NAME(ADD_POSTFIX)
	ADD_FUNC_NAME(TRIM_LEFT)
	ADD_FUNC_NAME(TRIM_RIGHT)
	ADD_FUNC_NAME(WRITE_TO_FILE)
    ADD_FUNC_NAME(STRINGTOBOOL)
	ADD_FUNC_NAME(GET_LENGTH)
	ADD_FUNC_NAME(IS_INT_EQUAL_TO)
	ADD_FUNC_NAME(IS_INT_MEMBER_OF)
	ADD_FUNC_NAME(IS_LESS_THAN)
	ADD_FUNC_NAME(IS_LESS_THAN_OR_EQUAL_TO)
	ADD_FUNC_NAME(IS_GREATER_THAN)
	ADD_FUNC_NAME(IS_GREATER_THAN_OR_EQUAL_TO)
	ADD_FUNC_NAME(ADD)
	ADD_FUNC_NAME(SUBTRACT)
	ADD_FUNC_NAME(TOSTRING)
    ADD_FUNC_NAME(PERCENTAGE)
	ADD_FUNC_NAME(GET_ITEM_COUNT)
    ADD_FUNC_NAME(LIST_FILTER)
    ADD_FUNC_NAME(LIST_GROUPBY)
    ADD_FUNC_NAME(LIST_GROUP_SEQUENCE_BY)
    ADD_FUNC_NAME(GET_INNER_ITEM_COUNT)
	ADD_FUNC_NAME(SEEK)
	ADD_FUNC_NAME(SEEK_TO_BEGIN)
	ADD_FUNC_NAME(SEEK_TO_END)
	ADD_FUNC_NAME(GET_CURR_ELEM)
	ADD_FUNC_NAME(IS_NULL)
	ADD_FUNC_NAME(IS_NOT_NULL)
	ADD_FUNC_NAME(BOOL_AND)
	ADD_FUNC_NAME(BOOL_OR)
    ADD_FUNC_NAME(SET_ENTITY_OBJECT)
    ADD_FUNC_NAME(SECONDS_TO_MONTHS)
    ADD_FUNC_NAME(SECONDS_TO_DAYS)
    ADD_FUNC_NAME(SECONDS_TO_YEARS)
    ADD_FUNC_NAME(GET_DIFFERENCE_BY_STRING)
    ADD_FUNC_NAME(STRING_TO_READABLE_DATETIME)
    ADD_FUNC_NAME(DATE_NOW)
    ADD_FUNC_NAME(STRING_TO_UNIX_TIME)
    ADD_FUNC_NAME(GET_NEXT_ELEM)
    ADD_FUNC_NAME(CHECK_NOT_NULL)
    ADD_FUNC_NAME(GET_UNIQUE_NODE_LIST_WITH_COUNT)
    ADD_FUNC_NAME(STRINGTOINTEGER)
    ADD_FUNC_NAME(GET_STRING)
    ADD_FUNC_NAME(GET_INTEGER)
    ADD_FUNC_NAME(STRINGTOBOOLEAN)
	ADD_FUNC_NAME(GET_COMMA)
	ADD_FUNC_NAME(GET_BOOLEAN)
	ADD_FUNC_NAME(SET_BOOL)
	ADD_FUNC_NAME(TO_FALSE)
	ADD_FUNC_NAME(TO_TRUE)
    ADD_FUNC_NAME(NEXT_SIBLING)
	ADD_FUNC_NAME(SORT_NODE_LIST)
    ADD_FUNC_NAME(EXTRACT_NODE_LIST_TOP)
	ADD_FUNC_NAME(GET_CUSTOM_OBJ)
	ADD_FUNC_NAME(CONVERT_TO_SENTENCE_CASE)
	ADD_FUNC_NAME(GET_DAY_OF_THE_WEEK_SHORT_STRING)
	ADD_FUNC_NAME(GET_DAY_STRING)
	ADD_FUNC_NAME(GET_MONTH_SHORT_STRING)
	ADD_FUNC_NAME(GET_TIME_24_HOUR_FORMAT)
	ADD_FUNC_NAME(GET_YEAR)
	ADD_FUNC_NAME(GET_OLDEST_DATE)
	ADD_FUNC_NAME(GET_LATEST_DATE)
	ADD_FUNC_NAME(ADD_PERIOD)
	ADD_FUNC_NAME(GET_UNIQUE_NODE_LIST_WITH_NODE_REF)
	ADD_FUNC_NAME(SET_INTEGER)
    ADD_FUNC_NAME(SET_ATTRIBUTES)
    ADD_FUNC_NAME(ADD_INNER_OBJ)
}
