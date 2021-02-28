// vue.config.js
module.exports = {
    // options...
    devServer: {
        proxy: {
            // '/api/': 'http://192.168.31.134:4000/api/',
            '/api': {
                target: 'http://192.168.31.134:4000/api',
                ws: true,
                changeOrigin: true
              },
            // "^/api": {
            //     target: "http://192.168.31.134:4000",
            //     changeOrigin: true,
            //     logLevel: "debug",
            //     // pathRewrite: { "^/api": "/" }
            //   }
        }
        //   proxy: 'http://backend.test/',
      }
  }