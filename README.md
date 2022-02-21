# goipinfo
IP Information API built with Go and stored on GCP using Terraform.
  
UPDATE:
Was having trouble with my google credentials being invalidated when trying to run more tests through Sunday night. I eventually got the project to push to my GCP console via terraform from my personal PC and have logs to verify, however when I tried to go down the Cloud Run/Docker route there was not a URL to retrieve, and I would need to add some more function to the script to allow that. I have updated the code to include some checks, error testing, and commented source at the top to include a neat video I found (Osint yay) and tried to use some of it as a source however the structure syntax I could not get running corectly. I believe the video may have been showing an older version of Go. 

TL;DR The code is functional if you provide an api key in the apiKey line from VT. 
