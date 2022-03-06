/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////
/////////////////////////  Defs.h ///////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////
/////// This file contains definitions that are used in all other modules ///////
/////////////////////////////////////////////////////////////////////////////////
////////////////////////// 22/08/2006 ///////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

// Remove if not stdafx.h is used
#define		MFC_PROJECT

//Error Codes
#define		ERR_NONE										0
#define		ERR_NODE_DOES_NOT_EXIST							1
#define		ERR_TRANSITION_ALREADY_EXISTS					2
#define		ERR_NO_SUCH_TRANSITION							3
#define		ERR_STARTING_NODE_ALREADY_EXISTS				4
#define		ERR_NODE_ALREADY_EXISTS							5
#define		ERR_STARTING_NODE_DOES_NOT_EXIST				6
#define		ERR_STRING_PARSING_FAILED						7
#define		ERR_NDFA_DOES_NOT_HAVE_FINAL_STATES				8
#define		ERR_NDFA_HAS_MORE_THAN_ONE_FINAL_STATE			9
#define		ERR_TRANSITIONS_EXIST_FROM_FINAL_STATE			10

#define		ERR_ENTITY_IS_NULL								100
#define		ERR_LIST_IS_EMPTY								101
#define		ERR_LEVEL_INCORRECT								102
#define		ERR_ELEMENT_DOES_NOT_EXIST						103
#define		ERR_ENTITY_DOES_NOT_BELONG_TO_TREE				104

//Special Characters
#define		EPSILON											'\0'

//log file names
#define		LOG_FILE_NAME									"Details.txt"

//Operation Argument Types
#define		OPERATION_ARGUMENT_TYPE_OPERATIONS_LIST			0
#define		OPERATION_ARGUMENT_TYPE_STRING					1


//Defs used in CEntityTreeLite
#define		ENTITY_TREE_LITE_INVALID_ENTITY					0
#define		ENTITY_TREE_LITE_INVALID_LEVEL					-5
#define		ENTITY_TREE_LITE_LOWEST_ENTITY_LEVEL			-1


////////////////////////////////////////////////////////////////////////////
/////// Entity Operation Codes /////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

//Operation code types
#define		ENTITY_OPERATION_TYPE_TRANSFORMATION						1
#define		ENTITY_OPERATION_TYPE_EVALUATE								2
#define		ENTITY_OPERATION_TYPE_CHANGE								3
#define		ENTITY_OPERATION_TYPE_GET									4

//Transformation OpCodes	:	range from -100 to -149
#define		ENTITY_OPERATION_LEFT_SIBLING								-100
#define		ENTITY_OPERATION_RIGHT_SIBLING								-101
#define		ENTITY_OPERATION_PARENT										-102
#define		ENTITY_OPERATION_FIRST_CHILD								-103
#define		ENTITY_OPERATION_LAST_CHILD									-104
#define		ENTITY_OPERATION_N_TH_CHILD									-105
#define		ENTITY_OPERATION_CURRENT_ENTITY								-106

//Evaluation OpCodes		:	range from -150 to -199
#define		ENTITY_OPERATION_MEMBER_OF									-150
#define		ENTITY_OPERATION_NOT_MEMBER_OF								-151
#define		ENTITY_OPERATION_TYPE_OF									-152
#define		ENTITY_OPERATION_CATEGORY_OF								-153
#define		ENTITY_OPERATION_FULL_STRING_MEMBER_OF						-154
#define		ENTITY_OPERATION_FULL_STRING_NOT_MEMBER_OF					-155
#define		ENTITY_OPERATION_LEFT_SPECIAL_STRING_MEMBER_OF				-156
#define		ENTITY_OPERATION_RIGHT_SPECIAL_STRING_MEMBER_OF				-157
#define		ENTITY_OPERATION_LEFT_SPECIAL_STRING_NOT_MEMBER_OF			-158
#define		ENTITY_OPERATION_RIGHT_SPECIAL_STRING_NOT_MEMBER_OF			-159
#define		ENTITY_OPERATION_EXIST										-160
#define		ENTITY_OPERATION_NOT_EXIST									-161
#define		ENTITY_OPERATION_LEFT_PART_OF_STRING_MEMBER_OF				-162
#define		ENTITY_OPERATION_RIGHT_PART_OF_STRING_MEMBER_OF				-163
#define		ENTITY_OPERATION_STRING_CONTAIN								-164
#define		ENTITY_OPERATION_LEFT_PART_OF_FULL_STRING_MEMBER_OF			-165
#define		ENTITY_OPERATION_RIGHT_PART_OF_FULL_STRING_MEMBER_OF		-166
#define		ENTITY_OPERATION_FULL_STRING_CONTAIN						-167
#define		ENTITY_OPERATION_LEFT_PART_OF_LEFT_STRING_MEMBER_OF			-168
#define		ENTITY_OPERATION_RIGHT_PART_OF_LEFT_STRING_MEMBER_OF		-169
#define		ENTITY_OPERATION_LEFT_STRING_CONTAIN						-170
#define		ENTITY_OPERATION_LEFT_PART_OF_RIGHT_STRING_MEMBER_OF		-171
#define		ENTITY_OPERATION_RIGHT_PART_OF_RIGHT_STRING_MEMBER_OF		-172
#define		ENTITY_OPERATION_RIGHT_STRING_CONTAIN						-173
#define		ENTITY_OPERATION_VALUE_MEMBER_OF							-174
#define		ENTITY_OPERATION_VALUE_NOT_MEMBER_OF						-175
#define		ENTITY_OPERATION_LEFT_PART_OF_VALUE_MEMBER_OF				-176
#define		ENTITY_OPERATION_RIGHT_PART_OF_VALUE_MEMBER_OF				-177
#define		ENTITY_OPERATION_VALUE_CONTAIN								-178
#define		ENTITY_OPERATION_LEFT_PART_OF_STRING_NOT_MEMBER_OF			-179
#define		ENTITY_OPERATION_RIGHT_PART_OF_STRING_NOT_MEMBER_OF			-180
#define		ENTITY_OPERATION_STRING_NOT_CONTAIN							-181
#define		ENTITY_OPERATION_LEFT_PART_OF_FULL_STRING_NOT_MEMBER_OF		-182
#define		ENTITY_OPERATION_RIGHT_PART_OF_FULL_STRING_NOT_MEMBER_OF	-183
#define		ENTITY_OPERATION_FULL_STRING_NOT_CONTAIN					-184
#define		ENTITY_OPERATION_LEFT_PART_OF_LEFT_STRING_NOT_MEMBER_OF		-185
#define		ENTITY_OPERATION_RIGHT_PART_OF_LEFT_STRING_NOT_MEMBER_OF	-186
#define		ENTITY_OPERATION_LEFT_STRING_NOT_CONTAIN					-187
#define		ENTITY_OPERATION_LEFT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF	-188
#define		ENTITY_OPERATION_RIGHT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF	-189
#define		ENTITY_OPERATION_RIGHT_STRING_NOT_CONTAIN					-190
#define		ENTITY_OPERATION_LEFT_PART_OF_VALUE_NOT_MEMBER_OF			-191
#define		ENTITY_OPERATION_RIGHT_PART_OF_VALUE_NOT_MEMBER_OF			-192
#define		ENTITY_OPERATION_VALUE_NOT_CONTAIN							-193

