#!/usr/bin/env python
import matplotlib
matplotlib.use('Agg')
import pylab
from pylab import genfromtxt;
mat0 = genfromtxt("plotfile.txt")
pylab.ylabel('No. of Transactions')
pylab.xlabel('Transaction Time (milliseconds)')
#pylab.title('Simulation Hyperledger Chaincode for Donor Registration')
pylab.plot(mat0[:,1],mat0[:,0])

pylab.savefig('10.png')

