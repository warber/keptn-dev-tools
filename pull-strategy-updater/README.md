# KEPTN Pull Strategy Updater (psu)


## What?
The Pull Strategy Updater (PSU) script downloads the latest keptn installation resources, changes all kubernetes image pull strategies to "Always" and repackages 
the installation resources. The modified resources are locally hosted on `http://localhost:8000/keptn-psu-tgz` and can be used in a `keptn install` command to 
reinstall keptn.

## Why?
Because as a developer you want to make sure you're installing the latest up to date images that were built from master branch on your dev cluster.
This script let you do that without manually touching every deployment descriptor or reinstalling the cluster itself.


## Build and run the docker image

```
docker build -t psu:latest .
docker run --rm -ti -p8000:8000 psu:latest
``` 

