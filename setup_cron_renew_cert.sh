#!/bin/bash

# Add the cron job to renew certificates
(crontab -l ; echo "1 0 * * * cd Project-Exam/ && chmod +x renew_certs.sh && ./renew_certs.sh >> renew_certs.log 2>&1") | crontab -