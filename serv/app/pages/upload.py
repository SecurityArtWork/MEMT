# -*- coding: utf-8 -*-
"""Module that controls the uploads view.
"""
from __future__ import absolute_import

import os
import hashlib

from flask import current_app as app
from flask import Blueprint
from flask import render_template
from flask import request
from flask import redirect
from flask import url_for
from flask import abort

from werkzeug import secure_filename

from app.forms.upload import UploadForm


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
            new_name = hashlib.sha256(data).hexdigest()
            if not os.path.isfile(os.path.join(app.config['BIN_UPLOAD_FOLDER'],new_name)):
                os.rename(os.path.join(app.config['TMP_UPLOAD_FOLDER'], filename), os.path.join(app.config['BIN_UPLOAD_FOLDER'], new_name))
            else:
                return redirect(url_for("detail.index", hash=new_name))
    return redirect(url_for('upload.landing', hash=new_name))

@bp.route("/landing", methods=["GET"])
def landing(filename=None):
    if filename:
        return render_template("upload/landing.html", filename)
    abort(404)

