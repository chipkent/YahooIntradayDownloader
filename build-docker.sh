#/usr/bin/env bash

docker build -t yahoointradaydownloader .

docker save -o yahoointradaydownloader.tar yahoointradaydownloader


