# gen key
openssl dsaparam -out dsaparam.pem 1024
openssl gendsa -out privkey.pem dsaparam.pem
openssl gendsa -out privkey.pem dsaparam.pem
