#!/bin/bash

# Convert JavaScript code in web/ to Go web handlers

# Define the web directory path
WEB_DIR="web/"

# Loop through each JavaScript file in the web directory
for FILE in $(find $WEB_DIR -type f -name "*.js"); do
    # Extract the file name without extension
    FILENAME=$(basename $FILE .js)
    CONTENT=$(cat $FILE | sed 's/`/\\`/g')
    
    # Define the output file path
    OUTPUT_FILE="src/web_handler_$FILENAME.go"

    # Create the output file
    echo "// THIS IS AN AUTOGENERATED FILE... See gen_web_handler.sh for details." > $OUTPUT_FILE
    echo "package qmq" >> $OUTPUT_FILE
    echo "" >> $OUTPUT_FILE
    echo "import (" >> $OUTPUT_FILE
    echo "    \"net/http\"" >> $OUTPUT_FILE
    echo "    \"fmt\"" >> $OUTPUT_FILE
    echo ")" >> $OUTPUT_FILE
    echo "" >> $OUTPUT_FILE
    echo "func Register_web_handler_$FILENAME() {" >> $OUTPUT_FILE

    # Convert the JavaScript code to a Go web handler function
    echo "" >> $OUTPUT_FILE
    echo "    http.HandleFunc(\"js/qmq/$FILENAME.js\", func(w http.ResponseWriter, r *http.Request) {" >> $OUTPUT_FILE
    echo "        fmt.Fprintf(w, \`$CONTENT\`)" >> $OUTPUT_FILE
    echo "    })" >> $OUTPUT_FILE
    echo "}" >> $OUTPUT_FILE
done

echo "Go web handlers generated successfully!"
