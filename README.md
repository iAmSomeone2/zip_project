# ZipProject

ZipProject is just a small Golang program for creating a zip archives of my assignments.

The default zip target in the Makefile zips all of the contents of the project directory
including object files, binaries, all folders, and generated test data. 

ZipProject allows the user to create a ".zipignore" file to indicate which files should
not be added to the zip archive.
