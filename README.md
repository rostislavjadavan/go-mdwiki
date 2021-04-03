![](logo/mdwiki_logo.png)

# mdwiki

## Golang Markdown Wiki

- super simple 😊
- one binary + config file
- no user accounts
- pages are stored as plain markdown files
- page versioning 
- github theme (thanks to https://github.com/sindresorhus/github-markdown-css)
- Echo used as underlying http framework (https://echo.labstack.com/)  
- using golang 1.16 (because of embed feature)
- inspired by go-bwiki

## Screenshots

#### Homepage
![](screenshots/homepage.png)

#### Editor
![](screenshots/edit_page.png)

#### Page versions
Every update will create new version of the page. 

![](screenshots/page_versions.png)

You can then view and restore any version you want.

![](screenshots/page_version.png)

#### Delete page
Deleted pages are moved to the trash. You can view them or restore them if you need.

![](screenshots/page_in_trash.png)

_screenshots made using:_ https://www.screely.com

## How to build

```
go build .
```

## Configuration
Application loads `config.yml` on start up.

```
host: localhost
port: 8080
storage: .storage
```

- `host` & `port` - web server configuration
- `storage` is directory where pages going to be stored

## Running as `systemd` service

# Service file example

```
[Unit]
Description=mdwiki
After=network.target
Wants=network-online.target

[Service]
Restart=always
Type=simple
WorkingDirectory=/opt
ExecStart=/opt/mdwiki

[Install]
WantedBy=multi-user.target
```

- update `ExecStart` and `WorkingDirectory` based on installation path
- working directory needs to be set because `mdwiki`
  expects `config.yml` to be in the same directory as executable
- place file in `/etc/systemd/system/mdwiki.service`

## Running service

- start service: `systemctl start mdwiki`
- stop service: `systemctl stop mdwiki`
- service status: `systemctl status mdwiki`
- enable service: `systemctl enable mdwiki` (`enable` will hook the specified unit into relevant places, so that it will automatically start on boot)
  

## Logo

Made using https://excalidraw.com

