#!/bin/bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
apt-get install build-essential libtool autotools-dev automake pkg-config bsdmainutils python3
apt-get install libssl-dev libevent-dev libboost-system-dev libboost-filesystem-dev libboost-chrono-dev libboost-test-dev libboost-thread-dev
apt-get install git
apt-get install sudo
