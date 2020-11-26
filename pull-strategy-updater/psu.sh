#!/bin/bash

trap cleanup INT

KEPTN_TGZ="keptn-0.1.0.tgz"
KEPTN_TGZ_MOD="keptn-psu.tgz"
TMP_DIR=`mktemp -d`
URL="https://storage.googleapis.com/keptn-installer/latest/${KEPTN_TGZ}"

cleanup() {
  echo "Removing ${TMP_DIR}"
  rm -rf ${TMP_DIR}
  echo "Removing ./${KEPTN_TGZ_MOD}"
  rm -f ./${KEPTN_TGZ_MOD}
}


echo "Temp directory: ${TMP_DIR}"
echo "Downloading charts from ${URL}"
wget -q -P ${TMP_DIR} ${URL}
WGET_ERR=$?

if [ ${WGET_ERR} -ne 0 ]; then
  echo "Unable to download charts from ${URL}"
  exit 1
fi

echo "Unpackinng keptn tgz..."
tar -C ${TMP_DIR} -xzf ${TMP_DIR}/${KEPTN_TGZ}

echo "Changing Image Pull Strategies to Always"
( shopt -s globstar dotglob;
    for file in ${TMP_DIR}/keptn/**; do
        if [[ -f $file ]] && [[ -w $file ]]; then
           # echo "processing ${file}"
            sed -i -- 's/IfNotPresent/Always/g' "$file"
        fi
    done
)

echo "Packing modifind keptn resources..."
tar czfC ./${KEPTN_TGZ_MOD} ${TMP_DIR} keptn

echo "Charts available at: http://localhost:8000/keptn-psu.tgz"
echo "Keptn install command: "
echo " "
echo "keptn install --chart-repo=http://localhost:8000/keptn-psu.tgz --endpoint-service-type=LoadBalancer --use-case=continuous-delivery"
echo " "
echo "End program with CTRL+C"

python3 -m http.server 2>&1 > /dev/null


