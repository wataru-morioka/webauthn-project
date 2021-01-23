module.exports = {
    client: {
      name: 'client',
      includes:  ["src/**/*.{ts,tsx,js,jsx,graphql,vue}"],
      tagName: 'gql',
      addTypename: true,
      service: {
        name: 'webauthn-project',
        localSchemaFile: 'schema.graphql'
      }
    },
  };