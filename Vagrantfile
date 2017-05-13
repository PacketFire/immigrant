Vagrant.configure(2) do |config|
  config.vm.define "mysql" do |mysql|
    mysql.vm.box = "debian/jessie64"
    mysql.vm.network "private_network", ip: "10.0.11.10"
    mysql.vm.hostname = "mysql"
    mysql.vm.provision "shell", path: "vagrant/provision/mysql/mysql.sh"
  end
end
