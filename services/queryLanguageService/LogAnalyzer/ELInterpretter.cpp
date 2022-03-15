//
//  ELInterpretter.cpp
//  LogAnalyzer
//
//  Created by Dileepa Jayathilaka on 12/29/13.
//  Copyright (c) 2013 99x Eurocenter. All rights reserved.
//

#include "ELInterpretter.h"
#include "DefFileReader.h"
#include "MetaData.h"
#include "ELParser.h"
#include "ELVariable.h"
#include "ELLineTemplate.h"
#include "ELBlockTemplate.h"
#include "MemMan.h"
#include "ELNodeWrapper.h"
#include "Utils.h"

MetaData *mdglob;

ELInterpretterResult::ELInterpretterResult()
{
    startNode = NULL;
}

ELInterpretterResult *ELInterpretter::EvaluateCase(std::string sDefFile)
{
    DefFileReader dfr;
    MetaData *md = dfr.Read(sDefFile);
    mdglob = md;
    ELParser p;
    ELParserResult res;
    ELInterpretterResult *ir = 0;
    MemoryManager::Inst.CreateObject(&ir);
    int start = Utils::getMilliCount();
    bool succ = p.ProcessScript(md->s_ScriptFile, md, res);
    ir->millisecondsForParsing = Utils::getMilliSpan(start);

    start = Utils::getMilliCount();
    WIDESTRING sLines = ProcessLinesInFile(md->s_LogFile, res.vec_LineTemplates, ir);
    ir->finalString = ProcessBlocks(sLines, res.vec_BlockTemplates, res.vec_LineTemplates, ir);
    ir->millisecondsForInterpreting = Utils::getMilliSpan(start);

    PrintInterpretterResult(ir);
    PrintInterpretterResultInJSON(ir);

    return ir;
}

void ELInterpretter::PrintInterpretterResult(ELInterpretterResult *ir)
{
    MOFSTREAM file;
    file.open(mdglob->s_ResultFile, std::ios::out | std::ios::trunc);
    file << _MSTR(Time spent parsing =) << SPACE << ir->millisecondsForParsing << SPACE << _MSTR(ms\n);
    file << _MSTR(Time spent interpreting =) << SPACE << ir->millisecondsForInterpreting << SPACE << _MSTR(ms\n);
    PNODE curr = ir->startNode;
    while (curr)
    {
        ELNodeWrapper *wrapper = ELNodeWrapper::mapNodeToWrapper[curr];
        wrapper->PrintNodeToFile(file);
        curr = curr->GetRightSibling();
    }
    file.close();
}

void ELInterpretter::PrintInterpretterResultInJSON(ELInterpretterResult *ir)
{
    MOFSTREAM jsonfile;
    jsonfile.open(mdglob->s_ResultJSONFile, std::ios::out | std::ios::trunc);
    PNODE curr = ir->startNode;
    int count =1;
    jsonfile<<_MSTR([);
    jsonfile<<_MSTR({);
                        while (curr) {
                            ELNodeWrapper* wrapper = ELNodeWrapper::mapNodeToWrapper[curr];
                            //wrapper->PrintNodeToFile(file);
                            wrapper->PrintNodeToJSONFile(jsonfile,count);
                            count++;
                            if(curr->GetRightSibling())
                            {
                                jsonfile <<",";
                            }
                            curr = curr->GetRightSibling();
                        }
                        //jsonfile<<_MSTR(\n);
                        jsonfile <<_MSTR(}]);
    jsonfile.close();
}

WIDESTRING ELInterpretter::ProcessLinesInFile(MSTRING sLogFile, VEC_ELLINETEMPLATE &vecLineTemplates, ELInterpretterResult *res)
{
    MIFSTREAM file(sLogFile.c_str());
    MSTRING sLine;
    WIDESTRING ret = EMPTY_WIDESTRING;
    res->startNode = NULL;
    PNODE currentNode = res->startNode;
    if (file.is_open())
    {
        while (!file.eof())
        {
            getline(file, sLine);
            bool matchingLineFound = false;
            ELLineTemplate *matchingLineTemplate = NULL;
            ELNodeWrapperInfo info;
            VEC_ELLINETEMPLATE::iterator ite = vecLineTemplates.begin();
            VEC_ELLINETEMPLATE::iterator iteEnd = vecLineTemplates.end();
            for (; ite != iteEnd; ++ite)
            {
                ELLineTemplate *lt = (*ite);
                if (lt->ParseLine(sLine))
                {
                    ret += lt->ch;
                    matchingLineFound = true;
                    matchingLineTemplate = lt;

                    info.type = ELNODE_TYPE_LINE;
                    info.name = lt->name;
                    info.value = sLine;
                    info.parserElement = lt;
                    break;
                }
            }
            if (!matchingLineFound)
            {
                info.type = ELNODE_TYPE_LINE;
                info.name = _MSTR(DEFAULT_LINE_TEMPLATE);
                info.value = sLine;
                ret += ELLineTemplate::defaultChar;
            }

            ELNodeWrapper *wrapper = new ELNodeWrapper(info);
            PNODE newNode = wrapper->GetNode();
            if (!currentNode)
            {
                currentNode = newNode;
                res->startNode = currentNode;
            }
            else
            {
                currentNode->SetRightSibling(newNode);
                newNode->SetLeftSibling(currentNode);
                currentNode = newNode;
            }

            if (matchingLineTemplate)
            {
                matchingLineTemplate->CreateNodesForLineElements(currentNode);
            }
        }
    }
    //std::wcout<<L"\nReturn string = "<<ret;
    //WIDECOUT<<_WIDESTR(\n)<<ret<<_WIDESTR(\n);
    return ret;
}

