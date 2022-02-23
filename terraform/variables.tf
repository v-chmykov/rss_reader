variable "PROJECT_ID" {
  description = "The project ID in Google Cloud to use for these resources."
}

variable "REGION" {
  description = "The region in Google Cloud where the resources will be deployed."
}

variable "FUNCTION_NAME" {
  description = "The name of the function to be deployed"
}

variable "ENTRY_POINT" {
  description = "The entrypoint where the function is called"
  default     = "RssReader"
}
