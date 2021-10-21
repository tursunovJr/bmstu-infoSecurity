from base64 import b32encode, b32decode
from rsa import RSA
import sys

if __name__ == '__main__':
        filename = "file.txt"
        with open(filename, 'rb') as file1:
            data = file1.read()
            rsa = RSA()
            print("\tP:", rsa.p, "\n\tQ:", rsa.q, "\n\tE:", rsa.e, "\n\tN:", rsa.n, "\n\tD:", rsa.d)
            str = b32encode(data)
            dta = str.decode("ascii")
            print("Encrypting...")
            enc = rsa.encrypt_string(dta)
            print("Encrypted: ",  enc)
            print("Decrypting...")
            dec = rsa.decrypt_string(enc)
            restored = b32decode(dec)
            print("Decrypted: ", restored)