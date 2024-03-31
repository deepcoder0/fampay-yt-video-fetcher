# Backend Assignment (Intern) | FamPay

## Project Goal

To make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

### Basic Requirements:

- ☑️ Server should call the YouTube API continuously in background (async) with some interval (say 10 seconds) for fetching the latest videos for a predefined search query and should store the data of videos (specifically these fields - Video title, description, publishing datetime, thumbnails URLs and any other fields you require) in a database with proper indexes.
- ☑️ A GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
- ☑️ A basic search API to search the stored videos using their title and description.
- ☑️ Dockerize the project.
- ☑️ It should be scalable and optimised.

### Bonus Points:

- ☑️ Add support for supplying multiple API keys so that if quota is exhausted on one, it automatically uses the next available key.
- Make a dashboard to view the stored videos with filters and sorting options (optional)
- ☑️ Optimise search api, so that it's able to search videos containing partial match for the search query in either video title or description.
  - Ex 1: A video with title _`How to make tea?`_ should match for the search query `tea how`

## Tech Stack Used

- Programming Language: Go (Golang)
- Web Framework: Echo
- Database: MongoDB
- Docker - Containerization for deployment
- YouTube Data API v3 - For fetching video data from YouTube-

## Instructions

### Running Instructions

- Clone the github repository :

```bash
  git clone https://github.com/deepcoder0/fampay-yt-video-fetcher.git
  cd fampay-yt-video-fetcher
```

- Create .env file and Add fields as per sampleENV file.
  The Ports I used:

  - SERVER_ADDRESS = 0.0.0.0:8080
  - MONGODB_URI = mongodb://mongodb:27017

- To build and run docker container :

```bash
  docker compose up --build -d
```

### API Instructions

- Access the API endpoints:
  - The API endpoints will be available at `http://localhost:8080`.

## API Endpoints

1. GET `/fetch/:topic` : A cron job has been set up to fetch YouTube videos without relying on an external API trigger. This job runs in every 10 seconds, retrieving video details and storing them in the database.

   ```
   curl --location 'http://localhost:8080/fetch/cricket'
   ```

2. GET `/search`: Retrieve stored video data in a paginated response sorted by publishing datetime in descending order.
   ```
   curl --location 'http://localhost:8080/search?topic=cricket?page=1?pageSize=5'
   ```
   - By default, topic is null, so if not provided, it will fetch all videos
   - By default, page = 1 and pageSize = 10 here.
3. GET `/search/:query`: Search stored video by search query, which tries to match query with video's Title or Description.

   ```
   curl --location 'http://localhost:8080/search/tea how'
   ```

   - By default, page = 1 and pageSize = 5 here.

4. GET `/healthcheck` : To check if the server is up!

   ```
   curl --location 'http://localhost:8080/healthcheck'
   ```

5. GET `/dbReadiness` : To check if the DB server is up!
   ```
   curl --location 'http://localhost:8080/dbReadiness'
   ```

#### To know more, I have attacked screenshots for all the cases in /screenshots. Please have a look.
