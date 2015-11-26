# -*- coding: utf-8 -*-
"""Module that controls the download view.
"""
from __future__ import absolute_import

from flask import Blueprint, render_template


bp = Blueprint('download', __name__, url_prefix='/downloads')


@bp.route('', methods=['GET'])
def index():
    return render_template('downloads/index.html')
