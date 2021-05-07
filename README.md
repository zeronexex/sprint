# Why
Say you have a file with 50k CNAMES and for some reason you want to split those CNAMES in smaller chunks. Sprint takes the count of the chunk, it creates files and writes those CNAMES chunks in files accordingly.

# Installation
You have to have Go installed in your machine

`go get github.com/rew1nter/sprint`

# Update

`go get -u github.com/rew1nter/sprint`

# Usage 

```
sprint -h

sprint <flags> <args>
	flags : 
		-c int
			The size of chunks.
		-fn string
			This name is will be used for created files with numbers added on right incrementally.
	args :
		First argument: is the path where the file you want to split exists
		Second argument: is the path where new chunked files will be created
	ex :
		sprint -c 2 -f raft_small /path/to/file/for/splitting /path/where/new/files/will/be/written
		sprint -c 10 -f "raft small.txt" /path/to/file/for/splitting /path/where/new/files/will/be/written
      
      Crafted with 🤍 by Rewinter
```
