# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import
import time

from bson.json_util import dumps

from flask import request, session
from flask_socketio import emit


from app.common import celery_namespace

from app.celery_tasks.realtime import rt_feed


thread = None

@socketio.on('connect', namespace=celery_namespace)
def connect(task_id):

    emit("connect", dumps(data), namespace=celery_namespace)


