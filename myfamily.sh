curl -s https://platform.zone01.gr/assets/superhero/all.json | \
jq -r '.[] | select(.id == '$HERO_ID') | .connections.relatives' | \
sed 's/"//g' |  \
if (($HERO_ID == 70))
then 
    tr '\n' '\\n' | sed 's/y\\/y/g' |sed 's/)\\/)\\n/g'
else
    sed 's/\\//g'
fi