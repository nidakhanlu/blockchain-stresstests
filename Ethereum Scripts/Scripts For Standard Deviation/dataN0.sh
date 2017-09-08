#!/bin/bash


N=1000
n=0
#m=0
c=3010000
#ifconfig enp0s3 down
now=$(date) 
echo $now >>time1000a.txt

        for((l=1;l<=$N;l++));
        do
            
               ((c+=1))
               export c
               j=$(source test5000)
               echo $j >>hash10000a.txt
               #if ! (($l % 10)); then
               #sleep 80
               #fi
              
       done


