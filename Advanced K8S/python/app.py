from flask import Flask, Response, request, Blueprint
import sys
from datetime import datetime
import json
import requests
import os

API_VERSION = 'v1'
SOURCE_URL = "http://api.pnd.gs/v1/sources/"
# Flask Blueprint
api = Blueprint('customEndpoints',__name__)

# Configure PORT
PORT = 5000 
if len(sys.argv) > 1:
    PORT = int(sys.argv[1])

# Create a flask APP object
app = Flask(__name__)

# **** UTIL FUNCTIONS :Begins****

def sendSuccessResponse(resp_data, status=200, mimetype="application/json"):
    return Response(response=resp_data, status=status, mimetype=mimetype)

def get(url):
    resp = requests.get(url)
    if resp.ok:
        return resp.json()

def get_news_source():
    content = get(SOURCE_URL)
    data = {
        "latest": [],
        "popular": []
    }
    if content:
        for item in content:
            data["latest"].append( item["endpoints"]["latest"] )
            data["popular"].append( item["endpoints"]["popular"] )
    return data

# **** UTIL FUNCTIONS :Ends****

# Add a base route
# "/"
@app.route("/", methods=['GET'])
def root():
    resp_data = {
        "status": "ok",
        "hostname": os.getenv("HOSTNAME")
        }
    resp_data = json.dumps(resp_data, indent=4)
    return sendSuccessResponse(resp_data)

# /api/v1/time
@api.route("/time", methods=['GET'])
# @accept("application/json")
def time():
    resp_data = { "time": str(datetime.now()) }
    resp_data = json.dumps(resp_data, indent=4)
    return sendSuccessResponse(resp_data)

# /api/v1/sources
@api.route("/source", methods=["GET"])
def getsources():
    resp_data = get_news_source()
    resp_data = json.dumps(resp_data, indent=4)
    return sendSuccessResponse(resp_data)
        


# Registering Blueprints to /
app.register_blueprint(api, url_prefix ='/api/' + API_VERSION)

# Start it on this "PORT"
print("Starting APP in {} PORT".format(PORT))
app.run(host='0.0.0.0', port=PORT, debug=False)
