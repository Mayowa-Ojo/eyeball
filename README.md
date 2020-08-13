## Eyeball :mag:

Eyeball is a very basic cli tool (written in Go) that displays stats about your code. Stats include number of files, number of lines per file, file size etc. Inspired by [tokei](https://github.com/XAMPPRocky/tokei) written in rust.

![eyeball magnifying glass illustration](https://www.nicepng.com/png/detail/87-872184_kommentit-eye-and-magnifying-glass.png)

#### Example :art:
![example image](https://res.cloudinary.com/devmayor/image/upload/v1597167075/Screenshot_from_2020-08-11_11-29-21.png)

#### Installation
> Go version: >= 1.13

Clone this repository and build the binary 
```shell
$ git clone https://github.com/Mayowa-Ojo/eyeball
$ cd eyeball
$ go build
```
> add eyeball to your $PATH

or run 
```shell
go run main.go <flags>
```


#### Usage :clipboard:
`-r || --root` specifies the root directory

`-e || --exclude` specifies folders to ignore

`$ eyeball [-r ./directory`]

`$ eyeball [-r ./directory] [-e .idea]`

#### Excluding Folders :bomb:
By default, eyeball ignores certain folders [.git, .vscode]. Use the -e || --exclude flag to specify folders to ignore. Separate multiple folders with a comma ','

`$ eyeball -r ./ -e folder1,folder2,folder3`

#### TODO :construction:
[] Sorting output
[] Grouping files by extension
[] More stats...
