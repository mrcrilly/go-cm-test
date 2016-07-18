Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"
  config.vm.define "go-cm-test"

  config.vm.provider :virtualbox do |v|
    v.memory = 1024
    v.cpus = 4
  end

  config.vm.network :public_network
  
  config.vm.synced_folder ".", "/home/vagrant/sync", disabled: true
  config.vm.synced_folder "C:/cygwin64-mcrilly/home/michael.crilly/code/go", "/vagrant", type: "virtualbox"
end
