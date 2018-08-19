# GeekMD: NoteClerk v0.3.1
NoteClerk is a micro-service dedicated to managing clinical notes from the Noted library. It uses protcol buffers and 
gRPC and so the API is relatively language agnostic.

|Branch|Build Status|Coverage %|
|:---:|:---:|:---:|
|Master|![master-build](https://travis-ci.org/geekmdio/noteclerk.svg?branch=master)| [![codecov-master](https://codecov.io/gh/geekmdio/noteclerk/branch/master/graph/badge.svg)](https://codecov.io/gh/geekmdio/noteclerk) |
|Development|![dev-build](https://travis-ci.org/geekmdio/noteclerk.svg?branch=development)| [![codecov-development](https://codecov.io/gh/geekmdio/noteclerk/branch/development/graph/badge.svg)](https://codecov.io/gh/geekmdio/noteclerk)  |


###RELEASE NOTES v0.3.0
- Updated to ehrproto v0.3.1
- See [repo](https://github.com/geekmdio/ehrprotorepo) for changes.


###SETUP
- Set the environmental variable NOTECLERK_ENVIRONMENT
- Run `setup.sh`; not tested in Windows, but should be able to run in Windows through Bash which is available in a Linux subsystem in Windows 10, or via the various terminal emulators.
    - Of note, you may run into problems creating folders for the config files and log files if your permissions are not set properly. Recommend running this server under limited user and keeping log in default directory.
    - Additionally, please note that each time `setup.sh` is run it will create a new config file for the existing environment.