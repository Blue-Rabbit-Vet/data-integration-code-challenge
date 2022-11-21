package com.integration.dao;

import com.integration.dto.Player;
import org.springframework.data.mongodb.repository.MongoRepository;

/*
    Lookup extends
 */
public interface PlayerDao extends MongoRepository<Player, String> {

    Player findByKey(String key);
}
