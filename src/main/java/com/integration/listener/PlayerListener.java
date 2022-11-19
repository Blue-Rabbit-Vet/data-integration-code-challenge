package com.integration.listener;

import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.integration.dto.Player;
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
                      Create Separate Consumer Factory instead of appending properties directly to listener
     */

    @KafkaListener(id="players", topics = "players2", groupId="testGroup", properties = {"value.deserializer:org.springframework.kafka.support.serializer.JsonDeserializer", "spring.json.trusted.packages:*"}
    ,containerFactory = "kafkaListenerContainerFactory")
    public void consumePlayer(final ConsumerRecord<String, Player> consumerRecord){
        Player player = consumerRecord.value();
        LOGGER.info("Received Player, firstName={}, lastName={}, number={}",player.getFirstName(), player.getLastName(), player.getNumber());

    }
}
