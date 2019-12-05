import sys
import math

print(sum([int(math.floor(int(l) / 3) - 2) for l in sys.stdin.readlines()]))
