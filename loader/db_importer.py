# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import sys
import os
import json
import random
import socket
import struct
import geoip2.database
from geoip2.errors import AddressNotFoundError

from datetime import datetime
from pymongo import MongoClient


DB_NAME = "memt"
UPLOAD_PATH = ""

IP = {
      "city": "unknown",
      "ip": "unknown",
      "country": "unknown",
      "iso_code": "unknown",
      "date": datetime.utcnow(),
      "geo": [0.0, 0.0]
}

reader = geoip2.database.Reader('/opt/dbs/GeoLite2-City.mmdb')

def main():
    if len(sys.argv) == 1:
        print("You must provide a file path as an argument.")
        sys.exit(-1)
    print("Connecting to DB {}...".format(DB_NAME))
    client = MongoClient('mongodb://localhost:27017/')

    with open(sys.argv[1]) as data_file:
        data = json.load(data_file)
        db = client[DB_NAME]
        for obj in data:
            obj["date"] = datetime.utcnow()
            obj['imagedir'] = os.path.join("..", "aux", "malware", "images", obj['imagedir'].split("/")[-1:][0])
            obj['artifactdir'] = os.path.join("..", "aux", "malware", "artifacts", obj['artifactdir'].split("/")[-1:][0])
            if not obj['symbols']:
                obj['symbols'] = []
            if not obj['imports']:
                obj['imports'] = []
            if not obj['sections']:
                obj['sections'] = []
            if not obj['sections']:
                obj['sections'] = []
            if not obj['mutations']:
                obj['mutations'] = []
            if "ipmeta" in obj:
                for ip in obj["ipmeta"]:
                    ip["date"] = datetime.utcnow()
            else:
                ip = socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff)))
                try:
                    response = reader.city(ip)
                except (AddressNotFoundError):
                    obj["ipmeta"] = [IP]
                else:
                    obj["ipmeta"] = [{
                        "city": response.city.name,
                        "ip": ip,
                        "country": response.country.name,
                        "iso_code": response.country.iso_code,
                        "date": datetime.utcnow(),
                        "geo": [response.location.longitude, response.location.latitude]
                    }]
            #print(obj)
            res = db.assets.insert_one(obj)
            print("ADDED: {}".format(res.inserted_id))
    reader.close()
    print("Importer has finished his job.")

if __name__ == '__main__':
    main()
