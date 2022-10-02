#!/bin/bash

# Update node_modules
cp -r /cache/node_modules/. /node/app/node_modules/

# Start the first process
(cd .. && vercel dev --token $(cat /run/secrets/vercel-token))

# Start the second process
ng serve --host 0.0.0.0 &

# # Wait for any process to exit
# wait -n # this will probably cause problems

# # Exit with status of process that exited first
# exit $?
