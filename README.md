# goipinfo
IP Information API built with Go and stored on GCP using Terraform.
  
UPDATE:
Decided to try and store the project on Cloud Run but had issues with Terraform refusing my credentials file and could not find any sources as to why. 
I was successful it uploading the whole project to my GCP console with terraform and have the verified logs. I realized when I got to this point the URL portion
of the assessment wasn't really obtained with this method. So, I have updated the code to include some necessary tests, uninclude some personal information and 
includes a .env_sample, and commented out blocks where I had trouble getting everything going. 
