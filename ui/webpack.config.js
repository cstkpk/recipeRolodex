const path = require('path');
const fs = require('fs');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const autoprefixer = require('autoprefixer');

const appDirectory = fs.realpathSync(process.cwd());

// Gets absolute path of file within app directory
const resolveAppPath = relativePath => path.resolve(appDirectory, relativePath);

const host = process.env.HOST || 'localhost';

// Required for babel-preset-react-app
process.env.NODE_ENV = 'development';

module.exports = {
  mode: 'development',
  
  // Entry point of app
  entry: resolveAppPath('src'),

  output: {
    // Development filename output
    filename: 'static/js/bundle.js',
    // Below two lines needed because of webpack v5 error: https://github.com/webpack/webpack/issues/11660
    chunkLoading: false,
    wasmLoading: false,
  },

  devServer: {
    contentBase: resolveAppPath('public'),
    compress: true,
    historyApiFallback: true,
    hot: true,
    host,
    port: 3000,
    publicPath: '/',
  },

  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.jsx', '.css', '.scss']
  },

  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        include: resolveAppPath('src'),
        loader: 'babel-loader',
        options: {
          presets: [
            // Preset includes JSX, TypeScript, and some ESnext features
            require.resolve('babel-preset-react-app'),
          ]
        }
      },
      {
        test: /\.tsx?$/,
        exclude: /node_modules/,
        use: ['babel-loader', 'ts-loader'],
      },
      {
        test: /\.scss$/,
        use: [
          { loader: 'style-loader' },
          {
            loader: 'css-loader',
            options: {
              importLoaders: 2,
              // localsConvention: 'asIs', // no longer supported?
              modules: true,
              sourceMap: true,
            },
          },
          {
            loader: 'postcss-loader',
            options: {
              postcssOptions: {
                plugins: [autoprefixer({grid: 'no-autoplace'})],
              }
            },
          },
          {
            loader: 'sass-loader',
          },
        ]
      },
      {
        test: /\.(png|jp(e*)g|svg)$/,
        use: [{
          loader: 'url-loader',
          options: {
            limit: 8000,
            name: 'images/[hash]-[name].[ext]'
          }
        }]
      }
    ],
  },

  plugins: [
    // Re-generate index.html with injected script tag
    // The injected script tag contains a src value of the filename output defined above
    new HtmlWebpackPlugin({
      inject: true,
      template: resolveAppPath('public/index.html'),
    }),
  ],
}