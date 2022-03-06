//
// Created by murtaza on 7/12/18.
//

#ifndef FLEXIBLECOMPUTERLANGUAGE_LOGGER_H
#define FLEXIBLECOMPUTERLANGUAGE_LOGGER_H

#include <thread>
#include "easylogging++.h"

class Logger
{

  public:
    static void ConfigureLogger()
    {
        // Console logger with color
        el::Loggers::addFlag(el::LoggingFlag::ColoredTerminalOutput);
        el::Configurations defaultConf;
        defaultConf.setToDefault();
        defaultConf.setGlobally(el::ConfigurationType::Filename, "logs/info.%datetime{%Y-%M-%d_%H-%m-%s}.log");
        el::Loggers::reconfigureLogger("default", defaultConf);
        std::thread logRotatorThread([]() {
            const std::chrono::seconds wakeUpDelta = std::chrono::hours(24);
            auto nextWakeUp = std::chrono::system_clock::now() + wakeUpDelta;

            while (true)
            {
                std::this_thread::sleep_until(nextWakeUp);
                nextWakeUp += wakeUpDelta;
                LOG(INFO) << "About to rotate log file!";
                auto L = el::Loggers::getLogger("default");
                if (L == nullptr)
                    LOG(ERROR) << "Oops, it is not called default!";
                else
                    L->reconfigure();
            }
        });

        logRotatorThread.detach();
    };
};

#endif //FLEXIBLECOMPUTERLANGUAGE_LOGGER_H
