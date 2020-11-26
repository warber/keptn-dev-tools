# KEPTN Pull Strategy Updater (psu)


## What?
The Pull Strategy Updater (PSU) script downloads the latest keptn installation resources, changes all kubernetes image pull strategies to "Always" and repackages 
the installation resources. The modified resources are locally hosted on `http://localhost:8000/keptn-psu-tgz` and can be used in a `keptn install` command to 
reinstall keptn.

## Why?
Because as a developer you want to make sure you're installing the latest up to date images that were built from master branch on your dev cluster.
This script let you do that without manually touching every deployment descriptor or reinstalling the cluster itself.

## Example:

```
./psu.sh 
Temp directory: /tmp/tmp.i8urqzG6gT
Downloading charts from https://storage.googleapis.com/keptn-installer/latest/keptn-0.1.0.tgz
Unpackinng keptn tgz...
Changing Image Pull Strategies to Always
Packing modifind keptn resources...
Charts available at: http://localhost:8000/keptn-psu.tgz
Keptn install command: 
 
keptn install --chart-repo=http://localhost:8000/keptn-psu.tgz --endpoint-service-type=LoadBalancer --use-case=continuous-delivery
 
End program with CTRL+C

```

## Required Dependencies
* wget
* python
