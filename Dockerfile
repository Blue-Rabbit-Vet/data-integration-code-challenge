FROM amazoncorretto:11
RUN mkdir -p /app/
COPY ../build/libs/DataIntegrationCodeChallenge-1.0.jar /app/app.jar
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]