//Change OpCodes			:	range from -200 to -249
#define		ENTITY_OPERATION_SET_VALUE									-200
#define		ENTITY_OPERATION_SET_TYPE									-201
#define		ENTITY_OPERATION_SET_CATEGORY								-202
#define		ENTITY_OPERATION_SET_LEFT_STRING							-203
#define		ENTITY_OPERATION_SET_RIGHT_STRING							-204

//Get OpCodes				:	range from -250 to -299
#define		ENTITY_OPERATION_GET_VALUE									-250
#define		ENTITY_OPERATION_GET_STRING									-251
#define		ENTITY_OPERATION_GET_STRING_WITH_SPECIAL_STRINGS			-252
#define		ENTITY_OPERATION_GET_LEFT_STRING							-253
#define		ENTITY_OPERATION_GET_RIGHT_STRING							-254

//Flags used in CRuleFactoryEvaluationArgument class
#define		RF_EVAL_ARG_NEW_TREE_FOR_OUTPUT						0x00000001L
#define		RF_EVAL_ARG_PASSED_ENTITIES							0x00000002L

//Defs used in CLargeString
#define		LARGE_STR_MAX_ELEMENT_SIZE									5000
#define		LARGE_STR_MAX_STRING_SIZE									30000

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////
/////////////////// File Reader Defs ////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

//String parse types
#define		PARSE_TYPE_ONLY_KEY_STRING									1
#define		PARSE_TYPE_ANY_STRING_FOLLOWED_BY_KEY_STRING				2
#define		PARSE_TYPE_ANY_STRING										3
#define		PARSE_TYPE_ANY_STRING_ALONE_OR_FOLLOWED_BY_KEY_STRING		4

//KeyString codes
#define		KEY_CODE_DUMMY												-3		//dummy key code for nodes with parse type = PARSE_TYPE_ANY_STRING
#define		KEY_CODE_WRONG_KEYWORD										-2
#define		KEY_CODE_IGNORE												-1		// string to be ignored
#define		KEY_CODE_COMMENT_START										0
#define		KEY_CODE_RULE_START											1
#define		KEY_CODE_RULE_SECTION_START									2
#define		KEY_CODE_RULE_SECTION_END									3
#define		KEY_CODE_ACTION_SECTION_START								4
#define		KEY_CODE_ACTION_SECTION_END									5
#define		KEY_CODE_OUTPUT_SECTION_START								6
#define		KEY_CODE_OUTPUT_SECTION_END									7
#define		KEY_CODE_SET_START											8
#define		KEY_CODE_SET_END											9
#define		KEY_CODE_SET_ELEMENT_SEPERATION								10
#define		KEY_CODE_FUNCTION_ARGUMENT_START							11
#define		KEY_CODE_FUNCTION_ARGUMENT_END								12
#define		KEY_CODE_FUNCTION_ARGUMENT_SEPERATION						13
#define		KEY_CODE_KEYWORD_SEPERATION									14
#define		KEY_CODE_CURRENT_NODE										15

#define		KEY_CODE_LEFT_SIBLING										51
#define		KEY_CODE_RIGHT_SIBLING										52
#define		KEY_CODE_PARENT												53
#define		KEY_CODE_N_TH_CHILD											54
#define		KEY_CODE_FIRST_CHILD										55
#define		KEY_CODE_LAST_CHILD											56

#define		KEY_CODE_IS_MEMBER_OF										100
#define		KEY_CODE_IS_NOT_MEMBER_OF									101
#define		KEY_CODE_IS_EXIST											102
#define		KEY_CODE_IS_NOT_EXIST										103
#define		KEY_CODE_IS_FULL_STRING_MEMBER_OF							104
#define		KEY_CODE_IS_FULL_STRING_NOT_MEMBER_OF						105
#define		KEY_CODE_IS_LEFT_STRING_MEMBER_OF							106
#define		KEY_CODE_IS_LEFT_STRING_NOT_MEMBER_OF						107
#define		KEY_CODE_IS_RIGHT_STRING_MEMBER_OF							108
#define		KEY_CODE_IS_RIGHT_STRING_NOT_MEMBER_OF						109
#define		KEY_CODE_IS_TYPE_OF											110
#define		KEY_CODE_IS_CATEGORY_OF										111
#define		KEY_CODE_IS_LEFT_PART_OF_STRING_MEMBER_OF					112
#define		KEY_CODE_IS_RIGHT_PART_OF_STRING_MEMBER_OF					113
#define		KEY_CODE_IS_STRING_CONTAIN									114
#define		KEY_CODE_IS_LEFT_PART_OF_FULL_STRING_MEMBER_OF				115
#define		KEY_CODE_IS_RIGHT_PART_OF_FULL_STRING_MEMBER_OF				116
#define		KEY_CODE_IS_FULL_STRING_CONTAIN								117
#define		KEY_CODE_IS_LEFT_PART_OF_LEFT_STRING_MEMBER_OF				118
#define		KEY_CODE_IS_RIGHT_PART_OF_LEFT_STRING_MEMBER_OF				119
#define		KEY_CODE_IS_LEFT_STRING_CONTAIN								120
#define		KEY_CODE_IS_LEFT_PART_OF_RIGHT_STRING_MEMBER_OF				121
#define		KEY_CODE_IS_RIGHT_PART_OF_RIGHT_STRING_MEMBER_OF			122
#define		KEY_CODE_IS_RIGHT_STRING_CONTAIN							123
#define		KEY_CODE_IS_VALUE_MEMBER_OF									124
#define		KEY_CODE_IS_VALUE_NOT_MEMBER_OF								125
#define		KEY_CODE_IS_LEFT_PART_OF_VALUE_MEMBER_OF					126
#define		KEY_CODE_IS_RIGHT_PART_OF_VALUE_MEMBER_OF					127
#define		KEY_CODE_IS_VALUE_CONTAIN									128
#define		KEY_CODE_IS_LEFT_PART_OF_STRING_NOT_MEMBER_OF				129
#define		KEY_CODE_IS_RIGHT_PART_OF_STRING_NOT_MEMBER_OF				130
#define		KEY_CODE_IS_STRING_NOT_CONTAIN								131
#define		KEY_CODE_IS_LEFT_PART_OF_FULL_STRING_NOT_MEMBER_OF			132
#define		KEY_CODE_IS_RIGHT_PART_OF_FULL_STRING_NOT_MEMBER_OF			133
#define		KEY_CODE_IS_FULL_STRING_NOT_CONTAIN							134
#define		KEY_CODE_IS_LEFT_PART_OF_LEFT_STRING_NOT_MEMBER_OF			135
#define		KEY_CODE_IS_RIGHT_PART_OF_LEFT_STRING_NOT_MEMBER_OF			136
#define		KEY_CODE_IS_LEFT_STRING_NOT_CONTAIN							137
#define		KEY_CODE_IS_LEFT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF			138
#define		KEY_CODE_IS_RIGHT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF		139
#define		KEY_CODE_IS_RIGHT_STRING_NOT_CONTAIN						140
#define		KEY_CODE_IS_LEFT_PART_OF_VALUE_NOT_MEMBER_OF				141
#define		KEY_CODE_IS_RIGHT_PART_OF_VALUE_NOT_MEMBER_OF				142
#define		KEY_CODE_IS_VALUE_NOT_CONTAIN								143


