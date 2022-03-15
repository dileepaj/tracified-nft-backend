//
//  DateTimeOperations.cpp
//  FlexibleComputerLanguage
//
//  Created by Murtaza Anverali on 5/31/18.
//  Copyright © 2018 Dileepa Jayathilaka. All rights reserved.
//

#include <climits>
#include "DateTimeOperations.h"

long DateTimeOperations::StringToUnix(std::string str)
{
    std::remove(str.begin(), str.end(), '"');
    std::string year = str.substr (0,4);
    std::string month = str.substr (5,2);
    std::string day = str.substr (8,2);
    std::string hours = str.substr (11,2);
    std::string minutes = str.substr (14,2);
    std::string seconds = str.substr (17,2);
    
    int iyear = atoi(year.c_str());
    int imonth = atoi(month.c_str());
    int iday = atoi(day.c_str());
    int ihours = atoi(hours.c_str());
    int iminutes = atoi(minutes.c_str());
    int iseconds = atoi(seconds.c_str());
    
    struct tm when;
    when.tm_year = iyear-1900;
    when.tm_mday = iday;
    when.tm_mon = imonth-1;
    when.tm_hour = ihours;
    when.tm_min = iminutes;
    when.tm_sec = iseconds;
    
    time_t now = mktime(&when);
    long date = now;
    return (date);
}

int DateTimeOperations::SecondsToMonths(long seconds) {
    int months = seconds/(30*3600*24);
    return months;
}

int DateTimeOperations::SecondsToDays(long seconds) {
    int days = seconds/(3600*24);
    return days;
}

int DateTimeOperations::SecondsToYears(long seconds) {
    int days = seconds/(3600*30*24*12);
    return days;
}

long DateTimeOperations::GetDifferenceByUnix(long date1, long date2) {
    time_t tdate1 = date1;
    time_t tdate2 = date2;
    long deff = difftime ( tdate1, tdate2 );
    return deff;
}

long DateTimeOperations::GetDifferenceByString(std::string date1, std::string date2) {
    time_t tdate1 = DateTimeOperations::StringToUnix(date1);
    time_t tdate2 = DateTimeOperations::StringToUnix(date2);
    long deff = difftime ( tdate1, tdate2 );
    return deff;
}

std::string DateTimeOperations::StringToReadable(std::string date) {
    time_t tdate = DateTimeOperations::StringToUnix(date);
    char* dt;
    
    // convert to tm struct for UTC
    tm *gmtm = gmtime(&tdate);
    dt = asctime(gmtm);
    dt[strlen(dt) - 1] = 0;
    return dt;
    
}

long DateTimeOperations::GetDateNow() {
    time_t date = time(0);
    return date;
}

std::string DateTimeOperations::GetDayOfTheWeekShortString(std::string date) {
    time_t tdate = DateTimeOperations::StringToUnix(date);
    char* dt;

    // convert to tm struct for UTC
    tm *gmtm = gmtime(&tdate);
    dt = asctime(gmtm);
    dt[strlen(dt) - 1] = 0;
    std::string datestring = dt;
    return datestring.substr(0,3);
}

std::string DateTimeOperations::GetDayString(std::string date) {
    time_t tdate = DateTimeOperations::StringToUnix(date);
    char* dt;

    // convert to tm struct for UTC
    tm *gmtm = gmtime(&tdate);
    dt = asctime(gmtm);
    dt[strlen(dt) - 1] = 0;
    std::string datestring = dt;
    return datestring.substr(8,2);
}

std::string DateTimeOperations::GetMonthShortString(std::string date) {
    time_t tdate = DateTimeOperations::StringToUnix(date);
    char* dt;

    // convert to tm struct for UTC
    tm *gmtm = gmtime(&tdate);
    dt = asctime(gmtm);
    dt[strlen(dt) - 1] = 0;
    std::string datestring = dt;
    return datestring.substr(4,3);
}

std::string DateTimeOperations::GetYear(std::string date) {
    time_t tdate = DateTimeOperations::StringToUnix(date);
    char* dt;

    // convert to tm struct for UTC
    tm *gmtm = gmtime(&tdate);
    dt = asctime(gmtm);
    dt[strlen(dt) - 1] = 0;
    std::string datestring = dt;
    return datestring.substr(20,5);
}

std::string DateTimeOperations::GetTime24HourFormat(std::string date) {
    time_t tdate = DateTimeOperations::StringToUnix(date);
    char* dt;

    // convert to tm struct for UTC
    tm *gmtm = gmtime(&tdate);
    dt = asctime(gmtm);
    dt[strlen(dt) - 1] = 0;
    std::string datestring = dt;
    return datestring.substr(11,5);
}

std::string DateTimeOperations::GetOldestDate(std::vector<std::string> date) {
    time_t oldesttdate = LLONG_MAX;
    std::string oldesttdateStr = "";
    for (int i = 0; i < date.size(); ++i) {
        time_t tdate = DateTimeOperations::StringToUnix(date.at(i));
        if (tdate <= oldesttdate)
        {
            oldesttdate = tdate;
            oldesttdateStr = date.at(i);
        }
    }
    return oldesttdateStr;
}

std::string DateTimeOperations::GetLatestDate(std::vector<std::string> date) {
    time_t latesttdate = 0;
    std::string latesttdateStr = "";
    for (int i = 0; i < date.size(); ++i) {
        time_t tdate = DateTimeOperations::StringToUnix(date.at(i));
        if (tdate >= latesttdate)
        {
            latesttdate = tdate;
            latesttdateStr = date.at(i);
        }
    }
    return latesttdateStr;
}
