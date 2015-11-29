# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import requests
import json

from app import celery
from app import __version__

from app import socketio

from app.common import celery_namespace
from app.utils import memt_dumps
from app.utils import memt_loads

MICROSERVICE = "http://localhost:31337"


@celery.task(name="memt.analysis", bind=True)
def analysis(self, data):
    try:
        tmp = memt_loads(data)
        sha256 = tmp["sha256"]
        del tmp["sha256"]
        data = memt_dumps(tmp)
    except (TypeError):
        pass

    headers = {'user-agent': 'MEMT-Server/{}'.format(__version__),
               'content-type': 'application/json'}
    r = requests.post(MICROSERVICE,
                       data=data,
                       headers=headers)
    resp_data = {"task_id": self.request.id,
                 "sha256": sha256}
    if r.status_code == 200:
        resp_data["ecode"] = 200
        socketio.emit("update", memt_dumps(resp_data), namespace=celery_namespace)
        return True
    resp_data["ecode"] = 500
    socketio.emit("update", memt_dumps(resp_data), namespace=celery_namespace)
    return False