#define		KEY_CODE_SET_VALUE											200
#define		KEY_CODE_SET_LEFT_STRING									201
#define		KEY_CODE_SET_RIGHT_STRING									202

#define		KEY_CODE_GET_STRING											250
#define		KEY_CODE_GET_STRING_WITH_SPECIAL_STRINGS					251
#define		KEY_CODE_GET_LEFT_STRING									252
#define		KEY_CODE_GET_RIGHT_STRING									253
#define		KEY_CODE_GET_VALUE											254

#define		KEY_CODE_RULE_FILENAME										300
#define		KEY_CODE_MASTER_RULE										301
#define		KEY_CODE_STRING_START										302
#define		KEY_CODE_STRING_END											303

#define		KEY_CODE_CATEGORY_NAME_START_STRING							350
#define		KEY_CODE_CATEGORY_NAME_END_STRING							351
#define		KEY_CODE_CATEGORY_VALUE_SEPERATOR_STRING					352
#define		KEY_CODE_IGNORE_STRING_SEPERATOR_STRING						353
#define		KEY_CODE_IGNORE_STRINGS										354
#define		KEY_CODE_CATEGORY_FILENAME									355

#define		KEY_CODE_BASE_LANGUAGE_FILENAME								375
#define		KEY_CODE_CATEGORY_SEQUENCE									376
#define		KEY_CODE_CATEGORY_SEPERATOR_STRING_IN_CATEGORY_SEQUENCE		377

//keycodes allocated for external classes
#define		KEY_CODE_RULE_FILENAME1										500
#define		KEY_CODE_RULE_FILENAME2										501
#define		KEY_CODE_RULE_FILENAME3										502
#define		KEY_CODE_RULE_FILENAME4										503
#define		KEY_CODE_RULE_FILENAME5										504
#define		KEY_CODE_GENERIC_FILENAME1									505
#define		KEY_CODE_GENERIC_FILENAME2									506
#define		KEY_CODE_GENERIC_FILENAME3									507
#define		KEY_CODE_GENERIC_FILENAME4									508
#define		KEY_CODE_GENERIC_FILENAME5									509
#define		KEY_CODE_GENERIC_ENTITY1									510
#define		KEY_CODE_GENERIC_ENTITY2									511
#define		KEY_CODE_GENERIC_ENTITY3									512
#define		KEY_CODE_GENERIC_ENTITY4									513
#define		KEY_CODE_GENERIC_ENTITY5									514
#define		KEY_CODE_GENERIC_ENTITY6									515
#define		KEY_CODE_GENERIC_ENTITY7									516
#define		KEY_CODE_GENERIC_ENTITY8									517
#define		KEY_CODE_GENERIC_ENTITY9									518
#define		KEY_CODE_GENERIC_ENTITY10									519
#define		KEY_CODE_GENERIC_ENTITY11									520
#define		KEY_CODE_GENERIC_ENTITY12									521
#define		KEY_CODE_GENERIC_ENTITY13									522
#define		KEY_CODE_GENERIC_ENTITY14									523
#define		KEY_CODE_GENERIC_ENTITY15									524
#define		KEY_CODE_GENERIC_ENTITY16									525
#define		KEY_CODE_GENERIC_ENTITY17									526
#define		KEY_CODE_GENERIC_ENTITY18									527
#define		KEY_CODE_GENERIC_ENTITY19									528
#define		KEY_CODE_GENERIC_ENTITY20									529
#define		KEY_CODE_GENERIC_ENTITY21									530
#define		KEY_CODE_GENERIC_ENTITY22									531
#define		KEY_CODE_GENERIC_ENTITY23									532
#define		KEY_CODE_GENERIC_ENTITY24									533
#define		KEY_CODE_GENERIC_ENTITY25									534

//KeyString Identifiers	: These strings are used in Defs file to identify key strings
//Keystrings used in rule files
#define		KEY_STRING_IGNORE											"IGNORE"		// string to be ignored
#define		KEY_STRING_COMMENT_START									"COMMENT_START"
#define		KEY_STRING_RULE_START										"RULE_START"
#define		KEY_STRING_RULE_SECTION_START								"RULE_SECTION_START"
#define		KEY_STRING_RULE_SECTION_END									"RULE_SECTION_END"
#define		KEY_STRING_ACTION_SECTION_START								"ACTION_SECTION_START"
#define		KEY_STRING_ACTION_SECTION_END								"ACTION_SECTION_END"
#define		KEY_STRING_OUTPUT_SECTION_START								"OUTPUT_SECTION_START"
#define		KEY_STRING_OUTPUT_SECTION_END								"OUTPUT_SECTION_END"
#define		KEY_STRING_SET_START										"SET_START"
#define		KEY_STRING_SET_END											"SET_END"
#define		KEY_STRING_SET_ELEMENT_SEPERATION							"SET_ELEMENT_SEPERATION"
#define		KEY_STRING_FUNCTION_ARGUMENT_START							"FUNCTION_ARGUMENT_START"
#define		KEY_STRING_FUNCTION_ARGUMENT_END							"FUNCTION_ARGUMENT_END"
#define		KEY_STRING_FUNCTION_ARGUMENT_SEPERATION						"FUNCTION_ARGUMENT_SEPERATION"
#define		KEY_STRING_KEYWORD_SEPERATION								"KEYWORD_SEPERATION"
#define		KEY_STRING_CURRENT_NODE										"CURRENT_NODE"

