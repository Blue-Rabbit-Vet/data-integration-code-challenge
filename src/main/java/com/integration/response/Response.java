package com.integration.response;

import org.springframework.http.HttpStatus;

public class Response {

    private HttpStatus status;
    private String message;

    public Response(final HttpStatus status, final String message) {
        this.status = status;
        this.message = message;
    }

    public Response(){}

    public HttpStatus getStatus() {
        return status;
    }

    public void setStatus(final HttpStatus status) {
        this.status = status;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(final String message) {
        this.message = message;
    }
}
