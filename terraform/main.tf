terraform {
  backend "gcs" {
    # Bucket is passed in via cli arg. Eg, terraform init -reconfigure -backend-configuration=dev.tfbackend
  }
}

provider "google" {
  project = var.PROJECT_ID
  region  = var.REGION
}

provider "google-beta" {
  project = var.PROJECT_ID
  region  = var.REGION
}
