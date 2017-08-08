# Install a yahoointradaydownloader. 

# 1) Create the image:
# docker build -t yahoointradaydownloader .
#
# 2) Put your ticker file in a local directory as tickers.txt
#
# 3) Run the new container to produce output in /local/path/prices.csv
# docker run --rm -v /local/path:/tmp yahoointradaydownloader
#
# If you want to get on the image to poke around:
# docker run --rm -ti -v /local/path:/tmp yahoointradaydownloader sh
#
#    Just keep in mind that's giving you a shell in a new instance of the image not 
#    connecting you to an already running container.

FROM alpine
MAINTAINER Chip Kent

ADD YahooIntradayDownloader_linux /usr/bin/YahooIntradayDownloader

CMD ["YahooIntradayDownloader", "/tmp/tickers.txt", "/tmp/prices.csv", "50"]

