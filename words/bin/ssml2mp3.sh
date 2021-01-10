#!/bin/bash

if [ -z $GOOGLE_APPLICATION_CREDENTIALS ]; then
	echo "Missing GOOGLE_APPLICATION_CREDENTIALS"
	exit 1
fi

mp3name="$(basename $1 .ssml).mp3"
ssml=$(jq -Rs . < $1)
read -r -d '' body << EOM
{
  "input":{
    "ssml":${ssml}
  },
  "voice":{
    "languageCode":"en-us",
    "name":"en-US-Wavenet-F",
    "ssmlGender":"FEMALE"
  },
  "audioConfig":{
    "audioEncoding":"MP3"
  }
}
EOM

curl -X POST \
-H "Authorization: Bearer "$(gcloud auth application-default print-access-token) \
-H "Content-Type: application/json; charset=utf-8" \
-d "${body}" \
https://texttospeech.googleapis.com/v1/text:synthesize | jq -r .audioContent | base64 -d > $mp3name


