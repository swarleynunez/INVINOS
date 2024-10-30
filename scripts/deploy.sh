nohup ganache-cli -d --db ganache-cli > log/ganache-cli.out 2>&1 &
nohup ipfs daemon > log/ipfs.out 2>&1 &
nohup ./invinos > log/invinos.out 2>&1 &
