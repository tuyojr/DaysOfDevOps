# Setup a CPU Usage Alarm using the AWS Management Console.

1. First we need to create a demo instance we can use to generate CPU usage. Follow the steps [**here**](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EC2_GetStarted.html) to create an EC2 instance.
2. Open the CloudWatch console at [**here**](https://console.aws.amazon.com/cloudwatch/)
3. Steps followed to create a CPU usage alarm are more detailed [here](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/US_AlarmAtThresholdEC2.html)
4. Choose Create alarm.
5. In step 1, choose Select metric.
6. Give your graph a title, for eg: CPU Utilization
7. Click Browse and select EC2 Metrics --> Per-Instance Metrics --> Paste your instance ID in the search box and select it --> Select CPUUtilization (under metric name) --> Select Metric
8. Under Specify metric and conditions, for Statistic choose Average, choose one of the predefined percentiles, or specify a custom percentile (for example, p95.45).
9. For Period, choose 5 minutes.
10. Under Conditions, specify the following:
    - For Threshold type, choose Static.
    - For Whenever CPUUtilization is, specify Greater. Under than..., specify the threshold that is to
    trigger the alarm to go to ALARM state if the CPU utilization exceeds this percentage. For example, 70. Click Next.
11. Under Notification, choose In alarm and select an SNS topic to notify when the alarm is in ALARM state. To have the alarm send multiple notifications for the same alarm state or for different alarm states, choose Add notification. To have the alarm not send notifications, choose Remove. Choose Next.
12. Enter a name and description for the alarm. Choose Next.
13. Under Preview and create, confirm that the information and conditions are what you want, then choose Create alarm.

## Setup CPU Usage Alarm using the AWS CLI

1. Create an alarm using the put-metric-alarm command

```sh
aws cloudwatch put-metric-alarm --alarm-name cpu-mon --alarm-description "Alarm when CPU exceeds 70 percent" --metric-name CPUUtilization --namespace AWS/EC2 --statistic Average --period 300 --threshold 70 --comparison-operator GreaterThanThreshold  --dimensions "Name=InstanceId,Value=i-12345678" --evaluation-periods 2 --alarm-actions arn:aws:sns:us-east-1:111122223333:MyTopic --unit Percent
```

2. Using the command line, we can test the Alarm by forcing an alarm state change using a set-alarm-state command

3. Change the alarm-state from INSUFFICIENT_DATA to OK

```sh
aws cloudwatch set-alarm-state --alarm-name "cpu-monitoring" --state-reason "initializing" --state-value OK
```

4. Change the alarm-state from OK to ALARM

```sh
aws cloudwatch set-alarm-state --alarm-name "cpu-monitoring" --state-reason "initializing" --state-value ALARM
```

5. Lastly, check if you have received an email notification about the alarm.

## Output
![cloud_watch](https://github.com/tuyojr/100DaysOfDevOps/blob/main/day_01_introduction_to_cloudwatch_metrics/cloud_watch.png)
