# -*- mode: ruby -*-
# vi: set ft=ruby :

$ip_file = "db_ip.txt"

Vagrant.configure("2") do |config|
    config.vm.box = 'digital_ocean'
    config.vm.box_url = "https://github.com/devopsgroup-io/vagrant-digitalocean/raw/master/box/digital_ocean.box"
    config.ssh.private_key_path = 'C:\Users\abc\.ssh\id_rsa'
    config.vm.synced_folder ".", "/vagrant", type: "rsync"

    # Web server VM definition
    config.vm.define "dbserver", primary: true do |server|
      server.vm.provider :digital_ocean do |provider|
        provider.ssh_key_name = ENV["SSH_KEY_NAME"]
        provider.token = ENV["DIGITAL_OCEAN_TOKEN"]
        provider.image = 'ubuntu-22-04-x64'
        provider.region = 'fra1'
        provider.size = 's-1vcpu-1gb'
        provider.privatenetworking = true
      end
      
      server.vm.hostname = "dbserver"

      server.vm.provision "shell", inline: <<-SHELL
        echo "Installing PostgreSQL"
        sudo apt-get update
        sudo apt-get install -y postgresql postgresql-contrib

        # Configure PostgreSQL: Set up user and permissions
        sudo -u postgres psql -c "CREATE USER root WITH SUPERUSER PASSWORD 'password';"

        # Create the 'minitwit' database owned by 'root' user
        sudo -u postgres createdb -O root minitwit

        # Restart PostgreSQL to apply changes
        sudo systemctl restart postgresql

        echo "PostgreSQL installation and setup complete."
      SHELL
    end

    config.vm.define "webserver", primary: false do |server|

      server.vm.provider :digital_ocean do |provider|
        provider.ssh_key_name = ENV["SSH_KEY_NAME"]
        provider.token = ENV["DIGITAL_OCEAN_TOKEN"]
        provider.image = 'ubuntu-22-04-x64'
        provider.region = 'fra1'
        provider.size = 's-1vcpu-1gb'
        provider.privatenetworking = true
      end
  
      server.vm.hostname = "webserver"
  
      server.vm.provision "shell", inline: <<-SHELL
        echo "Setting up the Go environment..."
        wget https://dl.google.com/go/go1.18.linux-amd64.tar.gz
        sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
        echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile
        source $HOME/.profile
        mkdir -p "$HOME/go/src" "$HOME/go/bin"
        echo "export GOPATH=$HOME/go" >> $HOME/.profile
        echo "export PATH=$PATH:$GOPATH/bin" >> $HOME/.profile
        source $HOME/.profile
        cp -r /vagrant/* $HOME/go/src/minitwit
        cd $HOME/go/src/minitwit
        /usr/local/go/bin/go mod download
        /usr/local/go/bin/go build -o minitwit ./src/main.go
        nohup ./minitwit > out.log 2>&1 &
        echo "Go application is running..."
        echo "Navigate in your browser to the server's IP"
        SHELL
    end
end