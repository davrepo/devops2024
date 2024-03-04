#!/bin/bash

# The URL of your homepage
url="http://104.248.43.157:8080/"

# Number of requests to make
requests=10

# Variable to accumulate total time
totalTime=0

echo "Loading $url $requests times..."

for ((i = 1; i <= requests; i++)); do
	# Make a request and extract the time it took
	time=$(curl -o /dev/null -s -w '%{time_total}\n' $url)

	# Add the time to totalTime
	totalTime=$(echo $totalTime + $time | bc)

	echo "Request $i: $time s"
done

# Calculate the average time
averageTime=$(echo "scale=3; $totalTime / $requests" | bc)

echo "Average load time: $averageTime s"
