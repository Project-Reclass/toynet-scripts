# Toynet Scripts

Scripts to allow a user to install, update, and retreive data for Toynet

# Dependencies

Toynet is run in Docker containers you'll need to have Docker installed on your OS first

You can download the latest version of docker from the official website: 
https://www.docker.com/get-started

To run the Golang scripts you'll also need to have go installed which you can do here: 
https://golang.org

# Running Toynet

On Windows you can run the powershell script to get started with toynet

Or on Linux you can run the bash script

Optionally, you can run the golang script for your specific OS/hardware

This can be accomplished by navigating to the go binary you need and running it in a terminal

e.g. ```./toynet-linux```

All scripts should make toynet available at: http://localhost:3000

# Updating Toynet

Toynet updates can be retrieved from this repository as it will remain up to date. 

You can delete the entire old version with ```rm -rf toynet-scripts```

And retrieve the updates from this repo with ```git clone https://github.com/Project-Reclass/toynet-scripts.git```

You can also accomplish this with: ```git pull origin master```

Optionally, you can run toynet yourself and pull the images directly from dockerhub

However you'll need to pull both the frontend and the backend with: 

```docker pull projectreclass/toynet-frontend:latest``` and 
````docker push projectreclass/toynet-backend:latest```
