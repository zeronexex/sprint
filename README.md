# Why
Say you have a file with 50k CNAMES and for some reason you want to split those CNAMES in smaller chunks. Sprint takes the count of the chunk, it creates files and writes those CNAMES chunks in files accordingly.

# Installation
You have to have Go installed in your machine

`go get github.com/zeronexex/sprint`

# Update

`go get -u github.com/zeronexex/sprint`

# Usage 

`splityet -h`

Crafted with 💜 by Zeron

`⚠ All flags must be set. Be sure to set them correctly`
   
    -chunk int
        
        Count of cnames each file will contain (except last one)
  
    -cpath string
        
        Absolute path of the file containing cnames
  
    -filename string
        
        This name is gonna be used for created files with numbers added on right incrementally
  
    -w string
        
        The absolute path of where split files will be created
