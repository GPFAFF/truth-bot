// connect to API
const Twitter = require('twitter');
const Sheet = require('./sheet');
require('dotenv').config();

const client = new Twitter({
  consumer_key: process.env.CONSUMER_KEY,
  consumer_secret: process.env.CONSUMER_SECRET,
  access_token_key: process.env.ACCESS_TOKEN,
  access_token_secret: process.env.ACCESS_TOKEN_SECRET
});

// pull tweet from spreadsheet
const mount = async () => {
    const sheet = new Sheet(process.env.SHEET);
    await sheet.load();
    const quotes = await sheet.getRows();
    const [status, name] = quotes[0]._rawData;

    client.post('statuses/update', {status: `${status} - ${name}`}, (error, _tweet, _response) => {
      if (error) throw error;
    });

    await quotes[0].delete();
    console.log("🎉 🎉 🎉 SUCCESS! 🎉 🎉 🎉");
};

mount();