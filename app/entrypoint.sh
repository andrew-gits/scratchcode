#!/bin/bash
echo "$@"
chown -R $USER_ID:$GROUP_ID /cache /node
echo exec runuser -u $(id -un $USER_ID) "$@"
exec runuser -u $(id -un $USER_ID) "$@"
