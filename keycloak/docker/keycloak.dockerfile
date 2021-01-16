FROM  jboss/keycloak
COPY ./keycloak/docker/standalone.xml /opt/jboss/keycloak/standalone/configuration/standalone.xml
COPY ./keycloak/docker/standalone-ha.xml /opt/jboss/keycloak/standalone/configuration/standalone-ha.xml
