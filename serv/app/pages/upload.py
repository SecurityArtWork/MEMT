# -*- coding: utf-8 -*-
"""Module that controls the uploads view.
"""
from __future__ import absolute_import

from flask import Blueprint, render_template


bp = Blueprint('upload', __name__, url_prefix='/upload')


@bp.route('', methods=['GET'])
def index():
    return render_template('upload/index.html')
