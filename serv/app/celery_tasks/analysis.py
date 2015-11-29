# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import requests
import json

from datetime import datetime
from app import celery
from app import __version__


MICROSERVICE = "http://localhost:31337"


@celery.task(name="memt.analysis", bind=True)
def analysis(self, data):
    headers = {'user-agent': 'MEMT-Server/{}'.format(__version__),
               'content-type': 'application/json'}
    r = requests.post(MICROSERVICE,
                       data=data,
                       headers=headers)
    if r.status_code == 200:
        return True
    return False

