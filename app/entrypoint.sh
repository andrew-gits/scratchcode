#!/bin/bash
echo "$@"
chown -R $USER_ID:$GROUP_ID /node
echo exec "$@"
exec "$@"
# echo exec runuser -u $(id -un $USER_ID) "$@"
# exec runuser -u $(id -un $USER_ID) "$@"
