FROM  jboss/keycloak
COPY ./standalone.xml /opt/jboss/keycloak/standalone/configuration/standalone.xml
COPY ./standalone-ha.xml /opt/jboss/keycloak/standalone/configuration/standalone-ha.xml