#define		KEY_STRING_LEFT_SIBLING										"LEFT_SIBLING"
#define		KEY_STRING_RIGHT_SIBLING									"RIGHT_SIBLING"
#define		KEY_STRING_PARENT											"PARENT"
#define		KEY_STRING_N_TH_CHILD										"N_TH_CHILD"
#define		KEY_STRING_FIRST_CHILD										"FIRST_CHILD"
#define		KEY_STRING_LAST_CHILD										"LAST_CHILD"

#define		KEY_STRING_IS_MEMBER_OF										"IS_MEMBER_OF"
#define		KEY_STRING_IS_NOT_MEMBER_OF									"IS_NOT_MEMBER_OF"
#define		KEY_STRING_IS_EXIST											"IS_EXIST"
#define		KEY_STRING_IS_NOT_EXIST										"IS_NOT_EXIST"
#define		KEY_STRING_IS_FULL_STRING_MEMBER_OF							"IS_FULL_STRING_MEMBER_OF"
#define		KEY_STRING_IS_FULL_STRING_NOT_MEMBER_OF						"IS_FULL_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_LEFT_STRING_MEMBER_OF							"IS_LEFT_STRING_MEMBER_OF"
#define		KEY_STRING_IS_LEFT_STRING_NOT_MEMBER_OF						"IS_LEFT_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_STRING_MEMBER_OF						"IS_RIGHT_STRING_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_STRING_NOT_MEMBER_OF					"IS_RIGHT_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_TYPE_OF										"IS_TYPE_OF"
#define		KEY_STRING_IS_CATEGORY_OF									"IS_CATEGORY_OF"
#define		KEY_STRING_IS_LEFT_PART_OF_STRING_MEMBER_OF					"IS_LEFT_PART_OF_STRING_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_STRING_MEMBER_OF				"IS_RIGHT_PART_OF_STRING_MEMBER_OF"
#define		KEY_STRING_IS_STRING_CONTAIN								"IS_STRING_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_FULL_STRING_MEMBER_OF			"IS_LEFT_PART_OF_FULL_STRING_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_FULL_STRING_MEMBER_OF			"IS_RIGHT_PART_OF_FULL_STRING_MEMBER_OF"
#define		KEY_STRING_IS_FULL_STRING_CONTAIN							"IS_FULL_STRING_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_LEFT_STRING_MEMBER_OF			"IS_LEFT_PART_OF_LEFT_STRING_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_LEFT_STRING_MEMBER_OF			"IS_RIGHT_PART_OF_LEFT_STRING_MEMBER_OF"
#define		KEY_STRING_IS_LEFT_STRING_CONTAIN							"IS_LEFT_STRING_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_RIGHT_STRING_MEMBER_OF			"IS_LEFT_PART_OF_RIGHT_STRING_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_RIGHT_STRING_MEMBER_OF			"IS_RIGHT_PART_OF_RIGHT_STRING_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_STRING_CONTAIN							"IS_RIGHT_STRING_CONTAIN"
#define		KEY_STRING_IS_VALUE_MEMBER_OF								"IS_VALUE_MEMBER_OF"
#define		KEY_STRING_IS_VALUE_NOT_MEMBER_OF							"IS_VALUE_NOT_MEMBER_OF"
#define		KEY_STRING_IS_LEFT_PART_OF_VALUE_MEMBER_OF					"IS_LEFT_PART_OF_VALUE_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_VALUE_MEMBER_OF					"IS_RIGHT_PART_OF_VALUE_MEMBER_OF"
#define		KEY_STRING_IS_VALUE_CONTAIN									"IS_VALUE_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_STRING_NOT_MEMBER_OF				"IS_LEFT_PART_OF_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_STRING_NOT_MEMBER_OF			"IS_RIGHT_PART_OF_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_STRING_NOT_CONTAIN							"IS_STRING_NOT_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_FULL_STRING_NOT_MEMBER_OF		"IS_LEFT_PART_OF_FULL_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_FULL_STRING_NOT_MEMBER_OF		"IS_RIGHT_PART_OF_FULL_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_FULL_STRING_NOT_CONTAIN						"IS_FULL_STRING_NOT_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_LEFT_STRING_NOT_MEMBER_OF		"IS_LEFT_PART_OF_LEFT_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_LEFT_STRING_NOT_MEMBER_OF		"IS_RIGHT_PART_OF_LEFT_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_LEFT_STRING_NOT_CONTAIN						"IS_LEFT_STRING_NOT_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF		"IS_LEFT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF		"IS_RIGHT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_STRING_NOT_CONTAIN						"IS_RIGHT_STRING_NOT_CONTAIN"
#define		KEY_STRING_IS_LEFT_PART_OF_VALUE_NOT_MEMBER_OF				"IS_LEFT_PART_OF_VALUE_NOT_MEMBER_OF"
#define		KEY_STRING_IS_RIGHT_PART_OF_VALUE_NOT_MEMBER_OF				"IS_RIGHT_PART_OF_VALUE_NOT_MEMBER_OF"
#define		KEY_STRING_IS_VALUE_NOT_CONTAIN								"IS_VALUE_NOT_CONTAIN"

#define		KEY_STRING_SET_VALUE										"SET_VALUE"
#define		KEY_STRING_SET_LEFT_STRING									"SET_LEFT_STRING"
#define		KEY_STRING_SET_RIGHT_STRING									"SET_RIGHT_STRING"

#define		KEY_STRING_GET_STRING										"GET_STRING"
#define		KEY_STRING_GET_STRING_WITH_SPECIAL_STRINGS					"GET_STRING_WITHOUT_SPECIAL_STRINGS"
#define		KEY_STRING_GET_LEFT_STRING									"GET_LEFT_STRING"
#define		KEY_STRING_GET_RIGHT_STRING									"GET_RIGHT_STRING"
#define		KEY_STRING_GET_VALUE										"GET_VALUE"

#define		KEY_STRING_RULE_FILENAME									"RULE_FILE_NAME"
#define		KEY_STRING_MASTER_RULE										"MASTER_RULE"
#define		KEY_STRING_STRING_START										"STRING_START"
#define		KEY_STRING_STRING_END										"STRING_END"

