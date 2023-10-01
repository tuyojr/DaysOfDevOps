#!/bin/bash

# Update the log file everytime the script is run
date >> /home/tuyojr/update_logs.txt

# Update the repository information
sudo apt update

# Check if there are any upgrades available
if [ $(sudo apt list --upgradable 2>/dev/null | grep -c 'upgradable') -gt 0 ]; then
    echo "Updates are available. Performing upgrades..."
    sudo apt dist-upgrade -y
else
    echo "No updates available."
fi
