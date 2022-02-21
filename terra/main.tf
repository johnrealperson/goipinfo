terraform {
    required_providers {
        google = {
            source = "hashicorp/google"
            version = "3.5.0"
        }
    }
}

provider "google" {
    credentials = file("YOUR_DOWNLOAD_FILE_CREDENTIALS_PATH_HERE.json")

    project = "YOUR PROJECT NAME HERE"
    region  = "YOUR REGION"
    zone    = "YOUR ZONE"
}

resource "google_compute_network" "vpc_network" {
name = "terraform-network"

    template { 
        spec {
            containers {
                image = "YOUR_DOCKER_CONTAINER_HERE"
            }
        }    
    }
}
