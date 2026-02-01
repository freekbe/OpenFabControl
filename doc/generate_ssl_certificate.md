To generate ssl certificate, run this command :
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -subj "/C=BE/ST=State/L=City/O=Organization/CN=localhost"
