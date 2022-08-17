from flask import Flask
from flask_smorest import Api
from flask_sqlalchemy import SQLAlchemy


app = Flask(__name__, static_url_path='')
app.config["API_TITLE"] = "Data Integration challenge"
app.config["API_VERSION"] = "v1"
app.config["OPENAPI_VERSION"] = "3.0.3"
app.config["OPENAPI_URL_PREFIX"] = "docs"
app.config["OPENAPI_RAPIDOC_PATH"] = "ui"
app.config["OPENAPI_RAPIDOC_URL"] = "https://cdn.jsdelivr.net/npm/rapidoc@9.3.3/dist/rapidoc-min.js"
api = Api(app)

app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///example.sqlite"
db = SQLAlchemy(app)

from blueprints.messages.messages import blp
api.register_blueprint(blp)

@app.route("/")
def home():
    db.create_all()
    return "<p>Hello, World!</p>"

