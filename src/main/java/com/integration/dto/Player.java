package com.integration.dto;

import org.bson.types.ObjectId;
import org.springframework.data.annotation.Id;

import java.util.UUID;


public class Player {

    private String firstName;
    private String lastName;
    private Integer number;

    @Id
    private ObjectId id;

    private String key;

    public Player(){
    }

    public String getKey() {
        return key;
    }

    private void setKey(String key) {
        this.key = key;
    }

    public String getFirstName() {
        return firstName;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getLastName() {
        return lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public Integer getNumber() {
        return number;
    }

    public void setNumber(Integer number) {
        this.number = number;
    }

    public String generateKey(){
        String generatedKey = this.getFirstName()+this.getLastName();
        String key = UUID.nameUUIDFromBytes(generatedKey.getBytes()).toString();
        setKey(key);
        return key;
    }
}
