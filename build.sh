#!/bin/bash

filename="./array_encode_decode.go"
echo "//" > $filename
echo "// @project GeniusRabbit" >> $filename
echo "// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2017" >> $filename
echo "//" >> $filename

cat ./array_encode_decode.go.tmp | genny gen \
  "GenType=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64" \
  >> $filename
