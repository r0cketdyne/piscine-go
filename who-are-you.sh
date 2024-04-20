json=$(curl -s https://platform.zone01.gr/assets/superhero/all.json)
id170=$(jq -r '.[] | select(.id == 70)' <<< "$json")
name=$(jq -r '.name' <<< "$id170")
printf "\"$name\"\n"