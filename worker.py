from time import time
from kafka import KafkaConsumer
import time

# This sleep is to solve a race condition where this was running before the broker was ready
time.sleep(30)

print("Worker Initiating")
consumer = KafkaConsumer('sample', bootstrap_servers=['broker:9092'], auto_offset_reset='earliest', group_id=None)
print("Worker Initiated")

for message in consumer:
    print("Message: ", message.value.decode('utf_8'))

print("Exited")