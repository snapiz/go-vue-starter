#!/bin/sh

# build all packages
lerna run build --parallel

# deploy services
sam package --template-file template.yml --output-template-file packaged.yml --s3-bucket ${1}-sam --region ${2}
sam deploy --template-file packaged.yml --stack-name ${1} --capabilities CAPABILITY_IAM --region ${2}

# deploy frontend
./scripts/s3-sync.sh ${1} ${2} web