//keystrings used in category files
#define		KEY_STRING_CATEGORY_NAME_START_STRING						"CAT_NAME_START_STRING"
#define		KEY_STRING_CATEGORY_NAME_END_STRING							"CAT_NAME_END_STRING"
#define		KEY_STRING_CATEGORY_VALUE_SEPERATOR_STRING					"CAT_VALUE_SEPERATOR_STRING"
#define		KEY_STRING_IGNORE_STRING_SEPERATOR_STRING					"IGNORE_STRING_SEPERATOR"
#define		KEY_STRING_IGNORE_STRINGS									"IGNORE_STRINGS"
#define		KEY_STRING_CATEGORY_FILENAME								"CATEGORY_FILE_NAME"

//keystrings used in base language files
#define		KEY_STRING_BASE_LANGUAGE_FILENAME							"BASE_FILE_NAME"
#define		KEY_STRING_CATEGORY_SEQUENCE								"CATEGORY_SEQUENCE"
#define		KEY_STRING_CATEGORY_SEPERATOR_STRING_IN_CATEGORY_SEQUENCE	"CATEGORY_SEPERATOR_IN_CATEGORY_SEQUENCE"

//keystrings allocated for the use of external classes
#define		KEY_STRING_RULE_FILENAME1									"RULE_FILE_NAME_1"
#define		KEY_STRING_RULE_FILENAME2									"RULE_FILE_NAME_2"
#define		KEY_STRING_RULE_FILENAME3									"RULE_FILE_NAME_3"
#define		KEY_STRING_RULE_FILENAME4									"RULE_FILE_NAME_4"
#define		KEY_STRING_RULE_FILENAME5									"RULE_FILE_NAME_5"
#define		KEY_STRING_GENERIC_FILENAME1								"GENERIC_FILE_NAME_1"
#define		KEY_STRING_GENERIC_FILENAME2								"GENERIC_FILE_NAME_2"
#define		KEY_STRING_GENERIC_FILENAME3								"GENERIC_FILE_NAME_3"
#define		KEY_STRING_GENERIC_FILENAME4								"GENERIC_FILE_NAME_4"
#define		KEY_STRING_GENERIC_FILENAME5								"GENERIC_FILE_NAME_5"
#define		KEY_STRING_GENERIC_ENTITY1									"GENERIC_ENTITY_1"
#define		KEY_STRING_GENERIC_ENTITY2									"GENERIC_ENTITY_2"
#define		KEY_STRING_GENERIC_ENTITY3									"GENERIC_ENTITY_3"
#define		KEY_STRING_GENERIC_ENTITY4									"GENERIC_ENTITY_4"
#define		KEY_STRING_GENERIC_ENTITY5									"GENERIC_ENTITY_5"
#define		KEY_STRING_GENERIC_ENTITY6									"GENERIC_ENTITY_6"
#define		KEY_STRING_GENERIC_ENTITY7									"GENERIC_ENTITY_7"
#define		KEY_STRING_GENERIC_ENTITY8									"GENERIC_ENTITY_8"
#define		KEY_STRING_GENERIC_ENTITY9									"GENERIC_ENTITY_9"
#define		KEY_STRING_GENERIC_ENTITY10									"GENERIC_ENTITY_10"
#define		KEY_STRING_GENERIC_ENTITY11									"GENERIC_ENTITY_11"
#define		KEY_STRING_GENERIC_ENTITY12									"GENERIC_ENTITY_12"
#define		KEY_STRING_GENERIC_ENTITY13									"GENERIC_ENTITY_13"
#define		KEY_STRING_GENERIC_ENTITY14									"GENERIC_ENTITY_14"
#define		KEY_STRING_GENERIC_ENTITY15									"GENERIC_ENTITY_15"
#define		KEY_STRING_GENERIC_ENTITY16									"GENERIC_ENTITY_16"
#define		KEY_STRING_GENERIC_ENTITY17									"GENERIC_ENTITY_17"
#define		KEY_STRING_GENERIC_ENTITY18									"GENERIC_ENTITY_18"
#define		KEY_STRING_GENERIC_ENTITY19									"GENERIC_ENTITY_19"
#define		KEY_STRING_GENERIC_ENTITY20									"GENERIC_ENTITY_20"
#define		KEY_STRING_GENERIC_ENTITY21									"GENERIC_ENTITY_21"
#define		KEY_STRING_GENERIC_ENTITY22									"GENERIC_ENTITY_22"
#define		KEY_STRING_GENERIC_ENTITY23									"GENERIC_ENTITY_23"
#define		KEY_STRING_GENERIC_ENTITY24									"GENERIC_ENTITY_24"
#define		KEY_STRING_GENERIC_ENTITY25									"GENERIC_ENTITY_25"


//FR Automata types
#define		FR_AUTOMATA_TYPE_READ_RULE									1
#define		FR_AUTOMATA_TYPE_READ_ACTION								2
#define		FR_AUTOMATA_TYPE_READ_KEY_STRING							3
#define		FR_AUTOMATA_TYPE_READ_OUTPUT								4

//File Reader Actions	:	range from -1000 to -1499
#define		FR_ACTION_DO_NOTHING										-1000
#define		FR_ACTION_CREATE_RULE										-1001
#define		FR_ACTION_SET_RULE_NAME										-1002
#define		FR_ACTION_START_RULE_SECTION								-1003
#define		FR_ACTION_END_RULE_SECTION									-1004
#define		FR_ACTION_START_ACTION_SECTION								-1005
#define		FR_ACTION_END_ACTION_SECTION								-1006
#define		FR_ACTION_START_OUTPUT_SECTION								-1007
#define		FR_ACTION_END_OUTPUT_SECTION								-1008

#define		FR_ACTION_LEFT_SIBLING										-1100
#define		FR_ACTION_RIGHT_SIBLING										-1101
#define		FR_ACTION_PARENT											-1102
#define		FR_ACTION_N_TH_CHILD										-1103
#define		FR_ACTION_FIRST_CHILD										-1104
#define		FR_ACTION_LAST_CHILD										-1105

#define		FR_ACTION_GET_STRING										-1150
#define		FR_ACTION_GET_LEFT_SPECIAL_STRING							-1151
#define		FR_ACTION_GET_RIGHT_SPECIAL_STRING							-1152
#define		FR_ACTION_GET_STRING_WITH_SPECIAL_STRINGS					-1153
#define		FR_ACTION_GET_VALUE											-1154

#define		FR_ACTION_SET_VALUE											-1200
#define		FR_ACTION_SET_LEFT_SPECIAL_STRING							-1201
#define		FR_ACTION_SET_RIGHT_SPECIAL_STRING							-1202
#define		FR_ACTION_SET_STRING_AS_ARGUMENT							-1203
#define		FR_ACTION_SET_OPCODE_AS_ARGUMENT							-1204

