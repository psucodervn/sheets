module.exports = {
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
