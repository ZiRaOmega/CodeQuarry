# Setting up SSL for PostgreSQL

To enable SSL encryption for your PostgreSQL database, follow these steps:

1. Generate SSL certificates:
    - You can either use self-signed certificates or obtain trusted certificates from a certificate authority (CA).
    - If you choose to use self-signed certificates, you can generate them using the `openssl` command-line tool.
    - Here's an example of how to generate self-signed certificates:
      ```
      openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout server.key -out server.crt
      ```
      This command will generate a private key (`server.key`) and a self-signed certificate (`server.crt`).

2. Configure PostgreSQL to use SSL:
    - Open the PostgreSQL configuration file (`postgresql.conf`).
    - Uncomment or add the following lines to enable SSL:
      ```
      ssl = on
      ssl_cert_file = '/path/to/server.crt'
      ssl_key_file = '/path/to/server.key'
      ssl_ca_file = '/path/to/root.crt'
      ```
      Replace `/path/to/server.crt`, `/path/to/server.key`, and `/path/to/root.crt` with the paths to your SSL certificates.

3. Restart PostgreSQL:
    - After making the configuration changes, restart the PostgreSQL service to apply the changes.

4. Verify SSL connection:
    - Connect to the PostgreSQL database using an SSL-enabled client.
    - Ensure that the connection is established over SSL by checking the SSL status.

That's it! You have successfully set up SSL for your PostgreSQL database. Remember to keep your SSL certificates secure and regularly update them as needed.

For more detailed instructions and additional configuration options, refer to the PostgreSQL documentation.