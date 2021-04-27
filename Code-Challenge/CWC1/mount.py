
import json
import subprocess

cmd = "df -a"
output = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE)
#data = output.stdout.read().strip().split("\n")
output = output.communicate()[0].strip().split("\n")

file = {}
for i in range(1,len(output)):
    l=output[i].strip().split(" ")
    l1=[]
    for i in l:
        if i == '':
            continue
        else:
            l1.append(i);
    file[l1[0]]=int(l1[2])/1024

print(file)

import React from 'react'

export default function AppFooter() {
    return (

        <div style={{height:"25px",backgroundColor:"#cc0000",textAlign:"center"}}>
            <footer>
                <p style={{color:"#ffffff",paddingTop:"2px",fontFamily:"HelveticaforTargetLight"}}>DojoTII@2020</p>

            </footer>
        </div>
    )
}


                  
