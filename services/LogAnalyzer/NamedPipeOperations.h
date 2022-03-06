//
//  NamedPipeOperations.h
//  Named Pipes C/C++
//
//  Created by Murtaza Anverali on 7/4/18.
//  Copyright © 2018 Murtaza Anverali. All rights reserved.
//
#include <iostream>

class NamedPipeOperations
{

  public:
    static std::string readFromPipe(FILE *stream);
    static int writeToPipe(FILE *stream, std::string s);
    static FILE *openPipeToRead(int fd);
    static int closeReadPipe(FILE *stream, int fd);
    static FILE *openPipeToWrite(int fd);
    static int closeWritePipe(FILE *stream, int fd);
};
