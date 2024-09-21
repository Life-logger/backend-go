const {defineConfig} = require('cypress');

module.exports = defineConfig({
  e2e: {
    baseUrl: 'http://localhost:22250',
    supportFile: false,
    screenshotOnRunFailure: false,
    video: false,
    responseTimeout: 150000, // 2m 30
  },
});

