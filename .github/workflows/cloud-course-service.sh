#!/bin/bash

echo "[Unit]
Description=Cloud Course Web App
After=cloud-final.service

[Service]
User=user1
Group=csye6225
WorkingDirectory=/opt/user1/appfiles
EnvironmentFile=.env
ExecStart=/opt/user1/appfiles/restful-api infinity
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=cloud-init.target" > /opt/cloud-course-webapp.service
sudo mv /opt/cloud-course-webapp.service ./builds/