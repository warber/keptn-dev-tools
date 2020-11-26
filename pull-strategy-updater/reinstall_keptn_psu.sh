#!/bin/bash

# exit when any command fails
set -e

echo "Uninstall keptn"
#keptn uninstall

echo "Installing keptn"
keptn install --chart-repo=http://localhost:8000/keptn-psu.tgz --endpoint-service-type=LoadBalancer --use-case=continuous-delivery


read -p "(Re)Installing Istio? " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
then
    bash <(curl -s https://raw.githubusercontent.com/keptn/examples/release-0.7.2/istio-configuration/configure-istio.sh)
fi

echo "Getting Keptn Endpoint"
KEPTN_ENDPOINT=http://$(kubectl -n keptn get ingress api-keptn-ingress -ojsonpath='{.spec.rules[0].host}')/api
echo $KEPTN_ENDPOINT
echo

echo "Getting Keptn API Token"
KEPTN_API_TOKEN=$(kubectl get secret keptn-api-token -n keptn -ojsonpath='{.data.keptn-api-token}' | base64 --decode)
echo $KEPTN_API_TOKEN
echo

echo "Authenticating keptn"
keptn auth --endpoint=$KEPTN_ENDPOINT --api-token=$KEPTN_API_TOKEN
echo

echo "Credentials for bridge:"
keptn configure bridge --output
