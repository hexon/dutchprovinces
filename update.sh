#!/bin/sh

wget -O - 'https://service.pdok.nl/kadaster/bestuurlijkegebieden/wfs/v1_0?request=GetFeature&service=WFS&version=1.1.0&outputFormat=application%2Fjson&typeName=bestuurlijkegebieden:Provinciegebied' | jq -c . > rijksdriehoeken.json
