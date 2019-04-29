### Quarkus : tests

Nécessite :

- JDK 8
- Apache Maven 3.5.3+

Sources de tests Quarkus : git clone https://github.com/quarkusio/quarkus-quickstarts.git

Création d'un projet :
’’’
mvn io.quarkus:quarkus-maven-plugin:0.13.2:create \
 -DprojectGroupId=org.acme \
 -DprojectArtifactId=getting-started \
 -DclassName="org.acme.quickstart.GreetingResource" \
 -Dpath="/hello"
’’’

Ce qui nous donne comme info :

’’’
[INFO] ========================================================================================
[INFO] Your new application has been created in /Users/JeanPhi/MesDocuments/Developpement/git/tests/quarkus/getting-started
[INFO] Navigate into this directory and launch your application with mvn compile quarkus:dev
[INFO] Your application will be accessible on http://localhost:8080
[INFO] ========================================================================================
’’’

Exécution en local : ./mvnw compile quarkus:dev

Packaging de l'application : ./mvnw package

Construction d'une image docker : docker build -f src/main/docker/Dockerfile -t quarkus-getting-started .

Création d'un conteneur : docker run -i --rm -p 8080:8080 quarkus-getting-started
