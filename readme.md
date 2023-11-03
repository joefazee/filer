# Filer

Chunk large files & join them back in Golang

## How to run the program

1. Download and compile the program (**You will need to init go module in the folder**)
2. Run `./filer -action=split -file=the-file-to-split -size=int-size-of-your-choice`
3. To join a file `./filer -action=join -file=the-file-to-split`

Watch the YouTube video tutorial on how to create this program and many more.  https://youtu.be/urB2EsjFmbs

**WARNING: Use This Program With Caution**

This program has the capability to delete files from your system. Please ensure that you operate it within a secure and controlled environment, such as a designated safe folder or a sandboxed testing area. The author provides no warranty and assumes no liability for any consequences that arise from the use of this software. By running this program, you acknowledge and accept the risks involved.
