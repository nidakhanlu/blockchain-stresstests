#!/usr/bin/env python
import numpy as np
from matplotlib import pyplot
from pylab import genfromtxt
from pylab import loadtxt
mat0 = np.loadtxt("10000nowifi.txt")
x=np.asarray(mat0)
mat1=np.genfromtxt("timestamp10000nowifi.txt")
#pyplot.plot(mat0,mat1,label="Timestamp")

pyplot.plot(x,mat1)
pyplot.xlabel('Number of Donor Registrations')
pyplot.ylabel('Transaction Timestamp in Seconds since Epoch')
pyplot.title('N = 10000, No Delay, No WiFi')
#pyplot.legend();
#pyplot.xticks(x)
#pyplot.show()
pyplot.savefig('Timestamp10000nowifi.png')
