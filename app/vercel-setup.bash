#!/bin/bash
cd /node
vercel --token $(cat /run/secrets/vercel-token)
