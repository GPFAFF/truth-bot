const { GoogleSpreadsheet } = require('google-spreadsheet');

module.exports = class Sheet {
  constructor(doc) {
    this.doc = new GoogleSpreadsheet(doc);
  }

  async load() {
    await this.doc.useServiceAccountAuth(require('./google-credentials.json'));
    await this.doc.loadInfo();
  }

  async getRows() {
    const sheet = this.doc.sheetsById[0];
    return await sheet.getRows();
  }
};