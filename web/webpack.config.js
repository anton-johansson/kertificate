const webpack = require('webpack');

module.exports = env => ({
    entry: './src/index.js',
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ['babel-loader']
            }
        ]
    },
    resolve: {
        extensions: ['*', '.js', '.jsx']
    },
    output: {
        path: __dirname + '/dist',
        publicPath: '/',
        filename: 'bundle.js'
    },
    plugins: [
        new webpack.HotModuleReplacementPlugin(),
        new webpack.DefinePlugin({'process.env.API_BASE': JSON.stringify(env.API_BASE)}),
    ],
    devServer: {
        contentBase: './dist',
        hot: true,
        port: 8000
    }
});
