#!/bin/bash

# Create the file lms.service with the specified content
cat << EOF > /etc/systemd/system/lms.service
[Unit]
Description=lms backend http server
After=mariadb.service

[Service]
Restart=always
User=root
WorkingDirectory=/root/library-management-system-backend/
ExecStart=/root/library-management-system-backend/cmd/lms
Requires=mariadb.service

[Install]
WantedBy=multi-user.target
EOF

# Reload the systemd daemon
sudo systemctl daemon-reload

# Enable the lms.service file so it starts on boot
sudo systemctl enable lms.service

# Reload the systemd daemon again to pick up the changes from enabling the service
sudo systemctl daemon-reload

# Start the lms.service
sudo systemctl start lms.service
