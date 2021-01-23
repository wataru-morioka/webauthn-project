workbox.core.setCacheNameDetails({prefix: "v2"});

// // First, import the library into the service worker global scope:
// importScripts('node_modules/sw-offline-google-analytics/build/importScripts/sw-offline-google-analytics.prod.v0.0.25.js');

// // Then, call goog.offlineGoogleAnalytics.initialize():
// // See https://github.com/GoogleChrome/workbox/tree/master/packages/workbox-google-analytics
// goog.offlineGoogleAnalytics.initialize();
workbox.skipWaiting();
workbox.clientsClaim();
/**
 * The workboxSW.precacheAndRoute() method efficiently caches and responds to
 * requests for URLs in the manifest.
 * See https://goo.gl/S9QRab
 */
self.__precacheManifest = [].concat(self.__precacheManifest || []);
workbox.precaching.suppressWarnings();
workbox.precaching.precacheAndRoute(self.__precacheManifest, {});
workbox.routing.registerRoute(/.*firebase.*/, workbox.strategies.networkFirst({ "cacheName":"api", plugins: [new workbox.expiration.Plugin({"maxAgeSeconds":86400,"purgeOnQuotaError":false})] }), 'GET');
workbox.routing.registerRoute(/\.(jpg|png|svg|woff|ttf|eot)/, workbox.strategies.cacheFirst({ "cacheName":"assets", plugins: [new workbox.expiration.Plugin({"maxAgeSeconds":86400,"purgeOnQuotaError":false})] }), 'GET');
workbox.routing.registerRoute(/\.html/, workbox.strategies.networkFirst({ "cacheName":"html", plugins: [new workbox.expiration.Plugin({"maxAgeSeconds":86400,"purgeOnQuotaError":false})] }), 'GET');
