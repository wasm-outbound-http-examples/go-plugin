GOPLUGIN_VERSION='0.8.0'
apt update
apt install -y protobuf-compiler
wget -O /tmp/go-plugin.deb https://github.com/knqyf263/go-plugin/releases/download/v${GOPLUGIN_VERSION}/go-plugin_${GOPLUGIN_VERSION}_linux-amd64.deb
dpkg -i /tmp/go-plugin.deb
rm /tmp/go-plugin.deb
