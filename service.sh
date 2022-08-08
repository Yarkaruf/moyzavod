echo "
[Unit]
Description=Zavod
StartLimitIntervalSec=0
After=postgresql.service
Requires=postgresql.service

[Service]
Type=simple
Restart=always
RestartSec=1
User=root
WorkingDirectory=$PWD
ExecStart=$PWD/zavod
[Install]
WantedBy=multi-user.target" > /etc/systemd/system/zavod.service

systemctl enable zavod.service
service zavod start
service zavod status
