const fs = require('fs');
let start = Date.now()
const allFileContents = fs.readFileSync('../input/GTFS/stop_times.txt', 'utf-8');
let lines = []
allFileContents.split(/\r?\n/).forEach(line =>  {
  lines.push(line);
});
let end = Date.now()
console.log(lines.length)

const used = process.memoryUsage().heapUsed / 1024 / 1024;
console.log(end - start);