variable "project_id" {
  type        = string
  default     = "notice-latest-program-version"
  description = "notice-latest-program-version"
}

variable "region" {
  type        = string
  description = "The region for the resources"
  default     = "asia-northeast1"
}

variable "location" {
  type        = string
  description = "The region for the resources"
  default     = "asia-northeast1-a"
}

variable "redis_instance_name" {
  type        = string
  default     = "notice-latest-program-version-memory"
  description = "The name of the MemoryStore instance"
}

variable "redis_tier" {
  type        = string
  default     = "BASIC"
  description = "The MemoryStore tier (BASIC or STANDARD)"
}

variable "redis_size_gb" {
  type        = number
  default     = 1
  description = "The MemoryStore memory size in GiB"
}

variable "redis_network" {
  type        = string
  default     = "notice-latest-program-redis-network"
  description = "The VPC network name for the MemoryStore instance"
}
