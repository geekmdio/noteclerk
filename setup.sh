#!/usr/bin/env bash

# NOTE: You must have the environmental variable NOTECLERK_ENVIRONMENT set.

VERSION="0.3.0"
LOG_DIR="${HOME}/.noteclerk/log"
LOG_PATH="${LOG_DIR}/server.log"
SERVER_PROTOCOL="tcp"
SERVER_IP="localhost"
SERVER_PORT="50051"
DB_IP="localhost"
DB_PORT="5433"
DB_USERNAME="USERNAME_REQUIRED"
DB_PASSWORD="PASSWORD_REQUIRED"
DB_NAME="noteclerk"
DB_SSL_MODE="disable"
CONFIG_DIRECOTRY="config"
CONFIG_FILE_PATH="${CONFIG_DIRECOTRY}/config.${NOTECLERK_ENVIRONMENT}.json"

get_user_input() {
    # Generate log directory and file based on input
    printf "Log file directory (default: ${LOG_DIR}): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         LOG_DIR=${USER_INPUT}
    fi
    mkdir -p ${LOG_DIR}
    LOG_PATH="${LOG_DIR}/server.log"
    touch ${LOG_PATH}

    printf "Server protocol (default: ${SERVER_PROTOCOL}): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         SERVER_PROTOCOL=${USER_INPUT}
    fi

    printf "Server IP Address (default: ${SERVER_IP}): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         SERVER_IP=${USER_INPUT}
    fi

    printf "Server port (default: ${SERVER_PORT}): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         SERVER_PORT=${USER_INPUT}
    fi

    printf "Database IP Address (default: ${DB_IP}): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         DB_IP=${USER_INPUT}
    fi

    printf "Database Username (required): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         DB_USERNAME=${USER_INPUT}
    fi

    printf "Database Password (required): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         DB_PASSWORD=${USER_INPUT}
    fi

    printf "Database Name (default: ${DB_NAME}): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         DB_NAME=${USER_INPUT}
    fi

    printf "Database SSL Mode (default: ${DB_SSL_MODE}): "
    read -r USER_INPUT
    if [ "${USER_INPUT}" != "" ]; then
         DB_SSL_MODE=${USER_INPUT}
    fi
}

write_to_config() {
    echo '{' >> ${CONFIG_FILE_PATH}
    echo '  "Version": "'${VERSION}'",' >> ${CONFIG_FILE_PATH}
    echo '  "LogPath": "'${LOG_PATH}'",' >> ${CONFIG_FILE_PATH}
    echo '  "ServerProtocol": "'${SERVER_PROTOCOL}'",' >> ${CONFIG_FILE_PATH}
    echo '  "ServerIp": "'${SERVER_IP}'",' >> ${CONFIG_FILE_PATH}
    echo '  "ServerPort": "'${SERVER_IP}'",' >> ${CONFIG_FILE_PATH}
    echo '  "DbIp": "'${DB_IP}'",' >> ${CONFIG_FILE_PATH}
    echo '  "DbPort": "'${DB_PORT}'",' >> ${CONFIG_FILE_PATH}
    echo '  "DbUsername": "'${DB_USERNAME}'",' >> ${CONFIG_FILE_PATH}
    echo '  "DbPassword": "'${DB_PASSWORD}'",' >> ${CONFIG_FILE_PATH}
    echo '  "DbName": "'${DB_NAME}'",' >> ${CONFIG_FILE_PATH}
    echo '  "DbSslMode": "'${DB_SSL_MODE}'"' >> ${CONFIG_FILE_PATH}
    echo '}' >> ${CONFIG_FILE_PATH}
}


main() {
    echo "NoteClerk v${VERSION} Setup"
    echo "==========================="
    # If the directory does not already exist, create the config directory
    echo "Checking for presence of ${CONFIG_DIRECOTRY} directory..."
    mkdir ${CONFIG_DIRECOTRY} 2> /dev/null
    if [ $? -ne 0 ]; then
        echo "Found!"
    else
        echo "Not found, creating ${CONFIG_DIRECOTRY}"
    fi

    get_user_input

    write_to_config

    echo "Setup complete!"
    echo ""
    echo "The following configuration file was written to ${CONFIG_FILE_PATH}."
    cat ${CONFIG_FILE_PATH}


    if [ "${DB_PASSWORD}" == "PASSWORD_REQUIRED" ] || [ "${DB_USERNAME}" == "USERNAME_REQUIRED" ]; then
        echo "NOTE: The username and/or password were not set. You will need to configure these manually at ${CONFIG_FILE_PATH}."
    fi
}

main