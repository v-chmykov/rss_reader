locals {
  root_dir = abspath("../src")
}

data "archive_file" "source" {
  type        = "zip"
  source_dir  = local.root_dir
  output_path = "/tmp/function.zip"
  excludes    = [
    "go.mod",
    "go.sum",
    "rss_reader_test.go"
  ]
}

resource "google_storage_bucket" "bucket" {
  name     = "${var.PROJECT_ID}--${lower(var.FUNCTION_NAME)}"
  location = var.REGION
}

resource "google_storage_bucket_object" "zip" {
  name   = "${data.archive_file.source.output_md5}.zip"
  bucket = google_storage_bucket.bucket.name
  source = data.archive_file.source.output_path
}

resource "google_cloudfunctions_function" "function" {
  available_memory_mb = "128"
  entry_point         = var.ENTRY_POINT
  ingress_settings    = "ALLOW_ALL"

  name                  = var.FUNCTION_NAME
  project               = var.PROJECT_ID
  region                = var.REGION
  runtime               = "go116"
  service_account_email = google_service_account.function-sa.email
  timeout               = 20
  trigger_http          = true
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = "${data.archive_file.source.output_md5}.zip"
}

resource "google_cloudfunctions_function_iam_member" "invoker" {
  project        = google_cloudfunctions_function.function.project
  region         = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}

resource "google_service_account" "function-sa" {
  account_id   = "function-sa"
  description  = "Controls the workflow for the cloud pipeline"
  display_name = "function-sa"
  project      = var.PROJECT_ID
}
