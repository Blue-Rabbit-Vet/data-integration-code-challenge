package com.integration.service;

import com.integration.dao.PlayerDao;
import com.integration.dto.Player;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class PlayerService {

    private final PlayerDao playerDao;

    private static final Logger LOGGER = LoggerFactory.getLogger(PlayerService.class);

    private final PublishKafkaService publishKafkaService2;

    @Autowired
    public PlayerService(PlayerDao playerDao, PublishKafkaService publishKafkaService2){
        this.playerDao = playerDao;
        this.publishKafkaService2 = publishKafkaService2;
    }

    // TODO: Lookup before creating
    public String createPlayer(Player player){
        Player existingPlayer = playerDao.findByKey(player.generateKey());
        if(existingPlayer == null){
            playerDao.save(player);
            publishKafkaService2.sendMessage("players2", player);
        }else{
            LOGGER.info("Player with key={} already existed", existingPlayer.getKey());
        }

        return player.getKey();
    }

    /*
        Function only updates number
     */
    public String updatePlayer(Player player){
        Player savedPlayer = playerDao.findByKey(player.generateKey());
        if(savedPlayer != null) {
            savedPlayer.setNumber(player.getNumber());
            playerDao.save(savedPlayer);
        }

        return player.getKey();
    }
}
