from app import db

class Message (db.Model):
    id = db.Column(db.Integer, primary_key=True)
    message = db.Column(db.String(), nullable=False)
    status = db.Column(db.String(), nullable=True)
    topic = db.Column(db.String(), nullable=True)