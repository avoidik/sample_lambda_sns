#!/usr/bin/env bash

if [ "$#" -ne 1 ]; then
  echo "Input ARN of SNS topic as parameter"
  exit 1
fi

export AWS_PROFILE="personal"

MSGS=0

while [ $MSGS -lt 5 ]; do
    aws sns publish --topic-arn "$1" --subject "TEST CREATE $MSGS" --message "{\"operation\":\"CREATE\",\"id\":${MSGS},\"data\":\"TEST CREATE ${MSGS}\"}"
    MSGS=$((MSGS+1))
done

MSGS=0

while [ $MSGS -lt 5 ]; do
    aws sns publish --topic-arn "$1" --subject "TEST DELETE $MSGS" --message "{\"operation\":\"DELETE\",\"id\":${MSGS},\"data\":\"TEST DELETE ${MSGS}\"}"
    MSGS=$((MSGS+1))
done
