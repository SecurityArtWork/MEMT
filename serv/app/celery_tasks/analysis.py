# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import requests
import json

from datetime import datetime
from pymongo import MongoClient

from config import Config
from app import __version__
from app import celery
from app import socketio

from app.utils import memt_dumps
from app.utils import memt_loads
from app.common import celery_namespace
from app.common import messages as msg

MICROSERVICE = "{}".format(Config.ANAL_SERVICE)
MONGO_CONN = "mongodb://{}:{}/{}".format(Config.MONGO_HOST,
                                         Config.MONGO_PORT,
                                         Config.MONGO_DBNAME)

@celery.task(name="memt.analysis", bind=True)
def analysis(self, data):
    headers = {'user-agent': 'MEMT-Server/{}'.format(__version__),
               'content-type': 'application/json'}
    r = requests.post(MICROSERVICE,
                       data=data,
                       headers=headers)
    resp_data = {"task_id": self.request.id}
    if r.status_code == 200:
        data = r.json()
        resp_data["ecode"] = 200
        if data["strain"] == "":
            obj = {"title": msg["feed_title_strain"],
                   "msg": msg["feed_msg_strain"],
                   "criticaly": "danger",
                   "sha256": data["sha256"],
                   "date": datetime.utcnow()
                   }
            client = MongoClient("mongodb://localhost:27017/memt")
            feed = client.db.feed

            print("INSERT FEED: {}".format(obj))
            a = feed.insert(obj)
            print("ID: {}".format(dir(a)))
        else:
            print("No strain found: {}".format(data["strain"]))
        return True
    return False

