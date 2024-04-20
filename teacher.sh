INTERVIEW=$(sed -n '179p' streets/Buckingham_Place| grep -oP '\d+')
echo "$INTERVIEW"
cat interviews/interview-$INTERVIEW
echo "$MAIN_SUSPECT"