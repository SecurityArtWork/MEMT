# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import requests
import json

from datetime import datetime
from app import celery
from app import __version__

from app import socketio

from app.extensions import mongo
from app.utils import memt_dumps
from app.utils import memt_loads
from app.common import celery_namespace
from app.common import messages as msg

MICROSERVICE = "http://localhost:31337"


@celery.task(name="memt.analysis", bind=True)
def analysis(self, data):
    headers = {'user-agent': 'MEMT-Server/{}'.format(__version__),
               'content-type': 'application/json'}
    r = requests.post(MICROSERVICE,
                       data=data,
                       headers=headers)
    resp_data = {"task_id": self.request.id,
                 "sha256": sha256}
    if r.status_code == 200:
        data = r.json()
        resp_data["ecode"] = 200
        feed = mongo.db.feed
        obj = {"title": msg["feed_title"],
               "msg": msg["feed_msg"],
               "criticaly": "default",
               "sha256": data["sha256"],
               "date": datetime.utcnow()}
        if not data["strain"]:
            obj["title"] = msg["feed_title_strain"],
            obj["criticaly"] = "danger",
        feed.insert(obj)
        return True
    return False

