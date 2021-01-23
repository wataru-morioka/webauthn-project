const path = require('path');
module.exports = {
    // publicPath: '/tenth-project/',
    // devServer: {
    //     proxy: {
    //         "/oauth2": {
    //         target: 'https://webauthn-app.auth.ap-northeast-1.amazoncognito.com',
    //       },
    //         "/o": {
    //         target: 'https://accounts.google.com',
    //       }
    //     }
    // },
    publicPath: '/',
    outputDir: 'docs',
    pwa: {
        workboxPluginMode: 'InjectManifest',
        workboxOptions: {
            // cacheId: 'v2',
            swSrc: __dirname + '/src/service-worker.js',
            swDest: path.join(__dirname, '/docs/service-worker.js'),
            // globDirectory: path.join(__dirname, '/docs/'),
            // globPatterns: ['*.{js,css}', 'img/**/*'],
            // globIgnores: [
            //     'index.html',
            //   ],
            // globPatterns: ['**/*.{js,css,png}'],
            // globIgnores: ['**/*.html'],
            // clientsClaim: true,
            // skipWaiting: true,
            // runtimeCaching: [
            //     {
            //         urlPattern: /.*firebase.*/,
            //         handler: 'networkFirst',
            //         options: {
            //             cacheName: 'api',
            //             expiration: {
            //                 maxAgeSeconds: 60 * 60 * 24,
            //             },
            //         }
            //     },
            //     {
            //         urlPattern: /\.(png|svg|woff|ttf|eot)/,
            //         handler: 'cacheFirst',
            //         options: {
            //             cacheName: 'assets',
            //             expiration: {
            //                 maxAgeSeconds: 60 * 60 * 24,
            //             },
            //         }
            //     },
            //     {
            //         urlPattern: /\.html/,
            //         handler: 'networkFirst',
            //         options: {
            //             cacheName: 'html',
            //             expiration: {
            //                 maxAgeSeconds: 60 * 60 * 24,
            //             },
            //         }
            //     }       
            // ]      
        }
      }
}
