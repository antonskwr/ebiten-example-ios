ebitenmobile bind -target ios -o ./xcode_project/EbitenTest/Mobile.framework ./mobile

exit_status=$?
if [ $exit_status -ne 0 ]; then
  echo "Error"
else
  echo "Compiled Successfully"
fi
exit $exit_status