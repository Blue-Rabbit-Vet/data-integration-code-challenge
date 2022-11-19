package com.integration.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class PublishKafkaService {

    private final KafkaTemplate<String, Object> kafkaTemplate;

    @Autowired
    PublishKafkaService(KafkaTemplate kafkaTemplate){
        this.kafkaTemplate = kafkaTemplate;
    }

    public void sendMessage(String topic, Object payload){
        kafkaTemplate.send(topic,payload);
    }
}
