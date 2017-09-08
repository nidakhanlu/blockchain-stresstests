#!/bin/bash


N=100


        for((l=1;l<=$N;l++));
        do
            
             
               #if ! (($l % 10)); then
               ifconfig enp0s3 down
               sleep 10
               ifconfig enp0s3 up
               sleep 80
               #fi

        done


