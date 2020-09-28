const fs = require('fs');

module.exports = {
  productionSourceMap: false,

  devServer: {
    port: 7203,
    // https: {
    //   key: fs.readFileSync('../certs/localhost.key'),
    //   cert: fs.readFileSync('../certs/localhost.cert'),
    // },
    disableHostCheck: true,
    proxy: {
      '/api': {
        target: 'http://localhost:7201',
        pathRewrite: { '^/api': '' },
      },
    },
    public: 'sheet.dev',
  },

  pluginOptions: {
    quasar: {
      importStrategy: 'manual',
      rtlSupport: false,
    },
  },

  transpileDependencies: ['quasar'],
};
