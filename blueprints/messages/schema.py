import marshmallow as ma

class MessageSchema(ma.Schema):
    id = ma.fields.Int(dump_only=True)
    message = ma.fields.String(required=True)
    topic = ma.fields.String(required=True)
