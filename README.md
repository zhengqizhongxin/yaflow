# yaflow
yet another go-dpi example

# Add repo
```
sudo apt-get install apt-transport-https
echo "deb https://packages.wand.net.nz $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/wand.list
sudo curl https://packages.wand.net.nz/keyring.gpg -o /etc/apt/trusted.gpg.d/wand.gpg
sudo apt-get update
```
# Dependencies 
```
sudo apt-get install liblinear4 liblinear-dev
sudo apt-get -y --force-yes install git gcc autoconf automake libtool libpcap-dev libtrace4 libtrace4-dev libprotoident libprotoident-dev

git clone --branch 3.2-stable https://github.com/ntop/nDPI/ /tmp/nDPI
cd /tmp/nDPI && ./autogen.sh && ./configure && make && sudo make install && cd -
```