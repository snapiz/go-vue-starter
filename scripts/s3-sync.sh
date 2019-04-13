#!/bin/sh
set -e

JQOUTPUTMAP='.StackResourceDetail.PhysicalResourceId'
FRONTENDBUCKET=`aws cloudformation describe-stack-resource --logical-resource-id=FrontendBucket --stack-name ${1} --region ${2} | jq -e --raw-output "${JQOUTPUTMAP}"`

cd apps/${3}/dist
aws s3 sync . s3://${FRONTENDBUCKET} --delete