#!/bin/bash

# Update the instance
printf "update instance...\n"
sudo apt -y update
printf ".. done\n"

# Install GNU Privacy Guard dependency software
printf "installing gnupg...\n"
sudo apt -y install gnupg
printf ".. done\n"

# Install jq for manipulating Json in Bash
printf "installing jq...\n"
sudo apt-get --assume-yes install jq
printf ".. done\n"

# Create and setup new user with less privileges
printf "creating user with home directory...\n"
sudo groupadd csye6225
sudo useradd -s /bin/false -g csye6225 -d /opt/user1 -m user1
sudo mkdir /opt/user1/appfiles
printf ".. done\n"

# Install cloud watch agent
printf "installing cloud watch agent...\n"
sudo wget https://amazoncloudwatch-agent.s3.amazonaws.com/debian/amd64/latest/amazon-cloudwatch-agent.deb
sudo dpkg -i -E amazon-cloudwatch-agent.deb
printf ".. done\n"

# Move files to correct location in the instance
printf "moving files to correct directories...\n"
sudo mv /tmp/restful-api /opt/user1/appfiles
sudo mv /tmp/cloud-course-webapp.service /etc/systemd/system/
sudo mv /tmp/users.csv /opt/user1/appfiles
sudo mv /tmp/tmp-cloudwatch-config.json /opt/user1/appfiles/
printf ".. done\n",

# Change ownership of entities
printf "Change ownership of entities...\n"
sudo chown user1:csye6225 /etc/systemd/system/cloud-course-webapp.service
sudo chown -R user1:csye6225 /opt/user1
sudo chown -R :csye6225 /opt/aws/
sudo chmod -R g+rwx /opt/aws/
sudo chmod 664 /etc/systemd/system/cloud-course-webapp.service
sudo systemctl daemon-reload
sudo systemctl enable cloud-course-webapp
printf ".. done\n"