# YahooIntradayDownloader
Download Yahoo Prices Intraday

# Compile 
The repository contains statically linked executables for Linux, Mac, and Windows.  These executables should run without problem.

To build the binaries, run:
```
./build.sh
```

To build a Docker container, run:
```
./build-docker.sh
```

# Run
To run the binaries:
```
./YahooIntradayDownloader_linux <tickerfile> <outputfile> <chunk_size>
```
where 
* `<tickerfile>` is the path to a file where each line contains a ticker to download data for
* `<outputfile>` is the path where a CSV file will be written with the prices
* `<chunk_size>` is an integer indicating how many tickers will be downloaded at once (e.g. 50)

To run the Docker container, create a directory that contains the ticker file.  The ticker file must be named `tickers.txt`.
```
mkdir downloads
cp /path/to/tickers.txt downloads
```
Then run the Docker container:
```
docker run --rm -v /path/to/downloads:/tmp yahoointradaydownloader
```
This command mounts your local download directory, `/path/to/downloads` in this example, to `/tmp` in the container.
The downloaded prices will appear in `prices.csv` in the download directory.  `/path/to/downloads/prices.csv` in this example.



