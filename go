#!/bin/bash

# URL of an image of the map of India with state boundaries
IMAGE_URL="https://upload.wikimedia.org/wikipedia/commons/thumb/7/7b/India_map_blank.svg/1024px-India_map_blank.svg.png"
IMAGE_FILE="india_map.png"

# Download the map image if it doesn't already exist
if [[ ! -f $IMAGE_FILE ]]; then
    echo "Downloading the map of India..."
    wget -O "$IMAGE_FILE" "$IMAGE_URL"
fi

# Check if feh or imgcat is available for displaying the image
if command -v feh >/dev/null 2>&1; then
    echo "Displaying the map using feh..."
    feh "$IMAGE_FILE"
elif command -v imgcat >/dev/null 2>&1; then
    echo "Displaying the map using imgcat..."
    imgcat "$IMAGE_FILE"
else
    echo "Neither 'feh' nor 'imgcat' is installed."
    echo "Please install 'feh' (Linux) or 'imgcat' (macOS/iTerm2) to display the image."
fi
