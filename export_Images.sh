#!/bin/bash

EXPORT_DIR="/mnt/d"
IMAGES_DIR="./images"

for filename in $(bash $EXPORT_DIR/get_fileName.sh)
do
	
	if [ -x "$IMAGES_DIR/$filename" ]
	then
		echo $filename
		mv $IMAGES_DIR/$filename $EXPORT_DIR/images
	else
		echo "Not Found"
	fi
done

