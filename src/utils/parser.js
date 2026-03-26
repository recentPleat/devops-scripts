const fs = require('fs');
const path = require('path');

const parser = {
  parse: function (filePath) {
    if (!fs.existsSync(filePath)) {
      throw new Error(`File ${filePath} not found`);
    }

    const fileContent = fs.readFileSync(filePath, 'utf8');
    const lines = fileContent.split('\n');

    const result = lines.reduce((acc, line) => {
      const trimmedLine = line.trim();
      if (trimmedLine && trimmedLine.startsWith('#')) {
        const comment = trimmedLine.slice(1);
        const [key, value] = comment.split('=');
        acc[key.trim()] = value.trim();
      }
      return acc;
    }, {});

    return result;
  },
};

module.exports = parser;