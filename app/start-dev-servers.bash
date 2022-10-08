#!/bin/bash

# If a Vercel Project is not linked this will start the cli so you can link to a project
if [ ! -f /node/.vercel/project.json ]; then
    echo "Please link to a project on Vercel for local development to work"
    ./vercel-setup.bash
    exit 0
fi

# install new packages if needed
npm install

# Start the first process
vercel dev --token $(cat /run/secrets/vercel-token) --yes

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?
