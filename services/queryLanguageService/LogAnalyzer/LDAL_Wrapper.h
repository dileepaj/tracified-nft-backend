//
// Created by tharindu on 8/27/2021.
//

#ifndef CODE2_LDAL_WRAPPER_H
#define CODE2_LDAL_WRAPPER_H
#include "CommonIncludes.h"
#include "TypeDefs.h"



class LDAL_Wrapper {
public:
  MSTRING GetLDALResult(MSTRING defFilePath);
  MSTRING GetTDPResult(MSTRING defFilePath);
  MSTRING GetLOGLDALResult(MSTRING defFilePath);
  MSTRING GetOTPResult(MSTRING defFilePath);
  MSTRING GetBuildResult(std::string defFilePath);
  MSTRING GetLOGLDALResultV2(MSTRING defFilePath,MSTRING query,MSTRING json);
  MSTRING GetTDPResultV2(MSTRING defFilePath,MSTRING query,MSTRING json);
  MSTRING GetOTPResultV2(MSTRING defFilePath,MSTRING query,MSTRING json);
};


#endif //CODE2_LDAL_WRAPPER_H
