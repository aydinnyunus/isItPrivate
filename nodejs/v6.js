const ip = require('ip');

const ipAddress = '::1';
console.log("Language: Node.js");
console.log(`=== Checking IP: ${ipAddress} ===`);
console.log(`Is Private: ${ip.isPrivate(ipAddress)}`);
console.log(`Is Loopback: ${ip.isLoopback(ipAddress)}`);
console.log();