#define		FR_ACTION_IS_MEMBER_OF										-1250
#define		FR_ACTION_IS_NOT_MEMBER_OF									-1251
#define		FR_ACTION_IS_FULL_STRING_MEMBER_OF							-1252
#define		FR_ACTION_IS_FULL_STRING_NOT_MEMBER_OF						-1253
#define		FR_ACTION_IS_LEFT_STRING_MEMBER_OF							-1254
#define		FR_ACTION_IS_LEFT_STRING_NOT_MEMBER_OF						-1255
#define		FR_ACTION_IS_RIGHT_STRING_MEMBER_OF							-1256
#define		FR_ACTION_IS_RIGHT_STRING_NOT_MEMBER_OF						-1257
#define		FR_ACTION_IS_EXIST											-1258
#define		FR_ACTION_IS_NOT_EXIST										-1259
#define		FR_ACTION_IS_TYPE_OF										-1260
#define		FR_ACTION_IS_CATEGORY_OF									-1261
#define		FR_ACTION_IS_LEFT_PART_OF_STRING_MEMBER_OF					-1262
#define		FR_ACTION_IS_RIGHT_PART_OF_STRING_MEMBER_OF					-1263
#define		FR_ACTION_IS_STRING_CONTAIN									-1264
#define		FR_ACTION_IS_LEFT_PART_OF_FULL_STRING_MEMBER_OF				-1265
#define		FR_ACTION_IS_RIGHT_PART_OF_FULL_STRING_MEMBER_OF			-1266
#define		FR_ACTION_IS_FULL_STRING_CONTAIN							-1267
#define		FR_ACTION_IS_LEFT_PART_OF_LEFT_STRING_MEMBER_OF				-1268
#define		FR_ACTION_IS_RIGHT_PART_OF_LEFT_STRING_MEMBER_OF			-1269
#define		FR_ACTION_IS_LEFT_STRING_CONTAIN							-1270
#define		FR_ACTION_IS_LEFT_PART_OF_RIGHT_STRING_MEMBER_OF			-1271
#define		FR_ACTION_IS_RIGHT_PART_OF_RIGHT_STRING_MEMBER_OF			-1272
#define		FR_ACTION_IS_RIGHT_STRING_CONTAIN							-1273
#define		FR_ACTION_IS_VALUE_MEMBER_OF								-1274
#define		FR_ACTION_IS_VALUE_NOT_MEMBER_OF							-1275
#define		FR_ACTION_IS_LEFT_PART_OF_VALUE_MEMBER_OF					-1276
#define		FR_ACTION_IS_RIGHT_PART_OF_VALUE_MEMBER_OF					-1277
#define		FR_ACTION_IS_VALUE_CONTAIN									-1278
#define		FR_ACTION_IS_LEFT_PART_OF_STRING_NOT_MEMBER_OF				-1279
#define		FR_ACTION_IS_RIGHT_PART_OF_STRING_NOT_MEMBER_OF				-1280
#define		FR_ACTION_IS_STRING_NOT_CONTAIN								-1281
#define		FR_ACTION_IS_LEFT_PART_OF_FULL_STRING_NOT_MEMBER_OF			-1282
#define		FR_ACTION_IS_RIGHT_PART_OF_FULL_STRING_NOT_MEMBER_OF		-1283
#define		FR_ACTION_IS_FULL_STRING_NOT_CONTAIN						-1284
#define		FR_ACTION_IS_LEFT_PART_OF_LEFT_STRING_NOT_MEMBER_OF			-1285
#define		FR_ACTION_IS_RIGHT_PART_OF_LEFT_STRING_NOT_MEMBER_OF		-1286
#define		FR_ACTION_IS_LEFT_STRING_NOT_CONTAIN						-1287
#define		FR_ACTION_IS_LEFT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF		-1288
#define		FR_ACTION_IS_RIGHT_PART_OF_RIGHT_STRING_NOT_MEMBER_OF		-1289
#define		FR_ACTION_IS_RIGHT_STRING_NOT_CONTAIN						-1290
#define		FR_ACTION_IS_LEFT_PART_OF_VALUE_NOT_MEMBER_OF				-1291
#define		FR_ACTION_IS_RIGHT_PART_OF_VALUE_NOT_MEMBER_OF				-1292
#define		FR_ACTION_IS_VALUE_NOT_CONTAIN								-1293


#define		FR_ACTION_FINISH_ARGUMENT									-1350
#define		FR_ACTION_ADD_STRING_ARGUMENT_AND_ADD_CURRENT_NODE			-1351
#define		FR_ACTION_SET_STRING_AS_ARGUMENT_AND_END_ARGUMENT_LIST		-1352
#define		FR_ACTION_START_RULE_ENTRY									-1353
#define		FR_ACTION_START_ACTION										-1354
#define		FR_ACTION_START_STRING_OUTPUT								-1355
#define		FR_ACTION_ADD_SUBSTRING_AND_END_STRING						-1356


//File Reader Action Types
#define		FR_ACTION_TYPE_GET_A_STRING									0
#define		FR_ACTION_TYPE_SET_A_STRING									1
#define		FR_ACTION_TYPE_MEMBERSHIP									2
#define		FR_ACTION_TYPE_EXISTANCE									3
#define		FR_ACTION_TYPE_NODE_TO_NODE_TRANSITION						4


//Rule File Reader states
#define		RULE_FILE_READER_STATE_READY_TO_START_RULE					1
#define		RULE_FILE_READER_STATE_READING_RULES						2
#define		RULE_FILE_READER_STATE_RULE_READING_OVER					3
#define		RULE_FILE_READER_STATE_READING_ACTIONS						4
#define		RULE_FILE_READER_STATE_ACTION_READING_OVER					5
#define		RULE_FILE_READER_STATE_READING_OUTPUT						6


/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////
/////////////////// File Handler Defs ///////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

#define		NO_OF_STRINGS_IN_TEMPLATE_RESULT_STRING_ARRAY				10		//Used in CFileReaderTemplateResultElement class
#define		FILE_READER_TEXT_FILE_READ_BUFFER_SIZE						5000

//FileReaderTemplate types
#define		FILE_READER_TEMPLATE_TYPE_GENERIC_STRING					0
#define		FILE_READER_TEMPLATE_TYPE_STRING_WITHIN_TAGS				1

//string types
#define		STRING_ENCODING_TYPE_ASCII									0
#define		STRING_ENCODING_TYPE_UNICODE								1

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////
/////////////////// Presentation Module Defs ////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

