# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import requests
import json

from flask_socketio import emit
from bson.json_util import dumps

from app import celery


@celery.task(name="memt.rt.feed", bind=True)
def rt_feed(self):
    pass

@celery.task(name="memt.rt.map", bind=True)
def rt_feed(self):
    pass

