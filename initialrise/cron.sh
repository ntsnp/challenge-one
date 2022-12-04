#$CWD = `pwd`
#$CRONJOB = "0 12 * * * /usr/bin/python3 `pwd`/main.py"
(crontab -l ; echo "0 12 * * * /usr/bin/python3 `pwd`/main.py -d `date +20%y%m%d-%H%M%S`.db") | crontab -


