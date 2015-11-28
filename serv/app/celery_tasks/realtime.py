# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import requests
import json

from flask_socketio import emit
from bson.json_util import dumps

from app import celery
from app.common import rt_map_namespace
from app.common import get_latest_geo


@celery.task(name="memt.rt.feed", bind=True)
def rt_feed(self):
    pass

@celery.task(name="memt.rt.map", bind=True)
def rt_feed(self):
    print(dir(self))
    data = get_latest_geo()
    emit('update', dumps(data), namespace=rt_map_namespace)

