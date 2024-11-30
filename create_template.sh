#!/bin/bash

# Function to display usage
usage() {
    echo "Usage: $0 -l <language>"
    echo "Options for <language>: go, typescript, rust, julia"
    exit 1
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        -l|--language)
            LANGUAGE=$2
            shift 2
            ;;
        *)
            usage
            ;;
    esac
done

#################
# SESSION AUTH
#################
# Load environment variables from the .env file
if [[ -f .env ]]; then
    source .env
    echo "Loaded environment variables from .env"
else
    echo "Warning: .env file not found. Please create it with your AOC_SESSION."
    exit 1
fi

#################
# DATE EXTRACTION
#################

# Validate the language parameter
if [[ -z "$LANGUAGE" ]]; then
    echo "Error: Language not specified."
    usage
fi

if [[ "$LANGUAGE" != "go" && "$LANGUAGE" != "typescript" && "$LANGUAGE" != "rust" && "$LANGUAGE" != "julia" ]]; then
    echo "Error: Invalid language specified."
    usage
fi

# Get the current date
current_date=$(date +%Y-%m-%d)  # Format: YYYY-MM-DD

# Extract the day from the current date
day=$(date +%d)  # Format: DD
year=$(date +%Y)  # Format: YYYY

# Display the results
echo "Current date: $current_date"
echo "Day of the month: $day"
echo "Selected programming language: $LANGUAGE"

#################
# TEMPLATE COPY
#################

# Create a folder named after the day of the month
target_folder="./$day"
if [[ ! -d "$target_folder" ]]; then
    mkdir "$target_folder"
    echo "Created folder: $target_folder"
else
    echo "Folder already exists: $target_folder"
fi

# Define the source folder for the selected language
source_folder="./templates/$LANGUAGE"

# Check if the source folder exists
if [[ ! -d "$source_folder" ]]; then
    echo "Error: Source folder '$source_folder' does not exist."
    exit 1
fi

#################
# DATA FETCHING
#################

# Copy files from the source folder to the target folder
cp -r "$source_folder/"* "$target_folder/"
echo "Copied files from '$source_folder' to '$target_folder'"

# TODO delete
day=12
year=2023

# THE USER SPECIFIC INPUT LARGE

# Fetch the input text from the Advent of Code URL
input_url="https://adventofcode.com/$year/day/$day/input"
input_file="$target_folder/in"

# Make the request and save the response to the input file
echo "Fetching input from: $input_url"
curl -s --cookie "session=$AOC_SESSION" "$input_url" -o "$input_file"

# Check if the input was successfully fetched
if [[ $? -eq 0 && -s "$input_file" ]]; then
    echo "Input saved to: $input_file"
else
    echo "Failed to fetch input or input is empty. Check your session token or the URL."
    rm -f "$input_file"  # Remove empty file if created
fi

# THE SMALL EXAMPLE INPUT

# Now fetch the content of the first <pre><code> element
code_url="https://adventofcode.com/$year/day/$day"
sm_file="$target_folder/sm"

# Fetch the HTML content from the page
echo "Fetching HTML content from: $code_url"
html_content=$(curl -s "$code_url")

# Extract the content inside the first <pre><code> block
sm_content=$(echo "$html_content" | awk '
    BEGIN { found = 0 }
    /<pre><code>/ && !found { 
        found = 1; 
        next 
    }
    found && /<\/code><\/pre>/ { exit }
    found { print }
')


# Save the extracted content into the "sm" file
if [[ -n "$sm_content" ]]; then
    echo "$sm_content" > "$sm_file"
    echo "Extracted content saved to: $sm_file"
else
    echo "Failed to extract content or no content found."
fi
