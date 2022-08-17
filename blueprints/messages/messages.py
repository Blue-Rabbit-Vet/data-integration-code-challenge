from email import message
from flask_smorest import Api, Blueprint, abort
from model import Message
from .schema import MessageSchema
from flask.views import MethodView
from kafka import KafkaProducer



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

        producer = KafkaProducer(bootstrap_servers='broker:9092')
        producer.send(new_data['topic'], new_data['message'].encode('UTF_8'))
        producer.flush()
        producer.close()
        return item
