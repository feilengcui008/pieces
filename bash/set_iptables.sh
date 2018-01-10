#!/bin/bash 

IPTABLES=/sbin/iptables

# clear filter
$IPTABLES -F
# clear filter user-defined
$IPTABLES -X
# preset rules 
$IPTABLES -P INPUT DROP
#$IPTABLES -P INPUT ACCEPT
$IPTABLES -P OUTPUT ACCEPT 
$IPTABLES -P FORWARD DROP

# internal NIC and loop device 
$IPTABLES -A INPUT -i lo  -j ACCEPT
$IPTABLES -A OUTPUT -o lo -j ACCEPT

$IPTABLES -A INPUT -i eth0  -j ACCEPT
$IPTABLES -A OUTPUT -o eth0 -j ACCEPT

# ssh
$IPTABLES -A INPUT -p tcp  --dport 22  -j ACCEPT
$IPTABLES -A INPUT -p tcp  --sport 22  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp --sport 22  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp --dport 22  -j ACCEPT

# nginx 
$IPTABLES -A INPUT  -p tcp  --dport 80  -j ACCEPT
$IPTABLES -A INPUT  -p tcp  --sport 80  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp  --dport 80  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp  --sport 80  -j ACCEPT

$IPTABLES -A INPUT  -p tcp  --dport 443  -j ACCEPT
$IPTABLES -A INPUT  -p tcp  --sport 443  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp  --dport 443  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp  --sport 443  -j ACCEPT

# dns 
$IPTABLES -A INPUT  -p udp -m udp --sport 53 -j ACCEPT
$IPTABLES -A INPUT  -p udp -m udp --dport 53 -j ACCEPT
$IPTABLES -A OUTPUT -p udp -m udp --sport 53 -j ACCEPT
$IPTABLES -A OUTPUT -p udp -m udp --dport 53 -j ACCEPT


# 4000
$IPTABLES -A INPUT  -p tcp  --dport 4000  -j ACCEPT
$IPTABLES -A INPUT  -p tcp  --sport 4000  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp  --dport 4000  -j ACCEPT
$IPTABLES -A OUTPUT -p tcp  --sport 4000  -j ACCEPT

