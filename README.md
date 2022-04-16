![](overview.png)

# Run with Docker Compose

```bash
# start + create
docker-compose up -d

# start
docker-compose start

# stop
docker-compose stop

# remove
docker-compose down

# build images
docker-compose build

# push images
docker-compose push
```

# Install docker on ec2

[source](https://thomasventurini.com/articles/install-docker-and-compose-in-ubuntu/)

```bash
sudo apt-get update \
&& sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common \
&& curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - \
&& sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable" \
&& sudo apt-get update \
&& sudo apt-get install docker-ce docker-ce-cli containerd.io

sudo groupadd docker \
&& sudo usermod -aG docker $USER \
&& newgrp docker # this enables the group without the need of a restart

sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
&& sudo chmod +x /usr/local/bin/docker-compose
```

Quick hack:

```
sudo chmod 777 /var/run/docker.sock
```

# Copy to ec2

```bash
cd ..
# sesuaikan keypair nya (i.e: laptop.pem)
rsync -r -e "ssh -i ~/laptop.pem" merdekaMock2 ubuntu@54.179.163.193:/home/ubuntu/app
```

# Run docker compose on server

```bash
docker-compose up -d
```

