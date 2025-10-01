# Sept 2025 Playing
A little toying before I try to make a go command line too for randomly generating 5 second clips from 

A tool that can 
- take a bunch of videos 
- cut them into a bunch of x second clips 
- then put each set of clips into a sub folder (2025/09/27 - 1)


## Usage

Run the program from the command line, passing the path to your input file as an argument. Wrap the path in quotes if it contains spaces:



## Usage

Run the program from the command line, passing the path to your input folder or file as an argument. Wrap the path in quotes if it contains spaces:

```bash
go run .\main.go "C:\\Users\\chris\\code projects\\input\\1\"
```

Accepts relative paths and exact file

```bash
go run .\main.go ".\\input\\1"
```

```bash
go run .\main.go ".\\input\\1\\video.m4v"
```