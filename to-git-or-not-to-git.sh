#! /bin/bash

curl https://adam-jerusalem.nd.edu/assets/superhero/all.json  | jq ' .[] | select( .id == 170) | .name '
curl https://adam-jerusalem.nd.edu/assets/superhero/all.json  | jq ' .[] | select( .id == 170) | .powerstats.power '
curl https://adam-jerusalem.nd.edu/assets/superhero/all.json  | jq ' .[] | select( .id == 170) | .appearance.gender '