WIDESTRING ELInterpretter::ProcessBlocks(WIDESTRING str, VEC_ELBLOCKTEMPLATE &blockTemplates, VEC_ELLINETEMPLATE &lineTemplates, ELInterpretterResult *res)
{
    WIDESTRING availableSymbols = EMPTY_WIDESTRING;
    VEC_ELLINETEMPLATE::iterator ite1 = lineTemplates.begin();
    VEC_ELLINETEMPLATE::iterator iteEnd1 = lineTemplates.end();
    for (; ite1 != iteEnd1; ++ite1)
    {
        availableSymbols += (*ite1)->ch;
    }

    blockTemplates = GetBlockTemplatesPreparedToHandleRecursiveDefs(blockTemplates);

    VEC_ELBLOCKTEMPLATE definiteBlocks;
    VEC_ELBLOCKTEMPLATE nondefiniteBlocks;
    MAP_WIDECHAR_ELBLOCKTEMPLATE mapBlocks;
    VEC_ELBLOCKTEMPLATE::iterator ite = blockTemplates.begin();
    VEC_ELBLOCKTEMPLATE::iterator iteEnd = blockTemplates.end();
    for (; ite != iteEnd; ++ite)
    {
        ELBlockTemplate *tmp = (*ite);
        if (tmp->IsDefinite())
        {
            definiteBlocks.push_back(tmp);
        }
        else
        {
            nondefiniteBlocks.push_back(tmp);
        }
        mapBlocks[tmp->ch] = tmp;
    }

    while (true)
    {
        WIDESTRING strBeforeIteration = str;
        bool shouldConsiderSequenceBlocks = true;
        ite = definiteBlocks.begin();
        iteEnd = definiteBlocks.end();
        for (; ite != iteEnd; ++ite)
        {
            ELBlockTemplate *tmp = (*ite);
            if (tmp->IsReadyToProcess(availableSymbols))
            {
                UnifyBlock(str, tmp, res);
                availableSymbols += tmp->ch;
                shouldConsiderSequenceBlocks = false;
            }
        }

        if (shouldConsiderSequenceBlocks)
        {
            ite = nondefiniteBlocks.begin();
            iteEnd = nondefiniteBlocks.end();
            for (; ite != iteEnd; ++ite)
            {
                ELBlockTemplate *tmp = (*ite);
                if (tmp->IsReadyToProcess(availableSymbols))
                {
                    UnifyBlock(str, tmp, res);
                    availableSymbols += tmp->ch;
                }
            }
        }

        // If no changes occured in the string, that means we have reached the end of parsing
        if (str == strBeforeIteration)
        {
            return str;
        }
    }
}

VEC_ELBLOCKTEMPLATE ELInterpretter::GetBlockTemplatesPreparedToHandleRecursiveDefs(VEC_ELBLOCKTEMPLATE &blockTemplates)
{
    VEC_ELBLOCKTEMPLATE ret;
    VEC_ELBLOCKTEMPLATE::iterator ite = blockTemplates.begin();
    VEC_ELBLOCKTEMPLATE::iterator iteEnd = blockTemplates.end();
    for (; ite != iteEnd; ++ite)
    {
        ELBlockTemplate *bt = (*ite);
        if (bt->isUnion)
        {
            VEC_BLOCKELEMENT::iterator ite2 = bt->elements.begin();
            VEC_BLOCKELEMENT::iterator iteEnd2 = bt->elements.end();
            for (; ite2 != iteEnd2; ++ite2)
            {
                ELBlockElement *be = (*ite2);
                ELBlockTemplate *newTemplate = 0;
                MemoryManager::Inst.CreateObject(&newTemplate);
                newTemplate->name = bt->name;
                newTemplate->ch = bt->ch;
                newTemplate->elements.push_back(be);
                newTemplate->isUnion = false;
                ret.push_back(newTemplate);
            }
            MemoryManager::Inst.DeleteObject(bt);
        }
        else
        {
            ret.push_back(bt);
        }
    }

    return ret;
}

