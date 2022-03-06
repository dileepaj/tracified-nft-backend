//
// Created by Isini-Dananjana on 2021-07-12.
//


#ifndef CODE_FCLWRAPPER_H
#define CODE_FCLWRAPPER_H
#include <string>

class FCLWrapper{
public:
    void RunELInterpretter(const char* defFilepath);
    std::string GetLDALResult(const char* defFilePath);
    std::string GetTDPResult(const char* defFilePath);
    std::string  GetLogLDALResult(const char* defFilePath);
    std::string  GetOTPResult(const char* defFilePath);
    std::string  GetBuildResult(const char* defFilePath);
    std::string  GetLogLDALResultV2(const char *defFilePath ,const char *query,const char *json);
};

#endif //CODE_FCLWRAPPER_H
