resource "cloudautomator_job" "example-create-image-job" {
  name           = "example-create-image-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "cron"
  cron_rule_value {
    hour          = "9"
    minutes       = "30"
    schedule_type = "monthly_day_of_week"
    monthly_day_of_week_schedule {
      friday = [-1]
    }
    national_holiday_schedule = "false"
    dates_to_skip = [
      "2023-01-01",
      "2023-01-02",
      "2023-01-03"
    ]
    start_timeout_minutes = "30"
    time_zone             = "Tokyo"
  }

  action_type = "create_image"
  create_image_action_value {
    region                 = "ap-northeast-1"
    specify_image_instance = "tag"
    tag_key                = "env"
    tag_value              = "develop"
    generation             = 10
    image_name             = "test-image"
    description            = "test image"
    reboot_instance        = "true"

    additional_tags {
      key   = "key-1"
      value = "value-1"
    }
    additional_tags {
      key   = "key-2"
      value = "value-2"
    }
    additional_tags {
      key   = "key-3"
      value = "value-3"
    }

    add_same_tag_to_snapshot            = "true"
    trace_status                        = "true"
    recreate_image_if_ami_status_failed = "true"
  }
}
