![](logo/mdwiki_logo.png)

# go-mdwiki

## Basic Golang Markdown Wiki

- super simple 😊
- one binary + config file
- no user accounts
- pages are stored as plain markdown files
- codejar for editing markdown (https://github.com/antonmedv/codejar) with tab support
- github theme (thanks to https://github.com/sindresorhus/github-markdown-css)
- using golang 1.16 embed
- inspired by go-bwiki

## How to run

```
go run .
```

## Configuration

```
host: localhost
port: 8080
storage: .storage
```

- host & port - web server properties
- storage is directory where pages going to be stored

## Screenshots

#### Page
![](screenshot.png)

#### Editor
![](screenshot_editor.png)


## Logo

Made using https://excalidraw.com

