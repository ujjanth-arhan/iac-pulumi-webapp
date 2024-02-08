#!/bin/bash
{
    echo "first_name,last_name,email,password"
    echo "john,doe,john.doe@example.com,abc123"
    echo "jane,doe,jane.doe@example.com,xyz456"
    echo "sample,sampel,ujjanthus@gmail.com,xyz456"
    echo "sample2,sampel2,arhan.u@northeastern.edu,xyz456"
    echo "Rishab,Agarwal,agarwal.risha@northeastern.edu,xyz456"
    echo "Souvik,Dinda,dinda.s@northeastern.edu,xyz456"
    echo "Sai,Dogiparthi,dogiparthi.sai@northeastern.edu,xyz456"
    echo "Raushan,Kumar,jha.rau@northeastern.edu,xyz456"
    echo "Aditya,Sawant,sawant.adit@northeastern.edu,xyz456"
    echo "Nagendra,babu,shakamuri.n@northeastern.edu,xyz456"
    echo "Shika,Singh,singh.shikha@northeastern.edu,xyz456"
} > /opt/users.csv

sudo cp /opt/users.csv ./builds/