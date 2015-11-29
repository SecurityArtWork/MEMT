# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import
import time

from flask_socketio import emit

from app.celery_tasks.analysis import analysis
from app.common import celery_namespace
from app.utils import memt_dumps

from app.celery_tasks.realtime import rt_feed


thread = None

@socketio.on('connect', namespace=celery_namespace)
def connect():
    emit("connect", memt_dumps({}), namespace=celery_namespace)


@socketio.on('landing', namespace=celery_namespace)
def update(task_id):
    task = analysis.AsyncResult(task_id)
    data = {"status": task.status}
    emit("update", memt_dumps(data), namespace=celery_namespace)
