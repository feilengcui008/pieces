#!/bin/bash 

if [ "$#" -lt 3 ]; then
    echo "usage: ./mysql_create_database.sh dbname username password"
    exit 0
fi 
ROOTUSER="root" 
ROOTPASSWORD="tan211211"
DBNAME=$1
USERNAME=$2
PASSWORD=$3
#show grants for 'name'@'localhost'
#grant select,insert,update,delete on *.* to 'name'@'localhost'
#revoke select on db.table from 'name'@'localhost'
#set password for 'name'@'%' = PASSWORD("newpass")
#drop user 'name'@'localhost'
#create table if not exists tablename (id int(10) unsigned auto_increment primary key, name varchar(200) not null ) engine=innodb default charset=utf8

sql_str="\"
    create database if not exists $DBNAME default charset utf8 collate utf8_general_ci;
    create user '$USERNAME'@'localhost' identified by '$PASSWORD';
    grant all on $DBNAME.* to '$USERNAME'@'localhost';
    flush privileges;
    \""

MYSQL_EXEC="/usr/bin/mysql -e $sql_str -u$ROOTUSER" 
if [ "$ROOTPASSWORD" != "" ]; then
    MYSQL_EXEC="$MYSQL_EXEC -p$ROOTPASSWORD" 
fi 
#echo $MYSQL_EXEC
eval $MYSQL_EXEC
