#!/bin/bash

aws ssm put-parameter --name PhoneNumber --type String --overwrite --value 44$1
