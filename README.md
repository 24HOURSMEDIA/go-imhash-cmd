# imhash command

Command line utility in GoLang to generate perceptual hashes from imag files.

## examples

Basic usage

---
command:
```
go run imhash.go hash ./resources/test1.jpg ./resources/test1_modified.jpg ./resources/test1_modified2.jpg
```
output:
```
8c1e07e8f86864f8
8c1e07e8f86860f8
8c1e07e8f868649c
```

---
command:
```
go run imhash.go distance ./resources/test1.jpg ./resources/test1_modified.jpg
```
output:
```
1
```

---
command:
```
go run imhash.go distance 8c1e07e8f86864f8 8c1e07e8f868649c -u
```
output:
```
3
```

---