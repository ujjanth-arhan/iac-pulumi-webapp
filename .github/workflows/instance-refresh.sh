#!/bin/bash

LATEST_IMAGE_ID=$(aws ec2 describe-images --executable-users self --query 'sort_by(Images, &CreationDate)[-1].ImageId' --output text)

LAUNCH_TEMPLATE_VERSION=$(aws ec2 describe-launch-templates --query 'sort_by(LaunchTemplates, &CreateTime)[-1].LatestVersionNumber' --output text)
if [ $LAUNCH_TEMPLATE_VERSION == "None" ]
then
    echo "No launch template to update. Please setup your infra and run again"
    exit 1
fi

LAUNCH_TEMPLATE_ID=$(aws ec2 describe-launch-templates --query 'sort_by(LaunchTemplates, &CreateTime)[-1].LaunchTemplateId' --output text)

aws ec2 create-launch-template-version --launch-template-id $LAUNCH_TEMPLATE_ID --source-version $LAUNCH_TEMPLATE_VERSION --launch-template-data "ImageId=$LATEST_IMAGE_ID"

AUTOSCALING_GROUP_NAME=$(aws autoscaling describe-auto-scaling-groups --query 'sort_by(AutoScalingGroups, &CreatedTime)[-1].AutoScalingGroupName' --output text)

CUR_TEMPLATE_VERSION=$(aws ec2 describe-launch-templates --query 'sort_by(LaunchTemplates, &CreateTime)[-1].LatestVersionNumber' --output text)

aws autoscaling update-auto-scaling-group --auto-scaling-group-name $AUTOSCALING_GROUP_NAME --launch-template "LaunchTemplateId=$LAUNCH_TEMPLATE_ID,Version=$CUR_TEMPLATE_VERSION"

REFRESH_ID=$(aws autoscaling start-instance-refresh --auto-scaling-group-name $AUTOSCALING_GROUP_NAME --query "InstanceRefreshId" --output text)

REFRESH_STATE="Pending"
while [ $REFRESH_STATE == "Pending" ] || [ $REFRESH_STATE == "InProgress" ]
do
    echo $REFRESH_STATE
    sleep 5
    REFRESH_STATE=$(aws autoscaling describe-instance-refreshes --auto-scaling-group-name $AUTOSCALING_GROUP_NAME --instance-refresh-ids $REFRESH_ID --query 'InstanceRefreshes[0].Status' --output text)
done

printf "Instance state refresh status:\n"
echo $REFRESH_STATE

if [ $REFRESH_STATE != "Successful" ]
then
    exit 1
fi