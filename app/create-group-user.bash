#!/bin/bash
if getent group $GROUP_ID; then
    echo "Group exits" && export GROUPNAME=$(getent group $GROUP_ID | cut -d: -f1)
else
    echo "Creating Group" && addgroup --gid $GROUP_ID $USERNAME
fi

if getent passwd $USER_ID; then
    echo "User exists" && export USERNAME=$(id -nu $USER_ID)
else
    echo "Creating user" && adduser --disabled-password --gecos '' --uid $USER_ID --gid $GROUP_ID $USERNAME
fi
