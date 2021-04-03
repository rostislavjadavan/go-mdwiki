# systemd service file

- update `ExecStart` and `WorkingDirectory` based on installation path
- working directory needs to be set because `mdwiki` 
  expects `config.yml` to be in the same directory as executable
- copy file to `/etc/systemd/system/mdwiki.service`
  
## Running service

- start service: `systemctl start mdwiki`
- stop service: `systemctl stop mdwiki`
- service status: `systemctl status mdwiki`
- enable service: `systemctl enable mdwiki` (`enable` will hook the specified unit into relevant places, so that it will automatically start on boot)
