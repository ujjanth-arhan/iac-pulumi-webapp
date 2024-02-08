#!/bin/bash

echo '{
    "agent": {
        "metrics_collection_interval": 10,
        "logfile": "/var/logs/amazon-cloudwatch-agent.log"
    },
    "logs": {
        "logs_collected": {
            "files": {
                "collect_list": [
                    {
                        "file_path": "/opt/user1/appfiles/log",
                        "log_group_name": "csye6225",
                        "log_stream_name": "webapp"
                    }
                ]
            }
        },
        "log_stream_name": "cloudwatch_log_stream"
    },
    "metrics":{
        "metrics_collected":{
            "statsd":{
                "service_address":":8125",
                "metrics_collection_interval":15,
                "metrics_aggregation_interval":10
            }
        }
    }
}'>/opt/tmp-cloudwatch-config.json
mv /opt/tmp-cloudwatch-config.json ./builds/