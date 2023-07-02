# What Is SNS?

Amazon Simple Notification Service (Amazon SNS) is a managed service that provides message delivery from publishers to subscribers (also known as producers and consumers).

Amazon SNS has no upfront costs. You pay based on the number of messages that you publish, the number of notifications that you deliver, and any additional API calls for managing topics and subscriptions. Delivery pricing varies by endpoint type. You can get started for free with the Amazon SNS free tier.

## Getting Started with Amazon SNS

> Step 1: Create a Topic

1. Sign in to the [Amazon SNS console](https://console.aws.amazon.com/sns/home).
2. In the left navigation pane, choose Topics --> choose Create topic.
3. By default, the console creates a FIFO topic. **Choose Standard** --> enter a Name --> Create topic.

> Step 2: Create a Subscription

1. In the left navigation pane, choose Subscriptions --> Create subscription.
2. For Topic ARN, choose the ARN of the topic that you created in the previous step.
3. For Protocol, choose Email.
4. For Endpoint, enter your email address.
5. Check your email inbox and choose Confirm subscription in the email from AWS Notifications. The sender ID is usually "<no-reply@sns.amazonaws.com>". If you don't see the email, check your spam folder.

![sns_confirmation_subscription](https://github.com/tuyojr/100DaysOfDevOps/blob/main/day_02_introduction_to_simple_notification_service_(SNS)/sns_confirm_subscription.png)

> Step 3: Publish a Message to the Topic

1. In the left navigation pane, choose Topics --> choose the topic that you created in the previous step.
2. Choose Publish message.
3. For Subject, enter a subject for your message.
4. For Message body, enter a message.
5. Choose Publish message.

![sns_publish_message](https://github.com/tuyojr/100DaysOfDevOps/blob/main/day_02_introduction_to_simple_notification_service_(SNS)/sns_publish_message.png)

> Step 4: Delete the Subscription and Topic

1. In the left navigation pane, choose Subscriptions --> choose the subscription that you created in the previous step.
2. Check the subscription --> Delete.
3. In the left navigation pane, choose Topics --> choose the topic that you created in the previous step.
4. Check the topic --> Delete.

## Setting up AWS SNS With Terraform

> other files like the `provider.tf` and `variable.tf` are in this directory.

```terraform
resource "aws_sns_topic" "shisui" {
    name            = "shisui"
    display_name    = "body_flicker"
}

resource "aws_sns_topic_subscription" "uchiha" {
    topic_arn = "${aws_sns_topic.shisui.arn}"
    protocol = "email"
    endpoint = "${var.alarms_email}"
}

output "sns_topic_arn" {
    value = "${aws_sns_topic.shisui.arn}"
}
output "sns_subscription_arn" {
    value = "${aws_sns_topic_subscription.uchiha.arn}"
}
```
