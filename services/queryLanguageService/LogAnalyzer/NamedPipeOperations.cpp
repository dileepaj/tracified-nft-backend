//
//  NamedPipeOperations.cpp
//  Named Pipes C/C++
//
//  Created by Murtaza Anverali on 7/4/18.
//  Copyright © 2018 Murtaza Anverali. All rights reserved.
//
#include <unistd.h>
#include <iostream>
#include "NamedPipeOperations.h"
#include <string.h>
#include <sys/types.h>
#include <stdio.h>

#define BUF_SIZE 1000000

std::string NamedPipeOperations::readFromPipe(FILE *stream)
{
    //    int BUF_SIZE = 4;
    //    char str1[BUF_SIZE];
    //    std::string resultString;
    char buff[BUF_SIZE];
    char *chrp;
    std::string resultString = "";
    //FILE *stream;
    int c;
    //stream = fdopen(fd, "r");
    // std::cout << "here 1 " << fd << std::endl;
    chrp = fgets(buff, BUF_SIZE, stream);
    std::string someStr(buff);
    //std::cout << "here 1 " << chrp << std::endl;
    while ((chrp != 0))
    {
        //std::cout << "here 2 " << chrp << std::endl;
        resultString = resultString + buff;
        chrp = fgets(buff, BUF_SIZE, stream);
    }
    // while ((c = fgetc (stream)) != EOF)
    // {
    //     // putchar (c);
    //     resultString = resultString + (char)c;
    // }
    // char** line;
    // size_t* n = 0;
    // getline(line, n, stream);
    //fclose(stream);

    //close(fd);
    //    while (1)
    //    {
    //        int read2result = read(fd, str1, BUF_SIZE);
    //        if (read2result < BUF_SIZE && read2result >= 0)
    //        {
    //            if (read2result != 0)
    //            {
    //                resultString += str1;
    //            }
    //            break;
    //        }
    //        resultString += str1;
    //        memset(str1, 0, sizeof(str1));
    //    }
    //    std::cout << "The result is " << resultString;
    return resultString;
};

int NamedPipeOperations::writeToPipe(FILE *stream, std::string s)
{
    //    int BUF_SIZE = 4;
    //    char str1[BUF_SIZE];
    //    char *sc;
    //    sc = (char *)s.c_str();
    //    int i, j = 0;
    //    for (i = 0; sc[i]; i++)
    //    {
    //        str1[i % 4] = sc[i];
    //        // std::cout << "i is " << i << "sc is " << sc[i] << std::endl;
    //        if (i % 4 == 3 && i != 0)
    //        {
    //            // std::cout << "arr1 is " << str2 << std::endl;
    //            int status = write(fd, str1, strlen(str1));
    //            // std::cout << "write status " << status << std::endl;
    //            // arr1 = new char[4];
    //            memset(str1, 0, sizeof(str1));
    //        }
    //        // printf("%c", sc[ i ]);
    //    }
    //    // std::cout << "i is " << i << std::endl;
    //    if (i % 4 < 4 && i % 4 != 0)
    //    {
    //        // std::cout << "arr1 is " << str2 << std::endl;
    //        int status = write(fd, str1, strlen(str1));
    //        // std::cout << "write status " << status << std::endl;
    //    }
    //FILE *stream;
    //stream = fdopen(fd, "w");
    // std::cout << "here 3 " << fd << std::endl;
    fprintf(stream, (char *)s.c_str());
    //fputs((char *)s.c_str(), stream);
    // std::cout << "here 4 " << stream << s << std::endl;
    //fclose(stream);
    //close(fd);
    return 0;
};
 
FILE *NamedPipeOperations::openPipeToRead(int fd)
{
    //std::cout << "openPipeToRead " << fd << std::endl;
    FILE *stream;
    stream = fdopen(fd, "r");
    //std::cout << "openPipeToRead 1 " << fd << std::endl;
    return stream;
};

int NamedPipeOperations::closeReadPipe(FILE *stream, int fd)
{
    //std::cout << "closeReadPipe " << fd << std::endl;
    fclose(stream);
    close(fd);
    //std::cout << "closeReadPipe 1 " << fd << std::endl;
    return 0;
}

FILE *NamedPipeOperations::openPipeToWrite(int fd)
{
    //std::cout << "openPipeToWrite " << fd << std::endl;
    FILE *stream;
    stream = fdopen(fd, "w");
    //std::cout << "openPipeToWrite 1 " << fd << std::endl;
    return stream;
};

int NamedPipeOperations::closeWritePipe(FILE *stream, int fd)
{
    //std::cout << "closeWritePipe " << fd << std::endl;
    fclose(stream);
    close(fd);
    //std::cout << "closeWritePipe 1 " << fd << std::endl;
    return 0;
}