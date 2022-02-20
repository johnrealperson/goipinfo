terraform {
    required_providers {
        google = {
            source = "hashicorp/google"
            version = "3.5.0"
        }
    }
}

provider "google" {
    credentials = file("~\Downloads\flash-flow-341718-c0afb0b97bc4.json")

    project = "<PROJECT_ID>"
    region  = "us-central1"
    zone    = "us-central1-c"
}

resource "google_compute_network" "vpc_network" {
name = "terraform-network"

    template { 
        spec {
            containers {
                image = "us.gcr.io/flash-flow-341718/gcf/us-central1/f33e13bb-fdd4-4c9f-8295-ba6a1c778d9d@sha256:64bc2f33b7d2e795114b9a8c0bd7c10b875b921432d3e4961e95b9cd0ec5e718"
            }
        }    
    }
}
