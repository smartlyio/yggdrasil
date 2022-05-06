# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|

    config.vm.box = "ubuntu/focal64"
    config.vm.network "forwarded_port", guest: 9901, host: 9901 #this is envoy managment port
    config.vm.provider "virtualbox" do |v|
    config.vm.synced_folder "../../../", "/yggdrasil"
    v.memory = 10240
    v.cpus = 4
    end

    config.vm.provision "shell", inline: <<-SHELL
       apt-get update
       apt-get dist-upgrade -y
       apt-get install -y htop ca-certificates curl gnupg lsb-release jq
       curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
       echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
       apt-get update && apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose
       curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
       echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
       apt-get update && apt-get install -y kubectl
       curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
       curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
     SHELL
  end
