//
// Created by Isini_Dananjana on 2021-07-12.
//


#include "FCLWrapper.h"
#include "ELInterpretter.h"
#include "CommonIncludes.h"
#include "CommonIncludes.h"
#include "LDAL_Wrapper.h"
#include <string>
//#include "LogJsonParser.h"
//#include "OTPParser.h"



void FCLWrapper::RunELInterpretter(const char *defFilepath) {
    ELInterpretter intp;
    intp.EvaluateCase(defFilepath);

}

std::string FCLWrapper::GetLDALResult(const char *defFilePath) {
    LDAL_Wrapper ldalWrapper ;
    return ldalWrapper.GetLDALResult(defFilePath);

}

std::string FCLWrapper::GetTDPResult(const char *defFilePath) {
    LDAL_Wrapper ldalWrapper;
    return  ldalWrapper.GetTDPResult(defFilePath);
}

std::string  FCLWrapper::GetLogLDALResult(const char *defFilePath) {
    LDAL_Wrapper ldalWrapper;
    return  ldalWrapper.GetLOGLDALResult(defFilePath);

}

std::string FCLWrapper::GetOTPResult(const char *defFilePath){
    LDAL_Wrapper LDAL_Wrapper;
    return LDAL_Wrapper.GetOTPResult(defFilePath);

}

std::string FCLWrapper::GetBuildResult(const char *defFilePath){
    LDAL_Wrapper LDAL_Wrapper;
    return LDAL_Wrapper.GetBuildResult(defFilePath);

}

std::string  FCLWrapper::GetLogLDALResultV2(const char *defFilePath ,const char *query,const char *json ) {
    LDAL_Wrapper ldalWrapper;
    return  ldalWrapper.GetLOGLDALResultV2(defFilePath,query,json);

}

