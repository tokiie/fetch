# Fetch HTML Tool

This tool fetches a webpage, downloads its html

## Features

- **Fetch Webpages:** Downloads HTML for a specified URL.
- **Metadata Printing:** Prints metadata for downloaded webpages, including the number of links, images, and the last fetch time.

## Prerequisites

- Docker
- Go (if you want to run the project outside of Docker)

### Build the Docker Image

To build the Docker image, run the following command in the directory where your `Dockerfile` and Go source files are located:

```bash
docker build -t fetch .
```

### Running the Docker Container

Once the image is built, you can run the container to fetch a webpage or print metadata.

```bash
docker run --rm -it fetch
```

## Running the Project Without Docker

If you have Go installed on your system and prefer to run the project directly

### Build the Go Application

```bash
go build -o fetch
```

### Fetch html for one or multiple URLs

```bash
./fetch https://www.example.com https://www.google.com
```

### Print Metadata for one or multiple URLs

```bash
./fetch -metadata https://www.example.com https://www.google.com
```

## How It Works

### Fetching a Webpage

The application:

1. Processes the provided URL to ensure it includes `https://` if no scheme is specified.
2. Downloads the HTML content of the webpage.

### Printing Metadata

The application:

1. Checks if the local directory exists. If not, it downloads the webpage.
2. Parses the saved HTML file.
3. Counts the number of `<a>` (link) and `<img>` (image) tags.
4. Prints the metadata, including the number of links, images, and the last fetch time.
