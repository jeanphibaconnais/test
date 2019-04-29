### Dashboard Nantastiques

Pour exécuter le projet en local : 
’’’
./mvnw compile quarkus:dev
’’’

Pour packager l'application (jdk 1.8) : 
’’’
./mvnw package
’’’

Pour déployer l'application dans un conteneur docker : 
’’’
docker build -f src/main/docker/Dockerfile  -t dashboard-nantastiques .
docker run -i --rm -p 8080:8080 dashboard-nantastiques
’’’

et l'application est disponible via : http://localhost:8008/supervision