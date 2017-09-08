#!/bin/bash


N=1000
n=0
#m=0
c=3010000
 

        for((l=1;l<=$N;l++));
        do
            
               ((c+=1))
               export c
               j=$(source test50001)
               echo $j >>hash10000a.txt
               #if ! (($l % 10)); then
               #sleep 80
               #fi

       done


