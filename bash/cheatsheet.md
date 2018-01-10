### bash script programming 

```
echo $#  # arguments number
echo $0  # exec filename
echo $1  # first argument
echo $2  # second argument 
echo "$*"  # all arguments as a whole string   
echo "$@"  # all arguments as an array
echo $!  # last executed cmd pid
echo $?  # exec status 
echo $$  # current bash script process pid 

# some configurations for output color
RED_COLOR='\E[1;31m'
GREEN_COLOR='\E[1;32m'
YELLOW_COLOR='\E[1;33m'
BLUE_COLOR='\E[1;34m'
END_COLOR='\E[0m'
echo -e "${RED_COLOR}installing basic utilities...${END_COLOR}"


if [ -e "/bin/bash" ]; then
  echo "yes"
else
  echo "no"
fi


read temp 
case $temp in
  1)
    echo "1"
    ;;
  2)
    echo "2"
    ;;
  [3-9])
    echo "3-9"
esac


for i in "$*"; do
  echo $i
done

for i in `ls`; do 
  echo $i
done

for (( i = 0; i < $#; i++ )); do
  echo $i
done 

x=1
while [[ $x -le $# ]]; do
  echo $x
  let x=x+1
done


function testFunc1 {
  echo "in testFunc1"
}
testFunc2() {
  echo "in testFunc2", $*
}
testFunc1
testFunc2 123 456


-e filename   =>  file exists 
-d dirname    =>  dir exists 
-f filename   =>  regular file
-L filename   =>  symbol link 
-r filename   =>  can be read 
-w filename   =>  can be written
-x filename   =>  can be executed 
filename1 -nt filename2   =>  newer than 
filename1 -ot filename2   =>  older than 

-z string        =>  null string 
-n string        =>  not null string 
string1=string2  =>  equal
string1!=string2 =>  not equal 

num1 -eq num2    =>  = 
num1 -ne num2    =>  != 
num1 -lt num2    =>  < 
num1 -le num2    =>  <=
num1 -gt num2    =>  >
num1 -ge num2    =>  >=

```
