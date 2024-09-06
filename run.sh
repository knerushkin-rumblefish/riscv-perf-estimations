#!/bin/bash

echo "Simulation Started"

declare -a arr=(5000 10000 20000 30000 40000 50000 60000 70000 80000 90000 100000 150000 200000 500000)

for i in "${arr[@]}"
do
    for k in {1..5}
    do
        echo "$i $k"
        echo "$i $k" >> $2
        { /usr/bin/time --format="%U" ./$1 "$i" ; } 2>> $2
    done
done

echo "Simulation Ended"
