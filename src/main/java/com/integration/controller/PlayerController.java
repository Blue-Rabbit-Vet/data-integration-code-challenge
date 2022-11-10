package com.integration.controller;

import com.integration.service.PlayerService;
import com.integration.dto.Player;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import com.integration.response.Response;


@Controller
public class PlayerController {

    private static final Logger LOGGER = LoggerFactory.getLogger(PlayerController.class);

    private final PlayerService playerService;

    @Autowired
    public PlayerController(PlayerService playerService) {
        this.playerService = playerService;
    }


    @PostMapping(value = "/addplayer", headers = "Accept=application/json")
    public ResponseEntity<Response> addPlayer(@RequestBody Player player){
        LOGGER.info("Received Player, firstName={}, lastName={}, number={}",player.getFirstName(), player.getLastName(), player.getNumber());

        playerService.createPlayer(player);

        return new ResponseEntity<>(new Response(HttpStatus.OK,"Successful").getStatus());
    }




}