void ELInterpretter::UnifyBlock(WIDESTRING &str, ELBlockTemplate *block, ELInterpretterResult *res)
{
    WIDESTRING::size_type pos = 0;
    WIDESTRING::size_type newpos;
    while (true)
    {
        WIDESTRING::size_type len = str.length();
        if (pos == len)
        {
            return;
        }
        if (block->TryUnify(str, pos, newpos))
        {
            // alter the node structure accordingly
            unsigned long lenAlterPart = newpos - pos;
            PNODE node = res->startNode;
            for (int i = 0; i < pos; ++i)
            {
                node = node->GetRightSibling();
            }
            ELNodeWrapperInfo info;
            info.type = ELNODE_TYPE_BLOCK;
            info.name = block->name;
            info.value = EMPTY_STRING;
            info.parserElement = block;
            ELNodeWrapper *wrapper = new ELNodeWrapper(info);
            PNODE newNode = wrapper->GetNode();
            PNODE left = node->GetLeftSibling();
            if (left)
            {
                left->SetRightSibling(newNode);
                newNode->SetLeftSibling(left);
            }
            else
            {
                res->startNode = newNode;
            }
            for (int i = 0; i < lenAlterPart; ++i)
            {
                PNODE temp = node->GetRightSibling();
                node->SetLeftSibling(NULL);
                node->SetRightSibling(NULL);
                newNode->AppendNode(node);
                node = temp;
            }
            newNode->SetRightSibling(node);
            node->SetLeftSibling(newNode);

            // alter the string
            WIDESTRING head = EMPTY_WIDESTRING;
            if (pos > 0)
            {
                head = str.substr(0, pos);
            }
            if (newpos == len)
            {
                str = (head + block->ch);
                return;
            }
            else
            {
                WIDESTRING tail = str.substr(newpos, len - newpos);
                str = (head + block->ch + tail);
                ++pos;
            }
        }
        else
        {
            ++pos;
        }
    }
}

VEC_ELLINEANNOTATION *ELInterpretter::AnnotateAgainstLineTemplate(MSTRING sDefFile, MSTRING sLineTemplateName)
{
    DefFileReader dfr;
    MetaData *md = dfr.Read(sDefFile);
    ELParser p;
    ELParserResult res;
    ELInterpretterResult *ir = 0;
    MemoryManager::Inst.CreateObject(&ir);
    p.ProcessScript(md->s_ScriptFile, md, res);
    ProcessLinesInFile(md->s_LogFile, res.vec_LineTemplates, ir);

    PNODE current = ir->startNode;
    VEC_ELLINEANNOTATION *ret = 0;
    MemoryManager::Inst.CreateObject(&ret);
    while (current)
    {
        if (sLineTemplateName == MSTRING(current->GetCustomString()))
        {
            ELLineAnnotation *an = 0;
            MemoryManager::Inst.CreateObject(&an);
            ret->push_back(an);
            MSTRING::size_type pos = 0;
            FillAnnotationElements(current, an, pos);
        }

        current = current->GetRightSibling();
    }
    return ret;
}

void ELInterpretter::FillAnnotationElements(PNODE node, ELLineAnnotation *an, MSTRING::size_type &startPos)
{
    PNODE currentChild = node->GetFirstChild();
    if (!currentChild)
    {
        bool shouldAnnotate = false;
        MSTRING::size_type len = MSTRING(node->GetValue()).length();
        MBYTE nodeType = node->GetNature();
        if (nodeType == ELNODE_TYPE_VARIABLE)
        {
            ELVariable *var = (ELVariable *)node->GetCustomObj();
            if (!var->IsConstant())
            {
                shouldAnnotate = true;
            }
        }

        if (shouldAnnotate)
        {
            ELLineAnnotationElement *lae = 0;
            MemoryManager::Inst.CreateObject(&lae);
            lae->start = startPos;
            lae->len = len;
            lae->name = node->GetCustomString();
            an->elements.push_back(lae);
        }

        startPos += len;
    }
    else
    {
        while (currentChild)
        {
            FillAnnotationElements(currentChild, an, startPos);
            currentChild = currentChild->GetRightSibling();
        }
    }
}
