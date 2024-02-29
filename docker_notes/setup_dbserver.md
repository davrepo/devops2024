# set up docker and get dbserver up
# Attn: Some commands might require '-y' flag if run as a script

# install http and curl packages
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common

# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

# Install docker
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Install psql
sudo apt update
sudo apt install postgresql postgresql-contrib

# Start postgres service
sudo systemctl start postgresql.service

# Build postgres docker container
docker build -t minitwit-postgres /vagrant/database

# Run docker container
docker run -d --name minitwit-postgres-instance -p 8080:5432 minitwit-postgres

