## Eyeball :mag:

Eyeball is a very basic cli tool (written in Go) that displays stats about your code. Stats include number of files, number of lines per file, file size etc. This is inspired by [tokei](https://github.com/XAMPPRocky/tokei) written in rust.

![image](https://www.nicepng.com/png/detail/87-872184_kommentit-eye-and-magnifying-glass.png)

#### Example :art:

#### Usage :clipboard:
`$ eyeball -r ./directory`

`$ eyeball -r ./directory -e .idea`

#### Excluding Folders :bomb:
By default, eyeball ignores certain folders [.git, .vscode]. Use the -e || --exclude flag to specify folders to ignore.

#### TODO :construction:
* sorting output
* grouping files with same extension
* more stats
