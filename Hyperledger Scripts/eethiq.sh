 #!/bin/bash
N=10000
n=0
     for((i=1;i<=$N;i++));
     do
         $(peer chaincode invoke -l golang -n mycc -c '{"Args": ["register", "nida"]}')    
         Input="$(peer chaincode query -l golang -n mycc -c '{"Args": ["query", "id"]}')"      
         echo -e "\r\n" $Input >>output.txt
     done
#sleep  5

#perl -lane '$F[2]/= 100000000; $F[4] /= 1000000;$F[2]+=$F[4]; print "@F[2]"' < output.txt 
sed 's/[^0-9]/ /g' output.txt >>test.txt  
sed 's/^ */ /g' test.txt >>prefinal.txt
sed '/^\s*$/d' prefinal.txt >>final.txt
perl -lane '$F[0]*= 1000; $F[1] /= 1000000;$F[0]+=$F[1]; print "@F[0]"' < final.txt >> file2.txt
    for((i=1;i<=$N;i++));
    do
     echo -e $i >> file1.txt 
    done  
paste -d' ' file1.txt file2.txt >>plotfile.txt
#sed 's/[ \t]\?/ /g' < test.txt
#sed 's/[ \t]\?,[ \t]\?/,/g; s/^[ \t]\+//g; s/[ \t]\+$//g' test.txt
#cat test.txt | 's/^ *//' 
python plotgraph.py
