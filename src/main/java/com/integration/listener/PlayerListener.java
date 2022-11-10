package com.integration.listener;

import org.apache.avro.generic.GenericRecord;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.support.Acknowledgment;
import org.springframework.stereotype.Component;

@Component
public class PlayerListener {

    private static final Logger LOGGER = LoggerFactory.getLogger(PlayerListener.class);
    /*
        IMPROVEMENTS: ACK the record instead of auto commit
     */

    @KafkaListener(id="player", topics = "player", groupId="testGroup")
    public void consumePlayer(final ConsumerRecord<String, GenericRecord> consumerRecord, Acknowledgment acknowledgment){

        acknowledgment.acknowledge();
    }
}
