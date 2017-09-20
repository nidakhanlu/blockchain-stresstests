#!/usr/bin/env python
# -*- coding: utf-8 -*-
from __future__ import unicode_literals
import numpy as np
import matplotlib.pyplot as plt

N = 5
syn_means = (6300)
syn_std = (1940)

ind = np.arange(N)  # the x locations for the groups
width = 0.35       # the width of the bars

fig, ax = plt.subplots()
rects1 = ax.bar( 1,syn_means, width, color='y', alpha=0.5,yerr=syn_std,error_kw=dict(ecolor='gray', lw=1, capsize=4, capthick=1))

asyn_means = (9290)
asyn_std = (310)
rects2 = ax.bar( 2,asyn_means, width, color='g',alpha=0.5, yerr=asyn_std,error_kw=dict(ecolor='gray', lw=1, capsize=4, capthick=1))

offline_means = (6530)
offline_std = (2300)
rects3 = ax.bar( 3,offline_means, width, color='b', alpha=0.5,yerr=offline_std,error_kw=dict(ecolor='gray', lw=1, capsize=4, capthick=1))


# add some text for labels, title and axes ticks
ax.set_ylabel('Mean of Mined Transactions')
ax.set_title('Comparison of Different Mechanisms of Transaction Submission')
ax.set_xticks(ind + width / 2)
ax.set_xticklabels(('','Synchronous', 'Asynchronous', 'Offline'))

#ax.legend((rects1[0], rects2[0]), ('', 'Women'))


def autolabel(rects):
    """
    Attach a text label above each bar displaying its height
    """
    for rect in rects:
        height = rect.get_height()
        ax.text(rect.get_x() + rect.get_width()/2., 1.05*height,
                '%d' % int(height),
                ha='center', va='bottom')

#autolabel(rects1)
#autolabel(rects2)
#autolabel(rects3)
plt.text(0.5, 6800,'$\sigma$ = 1940', color='black',alpha=1)
plt.text(0.5, 6400,'$\mu$=6300', color='black',alpha=1)
plt.text(1.5, 8800,'$\sigma$ = 310', color='black',alpha=1)
plt.text(1.5, 8400,'$\mu$=9290', color='black',alpha=1)
plt.text(2.6, 7000,'$\sigma$ = 2530', color='black',alpha=1)
plt.text(2.6, 6600,'$\mu$=6530', color='black',alpha=1)
plt.show()
