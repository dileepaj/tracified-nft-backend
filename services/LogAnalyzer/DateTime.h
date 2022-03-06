//
//  DateTime.h
//  FlexibleComputerLanguage
//
//  Created by Murtaza Anverali on 5/30/18.
//  Copyright © 2018 Dileepa Jayathilaka. All rights reserved.
//

#ifndef DateTime_h
#define DateTime_h

class DateTime : public Value<MSTRING, ENTITY_TYPE_DATETIME>
{
public:
    DateTime()
    : Value<MSTRING, ENTITY_TYPE_DATETIME>() {
        
    }
    
    DateTime(MSTRING str)
    : Value<MSTRING, ENTITY_TYPE_DATETIME>(str) {
        
    }
    
    virtual MSTRING ToString()
    {
        return val;
    }
};

#endif /* DateTime_h */
