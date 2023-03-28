#!/bin/bash

# Install needed packages. Please add to this list if you discover additional prerequisites
sudo yum group install -y "Development Tools"
sudo yum install -y \
    apr-devel \
    bison \
    bzip2-devel \
    cmake3 \
    flex \
    gcc \
    gcc-c++ \
    krb5-devel \
    libcurl-devel \
    libevent-devel \
    libkadm5 \
    libxml2-devel \
    libzstd-devel \
    openssl-devel \
    python39 \
    python39-devel \
    python39-psutil \
    python3-pip \
    perl-ExtUtils-MakeMaker.noarch \
    perl-ExtUtils-Embed.noarch \
    readline-devel \
    rsync \
    xerces-c-devel \
    zlib-devel 

sudo pip3.9 install pygresql
sudo pip3 install -r python-dependencies.txt

sudo tee -a /etc/sysctl.conf << EOF
kernel.shmmax = 5000000000000
kernel.shmmni = 32768
kernel.shmall = 40000000000
kernel.sem = 1000 32768000 1000 32768
kernel.msgmnb = 1048576
kernel.msgmax = 1048576
kernel.msgmni = 32768

net.core.netdev_max_backlog = 80000
net.core.rmem_default = 2097152
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216

vm.overcommit_memory = 2
vm.overcommit_ratio = 95
EOF

sudo sysctl -p

sudo mkdir -p /etc/security/limits.d
sudo tee -a /etc/security/limits.d/90-greenplum.conf << EOF
* soft nofile 1048576
* hard nofile 1048576
* soft nproc 1048576
* hard nproc 1048576
EOF

ulimit -n 65536 65536

