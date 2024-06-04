### Setup environement variable
 - First rename .env_empty to .env and change variable content 
 - Then rename docker-compose empty.yml to docker-compose.yml and change the variable name

## Go to app/websocket.go
- Add in the array all you're ip or domaine_name
  ```go
    var AllowedOrigins = []string{
	"https://localhost", // for local development
    }
    ```
## Generate your TLS cert using certbot
- First install certbot
- Then run the command
  ```bash
  sudo certbot certonly --standalone -d your_domain_name
  ```
- Then copy the cert and key to the folder cert
    ```bash
    sudo cp /etc/letsencrypt/archive/your_domain_name/fullchain.pem cert/cert.pem
    sudo cp /etc/letsencrypt/archive/your_domain_name/privkey.pem cert/key.pem
    ```
- Add theses commands to renew_certs.sh
- Then
  ```bash
  chmod +x renew_certs.sh
  chmod +x setup_cron_renew_cert.sh
  ```
- Add the cron job
  ```bash
  ./setup_cron_renew_cert.sh
  ```
## Go to main.go
- Change the path to your cert path
  ```go
  fmt.Println("Server is running on https://" + URL + ":443/")
	err = http.ListenAndServeTLS(":443", "./cert/fullchain1.pem", "./cert/privkey1.pem", nil)
	if err != nil {
		app.Log(app.ErrorLevel, "Error starting the server")
		log.Fatal("[DEBUG] ListenAndServe: ", err)
	}
    ```
## Generate your sitemap.xml
- go to  www.xml-sitemaps.com
- Generate your sitemap.xml and save it in public/sitemap.xml

## Next go to postgresql_config folder
- Rename *_empty.conf and update USER in pg_hba.conf
- Open postgre_ssl.md and follow instructions

## Rest of documentation is in doc/
- If you have problem while using the program go read the documentation !