#!/usr/bin/env bash
set -e

if [[ -z "$USER_ID" ]]; then
    echo "ERROR: please set USER_ID" >&2
    echo $USER_ID
    exit 1
fi
if [[ -z "$GROUP_ID" ]]; then
    echo "ERROR: please set GROUP_ID" >&2
    echo $GROUP_ID
    exit 1
fi

# Use this code if you want to create a new user account:
getent group $GROUP_ID || addgroup --gid $GROUP_ID user
getent passwd $USER_ID || adduser --disabled-password --gecos '' --uid $USER_ID --gid $GROUP_ID user

# -OR-
# Use this code if you want to modify an existing user account:
# groupmod --gid "$GROUP_ID" user
# usermod --uid "$USER_ID" user

# Drop privileges and execute next container command, or 'bash' if not specified.
# if [[ $# -gt 0 ]]; then
#     exec sudo -u -H app -- "$@"
# else
#     exec sudo -u -H app -- bash
# fi

echo Running \"exec "$@"\"
exec "$@"
