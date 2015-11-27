# -*- coding: utf-8 -*-
"""Module that controls the uploads view.
"""
from __future__ import absolute_import

import os

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
        form.malware.data.save(os.path.join(app.config['UPLOAD_FOLDER'], filename))
    return redirect(url_for('upload.landing', filename=filename))

@bp.route("/landing", methods=["GET"])
def landing(filename=None):
    if filename:
        return render_template("upload/landing.html", filename)
    abort(404)

