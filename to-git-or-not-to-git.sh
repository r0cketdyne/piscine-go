json=$(curl -s https://platform.zone01.gr/assets/superhero/all.json)
id170=$(jq -r '.[] | select(.id == 170)' <<< "$json")
appearance=$(jq -r '.appearance' <<< "$id170")
powerstats=$(jq -r '.powerstats' <<< "$id170")
name=$(jq -r '.name' <<< "$id170")
power=$(jq -r '.power' <<< "$powerstats")
gender=$(jq -r '.gender' <<< "$appearance")
printf "$name\n$power\n$gender\n"