# gen key
openssl dsaparam -out dsaparam.pem 1024
openssl gendsa -out privkey.pem dsaparam.pem
openssl gendsa -out privkey.pem dsaparam.pem

# sign
# openssl dgst -dss1 -sign dsa.key -out sig.data myfile.data

# verify
# openssl dgst -dss1 -verify dsa.key.pub -signature sig.data myfile.data

# sign with my program
# go build dsa.go

# ./dsa -action sign -key dsa.key -file myfile.data > sig.data
# ./dsa -action sign -key dsa.key -file myfile.data | penssl asn1parse -inform DER

# verify with my program
# ./dsa -action verify -pubkey dsa.key.pub -file myfile.data -signature sig.data