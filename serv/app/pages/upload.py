# -*- coding: utf-8 -*-
"""Module that controls the uploads view.
"""
from __future__ import absolute_import

import os
import hashlib
import geoip2.database

from geoip2.errors import AddressNotFoundError

from datetime import datetime

from flask import current_app as app
from flask import Blueprint
from flask import render_template
from flask import request
from flask import redirect
from flask import url_for
from flask import abort

from werkzeug import secure_filename

from app.forms.upload import UploadForm
from app.celery_tasks.analysis import analysis

from app.utils import memt_dumps


bp = Blueprint('upload', __name__, url_prefix='/upload')


@bp.route('', methods=['GET'])
def index():
    form = UploadForm()
    return render_template('upload/index.html', form=form)

@bp.route('/submit', methods=['POST'])
def submit():
    form = UploadForm()
    if form.validate_on_submit():
        filename = secure_filename(form.malware.data.filename)
        form.malware.data.save(os.path.join(app.config['TMP_UPLOAD_FOLDER'], filename))
        with open(os.path.join(app.config['TMP_UPLOAD_FOLDER'], filename), 'rb') as malware:
            data = malware.read()
            sha256 = hashlib.sha256(data).hexdigest()
            if os.path.isfile(os.path.join(app.config['BIN_UPLOAD_FOLDER'],sha256)):
                return redirect(url_for("detail.index", hash=sha256))
        ## Celery
        obj = {
            "path": os.path.join(app.config['TMP_UPLOAD_FOLDER'], filename),
            "sha256": sha256
        }
        reader = geoip2.database.Reader(app.config['MAXMAIN_DB_CITIES'])
        try:
            response = reader.city(request.remote_addr)
        except (AddressNotFoundError):
            obj["ipMeta"] = [{
                                "city": "unknown",
                                "ip": request.remote_addr,
                                "country": "unknown",
                                "iso_code": "unknown",
                                "date": datetime.utcnow(),
                                "geo": [0.0, 0.0]
                            }]
        else:
            obj["ipMeta"] = [{
                                "city": response.city.name,
                                "ip": request.remote_addr,
                                "country": response.country.name,
                                "iso_code": response.country.iso_code,
                                "date": datetime.utcnow(),
                                "geo": [response.location.longitude, response.location.latitude]
                            }]
        # Celery task
        task_id = analysis.delay(memt_dumps(obj))
        return redirect(url_for('upload.landing', hash=sha256, task_id=task_id.id))
    return redirect(url_for("index.index"))


@bp.route("/landing/<sha256:hash>/<task_id>", methods=["GET"])
def landing(hash, task_id):
    return render_template("upload/landing.html", hash=hash, task_id=task_id)


@bp.route("/landing/", methods=["GET"])
def landing_barra(filename=None):
    return redirect(url_for("upload.landing"))


@bp.route("/", methods=["GET"])
def index_barra(filename=None):
    return redirect(url_for("upload.index"))