//Entity write modes
#define		ENTITY_WRITE_MODE_COMPLETE_DESCRIPTION						1
#define		ENTITY_WRITE_MODE_STRING_WITH_SPECIAL_STRINGS				2
#define		ENTITY_WRITE_MODE_ONLY_STRING								3

//Output Tree write modes
#define		OUTPUT_TREE_WRITE_MODE_NONE									0
#define		OUTPUT_TREE_WRITE_MODE_FULL_TREE							1
#define		OUTPUT_TREE_WRITE_MODE_ROOT_STRING_WITH_SPECIAL_STRINGS		2
#define		OUTPUT_TREE_WRITE_MODE_ROOT_STRING							3

//Rule result write modes
#define		RULE_RESULT_WRITE_MODE_NONE									0
#define		RULE_RESULT_WRITE_MODE_ONLY_COUNT							1
#define		RULE_RESULT_WRITE_MODE_ONLY_ENTITIES						2
#define		RULE_RESULT_WRITE_MODE_COUNT_AND_ENTITIES					3

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////
/////////////////// Defs used by External classes////////////////////////////
/////////////////////////////////////////////////////////////////////////////

#define		LANGUAGE_FILE_WRITE_SOUND_LENGTH							30

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////
/////////////////// Defs used by New Tree classes////////////////////////////
/////////////////////////////////////////////////////////////////////////////

#define		NODE_ID_OFFSET												100
#define		NODE_ID_INVALID												0
#define		NODE_TYPE_GENERAL											0
#define		NODE_NATURE_GENERAL											0
#define		NODE_NATURE_ORGANIZING										1	// This category of nodes can have child nodes with weights and the childs will be sorted in the ascending order of their weights

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////
/////////////////// Common defs /////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

#ifndef WIDECHAR
#define		_MSTR(X)													#X
#define		EMPTY_STRING												""
#define		SPACE														" "
#define     QUOTE                                                       "\""
#define     COMMA                                                       ","
#define		_MATOI														atoi
#define		STRING_END_CHAR												'\0'
#else
#define     ISWIDECHAR
#define		_MSTR(X)													L#X
#define		EMPTY_STRING												L""
#define		SPACE														L" "
#define     QUOTE                                                       L"\""
#define     COMMA                                                       L","
#define		_MATOI														_wtoi
#define		STRING_END_CHAR												L'\0'
#endif

#ifdef PRODUCTION
#define		_WIDESTR(X)													L#X
#define     EMPTY_WIDESTRING                                            L""
#define     UNIDENTIFIED_LINE_TEMPLATE_CHAR                             L'a'
#define     LINE_TEMPLATE_CHAR_START                                    L'b'
#define     WIDECOUT                                                    std::wcout
#else
#define		_WIDESTR(X)													#X
#define     EMPTY_WIDESTRING                                            ""
#define     UNIDENTIFIED_LINE_TEMPLATE_CHAR                             'a'
#define     LINE_TEMPLATE_CHAR_START                                    'b'
#define     WIDECOUT                                                    std::cout
#endif


// Command types
// Node commands
#define		COMMAND_TYPE_INVALID										0
#define		COMMAND_TYPE_LEFT_SIBLING									1
#define		COMMAND_TYPE_RIGHT_SIBLING									2
#define		COMMAND_TYPE_PARENT											3
#define		COMMAND_TYPE_FIRST_CHILD									4
#define		COMMAND_TYPE_CHILDREN										5
#define		COMMAND_TYPE_CHILD_COUNT									6
#define		COMMAND_TYPE_GET_VALUE										7
#define		COMMAND_TYPE_GET_LVALUE										8
#define		COMMAND_TYPE_GET_RVALUE										9
#define		COMMAND_TYPE_GET_CUSTOM_STRING								10
#define		COMMAND_TYPE_GET_ID											11
#define		COMMAND_TYPE_GET_TYPE										12
#define		COMMAND_TYPE_GET_NATURE										13
#define		COMMAND_TYPE_GET_WEIGHT										14
#define		COMMAND_TYPE_GET_MIN_CHILD_WEIGHT							15
#define		COMMAND_TYPE_GET_MAX_CHILD_WEIGHT							16
#define		COMMAND_TYPE_SET_VALUE										25
#define		COMMAND_TYPE_SET_LVALUE										26
#define		COMMAND_TYPE_SET_RVALUE										27
#define		COMMAND_TYPE_SET_TYPE										28
#define		COMMAND_TYPE_SET_NATURE										29
#define		COMMAND_TYPE_SET_CUSTOM_STRING								30
#define		COMMAND_TYPE_SET_MIN_CHILD_WEIGHT							31
#define		COMMAND_TYPE_SET_MAX_CHILD_WEIGHT							32
#define		COMMAND_TYPE_SET_WEIGHT										33
#define		COMMAND_TYPE_EXPAND											45
#define		COMMAND_TYPE_ADD_NODE										46
#define		COMMAND_TYPE_ADD_NODE_WITH_WEIGHT							47
#define		COMMAND_TYPE_READ_FROM_FILE									48
#define		COMMAND_TYPE_GET_AGGREGATED_VALUE							49
#define		COMMAND_TYPE_GET_SUBTREE									50
#define		COMMAND_TYPE_IS_TYPE										51
#define		COMMAND_TYPE_IS_VALUE										52
#define		COMMAND_TYPE_GET_CHILD_OF_TYPE								53
#define		COMMAND_TYPE_LAST_CHILD										54
#define		COMMAND_TYPE_IS_HAVING_CUSTOM_STRING						55
#define		COMMAND_TYPE_GET_CHILD_NODE_BY_CUSTOM_STRING				56
#define		COMMAND_TYPE_SET_ID											57
#define     COMMAND_TYPE_FILTER_SUBTREE                                 58
#define     COMMAND_TYPE_GET_ENTITY_OBJECT                              59
#define     COMMAND_TYPE_SET_ENTITY_OBJECT                              60
#define     COMMAND_TYPE_CHECK_NOT_NULL                                 61
#define     COMMAND_TYPE_GET_STRING                                     62
#define     COMMAND_TYPE_GET_INTEGER                                    63
#define     COMMAND_TYPE_GET_BOOLEAN                                    64
#define     COMMAND_TYPE_GET_CUSTOM_OBJ                                 65
#define     COMMAND_TYPE_SET_ATTRIBUTES                                 66
#define     COMMAND_TYPE_GET_NODE_OBJ                                   67
#define     COMMAND_TYPE_ADD_INNER_OBJ                                  68

