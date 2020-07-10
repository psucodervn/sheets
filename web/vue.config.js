module.exports = {
  productionSourceMap: false,

  devServer: {
    disableHostCheck: true,
  },

  pluginOptions: {
    quasar: {
      importStrategy: 'manual',
      rtlSupport: false,
    },
  },

  transpileDependencies: [
    'quasar',
  ],
};
