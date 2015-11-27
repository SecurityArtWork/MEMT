# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

from flask_wtf import Form
from wtforms import FileField
from wtforms.validators import DataRequired


class UploadForm(Form):
    malware = FileField("Malware", validators=[DataRequired()])
