#!/bin/bash

# Check if an SSH key path is provided as an argument
if [ "$#" -ne 1 ]; then
	echo "Usage: $0 /path/to/ssh-key"
	exit 1
fi

# First argument is the path to the SSH key
SSH_KEY_PATH="$1"

# Define remote machine details
REMOTE_USER="root"
REMOTE_HOST="104.248.43.157"

# Define container names
CONTAINER1="minitwit-api-instance"
CONTAINER2="minitwit-app-instance"

# Define the filename to store Docker logs on the remote machine
REMOTE_LOG_FILE_1="/var/log/docker-analytics/docker_logs_api_$(date +%Y%m%d_%H%M%S).txt"
REMOTE_LOG_FILE_2="/var/log/docker-analytics/docker_logs_app_$(date +%Y%m%d_%H%M%S).txt"

# Define the local directory where you want to save the logs
LOCAL_PATH_TO_SAVE_LOGS="$(pwd)/docker-logs"

# SSH into the remote machine and execute Docker log commands
ssh -i "${SSH_KEY_PATH}" ${REMOTE_USER}@${REMOTE_HOST} <<EOF
docker logs ${CONTAINER1} > ${REMOTE_LOG_FILE_1}
docker logs ${CONTAINER2} >> ${REMOTE_LOG_FILE_2}
exit
EOF

# Copy the log file from the remote machine to your local machine
scp -i "${SSH_KEY_PATH}" ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_LOG_FILE_1} "${LOCAL_PATH_TO_SAVE_LOGS}"
scp -i "${SSH_KEY_PATH}" ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_LOG_FILE_2} "${LOCAL_PATH_TO_SAVE_LOGS}"

# Optionally, remove the log file from the remote machine after copying
ssh -i "${SSH_KEY_PATH}" ${REMOTE_USER}@${REMOTE_HOST} "rm ${REMOTE_LOG_FILE_1}"
ssh -i "${SSH_KEY_PATH}" ${REMOTE_USER}@${REMOTE_HOST} "rm ${REMOTE_LOG_FILE_2}"

echo "Logs copied to ${LOCAL_PATH_TO_SAVE_LOGS}/${REMOTE_LOG_FILE}"
