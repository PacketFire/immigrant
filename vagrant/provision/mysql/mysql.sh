#!/bin/bash

apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -yqq mysql-server
