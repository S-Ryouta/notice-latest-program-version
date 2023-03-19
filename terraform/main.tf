provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_storage_bucket" "bucket" {
  name          = "${var.project_id}-gcf"
  location      = var.region
  force_destroy = true
}

resource "google_storage_bucket_object" "object" {
  name   = "gcf.zip"
  bucket = google_storage_bucket.bucket.name
  source = "../gcf.zip"
}

resource "google_pubsub_topic" "topic" {
  name = "gcf-trigger"
}

resource "google_cloudfunctions_function" "function" {

  environment_variables = {
    REDIS_ADDR           = "your-redis-addr"
    REDIS_PASSWORD       = "your-redis-password"
    LINE_CHANNEL_SECRET  = "your-line-channel-secret"
    LINE_CHANNEL_TOKEN   = "your-line-channel-token"
    DISCORD_BOT_TOKEN    = "your-discord-bot-token"
    DISCORD_CHANNEL_ID   = "your-discord-channel-id"
  }

  name        = "check-update-version"
  description = "Check and notify for programming language version updates"
  runtime     = "go120"

  available_memory_mb   = 512
  timeout               = 60
  max_instances         = 1
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.object.name

  event_trigger {
    event_type = "google.pubsub.topic.publish"
    resource   = google_pubsub_topic.topic.id
  }

  entry_point = "CheckAndUpdateVersionHandler"
}

resource "google_cloud_scheduler_job" "job" {
  name             = "check-update-version-scheduler"
  description      = "Trigger the function to check and notify for programming language version updates every week on Monday at 18:00"
  schedule         = "*/30 * * * *"
  time_zone        = "UTC"
  attempt_deadline = "60s"

  pubsub_target {
    topic_name = google_pubsub_topic.topic.id
    data = base64encode("golang")
  }
}

resource "google_redis_instance" "memory_store_instance" {
  name          = var.redis_instance_name
  tier          = var.redis_tier
  memory_size_gb = var.redis_size_gb

  location_id = var.location
  region      = var.region
  authorized_network = google_compute_network.vpc_network.self_link
}

resource "google_compute_network" "vpc_network" {
  name                    = var.redis_network
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "vpc_subnetwork" {
  name          = "${var.redis_network}-subnetwork"
  region        = var.region
  network       = google_compute_network.vpc_network.self_link
  ip_cidr_range = "10.0.0.0/24"
}

resource "google_compute_firewall" "allow_internal" {
  name    = "allow-internal"
  network = google_compute_network.vpc_network.name

  allow {
    protocol = "tcp"
    ports    = ["6379"]
  }

  source_ranges = [google_compute_subnetwork.vpc_subnetwork.ip_cidr_range]
}

