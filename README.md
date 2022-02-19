# goipinfo
IP Information API built with Go and stored on GCP using Terraform.

Big OOPSIES I am aware of but had trouble configuring:
  Readily available API key in plaintext due to the fact I had trouble figuring out calling environment variables in Go.
  
  JSON data is not properly parsed ~ had issues figuring out how to label the data in structs that I specifically wanted to grab.
  Settled on grabbing the entire API response as well as Objects related to the IP.
  
  
