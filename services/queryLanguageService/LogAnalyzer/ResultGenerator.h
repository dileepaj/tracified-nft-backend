//
//  ResultGenerator.h
//  FlexibleComputerLanguage
//
//  Created by Murtaza Anverali on 7/5/18.
//  Copyright © 2018 Dileepa Jayathilaka. All rights reserved.
//

#ifndef ResultGenerator_h
#define ResultGenerator_h

#include "Node.h"

class ResultGenerator
{
  private:
    static std::string TypeFormatting(std::string result, int type);
    static std::string CreateArray(Node *result, int type);
    static std::string CreateObject(Node *result, int type);
    static std::string ValueFormatting(Node *result);

  public:
    static std::string CreateResult(Node *result);
};

#endif /* ResultGenerator_h */
