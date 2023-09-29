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
