const webpack = require('webpack');
module.exports = function override(config) {
    const fallback = config.resolve.fallback || {
        process: false
    };
    Object.assign(fallback, {
        "crypto": require.resolve("crypto-browserify"),
        "stream": require.resolve("stream-browserify"),
        "assert": require.resolve("assert"),
        "http": require.resolve("stream-http"),
        "https": require.resolve("https-browserify"),
        "os": require.resolve("os-browserify"),
        "url": require.resolve("url"),
        "path": require.resolve("path-browserify")
    })
    config.resolve.fallback = fallback;
    config.plugins = (config.plugins || []).concat([
        new webpack.ProvidePlugin({
            // process: 'process/browser.ts',
            Buffer: ['buffer', 'Buffer']
        })
    ])
    return config; }
