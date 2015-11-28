# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

from app import celery

@celery.task(name="memt.analysis")
def analysis(data):
    pass
