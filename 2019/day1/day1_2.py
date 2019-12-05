import sys
import math

def required(i):
    return int(math.floor(i / 3)) - 2

def fuel(m):
    l = []
    while m > 0:
        l.append(m)
        m = required(m)
    return sum(l[1:]) #initial value is mass


print(sum(fuel(int(l)) for l in sys.stdin.readlines()))
