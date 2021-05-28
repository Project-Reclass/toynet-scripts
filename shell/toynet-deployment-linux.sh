#!/bin/bash

Destination="${HOME}/Documents/usb-deployment"

echo ${Destination}

if [[ -d ${Destination} ]]
then    
    echo "${Destination} Exists!" 
else 
    echo "${Destination} does not exist" && mkdir -p ${Destination}
fi

echo "copying from usb-deployment to ${Destination}"  
cp -r ./toynet/ ${Destination}
cp -r ./dependencies/ ${Destination}