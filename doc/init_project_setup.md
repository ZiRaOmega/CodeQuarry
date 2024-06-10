# Setup CodeQuarry Project

## Environment Variable Setup
1. Rename `.env_empty` to `.env` and update the variable contents as needed.
2. Rename `docker-compose empty.yml` to `docker-compose.yml` and update the variable names.

## Configuring WebSocket
1. Go to `app/websocket.go`.
2. Add all your IPs or domain names to the `AllowedOrigins` array:
   ```go
   var AllowedOrigins = []string{
       "https://localhost", // for local development
   }
   ```

## Generate TLS Certificate using Certbot
1. Install Certbot:
   ```bash
   sudo apt-get install certbot
   ```
2. Run the command to generate your certificate:
   ```bash
   sudo certbot certonly --standalone -d your_domain_name
   ```
3. Copy the certificate and key to the `cert` folder:
   ```bash
   sudo cp /etc/letsencrypt/archive/your_domain_name/fullchain.pem cert/cert.pem
   sudo cp /etc/letsencrypt/archive/your_domain_name/privkey.pem cert/key.pem
   ```
4. Add these commands to `renew_certs.sh`:
   ```bash
   sudo certbot renew
   sudo cp /etc/letsencrypt/archive/your_domain_name/fullchain.pem cert/cert.pem
   sudo cp /etc/letsencrypt/archive/your_domain_name/privkey.pem cert/key.pem
   ```
5. Make the scripts executable:
   ```bash
   chmod +x renew_certs.sh
   chmod +x setup_cron_renew_cert.sh
   ```
6. Add the cron job to renew certificates automatically:
   ```bash
   ./setup_cron_renew_cert.sh
   ```

## Server Configuration in `main.go`
1. Update the certificate path:
   ```go
   fmt.Println("Server is running on https://" + URL + ":443/")
   err = http.ListenAndServeTLS(":443", "./cert/your_domain_name/cert.pem", "./cert/your_domain_name/key.pem", nil)
   if err != nil {
       app.Log(app.ErrorLevel, "Error starting the server")
       log.Fatal("[DEBUG] ListenAndServe: ", err)
   }
   ```

## Generate `sitemap.xml`
1. Go to [xml-sitemaps.com](https://www.xml-sitemaps.com/).
2. Generate your `sitemap.xml` and save it in `public/sitemap.xml`.

## PostgreSQL Configuration
1. In the `postgresql_config` folder, rename all files with the suffix `_empty.conf` to the same name without the suffix `_empty`.
   - For example, rename `pg_hba_empty.conf` to `pg_hba.conf`.
2. Update the USER in `pg_hba.conf`.
3. Open `postgre_ssl.md` and follow the instructions to configure SSL.

## Additional Documentation
- If you encounter problems while using the program, refer to the documentation in the `doc/` folder.
