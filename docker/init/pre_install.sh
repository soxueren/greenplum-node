#!/bin/bash
set -e

# cat <<EOF > /etc/sysctl.conf
# kernel.shmmax = 500000000
# kernel.shmmni = 4096
# kernel.shmall = 4000000000
# kernel.sem = 2000 4096000 2000 2048
# kernel.sysrq = 1
# kernel.core_uses_pid = 1
# kernel.msgmnb = 65536
# kernel.msgmax = 65536
# kernel.msgmni = 2048
# net.ipv4.tcp_syncookies = 1
# net.ipv4.ip_forward = 0
# net.ipv4.conf.default.accept_source_route = 0
# net.ipv4.tcp_tw_recycle = 1
# net.ipv4.tcp_max_syn_backlog = 4096
# net.ipv4.conf.all.arp_filter = 1
# net.ipv4.ip_local_port_range = 1025 65535
# net.core.netdev_max_backlog = 10000
# net.core.rmem_max = 2097152
# net.core.wmem_max = 2097152
# vm.overcommit_memory = 2
# EOF

#sysctl -p

cat <<EOF >  /etc/security/limits.conf
* soft nofile 65536
* hard nofile 65536
* soft nproc 131072
* hard nproc 131072
EOF

#/usr/sbin/blockdev --setra 16384 /dev/sda

#echo deadline > /sys/block/sda/queue/scheduler