// String commands
#define		COMMAND_TYPE_IS_STRING_EQUAL_TO								1000
#define		COMMAND_TYPE_IS_STRING_MEMBER_OF							1001
#define		COMMAND_TYPE_IS_HAVING_SUBSTRING							1002
#define		COMMAND_TYPE_IS_HAVING_LEFT_SUBSTRING						1003
#define		COMMAND_TYPE_IS_HAVING_RIGHT_SUBSTRING						1004
#define		COMMAND_TYPE_ADD_PREFIX										1100
#define		COMMAND_TYPE_ADD_POSTFIX									1101
#define		COMMAND_TYPE_TRIM_LEFT										1102
#define		COMMAND_TYPE_TRIM_RIGHT										1103
#define		COMMAND_TYPE_WRITE_TO_FILE									1104
#define		COMMAND_TYPE_GET_LENGTH										1200
#define		COMMAND_TYPE_STRINGTOINTEGER								1201
#define     COMMAND_TYPE_SECONDS_TO_MONTHS                              8001
#define     COMMAND_TYPE_SECONDS_TO_DAYS                                8002
#define     COMMAND_TYPE_SECONDS_TO_YEARS                               8003
#define     COMMAND_TYPE_GET_DIFFERENCE_BY_STRING                       8004
#define     COMMAND_TYPE_STRING_TO_READABLE_DATETIME                    8005
#define     COMMAND_TYPE_DATE_NOW                                       8006
#define     COMMAND_TYPE_STRING_TO_UNIX_TIME                            8007
#define     COMMAND_TYPE_STRINGTOBOOLEAN                                8008
#define     COMMAND_TYPE_STRINGTOBOOL                                   1202
#define     COMMAND_TYPE_GET_COMMA                                      8009
#define     COMMAND_TYPE_NEXT_SIBLING                                   8010
#define     COMMAND_TYPE_CONVERT_TO_SENTENCE_CASE                       8011
#define     COMMAND_TYPE_GET_DAY_OF_THE_WEEK_SHORT_STRING               8012
#define     COMMAND_TYPE_GET_DAY_STRING                                 8013
#define     COMMAND_TYPE_GET_MONTH_SHORT_STRING                         8014
#define     COMMAND_TYPE_GET_TIME_24_HOUR_FORMAT                        8015
#define     COMMAND_TYPE_GET_YEAR                                       8016
#define     COMMAND_TYPE_ADD_PERIOD                                     8017

// Int Commands
#define		COMMAND_TYPE_IS_INT_EQUAL_TO								2000
#define		COMMAND_TYPE_IS_INT_MEMBER_OF								2001
#define		COMMAND_TYPE_IS_LESS_THAN									2002
#define		COMMAND_TYPE_IS_LESS_THAN_OR_EQUAL_TO						2003
#define		COMMAND_TYPE_IS_GREATER_THAN								2004
#define		COMMAND_TYPE_IS_GREATER_THAN_OR_EQUAL_TO					2005
#define		COMMAND_TYPE_ADD											2100
#define		COMMAND_TYPE_SUBTRACT										2101
#define		COMMAND_TYPE_TOSTRING										2200
#define     COMMAND_TYPE_SET_INTEGER                                    2201
#define     COMMAND_TYPE_PERCENTAGE                                     2202


// Bool Commands
#define		COMMAND_TYPE_BOOL_AND										4000
#define		COMMAND_TYPE_BOOL_OR										4001
#define		COMMAND_TYPE_BOOLTOSTRING									4002
#define		COMMAND_TYPE_SET_BOOL									    4003
#define		COMMAND_TYPE_TO_FALSE									    4004
#define		COMMAND_TYPE_TO_TRUE									    4005

// DateTime Commands
#define     COMMAND_TYPE_DATETOSTRING                                   8000


// List commands
#define		COMMAND_TYPE_GET_ITEM_COUNT									5000
#define		COMMAND_TYPE_SEEK											5001
#define		COMMAND_TYPE_SEEK_TO_BEGIN									5002
#define		COMMAND_TYPE_SEEK_TO_END									5003
#define		COMMAND_TYPE_GET_CURR_ELEM									5004
#define		COMMAND_TYPE_GET_INNER_ITEM_COUNT							5005
#define     COMMAND_TYPE_LIST_FILTER                                    5006
#define     COMMAND_TYPE_LIST_GROUPBY                                   5007
#define     COMMAND_TYPE_LIST_GROUP_SEQUENCE_BY                         5008
#define     COMMAND_TYPE_GET_NEXT_ELEM                                  5009
#define     COMMAND_TYPE_GET_UNIQUE_NODE_LIST_WITH_COUNT                5010
#define     COMMAND_TYPE_SORT_NODE_LIST                                 5011
#define     COMMAND_TYPE_EXTRACT_NODE_LIST_TOP                          5012
#define     COMMAND_TYPE_GET_OLDEST_DATE                                5013
#define     COMMAND_TYPE_GET_LATEST_DATE                                5014
#define     COMMAND_TYPE_GET_UNIQUE_NODE_LIST_WITH_NODE_REF                5015

// Special Commands
#define		COMMAND_TYPE_ADDITIONAL_FUNCTION							10000
#define		COMMAND_TYPE_STORE_AS_VARIABLE								10001
#define		COMMAND_TYPE_WHILE											10002
#define		COMMAND_TYPE_DO												10003
#define		COMMAND_TYPE_BREAK											10004
#define		COMMAND_TYPE_IF												10005
#define		COMMAND_TYPE_IFNOT											10006
#define		COMMAND_TYPE_ENDIF											10007
#define		COMMAND_TYPE_CONTINUE										10008

// Entity level commands
#define		COMMAND_TYPE_IS_NULL										15000
#define		COMMAND_TYPE_IS_NOT_NULL									15001

// Entity types
#define		ENTITY_TYPE_NULL											0
#define		ENTITY_TYPE_INT												1
#define		ENTITY_TYPE_STRING											2
#define		ENTITY_TYPE_NODE											3
#define		ENTITY_TYPE_BOOL											4
#define		ENTITY_TYPE_LIST											5
#define		ENTITY_TYPE_EXECUTION_TEMPLATE								6
#define     ENTITY_TYPE                                                 7
#define     ENTITY_TYPE_DATETIME                                        8
#define		ENTITY_TYPE_INVALID											100

// Parser priority levels for LDEL variable types
#define     DEFAULT_PARSER_PRIORITY                                     1000
#define     PARSER_PRIORITY_LOW                                         500
#define     PARSER_PRIORITY_HIGH                                        1500

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

