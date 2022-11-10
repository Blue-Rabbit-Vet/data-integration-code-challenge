package com.integration.service;

import com.integration.dao.PlayerDao;
import com.integration.dto.Player;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class PlayerService {

    private final PlayerDao playerDao;

    @Autowired
    public PlayerService(PlayerDao playerDao){
        this.playerDao = playerDao;
    }


    public String createPlayer(Player player){
        playerDao.save(player);
        return "Some Key";
    }
}
