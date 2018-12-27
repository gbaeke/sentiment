# Command line tool to perform sentiment analysis with Cognitive Services container
Usage: ta --language en|nl --text "text to analyse" --url "http://localhost:5000"

- language: defaults to nl
- url: defaults to http://localhost:5000

This tool requires the Cognitive Services container from mcr.microsoft.com/azure-cognitive-services/sentiment

Use **docker pull mcr.microsoft.com/azure-cognitive-services/sentiment:latest** and run the container with -p 5000:5000 to map container port 5000 to host port 5000
