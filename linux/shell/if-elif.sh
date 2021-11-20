# !/bin/bash

age=10
if (($age > 18)); then
    echo "you are a grown up now"
elif (($age > 12)); then
    echo "you are so youthful"
else
    echo "go home! your under age"

fi

echo "
rdr pass on lo0 inet proto tcp from any to self port 4662 -> 127.0.0.1 port 4662
rdr pass on en0 inet proto tcp from any to any port 4662 -> 127.0.0.1 port 4662
rdr pass on en1 inet proto tcp from any to any port 4662 -> 127.0.0.1 port 4662
rdr pass on lo0 inet proto udp from any to self port 4672 -> 127.0.0.1 port 4672
rdr pass on en0 inet proto udp from any to any port 4672 -> 127.0.0.1 port 4672
rdr pass on en1 inet proto udp from any to any port 4672 -> 127.0.0.1 port 4672
rdr pass on lo0 inet proto udp from any to self port 4665 -> 127.0.0.1 port 4665
rdr pass on en0 inet proto udp from any to any port 4665 -> 127.0.0.1 port 4665
rdr pass on en1 inet proto udp from any to any port 4665 -> 127.0.0.1 port 4665
" | sudo pfctl -ef -
