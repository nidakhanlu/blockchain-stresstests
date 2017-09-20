#!/usr/bin/env python
from scipy import stats as st
import numpy as np
import scipy as sp
from matplotlib import pyplot

mat0 = np.loadtxt("offline1000.txt")
a=np.asarray(mat0)
#a=np.array([1,2,3,4])
#standard deviation
sd=np.std(a,ddof=1,dtype=np.float64)
print("Standard Deviation:")
print(sd)
#mean
mean=np.mean(a)
print("Mean:")
print(mean) 
z_critical = st.norm.ppf(q = 0.975)
print('z critical:')
print(z_critical)
margin_of_error = z_critical * (sd/np.sqrt(len(a)))
print("Margin of Error:")
print(margin_of_error)
cfd_left = mean - margin_of_error
cfd_right = mean + margin_of_error                    
print("Confidence interval:")
print(cfd_left)
print(cfd_right)
xdata=(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15)
pyplot.scatter(xdata,mat0)
pyplot.axhline(y=mean, color='r', linestyle='-')
right=cfd_right
left=cfd_left
pyplot.text(0.9, 0.75, 'Standard Deviation = 0.089', style='italic')
pyplot.text(0.9, 0.73, 'Mean = 0.965', color='r',style='italic')
pyplot.text(0.9, 0.71, 'Margin of Error = 0.045', style='italic')
pyplot.text(0.90, 1.02, 'Confidence Interval = 0.919 - 1.010', color='b',alpha=0.7,style='italic')
pyplot.xlabel('Samples of 10000 transactions')
pyplot.ylabel('Normalized Data for Mined Transactions')
pyplot.title('Offline: Total Number of Transactions = 40000')
# hatch confidence interval
#fill_between(xdata[left:right], ydata[left:right], facecolor='blue', alpha=0.5)
pyplot.fill_between(xdata, right, left,color='b',alpha=0.2)
pyplot.show()
