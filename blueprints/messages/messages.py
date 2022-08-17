from flask_smorest import Api, Blueprint, abort
from model import Message
from .schema import MessageSchema
from flask.views import MethodView
from flask import current_app


blp = Blueprint("messages", "messages", url_prefix="/api", description="Operations for messaging queue")

@blp.route("/api")
class Messages(MethodView):
    @blp.response(200, MessageSchema(many=True))
    def get(self):
        """List Messages"""
        return Message.query.all()

    @blp.arguments(MessageSchema)
    @blp.response(201, MessageSchema)
    def post(self, new_data):
        """Add a new Message"""
        from app import db
        item = Message(**new_data)
        db.session.add(item)
        db.session.commit()
        return item
