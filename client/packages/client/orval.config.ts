module.exports = {
    petstore: {
      output: {
        mode: 'tags-split',
        target: 'src/petstore.ts',
        schemas: 'src/model',
        client: 'react-query', 
      },
      input: {
        target: '../../../open-api-schema.yaml',
      },
    },
  };