echo "Enter prefix line: "
read PREFIX

today_date=`date +%Y-%m-%d`
dir=data/$PREFIX/scan_$today_date
mkdir -p $dir
echo $dir

while [ 1 ]
do
  time=`date +%H:%M:%S`
  echo "Scan at $time"
  /System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport -s > $dir/scan_$time
  sleep 30
done
