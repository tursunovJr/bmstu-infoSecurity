import sys
from Crypto.Util import number
from random import randrange
import random


class RSA(object):
    def __init__(self):
        # if not p:
        #     p = number.getPrime(bits)
        # if not q:
        #     q = number.getPrime(bits)
        # self.p = p if number.isPrime(p) else exit(-1)
        # self.q = q if number.isPrime(q) else exit(-1)
        self.p, self.q = self.getPrime(500)
        self.n = self.p * self.q
        self.phi = (self.p - 1) * (self.q - 1)
        self.e = self.get_e(self.phi)
        self.d = self.get_d(self.e, self.phi)

    def crypt(self, char, key):
        return char**key % self.n

    def encrypt_string(self, string):
        res = ""
        for char in string:
            ch = self.crypt(ord(char), self.d)
            res += chr(ch)
        return res

    def decrypt_string(self, string):
        res = ""
        for char in string:
            ch = self.crypt(ord(char), self.e)
            res += chr(ch)
        return res

    def getPrime(self, n):
        a, b = [], []
        for i in range(n):
            a.append(i+1)
        for i in range(1, n):
            if a[i] == False:
                continue
            for j in range((a[i]*2)-1, n, a[i]):
                if a[j] != False:
                    a[j] = False
        for i in range(0, n):
            if a[i] != False:
                b.append(a[i])
        return random.sample(b, 2)
        

    @staticmethod
    def gcdex(a, b):
        if b == 0:
            return a, 1, 0
        else:
            d, x, y = gcdex(b, a % b)
            return d, y, x - y * (a // b)

    @staticmethod
    def get_d(e, phi):
        return number.inverse(e, phi)

    @staticmethod
    def get_e(phi):
        while True:
            result = randrange(2, 255)
            modulus = number.GCD(result, phi)
            if modulus == 1:
                return result

    @staticmethod
    def get_greatest_common_divisor(a, b):
        while b != 0:
            a, b = b, a % b
        return a

    
    
