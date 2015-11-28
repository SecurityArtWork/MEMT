# -*- coding: utf-8 -*-
"""Module that controls the details view.
"""
from __future__ import absolute_import

from flask import Blueprint, render_template


bp = Blueprint('detail', __name__, url_prefix='/detail')


@bp.route('/<sha256:hash>', methods=['GET'])
def index(hash):
    return render_template('detail/index.html')
